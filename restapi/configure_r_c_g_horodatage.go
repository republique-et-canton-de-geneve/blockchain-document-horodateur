package restapi

import (
	"context"
	//	"crypto/rsa"
	"crypto/tls"
	//	"crypto/x509"
	"fmt"
	"github.com/Genova/bcp-genova/blockchain-document-horodateur/internal"
	"github.com/Genova/bcp-genova/blockchain-document-horodateur/restapi/operations"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"math/big"
	"net/http"
)

// This file is safe to edit. Once it exists it will not be overwritten

//go:generate swagger generate server --target .. --spec ../docs/rc-ge-ch.yml

var ethopts struct {
	WsURI      string `long:"ws-uri" env:"WS_URI" description:"Ethereum WS URI (e.g: ws://HOST:8546)"`
	PrivateKey string `long:"pkey" env:"PRIVATE_KEY" description:"hex encoded private key"`
	LockedAddress string `long:"locked-addr" env:"LOCKED_ADDR" description:"Ethereum address of the sole verifier (anchor emitter)"`
	ErrorThreshold big.Float`long:"error-threshold" env:"ERROR_THRESHOLD" description:""`
	WarningThreshold big.Float `long:"warning-threshold" env:"WARNING_THRESHOLD" description:""`
}

var serviceopts struct {
	DbDSN string `long:"db-dsn" env:"DB_DSN" description:"Database DSN (e.g: /tmp/test.sqlite)"`
}

func configureFlags(api *operations.RCGHorodatageAPI) {
	ethOpts := swag.CommandLineOptionsGroup{
		LongDescription:  "",
		ShortDescription: "Ethereum options",
		Options:          &ethopts,
	}
	serviceOpts := swag.CommandLineOptionsGroup{
		LongDescription:  "",
		ShortDescription: "Service options",
		Options:          &serviceopts,
	}
	api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ethOpts, serviceOpts}
	}

//func configureSAML() {
//	pblcKey, err := filepath.Abs("myservice.cert")
//	pvtKey, err := filepath.Abs("myservice.key")
//	fmt.Println(pblcKey, pvtKey)
//	keyPair, err := tls.LoadX509KeyPair(pblcKey, pvtKey)
//	fmt.Println(keyPair)
//	if err != nil {
//		panic(err) // TODO handle error
//	}
//	keyPair.Leaf, err = x509.ParseCertificate(keyPair.Certificate[0])
//	if err != nil {
//		panic(err) // TODO handle error
//	}
//	idpMetadataURL, err := url.Parse("http://ec2-18-184-234-216.eu-central-1.compute.amazonaws.com/ssorec.geneveid.ch_dgsi_blockchain.xml")
//	if err != nil {
//		panic(err) // TODO handle error
//	}
//	rootURL, err := url.Parse("http://127.0.0.1:8001/")
//	if err != nil {
//		panic(err) // TODO handle error
//	}
//	samlSP, _ := samlsp.New(samlsp.Options{
//		URL:            *rootURL,
//		Key:            keyPair.PrivateKey.(*rsa.PrivateKey),
//		Certificate:    keyPair.Leaf,
//		IDPMetadataURL: idpMetadataURL,
//	})
//	fmt.Println(samlSP)
//}

//func configureSAML(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, "Hello, %s!", samlsp.Token(r.Context()).Attributes.Get("cn"))
//}

func configureAPI(api *operations.RCGHorodatageAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError
	fmt.Println(ethopts.LockedAddress)
	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// s.api.Logger = log.Printf
	ctx := internal.NewDBToContext(context.Background(), serviceopts.DbDSN)
	ctx = internal.NewCCToContext(ctx, ethopts.WsURI)
	ctx = internal.NewBLKToContext(ctx, ethopts.WsURI, ethopts.PrivateKey)
	ctx = internal.NewMonitoringToContext(ctx, ethopts.WsURI, ethopts.LockedAddress, ethopts.ErrorThreshold, ethopts.WarningThreshold)

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.BinProducer = runtime.ByteStreamProducer()

	api.GetreceiptHandler = operations.GetreceiptHandlerFunc(func(params operations.GetreceiptParams) middleware.Responder {
		return internal.GetreceiptHandler(ctx, params)
	})
	api.DelreceiptsHandler = operations.DelreceiptsHandlerFunc(func(params operations.DelreceiptsParams) middleware.Responder {
		return internal.DelreceiptsHandler(ctx	, params)
	})
	api.ListtimestampedHandler = operations.ListtimestampedHandlerFunc(func(params operations.ListtimestampedParams) middleware.Responder {
		return internal.ListtimestampedHandler(ctx, params)
	})
	api.MonitoringHandler = operations.MonitoringHandlerFunc(func(params operations.MonitoringParams) middleware.Responder {
		return internal.MonitoringHandler(ctx, params)
	})
	api.ConfigureSAMLHandler = operations.ConfigureSAMLHandlerFunc(func(params operations.ConfigureSAMLParams) middleware.Responder {
		return internal.ConfigureSAMLHandler(ctx, params)
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(ctx, api.Serve(setupMiddlewares))
}


// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(ctx context.Context, handler http.Handler) http.Handler {
	return internal.UploadHandler(ctx, "/upload", handler)
}
