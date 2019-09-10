//Code is generated by ucloud code generator, don't modify it by hand, it will cause undefined behaviors.
//go:generate ucloud-gen-go-api UHost DescribeUHostTags

package uhost

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud/request"
	"github.com/ucloud/ucloud-sdk-go/ucloud/response"
)

// DescribeUHostTagsRequest is request schema for DescribeUHostTags action
type DescribeUHostTagsRequest struct {
	request.CommonBase

	// [公共参数] 地域。 参见 [地域和可用区列表](../summary/regionlist.html)
	// Region *string `required:"true"`

	// [公共参数] 可用区。参见 [可用区列表](../summary/regionlist.html)
	// Zone *string `required:"false"`

	// [公共参数] 项目ID。不填写为默认项目，子帐号必须填写。 请参考[GetProjectList接口](../summary/get_project_list.html)
	// ProjectId *string `required:"false"`

}

// DescribeUHostTagsResponse is response schema for DescribeUHostTags action
type DescribeUHostTagsResponse struct {
	response.CommonBase

	// 已有主机的业务组总个数
	TotalCount int

	// 业务组集合见 UHostTagSet
	TagSet []UHostTagSet
}

// NewDescribeUHostTagsRequest will create request of DescribeUHostTags action.
func (c *UHostClient) NewDescribeUHostTagsRequest() *DescribeUHostTagsRequest {
	req := &DescribeUHostTagsRequest{}

	// setup request with client config
	c.Client.SetupRequest(req)

	// setup retryable with default retry policy (retry for non-create action and common error)
	req.SetRetryable(true)
	return req
}

// DescribeUHostTags - 获取指定数据中心的业务组列表。
func (c *UHostClient) DescribeUHostTags(req *DescribeUHostTagsRequest) (*DescribeUHostTagsResponse, error) {
	var err error
	var res DescribeUHostTagsResponse

	err = c.Client.InvokeAction("DescribeUHostTags", req, &res)
	if err != nil {
		return &res, err
	}

	return &res, nil
}