package model

type Employes struct {
	ID        int64  `gorm:"column:employee_id"`
	Name      string `gorm:"column:employee_name"`
	Salary    int64  `gorm:"column:employee_salary"`
	Nric      string `gorm:"column:employee_nric"`
	IsActive  bool   `gorm:"column:is_active"`
	CreatedAt int64  `gorm:"column:created_at"`
	UpdatedAt int64  `gorm:"column:updated_at"`
}

type Salary struct {
	ID          int64 `gorm:"column:id"`
	BasicSalary int64 `gorm:"column:basic_salary"`
	Bonuses     int64 `gorm:"column:bonuses"`
	IdEmployee  int64 `gorm:"column:employee_id"`
}
