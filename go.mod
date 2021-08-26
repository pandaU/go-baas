module baas-fabric

go 1.14

require (
	github.com/Hyperledger-TWGC/tjfoc-gm v0.0.0-20210222084201-e65875425ad3
	github.com/Knetic/govaluate v3.0.1-0.20170926212237-aa73cfd04eeb+incompatible
	github.com/VoneChain-CS/fabric-sdk-go-gm v1.1.1
	github.com/VoneChain-CS/fabric-sdk-go-gm/cfssl v0.0.0-20201021101014-9a2abd087e1c
	github.com/cloudflare/cfssl v1.5.0
	github.com/gin-gonic/gin v1.7.4
	github.com/go-kit/kit v0.9.0
	github.com/go-playground/validator/v10 v10.9.0 // indirect
	github.com/golang/protobuf v1.5.2
	github.com/google/certificate-transparency-go v1.1.1
	github.com/grantae/certinfo v0.0.0-20170412194111-59d56a35515b
	github.com/hyperledger/fabric-lib-go v1.0.0
	github.com/hyperledger/fabric-protos-go v0.0.0-20200707132912-fee30f3ccd23
	github.com/json-iterator/go v1.1.11 // indirect
	github.com/mattn/go-isatty v0.0.13 // indirect
	github.com/miekg/pkcs11 v1.0.3
	github.com/mitchellh/mapstructure v1.3.3
	github.com/pkg/errors v0.8.1
	github.com/prometheus/client_golang v1.7.1
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.7.0
	github.com/tjfoc/gmsm v1.4.0
	github.com/tjfoc/gmtls v1.2.1
	github.com/ugorji/go v1.2.6 // indirect
	golang.org/x/crypto v0.0.0-20210817164053-32db794688a5
	golang.org/x/net v0.0.0-20210226172049-e18ecbb05110
	golang.org/x/sys v0.0.0-20210823070655-63515b42dcdf // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/grpc v1.40.0
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/yaml.v2 v2.4.0
)

replace (
	github.com/VoneChain-CS/fabric-sdk-go-gm v1.1.1 => ../github.com/hyperledger/fabric-gm-sdk-go-final
	github.com/VoneChain-CS/fabric-sdk-go-gm/cfssl v0.0.0-20201021101014-9a2abd087e1c => ../github.com/hyperledger/fabric-gm-sdk-go-final/cfssl
	github.com/hyperledger/fabric-sdk-go => ../github.com/hyperledger/fabric-gm-sdk-go-final/fabric-sdk-go
	github.com/spf13/cast v1.3.1 => ../github.com/hyperledger/fabric-gm-sdk-go-final/spf13/cast
	github.com/spf13/cobra => ../github.com/hyperledger/fabric-gm-sdk-go-final/spf13/cobra
	github.com/spf13/jwalterweatherman => ../github.com/hyperledger/fabric-gm-sdk-go-final/spf13/jwalterweatherman
	github.com/spf13/pflag => ../github.com/hyperledger/fabric-gm-sdk-go-final/spf13/pflag
	github.com/tjfoc/gmsm => ../github.com/hyperledger/fabric-gm-sdk-go-final/tjfoc/gmsm
	github.com/tjfoc/gmtls v1.2.1 => ../github.com/hyperledger/fabric-gm-sdk-go-final/tjfoc/gmtls
	google.golang.org/grpc v1.40.0 => google.golang.org/grpc v1.26.0
)
