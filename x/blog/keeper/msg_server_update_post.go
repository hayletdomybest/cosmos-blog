package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/james/blog/x/blog/types"
)

func (k msgServer) UpdatePost(goCtx context.Context, msg *types.MsgUpdatePost) (*types.MsgUpdatePostResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	err := k.Keeper.UpdatePost(ctx, msg)
	if err != nil {
		return nil, err
	}

	return &types.MsgUpdatePostResponse{}, nil
}
