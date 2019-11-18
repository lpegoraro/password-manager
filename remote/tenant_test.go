package remote

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var publicKey = "-----BEGIN PUBLIC KEY-----MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEAnwp/GrWnW9iufjIN0haZeWj2yI78/h0ywaOTnEXaX8epwoYLO3SCSbsqGWvA4xmevyZwKpKeOz0N9OI6Qivm6+Vy7/+OZ7nRYRomPGstHl1UPL6R1Gs+O7iaZjRvZtECIkEr+3IVbjlVBPqEwWXVPt/4M2WNkNURHcorVnA8w/FP3ti7h5AxMtIz8xhTtjjT7mP2wyO39QNA0VuX/Esk/LST1LfkdqKOmWR4i7J6xyBJyuF0hRPUAmC3ixXx64nZicrbj1Flnf/R3CTieJq90dPbNE/l6YxDHxliz6fTV3ctlYHZzhyDkOdAHCZzPcRQkSBYj96l3zuBnHkRWBBOPl1qewZZ4VnoXdkioekKC8zMXFaFCTCuYOnQ69/jsPqDBftoYfIfFLSEmfHG1vr+l/jqcvCYnKZcd2k9G4UTOWM62AMd/1TLZl2N1n6hrSyJJWZyvssQJps0e/Fb4ElZysEQut3eig368NzFDNGN+0Oq9sLhXfPckF+HcDPoCDshGkTq4g6g+Qv9f1XjzXH8hyL8orGSDXJF70bbdjilOuvPhWyzJixXyLDbxOGKBIabpXF8WCak4MF5t8YL99a5pfN/ig4hEmfTD8XEmFUcCcAp8l+OwrgnisLdLY5hfAl1jgpqt0jGBO1E/Nr1+G2WBa94Wfn5TwPRrcbgvc+TWlcCAwEAAQ==-----END PUBLIC KEY-----"

func TestConnect(t *testing.T) {
	request := &AddTenantReq{
		PubKey:     publicKey,
		TenantName: "testConnect",
		InitialConfiguration: &Configuration{
			Method:  "uuid",
			Factor:  4,
			Seed:    "GrWnW9iufjIN0haZ",
			Storage: "",
		},
	}
	testTenant := DefaultTenantClient{}
	response, err := testTenant.Connect(nil, request)
	assert.Nil(t, err)
	fingerprint := "\x03\xaf\x04\xa8\xbc\xe6\xf7~\"I\a\x85\xb0{\n\xce"
	assert.Equal(t, fingerprint, response.Fingerprint)
}
