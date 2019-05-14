//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UFS RemoveUFSVolumeMountPoint

package ufs

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// RemoveUFSVolumeMountPointRequest is request schema for RemoveUFSVolumeMountPoint action
type RemoveUFSVolumeMountPointRequest struct {
	request.CommonBase

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

	// 文件系统ID
	VolumeId *string `required:"true"`

	// Vpc ID
	VpcId *string `required:"true"`

	// Subnet ID
	SubnetId *string `required:"true"`
}

// RemoveUFSVolumeMountPointResponse is response schema for RemoveUFSVolumeMountPoint action
type RemoveUFSVolumeMountPointResponse struct {
	response.CommonBase
}

// NewRemoveUFSVolumeMountPointRequest will create request of RemoveUFSVolumeMountPoint action.
func (c *UFSClient) NewRemoveUFSVolumeMountPointRequest() *RemoveUFSVolumeMountPointRequest {
	req := &RemoveUFSVolumeMountPointRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// RemoveUFSVolumeMountPoint - 删除文件系统挂载点
func (c *UFSClient) RemoveUFSVolumeMountPoint(req *RemoveUFSVolumeMountPointRequest) (*RemoveUFSVolumeMountPointResponse, error) {
	var err error
	var res RemoveUFSVolumeMountPointResponse

	err = c.Client.InvokeAction("RemoveUFSVolumeMountPoint", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}