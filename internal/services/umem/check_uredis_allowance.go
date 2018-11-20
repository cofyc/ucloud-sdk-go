//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UMem CheckURedisAllowance

package umem

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// CheckURedisAllowanceRequest is request schema for CheckURedisAllowance action
type CheckURedisAllowanceRequest struct {
	request.CommonBase

	// 可用区。参见 [可用区列表](../summary/regionlist.html)
	Zone *string `required:"true"`

	// 创建实例的容量大小, 单位:GB 目前仅支持1/2/4/8/16/32六种规格
	Size *string `required:"true"`

	// 创建实例的数量，[1-10]
	Count *int `required:"true"`

	// 是否是跨机房URedis(默认false)
	RegionFlag *bool `required:"false"`

	//
	SlaveZone *string `required:"false"`
}

// CheckURedisAllowanceResponse is response schema for CheckURedisAllowance action
type CheckURedisAllowanceResponse struct {
	response.CommonBase

	// 可创建的数量
	Count int
}

// NewCheckURedisAllowanceRequest will create request of CheckURedisAllowance action.
func (c *UMemClient) NewCheckURedisAllowanceRequest() *CheckURedisAllowanceRequest {
	req := &CheckURedisAllowanceRequest{}

	// setup request with client config
	c.client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// CheckURedisAllowance - 检查主备Redis的资源是否足够创建新实例
func (c *UMemClient) CheckURedisAllowance(req *CheckURedisAllowanceRequest) (*CheckURedisAllowanceResponse, error) {
	var err error
	var res CheckURedisAllowanceResponse

	err = c.client.InvokeAction("CheckURedisAllowance", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}