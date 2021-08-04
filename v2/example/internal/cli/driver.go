package cli

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/yandex-cloud/ydb-go-sdk/v2"
	"github.com/yandex-cloud/ydb-go-sdk/v2/auth/iam"
	"github.com/yandex-cloud/ydb-go-sdk/v2/auth/metadata"
	"github.com/yandex-cloud/ydb-go-sdk/v2/internal/traceutil"
)

func ExportTLSConfig(flag *flag.FlagSet) func() *tls.Config {
	var rootCAs string
	flag.StringVar(&rootCAs,
		"root-ca", os.Getenv("YDB_SSL_ROOT_CERTIFICATES_FILE"),
		"path to the root certificates file",
	)
	return func() *tls.Config {
		if rootCAs == "" {
			return nil
		}
		c := new(tls.Config)
		c.RootCAs = mustReadRootCerts(rootCAs)
		return c
	}
}

func ExportDriverConfig(ctx context.Context, flag *flag.FlagSet) func(Parameters) *ydb.DriverConfig {
	var (
		config ydb.DriverConfig
		trace  bool
	)
	flag.BoolVar(&trace,
		"driver-trace", false,
		"trace all driver events",
	)
	flag.DurationVar(&config.DiscoveryInterval,
		"driver-discovery", 0,
		"driver's discovery interval",
	)
	return func(params Parameters) *ydb.DriverConfig {
		if trace {
			var dtrace ydb.DriverTrace
			traceutil.Stub(&dtrace, func(name string, args ...interface{}) {
				log.Printf(
					"[driver] %s: %+v",
					name, traceutil.ClearContext(args),
				)
			})
			config.Trace = dtrace
		}

		config.Database = params.Database
		config.Credentials = credentials(ctx)

		return &config
	}
}

func credentials(ctx context.Context) ydb.Credentials {
	if token := os.Getenv("YDB_TOKEN"); token != "" {
		return ydb.AuthTokenCredentials{
			AuthToken: token,
		}
	}
	if addr := os.Getenv("YDB_METADATA"); addr != "" {
		return &metadata.Client{
			Addr: addr,
		}
	}

	// iam (jwt)
	if pk, path := os.Getenv("SA_PRIVATE_KEY_FILE"), os.Getenv("SA_SERVICE_FILE"); pk != "" || path != "" {
		var opts []iam.ClientOption

		// with service account file
		if path != "" {
			opts = append(opts, iam.WithServiceFile(path))

			// with private key file, key id and issuer id
		} else {
			opts = append(opts,
				iam.WithPrivateKeyFile(pk),
				iam.WithKeyID(mustGetenv("SA_ACCESS_KEY_ID")),
				iam.WithIssuer(mustGetenv("SA_ID")),
			)
		}

		if e := os.Getenv("SA_ENDPOINT"); e != "" {
			opts = append(opts, iam.WithEndpoint(e))
		} else {
			opts = append(opts, iam.WithDefaultEndpoint()) // iam.api.cloud.yandex.net:443
		}

		if ca := os.Getenv("SSL_ROOT_CERTIFICATES_FILE"); ca != "" {
			opts = append(opts, iam.WithCertPoolFile(ca))
		} else {
			opts = append(opts, iam.WithSystemCertPool())
		}

		c, err := iam.NewClient(opts...)
		if err != nil {
			panic(fmt.Errorf("configure credentials error: %w", err))
		}
		return c
	}

	// iam metadata
	if url, ok := os.LookupEnv("IAM_METADATA"); ok {
		if url != "" {
			return iam.InstanceServiceAccountURL(ctx, url)
		}
		// use default endpoint
		return iam.InstanceServiceAccount(ctx)
	}
	return nil
}

func mustGetenv(name string) string {
	x := os.Getenv(name)
	if x == "" {
		panic(fmt.Sprintf("environment parameter is missing or empty: %q", name))
	}
	return x
}

func readRootCerts(path string) (*x509.CertPool, error) {
	p, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	roots, err := x509.SystemCertPool()
	if err != nil {
		return nil, err
	}
	if ok := roots.AppendCertsFromPEM(p); !ok {
		return nil, fmt.Errorf("parse pem error")
	}
	return roots, nil
}

func mustReadRootCerts(path string) *x509.CertPool {
	roots, err := readRootCerts(path)
	if err != nil {
		panic(fmt.Errorf("read root certs error: %w", err))
	}
	return roots
}
