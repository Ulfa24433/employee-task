package v1

import (
	"fmt"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"github.com/ulfa24433/employee-task/util/v1/envvar"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetPostgreClient() (db *gorm.DB, err error) {
	rdsHost, err := envvar.LoadEnvVar(envvar.RDSPostgresHost, true)
	if err != nil {
		log.Error(err)
		return
	}
	rdsUsername, err := envvar.LoadEnvVar(envvar.RDSPostgresUsername, true)
	if err != nil {
		log.Error(err)
		return
	}
	rdsPassword, err := envvar.LoadEnvVar(envvar.RDSPostgresPassword, true)
	if err != nil {
		log.Error(err)
		return
	}
	rdsDatabase, err := envvar.LoadEnvVar(envvar.RDSPostgresDatabase, true)
	if err != nil {
		log.Error(err)
		return
	}
	rdsSslCert, err := envvar.LoadEnvVar(envvar.RDSPostgresSslCert, true)
	if err != nil {
		log.Error(err)
		return
	}
	rdsPort, err := envvar.LoadEnvVar(envvar.RDSPostgresPort, true)
	if err != nil {
		log.Error(err)
		return
	}
	timezone, err := envvar.LoadEnvVar(envvar.Timezone, true)
	if err != nil {
		log.Error(err)
		return
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		rdsHost, rdsUsername, rdsPassword, rdsDatabase, rdsPort, rdsSslCert, timezone)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Error(err)
		return
	}
	return
}
