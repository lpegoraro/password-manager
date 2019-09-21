package remote

import (
	context "context"

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
	var newTenant = DefaultTenantClient{
		TenantKey:  tenantName,
		TenantPort: GetUnusedPort(),
		crypto: &encryption.DefaultCrytoHelper{
			PemCert: publicKey,
		},
	}
	var fingerprint, err = newTenant.crypto.EncodeFingerprint(tenantName)
	if err != nil {
		panic("Error generating fingerprint, please check the public key")
	}
	return &TenantValue{
		Port:        newTenant.TenantPort,
		Fingerprint: fingerprint,
	}

}

func GetUnusedPort() string {
	return "7894"
}
