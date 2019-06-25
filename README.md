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

INSERT into employee(Name,Phone,Email,Address,CreatedDate,UpdatedDate)
VALUES('dadang','888121222','dadang@gmail.com','Jalan raya BSD no 214',now(),now());

INSERT into employee(Name,Phone,Email,Address,CreatedDate,UpdatedDate)
VALUES('dudung','88888888','dudung@gmail.com','Jalan raya BSD no 212',now(),now())

# Test Curl

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


