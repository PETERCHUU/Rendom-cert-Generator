package tlsgen

import (
	"crypto/x509"
	"linuxWebUI/model/security/tlsgen/filegen"
	"os"
)

type TlsSet interface {
	GenerateKey()
}

func NewFileGen(ca *x509.Certificate) TlsSet {
	netinterface := &filegen.Filessl{
		Ca: ca,
	}
	return netinterface
}
func NewUseKey() (cert, key []byte) {
	var err error
	cert, err = os.ReadFile("cert.pem")
	if err != nil {
		print(err.Error())
		return
	}
	key, err = os.ReadFile("key.pem")
	if err != nil {
		print(err.Error())
		return
	}
	return cert, key
}
func NewbyteGen(ca *x509.Certificate) TlsSet {
	netinterface := &filegen.Filessl{
		Ca: ca,
	}
	return netinterface
}
