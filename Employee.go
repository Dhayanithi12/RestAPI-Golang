
package main

import (
	"database/sql"
	"entities"
	
)

type Employee struct {
	Db *sql.DB
}

func main () {

	const (
        DB_USER     = "postgres"
        DB_PASSWORD = "dhaya"
        DB_NAME     = "User"
      )
      
      // DB set up
      func setupDB() *sql.DB {
        dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
        db, err := sql.Open("postgres", dbinfo)
        return DB


	var exists bool
    row := db.QueryRow("SELECT EXISTS(SELECT 1 FROM ...)")
    if err := row.Scan(&exists); err != nil {
        return err
    } 
    else if !exists {
    if err := db.Exec("INSERT ..."); err != nil {
        return err
    }
   }
	rows, err := EmployeeModel.Db.Query("select * from Employee limit 5, 3", offset, count)
	if err != nil {
		return nil, err
	} else {
		Employee_details := []entities.Employee{}
		for rows.Next() {
			var Name string 
			var Employee_Id int
			var Employee_Age int
			
			err2 := rows.Scan(&Name, &Employee_Id, &Employee_Age)
			if err2 != nil {
				return nil, err2
			} else {
				Employee_details := entities.Employee{Name, Employee_Id, Employee_age}
				Employee_details = append(Employee_details, Employee_details)
			}
		}
		return Employee_details, nil
		
	}
}
