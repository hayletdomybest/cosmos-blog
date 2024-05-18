package keeper

import (
	"encoding/binary"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdktypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/james/blog/common"
	"github.com/james/blog/tools"
	blogtypes "github.com/james/blog/x/blog/types"
)

func (k Keeper) AppendPost(ctx sdktypes.Context, post *blogtypes.Post) uint64 {
	adapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))

	store := prefix.NewStore(adapter, blogtypes.KeyPrefix(blogtypes.PostKey))

	count := k.GetPostCount(ctx)

	post.Id = count + 1
	k.SetPostCount(ctx, post.Id)

	value := k.cdc.MustMarshal(post)
	store.Set(tools.Uint64ToBytes(post.Id), value)

	return post.Id
}

func (k Keeper) UpdatePost(ctx sdktypes.Context, req *blogtypes.MsgUpdatePost) error {
	adapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))

	store := prefix.NewStore(adapter, blogtypes.KeyPrefix(blogtypes.PostKey))

	bz := store.Get(tools.Uint64ToBytes(req.Id))
	if bz == nil {
		return common.ErrNotFound
	}

	var post blogtypes.Post
	k.cdc.Unmarshal(bz, &post)

	post.Creator = req.Creator
	post.Body = req.Body
	post.Title = req.Title

	store.Set(tools.Uint64ToBytes(req.Id), k.cdc.MustMarshal(&post))

	return nil
}

func (k Keeper) DeletePost(ctx sdktypes.Context, req *blogtypes.MsgDeletePost) error {
	adapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))

	store := prefix.NewStore(adapter, blogtypes.KeyPrefix(blogtypes.PostKey))

	existed := store.Has(tools.Uint64ToBytes(req.Id))
	if !existed {
		return common.ErrNotFound
	}

	count := k.GetPostCount(ctx)
	store.Delete(tools.Uint64ToBytes(req.Id))
	k.SetPostCount(ctx, count-1)

	return nil
}

func (k Keeper) GetPostCount(ctx sdktypes.Context) uint64 {
	adapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))

	store := prefix.NewStore(adapter, []byte{})

	bz := store.Get(blogtypes.KeyPrefix(blogtypes.GetPostCountKey))

	return binary.BigEndian.Uint64(bz)
}

func (k Keeper) SetPostCount(ctx sdktypes.Context, count uint64) {
	adapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))

	store := prefix.NewStore(adapter, []byte{})

	store.Set(blogtypes.KeyPrefix(blogtypes.GetPostCountKey), tools.Uint64ToBytes(count))
}
