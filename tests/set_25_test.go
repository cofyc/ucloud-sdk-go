// Code is generated by ucloud-model, DO NOT EDIT IT.

package tests

import (
	"testing"
	"time"

	"github.com/ucloud/ucloud-sdk-go/internal/utest"
)

func TestSet25(t *testing.T) {
	t.Parallel()

	ctx := utest.NewTestContext()
	ctx.T = t
	ctx.Vars = map[string]interface{}{}

	ctx.SetVar("Region", "cn-sh2")
	ctx.SetVar("Zone", "cn-sh2-02")

	ctx.SetVar("Region", "cn-sh2")
	ctx.SetVar("Zone", "cn-sh2-02")

	testSet25CreateVPC00(&ctx)
	testSet25CreateSubnet01(&ctx)
	testSet25GetVPNGatewayPrice02(&ctx)
	testSet25CreateVPNGateway03(&ctx)
	testSet25AllocateEIP04(&ctx)
	testSet25BindEIP05(&ctx)
	testSet25DescribeVPNGateway06(&ctx)
	testSet25GetVPNGatewayUpgradePrice07(&ctx)
	testSet25UpdateVPNGateway08(&ctx)
	testSet25CreateRemoteVPNGateway09(&ctx)
	testSet25DescribeRemoteVPNGateway10(&ctx)
	testSet25CreateVPNTunnel11(&ctx)
	testSet25DescribeVPNTunnel12(&ctx)
	testSet25UpdateVPNTunnelAttribute13(&ctx)
	testSet25DeleteVPNGateway14(&ctx)
	testSet25DeleteRemoteVPNGateway15(&ctx)
	testSet25DeleteVPNTunnel16(&ctx)
	testSet25DeleteVPNGateway17(&ctx)
	testSet25DeleteRemoteVPNGateway18(&ctx)
	testSet25ReleaseEIP19(&ctx)
	testSet25DeleteSubnet20(&ctx)
	testSet25DeleteVPC21(&ctx)
}

func testSet25CreateVPC00(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := vpcClient.NewCreateVPCRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "Network", "192.168.0.0/16"))

	ctx.NoError(utest.SetReqValue(req, "Name", "ipsecvpn-vpc"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return vpcClient.CreateVPC(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", 0, "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Error(err)

	}

	ctx.Vars["vpc_id"] = ctx.Must(utest.GetValue(resp, "VPCId"))
}

func testSet25CreateSubnet01(ctx *utest.TestContext) {
	time.Sleep(time.Duration(10) * time.Second)

	req := vpcClient.NewCreateSubnetRequest()

	ctx.NoError(utest.SetReqValue(req, "VPCId", ctx.GetVar("vpc_id")))

	ctx.NoError(utest.SetReqValue(req, "SubnetName", "ipsecvpn-subnet"))

	ctx.NoError(utest.SetReqValue(req, "Subnet", "192.168.11.0"))

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return vpcClient.CreateSubnet(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", 0, "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Error(err)

	}

	ctx.Vars["subnet_id"] = ctx.Must(utest.GetValue(resp, "SubnetId"))
}

func testSet25GetVPNGatewayPrice02(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := iipsecvpnClient.NewGetVPNGatewayPriceRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "Grade", "Standard"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return iipsecvpnClient.GetVPNGatewayPrice(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", 0, "str_eq"),
		},
		MaxRetries:    0,
		RetryInterval: 0 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Error(err)

	}

}

func testSet25CreateVPNGateway03(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := ipsecvpnClient.NewCreateVPNGatewayRequest()

	ctx.NoError(utest.SetReqValue(req, "VPNGatewayName", "auto_apitest"))

	ctx.NoError(utest.SetReqValue(req, "VPCId", ctx.GetVar("vpc_id")))

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "Grade", "Standard"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return ipsecvpnClient.CreateVPNGateway(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", 0, "str_eq"),
		},
		MaxRetries:    0,
		RetryInterval: 0 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Error(err)

	}

	ctx.Vars["vpngw_id"] = ctx.Must(utest.GetValue(resp, "VPNGatewayId"))
}

func testSet25AllocateEIP04(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := unetClient.NewAllocateEIPRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "OperatorName", "Bgp"))

	ctx.NoError(utest.SetReqValue(req, "Bandwidth", 2))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return unetClient.AllocateEIP(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", 0, "str_eq"),
		},
		MaxRetries:    0,
		RetryInterval: 0 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Error(err)

	}

	ctx.Vars["eip_id"] = ctx.Must(utest.GetValue(resp, "EIPSet.0.EIPId"))
}

func testSet25BindEIP05(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := unetClient.NewBindEIPRequest()

	ctx.NoError(utest.SetReqValue(req, "ResourceType", "vpngw"))

	ctx.NoError(utest.SetReqValue(req, "ResourceId", ctx.GetVar("vpngw_id")))

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "EIPId", ctx.GetVar("eip_id")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return unetClient.BindEIP(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", 0, "str_eq"),
		},
		MaxRetries:    0,
		RetryInterval: 0 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Error(err)

	}

}

func testSet25DescribeVPNGateway06(ctx *utest.TestContext) {
	time.Sleep(time.Duration(5) * time.Second)

	req := ipsecvpnClient.NewDescribeVPNGatewayRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return ipsecvpnClient.DescribeVPNGateway(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", 0, "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Error(err)

	}

}

func testSet25GetVPNGatewayUpgradePrice07(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := iipsecvpnClient.NewGetVPNGatewayUpgradePriceRequest()

	ctx.NoError(utest.SetReqValue(req, "VPNGatewayId", ctx.GetVar("vpngw_id")))

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "Grade", "Enhanced"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return iipsecvpnClient.GetVPNGatewayUpgradePrice(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", 0, "str_eq"),
		},
		MaxRetries:    0,
		RetryInterval: 0 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Error(err)

	}

}

func testSet25UpdateVPNGateway08(ctx *utest.TestContext) {
	time.Sleep(time.Duration(5) * time.Second)

	req := ipsecvpnClient.NewUpdateVPNGatewayRequest()

	ctx.NoError(utest.SetReqValue(req, "VPNGatewayId", ctx.GetVar("vpngw_id")))

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "Grade", "Enhanced"))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return ipsecvpnClient.UpdateVPNGateway(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", 0, "str_eq"),
		},
		MaxRetries:    10,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Error(err)

	}

}

func testSet25CreateRemoteVPNGateway09(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := ipsecvpnClient.NewCreateRemoteVPNGatewayRequest()

	ctx.NoError(utest.SetReqValue(req, "RemoteVPNGatewayName", "auto_apitest"))

	ctx.NoError(utest.SetReqValue(req, "RemoteVPNGatewayAddr", "10.1.1.0"))

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return ipsecvpnClient.CreateRemoteVPNGateway(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", 0, "str_eq"),
		},
		MaxRetries:    0,
		RetryInterval: 0 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Error(err)

	}

	ctx.Vars["remote_vpngw_id"] = ctx.Must(utest.GetValue(resp, "RemoteVPNGatewayId"))
}

func testSet25DescribeRemoteVPNGateway10(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := ipsecvpnClient.NewDescribeRemoteVPNGatewayRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return ipsecvpnClient.DescribeRemoteVPNGateway(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", 0, "str_eq"),
		},
		MaxRetries:    0,
		RetryInterval: 0 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Error(err)

	}

}

func testSet25CreateVPNTunnel11(ctx *utest.TestContext) {
	time.Sleep(time.Duration(5) * time.Second)

	req := ipsecvpnClient.NewCreateVPNTunnelRequest()

	ctx.NoError(utest.SetReqValue(req, "VPNTunnelName", "auto_apitest"))

	ctx.NoError(utest.SetReqValue(req, "VPNGatewayId", ctx.GetVar("vpngw_id")))

	ctx.NoError(utest.SetReqValue(req, "RemoteVPNGatewayId", ctx.GetVar("remote_vpngw_id")))

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "IPSecRemoteSubnets", "10.1.1.0/24"))

	ctx.NoError(utest.SetReqValue(req, "IPSecProtocol", "ah"))

	ctx.NoError(utest.SetReqValue(req, "IPSecPFSDhGroup", 15))

	ctx.NoError(utest.SetReqValue(req, "IPSecLocalSubnetIds", ctx.GetVar("subnet_id")))

	ctx.NoError(utest.SetReqValue(req, "IKEPreSharedKey", "test"))

	ctx.NoError(utest.SetReqValue(req, "IKEExchangeMode", "main"))

	ctx.NoError(utest.SetReqValue(req, "IKEDhGroup", 15))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return ipsecvpnClient.CreateVPNTunnel(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", 0, "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Error(err)

	}

	ctx.Vars["vpn_tunnel_id"] = ctx.Must(utest.GetValue(resp, "VPNTunnelId"))
}

func testSet25DescribeVPNTunnel12(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := ipsecvpnClient.NewDescribeVPNTunnelRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return ipsecvpnClient.DescribeVPNTunnel(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", 0, "str_eq"),
		},
		MaxRetries:    0,
		RetryInterval: 0 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Error(err)

	}

}

func testSet25UpdateVPNTunnelAttribute13(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := ipsecvpnClient.NewUpdateVPNTunnelAttributeRequest()

	ctx.NoError(utest.SetReqValue(req, "VPNTunnelId", ctx.GetVar("vpn_tunnel_id")))

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return ipsecvpnClient.UpdateVPNTunnelAttribute(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", 0, "str_eq"),
		},
		MaxRetries:    0,
		RetryInterval: 0 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Error(err)

	}

}

func testSet25DeleteVPNGateway14(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := ipsecvpnClient.NewDeleteVPNGatewayRequest()

	ctx.NoError(utest.SetReqValue(req, "VPNGatewayId", ctx.GetVar("vpngw_id")))

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return ipsecvpnClient.DeleteVPNGateway(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", 66007, "str_eq"),
		},
		MaxRetries:    0,
		RetryInterval: 0 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Error(err)

	}

}

func testSet25DeleteRemoteVPNGateway15(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := ipsecvpnClient.NewDeleteRemoteVPNGatewayRequest()

	ctx.NoError(utest.SetReqValue(req, "RemoteVPNGatewayId", ctx.GetVar("remote_vpngw_id")))

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return ipsecvpnClient.DeleteRemoteVPNGateway(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", 66032, "str_eq"),
		},
		MaxRetries:    0,
		RetryInterval: 0 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Error(err)

	}

}

func testSet25DeleteVPNTunnel16(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := ipsecvpnClient.NewDeleteVPNTunnelRequest()

	ctx.NoError(utest.SetReqValue(req, "VPNTunnelId", ctx.GetVar("vpn_tunnel_id")))

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return ipsecvpnClient.DeleteVPNTunnel(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", 0, "str_eq"),
		},
		MaxRetries:    0,
		RetryInterval: 0 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Error(err)

	}

}

func testSet25DeleteVPNGateway17(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := ipsecvpnClient.NewDeleteVPNGatewayRequest()

	ctx.NoError(utest.SetReqValue(req, "VPNGatewayId", ctx.GetVar("vpngw_id")))

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return ipsecvpnClient.DeleteVPNGateway(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", 0, "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Error(err)

	}

}

func testSet25DeleteRemoteVPNGateway18(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := ipsecvpnClient.NewDeleteRemoteVPNGatewayRequest()

	ctx.NoError(utest.SetReqValue(req, "RemoteVPNGatewayId", ctx.GetVar("remote_vpngw_id")))

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return ipsecvpnClient.DeleteRemoteVPNGateway(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", 0, "str_eq"),
		},
		MaxRetries:    0,
		RetryInterval: 0 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Error(err)

	}

}

func testSet25ReleaseEIP19(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := unetClient.NewReleaseEIPRequest()

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	ctx.NoError(utest.SetReqValue(req, "EIPId", ctx.GetVar("eip_id")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return unetClient.ReleaseEIP(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", 0, "str_eq"),
		},
		MaxRetries:    0,
		RetryInterval: 0 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Error(err)

	}

}

func testSet25DeleteSubnet20(ctx *utest.TestContext) {
	time.Sleep(time.Duration(0) * time.Second)

	req := vpcClient.NewDeleteSubnetRequest()

	ctx.NoError(utest.SetReqValue(req, "SubnetId", ctx.GetVar("subnet_id")))

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return vpcClient.DeleteSubnet(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", 0, "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Error(err)

	}

}

func testSet25DeleteVPC21(ctx *utest.TestContext) {
	time.Sleep(time.Duration(10) * time.Second)

	req := vpcClient.NewDeleteVPCRequest()

	ctx.NoError(utest.SetReqValue(req, "VPCId", ctx.GetVar("vpc_id")))

	ctx.NoError(utest.SetReqValue(req, "Region", ctx.GetVar("Region")))

	testCase := utest.TestCase{
		Invoker: func() (interface{}, error) {
			return vpcClient.DeleteVPC(req)
		},
		Validators: []utest.TestValidator{
			ctx.NewValidator("RetCode", 0, "str_eq"),
		},
		MaxRetries:    3,
		RetryInterval: 1 * time.Second,
		T:             ctx.T,
	}

	resp, err := testCase.Run()
	if resp == nil || err != nil {

		ctx.T.Error(err)

	}

}