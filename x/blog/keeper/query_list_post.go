package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/james/blog/x/blog/types"
	blogtypes "github.com/james/blog/x/blog/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) ListPost(goCtx context.Context, req *types.QueryListPostRequest) (*types.QueryListPostResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	adapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(goCtx))
	store := prefix.NewStore(adapter, blogtypes.KeyPrefix(blogtypes.PostKey))

	var posts []types.Post

	pageRes, err := query.Paginate(store, req.Pagination, func(key, value []byte) error {
		var post types.Post
		if err := k.cdc.Unmarshal(value, &post); err != nil {
			return err
		}

		posts = append(posts, post)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryListPostResponse{
		Posts:      posts,
		Pagination: pageRes,
	}, nil
}
