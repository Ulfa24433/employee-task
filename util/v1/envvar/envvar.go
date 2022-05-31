package envvar

import (
	"fmt"
	"os"
)

const (
	RDSPostgresHost     = "RDS_POSTGRES_HOST"
	RDSPostgresPort     = "RDS_POSTGRES_PORT"
	RDSPostgresDatabase = "RDS_POSTGRES_DATABASE"
	RDSPostgresUsername = "RDS_POSTGRES_USERNAME"
	RDSPostgresPassword = "RDS_POSTGRES_PASSWORD"
	RDSPostgresSslCert  = "RDS_POSTGRES_SSL_CERT"
	Port                = "PORT"
	Timezone            = "TIMEZONE"
	LenNric             = "LEN_NRIC"
)

func LoadEnvVar(key string, isRequired bool) (out string, err error) {
	if key == "" {
		err = fmt.Errorf("missing key")
		return
	}
	out = os.Getenv(key)
	if out == "" && isRequired {
		err = fmt.Errorf("Missing ENVVAR %s", key)
		return
	}
	return
}
