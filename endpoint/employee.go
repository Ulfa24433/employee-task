package endpoint

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	modelV1 "github.com/ulfa24433/employee-task/db/v1/model"
	utilV1 "github.com/ulfa24433/employee-task/util/v1"
	"github.com/ulfa24433/employee-task/util/v1/envvar"
)

type AddEmployeeRequest struct {
	Name     string `json:"employee_name"`
	Salary   int    `json:"employee_salary"`
	Nric     string `json:"employee_nric"`
	IsActive bool   `json:"is_active"`
}

func AddEmployee(c *gin.Context) {
	payload := &AddEmployeeRequest{}
	err := c.BindJSON(&payload)
	if err != nil {
		log.Error(err)
		utilV1.CallServerError(c, "invalid request", err)
		return
	}

	maxNric, err := envvar.LoadEnvVar(envvar.LenNric, true)
	if err != nil {
		log.Error(err)
		return
	}
	//convert string to integer
	i, _ := strconv.Atoi(maxNric)

	if payload.Name == "" || payload.Nric == "" || len(payload.Nric) != i {
		err = fmt.Errorf("Invalid request")
		log.Error(err)
		utilV1.CallServerError(c, "invalid request", err)
		return
	}

	//init connection to postgres
	db, err := utilV1.GetPostgreClient()
	if err != nil {
		log.Error(err)
		utilV1.CallServerError(c, "an error occure", err)
		return
	}

	tx := db.Begin()
	defer tx.Rollback()

	//Check existing record by NRIC
	emp := []*modelV1.Employes{}
	err = tx.
		Model(&emp).
		Where("employee_nric = ?", payload.Nric).
		Find(&emp).
		Error
	if err != nil {
		log.Error(err)
		utilV1.CallServerError(c, "An error occure", err)
		return
	}

	if len(emp) != 0 {
		err = fmt.Errorf("NRIC already exist")
		log.Error(err)
		utilV1.CallServerError(c, "NRIC already exist", err)
		return
	}

	//add employee to database
	employee := &modelV1.Employes{

		Name:      payload.Name,
		Salary:    int64(payload.Salary),
		Nric:      payload.Nric,
		IsActive:  true,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: 0,
	}

	err = tx.Create(&employee).Error
	if err != nil {
		log.Error(err)
		utilV1.CallServerError(c, "An error occure", err)
		return
	}
	tx.Commit()
	utilV1.CallSuccess0K(c, "success", nil)

}

func ListEmployee(c *gin.Context) {

	//init connection to postgres
	db, err := utilV1.GetPostgreClient()
	if err != nil {
		log.Error(err)
		utilV1.CallServerError(c, "an error occure", err)
		return
	}

	emp := []*modelV1.Employes{}
	err = db.
		Model(&emp).
		Where("is_active = ?", true).
		Limit(10).
		Find(&emp).
		Error
	if err != nil {
		log.Error(err)
		utilV1.CallServerError(c, "An error occure", err)
		return
	}
	utilV1.CallSuccess0K(c, "success", emp)
}
