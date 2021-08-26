package service

import (
	"baas-fabric/util"
	"fmt"
	"github.com/VoneChain-CS/fabric-sdk-go-gm/pkg/gateway"
)

var network  *gateway.Network

func QueryChainCode(channel string, ccid string, funcName string, args ...string)(respJson string, err error)  {
	respJson = ""
	if err !=nil {
		return respJson, fmt.Errorf("初始化network失败：%s",err)
	}
	concat := network.GetContract(ccid)

	resp , err:=concat.EvaluateTransaction(funcName, args...)
	if err !=nil {
		return respJson, fmt.Errorf("调用chaincode：%s,func：%s ,err ：%s", ccid, funcName ,err)
	}
	respJson = string(resp)
	return  respJson ,nil
}


func InvokeChainCode(channel string, ccid string, funcName string, args ...string)(respJson bool, err error)  {
	respJson = false
	if err !=nil {
		return respJson, fmt.Errorf("初始化network失败：%s",err)
	}
	concat := network.GetContract(ccid)

	_ , err =concat.SubmitTransaction(funcName, args...)
	if err !=nil {
		return respJson, fmt.Errorf("调用chaincode：%s,func：%s ,err ：%s", ccid, funcName ,err)
	}

	return  respJson ,nil
}

func init() {
	network = util.Gatewaysdk
}





