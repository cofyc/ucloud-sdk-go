// Code is generated by ucloud-model, DO NOT EDIT IT.

package uaccount

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// UAccount API Schema

// CreateCharacterRequest is request schema for CreateCharacter action
type CreateCharacterRequest struct {
	request.CommonBase

	// 角色对产品的权限(增)
	Add []string `required:"false"`

	// 角色描述
	CharacterDescription *string `required:"false"`

	// 角色名称，不得与现有角色重名
	CharacterName *string `required:"true"`

	// 角色对产品的权限(删)
	Del []string `required:"false"`

	// 角色对产品的权限(查)
	Get []string `required:"false"`

	// 角色对产品的权限(改)
	Mod []string `required:"false"`
}

// CreateCharacterResponse is response schema for CreateCharacter action
type CreateCharacterResponse struct {
	response.CommonBase

	// 角色ID
	CharacterId string
}

// NewCreateCharacterRequest will create request of CreateCharacter action.
func (c *UAccountClient) NewCreateCharacterRequest() *CreateCharacterRequest {
	req := &CreateCharacterRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(false)
	return req
}

/*
API: CreateCharacter

创建角色
*/
func (c *UAccountClient) CreateCharacter(req *CreateCharacterRequest) (*CreateCharacterResponse, error) {
	var err error
	var res CreateCharacterResponse

	reqCopier := *req

	err = c.Client.InvokeAction("CreateCharacter", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// CreateProjectRequest is request schema for CreateProject action
type CreateProjectRequest struct {
	request.CommonBase

	//
	ParentId *string `required:"false"`

	//
	ProjectName *string `required:"true"`
}

// CreateProjectResponse is response schema for CreateProject action
type CreateProjectResponse struct {
	response.CommonBase

	//
	ProjectId string
}

// NewCreateProjectRequest will create request of CreateProject action.
func (c *UAccountClient) NewCreateProjectRequest() *CreateProjectRequest {
	req := &CreateProjectRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(false)
	return req
}

/*
API: CreateProject
*/
func (c *UAccountClient) CreateProject(req *CreateProjectRequest) (*CreateProjectResponse, error) {
	var err error
	var res CreateProjectResponse

	reqCopier := *req

	err = c.Client.InvokeAction("CreateProject", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// DescribeCharacterListRequest is request schema for DescribeCharacterList action
type DescribeCharacterListRequest struct {
	request.CommonBase

	// 角色列表的最大数量，默认为20
	Limit *int `required:"false"`

	// 角色列表的偏移量，默认为0
	Offset *int `required:"false"`
}

// DescribeCharacterListResponse is response schema for DescribeCharacterList action
type DescribeCharacterListResponse struct {
	response.CommonBase

	// JSON格式的角色列表实例，每项参数参见下面的 ResponseItem
	CharacterSet []CharacterSet

	// 角色总数
	TotalCount int
}

// NewDescribeCharacterListRequest will create request of DescribeCharacterList action.
func (c *UAccountClient) NewDescribeCharacterListRequest() *DescribeCharacterListRequest {
	req := &DescribeCharacterListRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: DescribeCharacterList

获取角色列表
*/
func (c *UAccountClient) DescribeCharacterList(req *DescribeCharacterListRequest) (*DescribeCharacterListResponse, error) {
	var err error
	var res DescribeCharacterListResponse

	reqCopier := *req

	err = c.Client.InvokeAction("DescribeCharacterList", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// DescribeMemberListRequest is request schema for DescribeMemberList action
type DescribeMemberListRequest struct {
	request.CommonBase

	// [公共参数] 项目ID，请参考[GetProjectList接口](../summary/get_project_list.html)。不填写为查询所有项目。
	// ProjectId *string `required:"false"`

	// 成员列表的最大数量，默认为200
	Limit *string `required:"false"`

	// 成员列表的偏移量，默认为0
	Offset *string `required:"false"`
}

// DescribeMemberListResponse is response schema for DescribeMemberList action
type DescribeMemberListResponse struct {
	response.CommonBase

	// JSON格式的成员列表实例
	MemberSet []MemberInfo

	// 成员总数
	TotalCount int
}

// NewDescribeMemberListRequest will create request of DescribeMemberList action.
func (c *UAccountClient) NewDescribeMemberListRequest() *DescribeMemberListRequest {
	req := &DescribeMemberListRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: DescribeMemberList

获取成员列表，限主账号使用。
*/
func (c *UAccountClient) DescribeMemberList(req *DescribeMemberListRequest) (*DescribeMemberListResponse, error) {
	var err error
	var res DescribeMemberListResponse

	reqCopier := *req

	err = c.Client.InvokeAction("DescribeMemberList", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// GetNetworkMaskRequest is request schema for GetNetworkMask action
type GetNetworkMaskRequest struct {
	request.CommonBase
}

// GetNetworkMaskResponse is response schema for GetNetworkMask action
type GetNetworkMaskResponse struct {
	response.CommonBase

	// 接口返回数据
	Data NetworkMask

	// 接口信息，成功时为`success`，错误时显示具体错误信息。
	Message string
}

// NewGetNetworkMaskRequest will create request of GetNetworkMask action.
func (c *UAccountClient) NewGetNetworkMaskRequest() *GetNetworkMaskRequest {
	req := &GetNetworkMaskRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: GetNetworkMask

查询登录与API调用的网络掩码
*/
func (c *UAccountClient) GetNetworkMask(req *GetNetworkMaskRequest) (*GetNetworkMaskResponse, error) {
	var err error
	var res GetNetworkMaskResponse

	reqCopier := *req

	err = c.Client.InvokeAction("GetNetworkMask", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// GetProjectListRequest is request schema for GetProjectList action
type GetProjectListRequest struct {
	request.CommonBase

	// 是否是财务账号（Yes：是，No：否）
	IsFinance *string `required:"false"`
}

// GetProjectListResponse is response schema for GetProjectList action
type GetProjectListResponse struct {
	response.CommonBase

	// 项目总数
	ProjectCount int

	// JSON格式的项目列表实例
	ProjectSet []ProjectListInfo
}

// NewGetProjectListRequest will create request of GetProjectList action.
func (c *UAccountClient) NewGetProjectListRequest() *GetProjectListRequest {
	req := &GetProjectListRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: GetProjectList

获取项目列表
*/
func (c *UAccountClient) GetProjectList(req *GetProjectListRequest) (*GetProjectListResponse, error) {
	var err error
	var res GetProjectListResponse

	reqCopier := *req

	err = c.Client.InvokeAction("GetProjectList", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// GetRegionRequest is request schema for GetRegion action
type GetRegionRequest struct {
	request.CommonBase
}

// GetRegionResponse is response schema for GetRegion action
type GetRegionResponse struct {
	response.CommonBase

	// 各数据中心信息
	Regions []RegionInfo
}

// NewGetRegionRequest will create request of GetRegion action.
func (c *UAccountClient) NewGetRegionRequest() *GetRegionRequest {
	req := &GetRegionRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: GetRegion

获取用户在各数据中心的权限等信息
*/
func (c *UAccountClient) GetRegion(req *GetRegionRequest) (*GetRegionResponse, error) {
	var err error
	var res GetRegionResponse

	reqCopier := *req

	err = c.Client.InvokeAction("GetRegion", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// GetUserInfoRequest is request schema for GetUserInfo action
type GetUserInfoRequest struct {
	request.CommonBase
}

// GetUserInfoResponse is response schema for GetUserInfo action
type GetUserInfoResponse struct {
	response.CommonBase

	//
	DataSet []UserInfo
}

// NewGetUserInfoRequest will create request of GetUserInfo action.
func (c *UAccountClient) NewGetUserInfoRequest() *GetUserInfoRequest {
	req := &GetUserInfoRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(false)
	return req
}

/*
API: GetUserInfo
*/
func (c *UAccountClient) GetUserInfo(req *GetUserInfoRequest) (*GetUserInfoResponse, error) {
	var err error
	var res GetUserInfoResponse

	reqCopier := *req

	err = c.Client.InvokeAction("GetUserInfo", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// InviteSubaccountRequest is request schema for InviteSubaccount action
type InviteSubaccountRequest struct {
	request.CommonBase

	// 是否有财务权限(true:是,false:否,默认为否)
	IsFinance *string `required:"true"`

	// 受邀成员邮箱地址，不得重复
	UserEmail *string `required:"true"`

	// 受邀成员姓名
	UserName *string `required:"true"`

	// 受邀成员手机号码
	UserPhone *string `required:"true"`
}

// InviteSubaccountResponse is response schema for InviteSubaccount action
type InviteSubaccountResponse struct {
	response.CommonBase
}

// NewInviteSubaccountRequest will create request of InviteSubaccount action.
func (c *UAccountClient) NewInviteSubaccountRequest() *InviteSubaccountRequest {
	req := &InviteSubaccountRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: InviteSubaccount

邀请子帐号成员
*/
func (c *UAccountClient) InviteSubaccount(req *InviteSubaccountRequest) (*InviteSubaccountResponse, error) {
	var err error
	var res InviteSubaccountResponse

	reqCopier := *req

	err = c.Client.InvokeAction("InviteSubaccount", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// ModifyCharacterRequest is request schema for ModifyCharacter action
type ModifyCharacterRequest struct {
	request.CommonBase

	// 角色权限(增)
	Add []string `required:"false"`

	// 角色描述
	CharacterDescription *string `required:"false"`

	// 角色ID
	CharacterId *string `required:"true"`

	// 新角色名称
	CharacterName *string `required:"false"`

	// 角色权限(删)
	Del []string `required:"false"`

	// 角色权限(查)
	Get []string `required:"false"`

	// 角色权限(改)
	Mod []string `required:"false"`
}

// ModifyCharacterResponse is response schema for ModifyCharacter action
type ModifyCharacterResponse struct {
	response.CommonBase
}

// NewModifyCharacterRequest will create request of ModifyCharacter action.
func (c *UAccountClient) NewModifyCharacterRequest() *ModifyCharacterRequest {
	req := &ModifyCharacterRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: ModifyCharacter

修改角色
*/
func (c *UAccountClient) ModifyCharacter(req *ModifyCharacterRequest) (*ModifyCharacterResponse, error) {
	var err error
	var res ModifyCharacterResponse

	reqCopier := *req

	err = c.Client.InvokeAction("ModifyCharacter", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// ModifyProjectRequest is request schema for ModifyProject action
type ModifyProjectRequest struct {
	request.CommonBase

	//
	ProjectName *string `required:"true"`
}

// ModifyProjectResponse is response schema for ModifyProject action
type ModifyProjectResponse struct {
	response.CommonBase
}

// NewModifyProjectRequest will create request of ModifyProject action.
func (c *UAccountClient) NewModifyProjectRequest() *ModifyProjectRequest {
	req := &ModifyProjectRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(false)
	return req
}

/*
API: ModifyProject
*/
func (c *UAccountClient) ModifyProject(req *ModifyProjectRequest) (*ModifyProjectResponse, error) {
	var err error
	var res ModifyProjectResponse

	reqCopier := *req

	err = c.Client.InvokeAction("ModifyProject", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// RemoveMemberFromProjectRequest is request schema for RemoveMemberFromProject action
type RemoveMemberFromProjectRequest struct {
	request.CommonBase

	// [公共参数] 项目ID，请参考[GetProjectList接口](../summary/get_project_list.html)的描述。不填写为默认项目，子帐号必须填写。
	// ProjectId *string `required:"true"`

	// 需要被移除成员Email
	MemberEmail *string `required:"true"`
}

// RemoveMemberFromProjectResponse is response schema for RemoveMemberFromProject action
type RemoveMemberFromProjectResponse struct {
	response.CommonBase
}

// NewRemoveMemberFromProjectRequest will create request of RemoveMemberFromProject action.
func (c *UAccountClient) NewRemoveMemberFromProjectRequest() *RemoveMemberFromProjectRequest {
	req := &RemoveMemberFromProjectRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: RemoveMemberFromProject

从项目中移除成员
*/
func (c *UAccountClient) RemoveMemberFromProject(req *RemoveMemberFromProjectRequest) (*RemoveMemberFromProjectResponse, error) {
	var err error
	var res RemoveMemberFromProjectResponse

	reqCopier := *req

	err = c.Client.InvokeAction("RemoveMemberFromProject", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// SetNetworkMaskRequest is request schema for SetNetworkMask action
type SetNetworkMaskRequest struct {
	request.CommonBase

	// API调用网络掩码，多个IP以英文逗号分隔。默认空字符串，不限制登录IP。
	APINetworkMask *string `required:"false"`

	// 短信验证码
	Code *string `required:"true"`

	// 登录网络掩码，多个IP以英文逗号分隔。默认空字符串，不限制登录IP。
	LoginNetworkMask *string `required:"false"`
}

// SetNetworkMaskResponse is response schema for SetNetworkMask action
type SetNetworkMaskResponse struct {
	response.CommonBase

	// 接口信息，成功时为`success`，错误时显示具体错误信息。
	Message string
}

// NewSetNetworkMaskRequest will create request of SetNetworkMask action.
func (c *UAccountClient) NewSetNetworkMaskRequest() *SetNetworkMaskRequest {
	req := &SetNetworkMaskRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: SetNetworkMask

设置登录与API调用的网络掩码
*/
func (c *UAccountClient) SetNetworkMask(req *SetNetworkMaskRequest) (*SetNetworkMaskResponse, error) {
	var err error
	var res SetNetworkMaskResponse

	reqCopier := *req

	err = c.Client.InvokeAction("SetNetworkMask", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// TerminateCharacterRequest is request schema for TerminateCharacter action
type TerminateCharacterRequest struct {
	request.CommonBase

	// 角色ID，使用[DescribeCharacterList接口](../summary/describe_character_list.html) 获取角色ID
	CharacterId *string `required:"true"`
}

// TerminateCharacterResponse is response schema for TerminateCharacter action
type TerminateCharacterResponse struct {
	response.CommonBase
}

// NewTerminateCharacterRequest will create request of TerminateCharacter action.
func (c *UAccountClient) NewTerminateCharacterRequest() *TerminateCharacterRequest {
	req := &TerminateCharacterRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: TerminateCharacter

删除用户角色管理列表中的指定角色
*/
func (c *UAccountClient) TerminateCharacter(req *TerminateCharacterRequest) (*TerminateCharacterResponse, error) {
	var err error
	var res TerminateCharacterResponse

	reqCopier := *req

	err = c.Client.InvokeAction("TerminateCharacter", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// TerminateMemberRequest is request schema for TerminateMember action
type TerminateMemberRequest struct {
	request.CommonBase

	// 用户邮箱
	MemberEmail *string `required:"true"`
}

// TerminateMemberResponse is response schema for TerminateMember action
type TerminateMemberResponse struct {
	response.CommonBase
}

// NewTerminateMemberRequest will create request of TerminateMember action.
func (c *UAccountClient) NewTerminateMemberRequest() *TerminateMemberRequest {
	req := &TerminateMemberRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: TerminateMember

删除管理员人员管理页面的指定子账号
*/
func (c *UAccountClient) TerminateMember(req *TerminateMemberRequest) (*TerminateMemberResponse, error) {
	var err error
	var res TerminateMemberResponse

	reqCopier := *req

	err = c.Client.InvokeAction("TerminateMember", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}

// TerminateProjectRequest is request schema for TerminateProject action
type TerminateProjectRequest struct {
	request.CommonBase

	// [公共参数] 项目ID，请参考[GetProjectList接口](../summary/get_project_list.html)的描述。
	// ProjectId *string `required:"true"`

}

// TerminateProjectResponse is response schema for TerminateProject action
type TerminateProjectResponse struct {
	response.CommonBase
}

// NewTerminateProjectRequest will create request of TerminateProject action.
func (c *UAccountClient) NewTerminateProjectRequest() *TerminateProjectRequest {
	req := &TerminateProjectRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

/*
API: TerminateProject

删除项目
*/
func (c *UAccountClient) TerminateProject(req *TerminateProjectRequest) (*TerminateProjectResponse, error) {
	var err error
	var res TerminateProjectResponse

	reqCopier := *req

	err = c.Client.InvokeAction("TerminateProject", &reqCopier, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}
