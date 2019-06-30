# golang-simple-crud

# Database
Create database golang_crud;

create table employee(
Id int not null PRIMARY key AUTO_INCREMENT,
Name varchar(40),
Phone int(13),
Email varchar(70),
Address varchar(100),
CreatedDate timestamp,
UpdatedDate timestamp
);

INSERT into employee(Name,Phone,Email,Address,CreatedDate)
VALUES('dadang','888121222','dadang@gmail.com','Jalan raya BSD no 214',now());

INSERT into employee(Name,Phone,Email,Address,CreatedDate,UpdatedDate)
VALUES('dudung','88888888','dudung@gmail.com','Jalan raya BSD no 212',now());

# Test Curl

## New Employee
curl -X POST \
  http://localhost:3000/api/employeeSave \
  -H 'Content-Type: application/json' \
  -H 'Postman-Token: 01008f78-804d-4ce1-b25c-b2c6a3ad746e' \
  -H 'cache-control: no-cache' \
  -d '{

"name":"dadang",
"email":"dadang@gmial.com",
"phone":"88881111222",
"address":"Jl. Pulau seribu no 212, Tangsel"
}'

## Update Employee

curl -X POST \
  http://localhost:3000/api/employeeUpdate \
  -H 'Accept: */*' \
  -H 'Cache-Control: no-cache' \
  -H 'Connection: keep-alive' \
  -H 'Content-Type: application/json' \
  -d '{
"id":1,
"name":"dudung",
"email":"dudung@gmial.com",
"phone":"88881111222",
"address":"Jl. Pulau seribu no 212, Tangsel"
}'



## Get Employee List
curl -X POST \
  http://localhost:3000/api/employee \
  -H 'Cache-Control: no-cache' \
  -H 'Connection: keep-alive' \
  -H 'Content-Type: application/json' \
  
  ## Get Employee Detail
  curl -X POST \
  http://localhost:3000/api/employeeDetail \
  -H 'Content-Type: application/json' \
  -d '{
  "Id":"1"
  }'


