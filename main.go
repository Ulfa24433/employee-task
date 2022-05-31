package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	endpointV1 "github.com/ulfa24433/employee-task/endpoint"
	"github.com/ulfa24433/employee-task/util/v1/envvar"
)

func main() {
	//initialize gin

	r := gin.Default()
	v1 := r.Group("/v1")

	//endpoint to add employee data
	employe := v1.Group("/employee")
	salary := v1.Group("/salary")
	employe.POST("/add", endpointV1.AddEmployee)

	//endpoint to get employee lists
	employe.GET("/list", endpointV1.ListEmployee)

	salary.POST("/add", endpointV1.AddSalary)

	port, err := envvar.LoadEnvVar(envvar.Port, true)
	if err != nil {
		log.Error(err)
		return
	}

	err = r.Run(fmt.Sprintf(":%s", port))
	if err != nil {
		log.Error(err)

	}

}
