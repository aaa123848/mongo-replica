package application

import (
	"context"
	"mongotest/mongotool"
)

var mt mongotool.MongoTool = mongotool.MongoTool{}
var ctx context.Context = context.Background()
