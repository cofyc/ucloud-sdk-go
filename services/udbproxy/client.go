// Code is generated by ucloud-model, DO NOT EDIT IT.

package udbproxy

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
)

// UDBProxyClient is the client of UDBProxy
type UDBProxyClient struct {
	*ucloud.Client
}

// NewClient will return a instance of UDBProxyClient
func NewClient(config *ucloud.Config, credential *auth.Credential) *UDBProxyClient {
	meta := ucloud.ClientMeta{Product: "UDBProxy"}
	client := ucloud.NewClientWithMeta(config, credential, meta)
	return &UDBProxyClient{
		client,
	}
}