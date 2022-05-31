# Documentation

## 1. Clone or extract file

extract this zip file into your golang development environment

## 2. Installation

setup your local machine by install 

- Docker 

- gin-gonic

- gorm

## 3. Setup Envvar

in your local machine, make sure that .env file is already created. this is to load certain values for database connection and run the server in port 50050.
we can refer to .env.schema to see all envvar that we use to run application and fill up with the below value.

```
  RDS_POSTGRES_HOST=localhost
  RDS_POSTGRES_PORT=5432
  RDS_POSTGRES_DATABASE=emp_test
  RDS_POSTGRES_USERNAME=root
  RDS_POSTGRES_PASSWORD=secret
  RDS_POSTGRES_SSL_CERT=disable
  TIMEZONE=Asia/Singapore
  PORT=50050
```
## 4. Run Command

once everyhing is installed please execute below command. 
for futher detail, please go to ```Makefile``` in root folder. it will contain all command that will be used to use the application.

- ```sudo make postgres```
  this command is to setup database in docker with name postgres12.
  make sure the images is is created and container is up with command ```sudo docker images``` and ```sudo docker ps -a```

- ```sudo make createdb```
  this command is to create database

- ```sudo make dropdb```
  this command is to drop database

- ```sudo make migrateup```
  this command is to apply data migration. 
  this will create model and seed database value after table creation.

- ```sudo make migratedown```
  this command is to delete our table that we've created. 

- ```sudo make gotodb```
  this command is to go our database. we may check whether the table is created by excute ```\dt``` to find table list

## 5. Run Server

- execute in terminal ```go run main.go``` or simply hit f5 in to run the server. server will be run on localhost:50050


- from your browser or postman, please run below endpoint
    - ```http://localhost:50050/v1/employee/add``` to add more country.

      sample curl to add more country

      ```
        curl --location --request POST 'http://localhost:50050/v1/country/add' \
        --header 'Content-Type: application/json' \
        --data-raw '{
            "country_code": "SG",
            "country_name": "Singapore"
          }'

      ```

    - ```http://localhost:50050/v1/property/add``` to add property. this endpoint required to pass body request with PropertyName, PropertyAddress, CountryCode

      sample curl to add more property

      ```
        curl --location --request POST 'http://localhost:50050/v1/property/add' \
        --header 'Content-Type: application/json' \
        --data-raw '{
            "property_name": "Building D",
            "property_address": "Buiding D Street D No 120 Block D",
            "country_code": "SG"
          }'

      ```

    - ```http://localhost:50050/v1/employee/list``` to get list property by country code. we should specify country_code in query parameter. for example "SG"

      sample curl to get list property

      ```
        curl --location --request GET 'http://localhost:50050/v1/property/list?country_code=SG'
      
      ```



