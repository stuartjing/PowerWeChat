package response

import (
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/power"
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/response"
)

type ResponseMomentGetMomentComments struct {
	response.ResponseWork

	CommentList []*power.HashMap `json:"comment_list"`
	LikeList    []*power.HashMap `json:"like_list"`
	NextCursor  string           `json:"next_cursor"`
}
