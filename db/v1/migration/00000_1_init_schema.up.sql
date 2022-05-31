CREATE TABLE "employes" (
    "employee_id" SERIAL PRIMARY KEY,
    "employee_name" VARCHAR(255) NOT NULL,
    "employee_salary" int,
    "employee_nric" VARCHAR (20) NOT NULL,
    "is_active" BOOLEAN NOT NULL,
    "created_at" int,
    "updated_at" int
);

CREATE TABLE "salaries" (
    "id" SERIAL PRIMARY KEY,
    "basic_salary" int,
    "bonuses" int,
    "employee_id" int,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    foreign key(employee_id) references employes(employee_id)
  
);