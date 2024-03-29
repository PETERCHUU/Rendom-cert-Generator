package controller

import (
	"crypto/tls"
	"crypto/x509"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/quic-go/quic-go/http3"
)

func Run(certPEMBlock, keyPEMBlock []byte) {
	router := gin.New()

	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "OK"})
	})

	// Listen and Server in https://127.0.0.1:8080 []byte{cert.pem}, []byte{key.pem}
	err := ListenAndServeQUICC(":443", certPEMBlock, keyPEMBlock, router)
	log.Fatal(err)
}

func TrustTheCert(certPEMBlock, keyPEMBlock []byte) (rootCAs *x509.CertPool) {
	rootCAs, _ = x509.SystemCertPool()

	// 	// handle case where rootCAs == nil and create an empty pool...
	if ok := rootCAs.AppendCertsFromPEM(certPEMBlock); !ok {
		print("secc append")
	}
	// Trust the augmented cert pool in our client
	config := &tls.Config{
		InsecureSkipVerify: false,
		RootCAs:            rootCAs,
	}
	tr := &http.Transport{TLSClientConfig: config}
	client := &http.Client{Transport: tr}

	// Uses local self-signed cert
	req, _ := http.NewRequest(http.MethodGet, "https://localhost/api/version", nil)
	resp, err := client.Do(req)
	// Handle resp and err
	print(resp, err)

	// Still works with host-trusted CAs!
	req, _ = http.NewRequest(http.MethodGet, "https://example.com/", nil)
	resp, err = client.Do(req)
	// Handle resp and err
	print(resp, err)
	return
}

// listenAndServeTLS(":8443", []byte{cert.pem}, []byte{key.pem}, gin.engine)
func ListenAndServeTLS(addr string, certPEMBlock, keyPEMBlock []byte, handler http.Handler) error {
	if addr == "" {
		addr = ":https"
	}
	config := &tls.Config{
		//	InsecureSkipVerify: false,
		// 	RootCAs: rootCAs
	}
	if config.NextProtos == nil {
		config.NextProtos = []string{"http/2"}
	}

	var err error
	config.Certificates = make([]tls.Certificate, 1)
	config.Certificates[0], err = tls.X509KeyPair(certPEMBlock, keyPEMBlock)
	if err != nil {
		return err
	}
	srv := &http.Server{Addr: addr, Handler: handler, TLSConfig: config}
	// rootCAs := TrustTheCert(certPEMBlock, keyPEMBlock)

	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	tlsListener := tls.NewListener(tcpKeepAliveListener{ln.(*net.TCPListener)}, config)
	return srv.Serve(tlsListener)
}

// tcpKeepAliveListener sets TCP keep-alive timeouts on accepted
// connections. It's used by ListenAndServe and ListenAndServeTLS so
// dead TCP connections (e.g. closing laptop mid-download) eventually
// go away.
type tcpKeepAliveListener struct {
	*net.TCPListener
}

func (ln tcpKeepAliveListener) Accept() (c net.Conn, err error) {
	tc, err := ln.AcceptTCP()
	if err != nil {
		return
	}
	tc.SetKeepAlive(true)
	tc.SetKeepAlivePeriod(3 * time.Minute)
	return tc, nil
}

func ListenAndServeQUICC(addr string, certPEMBlock, keyPEMBlock []byte, handler http.Handler) error {
	if addr == "" {
		addr = ":https"
	}
	var err error
	certs := make([]tls.Certificate, 1)
	certs[0], err = tls.X509KeyPair(certPEMBlock, keyPEMBlock)
	if err != nil {
		return err
	}
	config := &tls.Config{
		Certificates: certs,
	}
	server := http3.Server{
		Addr:      addr,
		Handler:   handler,
		TLSConfig: config,
	}

	return server.ListenAndServe()
	// We currently only use the cert-related stuff from tls.Config,
	// so we don't need to make a full copy.

}
