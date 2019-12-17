package remote

import (
	context "context"
	fmt "fmt"
	"net"

	"github.com/lpegoraro/password-manager/encryption"
)

type DefaultTenantClient struct {
	TenantKey  string
	TenantPort string
	crypto     *encryption.DefaultCrytoHelper
}

func (dtc *DefaultTenantClient) Connect(ctx context.Context, req *AddTenantReq) (*TenantValue, error) {
	response := ServeNewTenant(req.TenantName, req.PubKey)
	return response, nil
}

func ServeNewTenant(tenantName string, publicKey string) *TenantValue {
	fmt.Println("Serving to: " + tenantName + ", " + publicKey)
	var newTenant = DefaultTenantClient{
		TenantKey:  tenantName,
		TenantPort: GetUnusedPort(),
		crypto: &encryption.DefaultCrytoHelper{
			PemCert: publicKey,
		},
	}
	var fingerprint, err = newTenant.crypto.EncodeFingerprint(tenantName + publicKey)
	if err != nil {
		panic("Error generating fingerprint, please check the public key")
	}
	serve(newTenant)

	return &TenantValue{
		Port:        newTenant.TenantPort,
		Fingerprint: string(fingerprint[:16]),
	}
}

func serve(tenant DefaultTenantClient) {
	ln, err := net.Listen("tcp", ":7894")
	if err != nil {
		fmt.Println("Error while Listening: ", err)
	}
	if ln != nil {
		fmt.Println("Success Listening to port 7894")
	}
	select {}
}

func GetUnusedPort() string {
	return "7894"
}
