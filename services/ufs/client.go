// Code is generated by ucloud-model, DO NOT EDIT IT.

package ufs

import (
	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
)

// UFSClient is the client of UFS
type UFSClient struct {
	*ucloud.Client
}

// NewClient will return a instance of UFSClient
func NewClient(config *ucloud.Config, credential *auth.Credential) *UFSClient {
	meta := ucloud.ClientMeta{Product: "UFS"}
	client := ucloud.NewClientWithMeta(config, credential, meta)
	return &UFSClient{
		client,
	}
}
