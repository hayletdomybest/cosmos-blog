package keeper

import (
	"github.com/james/blog/x/blog/types"
)

var _ types.QueryServer = Keeper{}
