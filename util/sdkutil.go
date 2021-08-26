package util

import (
	"fmt"
	"github.com/VoneChain-CS/fabric-sdk-go-gm/pkg/client/resmgmt"
	"github.com/VoneChain-CS/fabric-sdk-go-gm/pkg/core/config"
	"github.com/VoneChain-CS/fabric-sdk-go-gm/pkg/fabsdk"
	"github.com/VoneChain-CS/fabric-sdk-go-gm/pkg/gateway"
	"io/ioutil"
	"os"
	"path/filepath"
)

const (
	CHANNEL = "mychannel"
	WALLET = "config/wallet"
	ADMIN = "Admin"
	ADMINSECRET = "org1adminpw"
	MSP_ID = "Org1MSP"
	CERT_PATH = "config/organizations/peerOrganizations/org1.xxzx.com/users/Admin@org1.xxzx.com/msp"
	CONFIG_PTH = "config/config_network.yaml"
	ORG1 = "org1"
	ORDER_ORG = "OrdererOrg"
)

var (
	Orgsdk *resmgmt.Client
	Channelsdk *resmgmt.Client
	Gatewaysdk  *gateway.Network
	Sdk *fabsdk.FabricSDK
)

func InitSDK()  {
	c := config.FromFile(CONFIG_PTH)
	sdk, err := fabsdk.New(c)
	if err != nil {
		os.Exit(1)
		panic("初始化org sdk失败")
	}
	Sdk = sdk
}

func InitOrgAdminSDK()  {

	//defer sdk.Close()
	adminContext := Sdk.Context(fabsdk.WithUser(ADMIN), fabsdk.WithOrg(ORG1))

	// Org resource management client
	Orgsdk, _ = resmgmt.New(adminContext)
}

func InitOrderAdminSDK(){
	adminContext := Sdk.Context(fabsdk.WithUser(ADMIN), fabsdk.WithOrg(ORDER_ORG))

	// Org resource management client
	Channelsdk, _ = resmgmt.New(adminContext)
}

func InitGatewaySDK(channel string) {
	wallet, err := gateway.NewFileSystemWallet(WALLET)
	if !wallet.Exists(ADMIN) {
		err = PopulateWallet(wallet)
		if err != nil {
			panic("初始化gateway sdk 失败")
		}
	}
	gw, err := gateway.Connect(
		gateway.WithConfig(config.FromFile(CONFIG_PTH)),
		gateway.WithIdentity(wallet, ADMIN),
	)
	if err != nil {
		panic("初始化gateway sdk 失败")
	}
	defer gw.Close()

	Gatewaysdk, err = gw.GetNetwork(channel)

}

func PopulateWallet(wallet *gateway.Wallet) error {

	credPath := CERT_PATH

	certPath := filepath.Join(credPath, "signcerts", "cert.pem")
	// read the certificate pem
	cert, err := ioutil.ReadFile(filepath.Clean(certPath))
	if err != nil {
		return err
	}

	keyDir := filepath.Join(credPath, "keystore")
	// there's a single file in this dir containing the private key
	files, err := ioutil.ReadDir(keyDir)
	if err != nil {
		return err
	}
	if len(files) != 1 {
		return fmt.Errorf("keystore folder should have contain one file")
	}
	keyPath := filepath.Join(keyDir, files[0].Name())
	key, err := ioutil.ReadFile(filepath.Clean(keyPath))
	if err != nil {
		return err
	}

	identity := gateway.NewX509Identity(MSP_ID, string(cert), string(key))

	return wallet.Put(ADMIN, identity)
}

func init() {
	//该方法必须放第一  属于后续依赖
	InitSDK()
    InitOrgAdminSDK()
    InitGatewaySDK(CHANNEL)
    InitOrderAdminSDK()
}