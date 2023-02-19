package tlsgen

import (
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"log"
	"math/big"
	"net"
	"os"
	"time"
)

var ca = &x509.Certificate{
	SerialNumber: big.NewInt(1653),
	// Issuer: pkix.Name{
	// 	Organization:       []string{"Company, INC."},
	// 	OrganizationalUnit: []string{"sdd"},
	// 	Country:            []string{"US"},
	// 	Province:           []string{""},
	// 	Locality:           []string{"San Francisco"},
	// 	StreetAddress:      []string{"Golden Gate Bridge"},
	// 	PostalCode:         []string{"94016"},
	// },
	Subject: pkix.Name{
		Organization:       []string{"Company, INC."},
		OrganizationalUnit: []string{"sdd"},
		Country:            []string{"US"},
		Province:           []string{""},
		Locality:           []string{"San Francisco"},
		StreetAddress:      []string{"Golden Gate Bridge"},
		PostalCode:         []string{"94016"},
	},
	IPAddresses:           []net.IP{net.IPv4(127, 0, 0, 1), net.IPv6loopback},
	NotBefore:             time.Now(),
	NotAfter:              time.Now().AddDate(10, 0, 0),
	BasicConstraintsValid: true,
	IsCA:                  true,
	// SignatureAlgorithm:    x509.ECDSAWithSHA256,
	// PublicKeyAlgorithm:    x509.Ed25519,
	ExtKeyUsage: []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
	KeyUsage:    x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
}

func TrustTheCert(cert []byte) {
	rootCAs, _ := x509.SystemCertPool()
	// 	// handle case where rootCAs == nil and create an empty pool...
	if ok := rootCAs.AppendCertsFromPEM(cert); !ok {
		print("secc append")
	}

	// 	config := &tls.Config{
	// 		InsecureSkipVerify: *insecure,
	// 		RootCAs:            rootCAs,
	// 	}
	// 	tr := &http.Transport{TLSClientConfig: config}
	// 	client := &http.Client{Transport: tr}
}

// input Ca , Cert , KeyStyle
// output publicKeyblock ,privitekeyblock
func GenerateKey() ([]byte, []byte) {

	cert := &x509.Certificate{
		SerialNumber: big.NewInt(1024),
		Subject: pkix.Name{
			Organization: []string{"Acme Co"},
		},
		NotBefore: time.Now(),
		NotAfter:  time.Now().Add(time.Hour * 24 * 180),
		// SignatureAlgorithm: x509.ECDSAWithSHA256,
		// PublicKeyAlgorithm: x509.Ed25519,

		KeyUsage:              x509.KeyUsageDigitalSignature, // if rsa  |x509.KeyUsageKeyEncipherment,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
	}

	// capriv, err := rsa.GenerateKey(rand.Reader, *rsaBits)
	certpriv, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	certpubl := &certpriv.PublicKey
	//certpubl, certpriv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		log.Println("key generate failed", err)
		return nil, nil
	}

	certCert, err := x509.CreateCertificate(rand.Reader, cert, ca, certpubl, certpriv)
	if err != nil {
		log.Println("create cert failed", err)
		return nil, nil
	}

	/*
	   hosts := strings.Split(*host, ",")
	   for _, h := range hosts {
	   	if ip := net.ParseIP(h); ip != nil {
	   		template.IPAddresses = append(template.IPAddresses, ip)
	   	} else {
	   		template.DNSNames = append(template.DNSNames, h)
	   	}
	   }
	   if *isCA {
	   	template.IsCA = true
	   	template.KeyUsage |= x509.KeyUsageCertSign
	   }
	*/
	certpem := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certCert})

	keypem := pem.EncodeToMemory(pemBlockForKey(certpriv))

	certDERBlock, publicCert := pem.Decode(certpem)
	print(certDERBlock.Type, "\n")
	if publicCert != nil {
		print("publicCert nil\n")
	}
	//privitKey = append([]byte{10}, privitKey...)
	_, err = tls.X509KeyPair(certpem, keypem)
	if err != nil {
		print(err.Error())
	}

	return certpem, keypem
}

type private interface {
	*rsa.PrivateKey | *ecdsa.PrivateKey | ed25519.PrivateKey
}

// *rsa.PrivateKey | *ecdsa.PrivateKey | ed25519.PrivateKey
func pemBlockForKey[T private](priv T) *pem.Block {
	switch k := any(priv).(type) {
	case *rsa.PrivateKey:
		return &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(k)}
	case *ecdsa.PrivateKey:
		b, err := x509.MarshalECPrivateKey(k)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to marshal ECDSA private key: %v", err)
			os.Exit(2)
		}
		return &pem.Block{Type: "EC PRIVATE KEY", Bytes: b}
	case ed25519.PrivateKey:
		b, err := x509.MarshalPKCS8PrivateKey(k)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to marshal ed25519 private key: %v", err)
			os.Exit(2)
		}
		return &pem.Block{Type: "PRIVATE KEY", Bytes: b}
	default:
		return nil
	}
}
