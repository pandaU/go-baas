package service

import (
	"baas-fabric/util"
	"github.com/stretchr/testify/require"
)

func CreateChannel()  {
	mspClient, err := mspclient.New(util.Sdk.Context(), mspclientgo .WithOrg(util.ORG1))
	if err != nil {
		t.Fatal(err)
	}
	adminIdentity, err := mspClient.GetSigningIdentity(orgAdmin)
	if err != nil {
		t.Fatal(err)
	}
	req := resmgmt.SaveChannelRequest{ChannelID: channelID,
		ChannelConfigPath: integration.GetChannelConfigTxPath(channelID + ".tx"),
		SigningIdentities: []msp.SigningIdentity{adminIdentity}}
	txID, err := resMgmtClient.SaveChannel(req, resmgmt.WithRetry(retry.DefaultResMgmtOpts), resmgmt.WithOrdererEndpoint("orderer.example.com"))
	require.Nil(t, err, "error should be nil")
	require.NotEmpty(t, txID, "transaction ID should be populated")
}