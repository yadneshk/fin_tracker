package database

import (
	"fmt"
    "log"
	"database/sql"
	// "encoding/json"
	_ "github.com/go-sql-driver/mysql"
)

func execDatabaseQuery(query string, dbName string) {
    dbUser := "root"
    dbPass := "my-secret-pw"
    dbHost := "127.0.0.1"
    dbPort := "3306"

    // Create a connection string
    connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

    // Open a database connection
    db, err := sql.Open("mysql", connectionString)
    if err != nil {
        log.Fatal(err)
    }
    log.Println("MySQL database opened!")
    defer db.Close()

    // Check if the database connection is working
    err = db.Ping()
    if err != nil {
        if err.Error() == fmt.Sprintf("Error 1049 (42000): Unknown database '%s'", dbName) {
            connectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/", dbUser, dbPass, dbHost, dbPort)
            db, err = sql.Open("mysql", connectionString)
        } else {
            log.Fatal(err)
        }
    }
    log.Println("Connected to MySQL database!")

    _, err = db.Exec(query)
    if err != nil {
        log.Fatal(err)
    }

    log.Println("Query created successfully!\n")
}

func prepareDatabase() {
    // Create a new database
    dbName := "fin_tracker"
    createDBStatement := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", dbName)
    execDatabaseQuery(createDBStatement, dbName)
    log.Printf("Database '%s' created successfully!\n", dbName)

    // Create a new database
    tableName := "expenses"
    createTableStatement := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (txn_id int primary key auto_increment, txn_date date, particulars varchar(255), debit float(2), credit float(2))", tableName)
    execDatabaseQuery(createTableStatement, dbName)
    log.Printf("Table '%s' created successfully!\n", tableName)

}


func InsertJSONData() {
    prepareDatabase()

	// // Define your SQL insert statement
	// insertStatement := "INSERT INTO your_table_name (json_column) VALUES (?)"

	// // Prepare the statement
	// stmt, err := db.Prepare(insertStatement)
	// if err != nil {
	// 	return err
	// }
	// defer stmt.Close()

    // // Execute the statement with the JSON data string
    // _, err = stmt.Exec(string(jsonData))
    // if err != nil {
    //     return err
    // }

    // return nil
	
}
