package response

import (
	"github.com/stuartjing/PowerWeChat/v3/src/kernel/response"
)

type ResponseUserBatchJobs struct {
	response.ResponseWork

	JobID string `json:"jobid"`
}
