package service

import (
	"baas-fabric/internal/github.com/hyperledger/fabric/common/policydsl"
	"baas-fabric/util"
	"fmt"
	"github.com/VoneChain-CS/fabric-sdk-go-gm/pkg/client/resmgmt"
	"github.com/VoneChain-CS/fabric-sdk-go-gm/pkg/common/errors/retry"
	lcpackager "github.com/VoneChain-CS/fabric-sdk-go-gm/pkg/fab/ccpackager/lifecycle"
	"github.com/hyperledger/fabric-protos-go/common"
	pb "github.com/hyperledger/fabric-protos-go/peer"
)

const (
	PEER1 = "peer0.org1.xxzx.com"
	ORDER_ENDPOINT = "orderer.xxzx.com"
)

type ChainCode struct {
	ChannelId string
	Label string
	Version string
    Seq int64
	Ccp string
	Path string
	Init bool
    Type pb.ChaincodeSpec_Type
}
var adminClient *resmgmt.Client

func DeployCC(chaincode *ChainCode) (bool, error) {
	packageName :=chaincode.Label + "_" + chaincode.Version
	resp := false
	ccpkg ,err :=packageCC(chaincode.Path, chaincode.Type,packageName)
	if err !=nil {
        return false, fmt.Errorf("智能合约打包失败，%s", err)
	}
	label := chaincode.Label
	installCC(label,ccpkg,adminClient)
	packageID := lcpackager.ComputePackageID(packageName, ccpkg)
    err = approveCC(chaincode.Init,chaincode.ChannelId,label,packageID,chaincode.Ccp,chaincode.Version,chaincode.Seq,adminClient)
	if err !=nil {
		return resp,err
	}
	err = commitCC(chaincode.ChannelId,chaincode.Version,label,chaincode.Init,chaincode.Ccp,chaincode.Seq,adminClient)
	if err !=nil {
		return resp,err
	}
	defer func()  {
		if errs := recover(); errs != nil {
			fmt.Printf("%s\n", errs)
			panic("部署智能合约失败")
		}
	}()
    resp = true

	return resp, nil
}

func packageCC(path string, tp pb.ChaincodeSpec_Type, label string) ([]byte, error) {
	if len(path) <1  || tp == 0 || len(label) < 1{
        return nil,fmt.Errorf("package 参数不合法或者缺失")
	}
	desc := &lcpackager.Descriptor{
		Path:  path,
		Type:  tp,
		Label: label,
	}
	ccPkg, err := lcpackager.NewCCPackage(desc)
	if err != nil {
		fmt.Print(err)
	}
	return ccPkg, nil
}

func installCC(label string, ccPkg []byte, orgResMgmt *resmgmt.Client) {
	installCCReq := resmgmt.LifecycleInstallCCRequest{
		Label:   label,
		Package: ccPkg,
	}

	_, err := orgResMgmt.LifecycleInstallCC(installCCReq, resmgmt.WithRetry(retry.DefaultResMgmtOpts))
	if err != nil {
		panic("install 部署智能合约失败")
	}
}

func approveCC(init bool,channelId string,label string, packageID string,ccp string,version string, seq int64,orgResMgmt *resmgmt.Client) error {
	var ccPolicy *common.SignaturePolicyEnvelope
	if len(ccp) > 0 {
		ccPolicy, _ = policydsl.FromString(ccp)
	}
	approveCCReq := resmgmt.LifecycleApproveCCRequest{
		Name:      label,
		Version:   version,
		PackageID: packageID,
		Sequence:  seq,
		/*EndorsementPlugin: "escc",
		ValidationPlugin:  "vscc",*/
		SignaturePolicy:   ccPolicy,
		InitRequired: init,
	}

	_, err := orgResMgmt.LifecycleApproveCC(channelId, approveCCReq, resmgmt.WithTargetEndpoints(PEER1), resmgmt.WithOrdererEndpoint(ORDER_ENDPOINT), resmgmt.WithRetry(retry.DefaultResMgmtOpts))

	if err != nil {
		return fmt.Errorf("组织审批失败,err：%s",err)
	}
	return nil
}

func queryInstalled(label string, packageID string, orgResMgmt *resmgmt.Client) {
	resp, err := orgResMgmt.LifecycleQueryInstalledCC(resmgmt.WithTargetEndpoints(PEER1), resmgmt.WithRetry(retry.DefaultResMgmtOpts))
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(resp)
}

func queryApprovedCC(channelId,label string, seq int64,orgResMgmt *resmgmt.Client) {
	queryApprovedCCReq := resmgmt.LifecycleQueryApprovedCCRequest{
		Name:     label,
		Sequence: seq,
	}
	resp, err := orgResMgmt.LifecycleQueryApprovedCC(channelId, queryApprovedCCReq, resmgmt.WithTargetEndpoints(PEER1), resmgmt.WithRetry(retry.DefaultResMgmtOpts))
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(resp)
}

func getInstalledCCPackage(packageID string, orgResMgmt *resmgmt.Client) []byte {
	resp, err := orgResMgmt.LifecycleGetInstalledCCPackage(packageID, resmgmt.WithTargetEndpoints(PEER1), resmgmt.WithRetry(retry.DefaultResMgmtOpts))
	if err != nil {
		fmt.Print(err)
		return nil
	}
	return resp
}

func commitCC(channelId,version,label string,init bool,ccp string,seq int64,orgResMgmt *resmgmt.Client) error{
	var ccPolicy *common.SignaturePolicyEnvelope
	if len(ccp) > 0 {
		ccPolicy, _ = policydsl.FromString(ccp)
	}
	req := resmgmt.LifecycleCommitCCRequest{
		Name:     label,
		Version:  version,
		Sequence: seq,
		/*EndorsementPlugin: "escc",
		ValidationPlugin:  "vscc",*/
		SignaturePolicy:   ccPolicy,
		InitRequired: init,
	}
	txnID, err := orgResMgmt.LifecycleCommitCC(channelId, req, resmgmt.WithRetry(retry.DefaultResMgmtOpts),
		resmgmt.WithTargetEndpoints(PEER1),
		resmgmt.WithOrdererEndpoint(ORDER_ENDPOINT))
	if err != nil  {
		return err
	}
	if len(txnID) < 1{
		return fmt.Errorf("txnid 为空")
	}
	return nil
}
func init() {
	adminClient = util.Orgsdk
}