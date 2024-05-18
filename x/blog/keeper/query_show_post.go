package keeper

import (
	"context"
	"fmt"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdktypes "github.com/cosmos/cosmos-sdk/types"
	"github.com/james/blog/tools"
	"github.com/james/blog/x/blog/types"
	blogtypes "github.com/james/blog/x/blog/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ShowPost(goCtx context.Context, req *types.QueryShowPostRequest) (*types.QueryShowPostResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdktypes.UnwrapSDKContext(goCtx)
	adapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))

	store := prefix.NewStore(adapter, blogtypes.KeyPrefix(blogtypes.PostKey))
	bz := store.Get(tools.Uint64ToBytes(req.Id))

	if bz == nil {
		errorStr := fmt.Sprintf("post with ID %d not found", req.Id)
		return nil, status.Error(codes.NotFound, errorStr)
	}

	var post blogtypes.Post
	if err := k.cdc.Unmarshal(bz, &post); err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryShowPostResponse{
		Post: post,
	}, nil
}
