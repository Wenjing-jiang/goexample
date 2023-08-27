package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	//_ "github.com/mattn/go-sqlite3"
)

func main() {
	// 连接 MySQL 数据库
	mysqlInfo := "root:1999@tcp(localhost:3306)/wltx?charset=utf8"
	db, err := sql.Open("mysql", mysqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// 创建新表
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS employees (
			first_name VARCHAR(255),
			last_name VARCHAR(255),
			age INT
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("New table created successfully!")

	// 插入数据
	insert := "INSERT INTO Employees(first_name, last_name, age) VALUES (?, ?, ?)"
	result, err := db.Exec(insert, "Alice", "Smith", 30)
	if err != nil {
		fmt.Println("Error inserting data:", err)
	} else {
		fmt.Println("Data inserted successfully")
	}
	// 在这里可以执行数据库操作
	fmt.Println("Connected to MySQL database!")

	// 更新数据
	update := "UPDATE Employees SET age = ? WHERE first_name = ? AND last_name = ?"
	result, err = db.Exec(update, 31, "Alice", "Smith")
	if err != nil {
		fmt.Println("Error updating data:", err)
	} else {
		fmt.Println("Data updated successfully")
		rowsAffected, _ := result.RowsAffected()
		fmt.Println("Rows affected:", rowsAffected)
	}

	// 查询数据
	query := "SELECT * FROM Employees"
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("Error querying data:", err)
		return
	}
	defer rows.Close()

	fmt.Println("Fetching data from Employees table:")
	var id, age int
	var firstName, lastName string
	for rows.Next() {
		err := rows.Scan(&id, &firstName, &lastName, &age)
		if err != nil {
			fmt.Println("Error scanning data:", err)
		} else {
			fmt.Printf("ID: %d, Name: %s %s, Age: %d\n", id, firstName, lastName, age)
		}
	}

	// 删除数据
	delete := "DELETE FROM Employees WHERE first_name = ? AND last_name = ?"
	result, err = db.Exec(delete, "Alice", "Smith")
	if err != nil {
		fmt.Println("Error deleting data:", err)
	} else {
		fmt.Println("Data deleted successfully")
		rowsAffected, _ := result.RowsAffected()
		fmt.Println("Rows affected:", rowsAffected)
	}
}
