package models

import (
	"fmt"
	conf "golang-simple-crud/config"
	dts "golang-simple-crud/datastruct"
)

//SaveEmployee .
func SaveEmployee(conn *conf.Connection, name string, phone string, email, address string) ([]dts.EmployeeSave, error) {
	arrEmployeeSave := []dts.EmployeeSave{}
	strEmployeeSave := dts.EmployeeSave{}

	tx, err := conn.Begin()
	if err != nil {
		return arrEmployeeSave, err
	}

	stmt, err := tx.Prepare(`insert into employee (Name,Phone,Email,Address,CreatedDate)
								   values (?,?,?,?,now())`)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(name, phone, email, address)

	if err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()

	rows, err := res.LastInsertId()

	strEmployeeSave.ID = rows
	arrEmployeeSave = append(arrEmployeeSave, strEmployeeSave)

	if err != nil {
		return nil, err
	}

	return arrEmployeeSave, nil
}

//UpdateEmployee .
func UpdateEmployee(conn *conf.Connection, name string, phone string, email, address string, id int) (int64, error) {
	tx, err := conn.Begin()
	if err != nil {
		return 0, err
	}

	stmt, err := tx.Prepare(`update employee set Name=?,Phone=?,Email=?,Address=?,UpdatedDate=now()
							  where Id=?`)

	if err != nil {
		tx.Rollback()
		return 0, err
	}

	defer stmt.Close()

	//if _, err := stmt.Exec(name, phone, email, address, id); err != nil {
	//	tx.Rollback()
	//		return err
	//}

	res, err := stmt.Exec(name, phone, email, address, id)

	if err != nil {
		tx.Rollback()
		return 0, err
	}

	tx.Commit()

	rows, err := res.RowsAffected()

	if err != nil {
		return 0, err
	}

	fmt.Println(rows)
	return rows, nil
}

//GetEmployee as
func GetEmployee(conn *conf.Connection) ([]dts.Employee, error) {
	arrEmployee := []dts.Employee{}
	strEmployee := dts.Employee{}
	strQuery := `select * from employee`
	rows, err := conn.Query(strQuery)

	conf.Logf("Query Employee : %s", strQuery)

	defer rows.Close()
	if err != nil {
		return nil, err
	}

	for rows.Next() {

		err := rows.Scan(
			&strEmployee.ID,
			&strEmployee.Name,
			&strEmployee.Phone,
			&strEmployee.Email,
			&strEmployee.Address,
			&strEmployee.CreatedDate,
			&strEmployee.UpdatedDate,
		)

		if err != nil {
			return nil, err
		}

		arrEmployee = append(arrEmployee, strEmployee)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return arrEmployee, nil
}

//GetEmployeeDetail ..
func GetEmployeeDetail(conn *conf.Connection, ID string) ([]dts.EmployeeDetail, error) {
	arrEmployeeDetail := []dts.EmployeeDetail{}
	strEmployeeDetail := dts.EmployeeDetail{}

	strQuery := `select * from employee where Id=?`

	rows, err := conn.Query(strQuery, ID)

	conf.Logf("Query Employee  detail: %s %s", strQuery, ID)

	defer rows.Close()
	if err != nil {
		return nil, err
	}

	for rows.Next() {

		err := rows.Scan(&strEmployeeDetail.ID,
			&strEmployeeDetail.Name,
			&strEmployeeDetail.Phone,
			&strEmployeeDetail.Email,
			&strEmployeeDetail.Address,
			&strEmployeeDetail.CreatedDate,
			&strEmployeeDetail.UpdatedDate,
		)

		if err != nil {
			return nil, err
		}

		if err = rows.Err(); err != nil {
			return nil, err
		}

		arrEmployeeDetail = append(arrEmployeeDetail, strEmployeeDetail)
	}

	return arrEmployeeDetail, nil
}
