package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/james/blog/x/blog/types"
)

func (k msgServer) CreatePost(goCtx context.Context, msg *types.MsgCreatePost) (*types.MsgCreatePostResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	post := &types.Post{
		Title:   msg.Title,
		Body:    msg.Body,
		Creator: msg.Creator,
	}

	id := k.AppendPost(ctx, post)

	return &types.MsgCreatePostResponse{
		Id: id,
	}, nil
}
