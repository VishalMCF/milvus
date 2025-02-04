package writebuffer

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"github.com/milvus-io/milvus-proto/go-api/v2/commonpb"
	"github.com/milvus-io/milvus-proto/go-api/v2/msgpb"
	"github.com/milvus-io/milvus-proto/go-api/v2/schemapb"
	"github.com/milvus-io/milvus/internal/datanode/allocator"
	"github.com/milvus-io/milvus/internal/datanode/metacache"
	"github.com/milvus-io/milvus/internal/datanode/syncmgr"
	"github.com/milvus-io/milvus/pkg/common"
	"github.com/milvus-io/milvus/pkg/util/merr"
	"github.com/milvus-io/milvus/pkg/util/paramtable"
	"github.com/milvus-io/milvus/pkg/util/tsoutil"
)

type ManagerSuite struct {
	suite.Suite
	collID      int64
	channelName string
	collSchema  *schemapb.CollectionSchema
	syncMgr     *syncmgr.MockSyncManager
	metacache   *metacache.MockMetaCache
	allocator   *allocator.MockAllocator

	manager *bufferManager
}

func (s *ManagerSuite) SetupSuite() {
	paramtable.Get().Init(paramtable.NewBaseTable())
	s.collID = 100
	s.collSchema = &schemapb.CollectionSchema{
		Name: "test_collection",
		Fields: []*schemapb.FieldSchema{
			{FieldID: common.RowIDField, DataType: schemapb.DataType_Int64, Name: common.RowIDFieldName},
			{FieldID: common.TimeStampField, DataType: schemapb.DataType_Int64, Name: common.TimeStampFieldName},
			{
				FieldID: 100, Name: "pk", DataType: schemapb.DataType_Int64, IsPrimaryKey: true,
			},
			{
				FieldID: 101, Name: "vector", DataType: schemapb.DataType_FloatVector,
				TypeParams: []*commonpb.KeyValuePair{
					{Key: common.DimKey, Value: "128"},
				},
			},
		},
	}

	s.channelName = "by-dev-rootcoord-dml_0_100_v0"
}

func (s *ManagerSuite) SetupTest() {
	s.syncMgr = syncmgr.NewMockSyncManager(s.T())
	s.metacache = metacache.NewMockMetaCache(s.T())
	s.metacache.EXPECT().Collection().Return(s.collID).Maybe()
	s.metacache.EXPECT().Schema().Return(s.collSchema).Maybe()
	s.allocator = allocator.NewMockAllocator(s.T())

	mgr := NewManager(s.syncMgr)
	var ok bool
	s.manager, ok = mgr.(*bufferManager)
	s.Require().True(ok)
}

func (s *ManagerSuite) TestRegister() {
	manager := s.manager

	storageCache, err := metacache.NewStorageV2Cache(s.collSchema)
	s.Require().NoError(err)

	err = manager.Register(s.channelName, s.metacache, storageCache, WithIDAllocator(s.allocator))
	s.NoError(err)

	err = manager.Register(s.channelName, s.metacache, storageCache, WithIDAllocator(s.allocator))
	s.Error(err)
	s.ErrorIs(err, merr.ErrChannelReduplicate)
}

func (s *ManagerSuite) TestFlushSegments() {
	manager := s.manager
	s.Run("channel_not_found", func() {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		err := manager.FlushSegments(ctx, s.channelName, []int64{1, 2, 3})
		s.Error(err, "FlushSegments shall return error when channel not found")
	})

	s.Run("normal_flush", func() {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		wb := NewMockWriteBuffer(s.T())

		s.manager.mut.Lock()
		s.manager.buffers[s.channelName] = wb
		s.manager.mut.Unlock()

		wb.EXPECT().FlushSegments(mock.Anything, mock.Anything).Return(nil)

		err := manager.FlushSegments(ctx, s.channelName, []int64{1})
		s.NoError(err)
	})
}

func (s *ManagerSuite) TestBufferData() {
	manager := s.manager
	s.Run("channel_not_found", func() {
		err := manager.BufferData(s.channelName, nil, nil, nil, nil)
		s.Error(err, "BufferData shall return error when channel not found")
	})

	s.Run("normal_buffer_data", func() {
		wb := NewMockWriteBuffer(s.T())

		s.manager.mut.Lock()
		s.manager.buffers[s.channelName] = wb
		s.manager.mut.Unlock()

		wb.EXPECT().BufferData(mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

		err := manager.BufferData(s.channelName, nil, nil, nil, nil)
		s.NoError(err)
	})
}

func (s *ManagerSuite) TestGetCheckpoint() {
	manager := s.manager
	s.Run("channel_not_found", func() {
		_, _, err := manager.GetCheckpoint(s.channelName)
		s.Error(err, "FlushSegments shall return error when channel not found")
	})

	s.Run("normal_checkpoint", func() {
		wb := NewMockWriteBuffer(s.T())

		manager.mut.Lock()
		manager.buffers[s.channelName] = wb
		manager.mut.Unlock()

		pos := &msgpb.MsgPosition{ChannelName: s.channelName, Timestamp: tsoutil.ComposeTSByTime(time.Now(), 0)}
		wb.EXPECT().GetCheckpoint().Return(pos)
		wb.EXPECT().GetFlushTimestamp().Return(nonFlushTS)
		result, needUpdate, err := manager.GetCheckpoint(s.channelName)
		s.NoError(err)
		s.Equal(pos, result)
		s.False(needUpdate)
	})

	s.Run("checkpoint_need_update", func() {
		wb := NewMockWriteBuffer(s.T())

		manager.mut.Lock()
		manager.buffers[s.channelName] = wb
		manager.mut.Unlock()

		cpTimestamp := tsoutil.ComposeTSByTime(time.Now(), 0)

		pos := &msgpb.MsgPosition{ChannelName: s.channelName, Timestamp: cpTimestamp}
		wb.EXPECT().GetCheckpoint().Return(pos)
		wb.EXPECT().GetFlushTimestamp().Return(cpTimestamp - 1)
		result, needUpdate, err := manager.GetCheckpoint(s.channelName)
		s.NoError(err)
		s.Equal(pos, result)
		s.True(needUpdate)
	})
}

func (s *ManagerSuite) TestRemoveChannel() {
	manager := NewManager(s.syncMgr)

	s.Run("remove_not_exist", func() {
		s.NotPanics(func() {
			manager.RemoveChannel(s.channelName)
		})
	})

	s.Run("remove_channel", func() {
		storageCache, err := metacache.NewStorageV2Cache(s.collSchema)
		s.Require().NoError(err)
		err = manager.Register(s.channelName, s.metacache, storageCache, WithIDAllocator(s.allocator))
		s.Require().NoError(err)

		s.NotPanics(func() {
			manager.RemoveChannel(s.channelName)
		})
	})
}

func TestManager(t *testing.T) {
	suite.Run(t, new(ManagerSuite))
}
