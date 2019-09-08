package encryption

type CryptoHelper interface {
	DecryptMessage(crytoData string) (string, error)
	EncodeFingerprint(dataToEncode string) (string, error)
}

type DefaultCrytoHelper struct {
	publicKey, signingMethod string
}

func (dch DefaultCrytoHelper) DecryptMessage(cryptoData string) (string, error) {

	return "", nil
}

func (dch DefaultCrytoHelper) EncodeFingerprint(dataToEncode string) (string, error) {

	return "", nil
}
