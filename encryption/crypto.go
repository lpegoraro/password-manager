package encryption

import (
	"crypto"
	"crypto/rsa"
)

type CryptoHelper interface {
	Register(signingMethod string) SigningMethodRSA
	DecryptMessage(crytoData string) (string, error)
	EncodeFingerprint(dataToEncode string) (string, error)
}

type SigningMethodRSA struct {
	Name string
	Hash crypto.Hash
}

var (
	SigningMethodRS256 *SigningMethodRSA
	SigningMethodRS384 *SigningMethodRSA
	SigningMethodRS512 *SigningMethodRSA
)

type DefaultCrytoHelper struct {
	publicKey     rsa.PublicKey
	signingMethod SigningMethodRSA
	Hash          crypto.Hash
	PemCert       string
}

func (dch DefaultCrytoHelper) DecryptMessage(cryptoData string) (string, error) {

	return "", nil
}

func (dch DefaultCrytoHelper) EncodeFingerprint(dataToEncode string) (string, error) {

	return "", nil
}

func (dch DefaultCrytoHelper) Register(signingMethod string) (*SigningMethodRSA, error) {
	var err error
	switch signingMethod {
	case "RS256":
		return &SigningMethodRSA{"RS256", crypto.SHA256}, nil
	case "RS384":
		return &SigningMethodRSA{"RS384", crypto.SHA384}, nil
	case "RS512":
		return &SigningMethodRSA{"RS512", crypto.SHA512}, nil
	default:
		return nil, err
	}
}

func (m *SigningMethodRSA) Alg() string {
	return m.Name
}
