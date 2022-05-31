package endpoint

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	modelV1 "github.com/ulfa24433/employee-task/db/v1/model"
	utilV1 "github.com/ulfa24433/employee-task/util/v1"
)

type AddSalaryRequest struct {
	BasicSalary int64 `json:"basic_salary"`
	Bonus       int64 `json:"bonuses"`
	EmployeeId  int64 `json:"employee_id"`
}

func AddSalary(c *gin.Context) {
	payload := &AddSalaryRequest{}
	err := c.BindJSON(&payload)
	if err != nil {
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

	//Check employee id
	emp := &modelV1.Employes{}
	err = tx.
		Model(&emp).
		Where("employee_id = ?", payload.EmployeeId).
		First(&emp).
		Error
	if err != nil {
		log.Error(err)
		utilV1.CallServerError(c, "An error occure", err)
		return
	}

	if emp.ID == 0 {
		err = fmt.Errorf("Employee id not found")
		log.Error(err)
		utilV1.CallServerError(c, "An error occured", err)
		return

	}

	//Add data to Salary table
	salary := &modelV1.Salary{
		BasicSalary: payload.BasicSalary,
		Bonuses:     payload.Bonus,
		IdEmployee:  payload.EmployeeId,
	}
	err = tx.Create(&salary).Error
	if err != nil {
		log.Error(err)
		utilV1.CallServerError(c, "An error occure", err)
		return
	}

	//update field employee_salary in table employes
	now := time.Now().Unix()
	var totalSalary int64 = payload.BasicSalary + payload.Bonus
	updates := map[string]interface{}{
		"updated_at": now,
	}

	if totalSalary != 0 {
		updates["salary_employee"] = totalSalary
	}

	err = tx.
		Model(&emp).
		Where("employee_id = ?", payload.EmployeeId).
		Updates(updates).
		Error
	if err != nil {
		log.Error(err)
		utilV1.CallServerError(c, "An error occure", err)
		return
	}
	tx.Commit()
	utilV1.CallSuccess0K(c, "success", nil)
}
