package remote


import (
	context "context"

	grpc "google.golang.org/grpc"
)

type DefaultTenantClient struct {
	TenantKey  string
	TenantPort string
	crypto     DefaultCrytoHelper
}

func (dtc *DefaultTenantClient) Connect(ctx context.Context, req *AddTenantReq) (*TenantValue, error) {}
	var newTenantValue TenantValue
	newTenantValue = TenantValue{
		Port:dtc.TenantPort,
		Fingerprint:"",
	}
	return newTenantValue, nil
}	

func ServeNewTenant(tenantName, publicKey string) {
	var newTenant = DefaultTenantClient{
		TenantKey: tenantName,
		TenantPort: GetUnusedPort(),
		crypto: DefaultCrytoHelper{
			PemCert: publicKey,
		},
	}
}

func GetUnusedPort() string {
	return "7894"
}