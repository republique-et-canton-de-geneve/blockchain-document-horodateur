package internal

import (
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/crewjam/saml/samlsp"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"net/http"
	"net/url"
	"path/filepath"
)

func configureSAML() {
	pblcKey, err := filepath.Abs("myservice.cert")
	pvtKey, err := filepath.Abs("myservice.key")
	fmt.Println(pblcKey, pvtKey)
	keyPair, err := tls.LoadX509KeyPair(pblcKey, pvtKey)
	fmt.Println(keyPair)
	if err != nil {
		panic(err) // TODO handle error
	}
	keyPair.Leaf, err = x509.ParseCertificate(keyPair.Certificate[0])
	if err != nil {
		panic(err) // TODO handle error
	}
	idpMetadataURL, err := url.Parse("http://ec2-18-184-234-216.eu-central-1.compute.amazonaws.com/ssorec.geneveid.ch_dgsi_blockchain.xml")
	if err != nil {
		panic(err) // TODO handle error
	}
	rootURL, err := url.Parse("http://127.0.0.1:8001/")
	if err != nil {
		panic(err) // TODO handle error
	}
	samlSP, _ := samlsp.New(samlsp.Options{
		URL:            *rootURL,
		Key:            keyPair.PrivateKey.(*rsa.PrivateKey),
		Certificate:    keyPair.Leaf,
		IDPMetadataURL: idpMetadataURL,
	})
	fmt.Println(samlSP)
	app := http.HandlerFunc(hello)
	http.Handle("/hello", samlSP.RequireAccount(app))
	http.Handle("/saml/", samlSP)
	http.ListenAndServe(":8080", nil)
	return
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", samlsp.Token(r.Context()).Attributes.Get("cn"))
}