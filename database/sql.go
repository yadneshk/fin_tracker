package database

import (
	"fmt"
    "log"
    "time"
	"database/sql"
    "strings"
	// "encoding/json"
	_ "github.com/go-sql-driver/mysql"
)

func convertToMySQLData(expense []string) ([]string, error) {

    exp := make([]string, 4)

	// Parse the input date
	parsedDate, err := time.Parse("02-01-2006", expense[0])
	if err != nil {
		return []string{}, err
	}

	// Format it as MySQL date
	exp[0] = parsedDate.Format("2006-01-02")

    exp[1] = expense[2]

    exp[2] = strings.Trim(expense[3], " ")

    exp[3] = strings.Trim(expense[4], " ")


	return exp, nil
}

func InsertExpenses(expenses [][]string) {
    // Replace these with your MySQL database credentials
    dbUser := "root"
    dbPass := "my-secret-pw"
    dbName := "fin_tracker"
    dbHost := "localhost" // or the address of your MySQL server
    dbPort := "3306"

    // Create a connection string
    connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)

    // Open a database connection
    db, err := sql.Open("mysql", connectionString)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Check if the database connection is working
    err = db.Ping()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Connected to MySQL database!")

    // Insert each element of the string array into the database
    for _, value := range expenses {

        expense, err := convertToMySQLData(value)
        if err != nil {
            log.Fatal(err)
        }

        err = insertValue(db, expense)

        if err != nil {
            log.Fatal(err)
        }
    }

    fmt.Println("String array inserted successfully!")
}

func newNullString(s string) sql.NullString {
    if len(s) == 0 {
        return sql.NullString{}
    }
    return sql.NullString{
         String: s,
         Valid: true,
    }
}

func insertValue(db *sql.DB, values []string) error {

    // Define your SQL insert statement
    insertStatement := "INSERT INTO expenses (txn_date, particulars, debit, credit) VALUES (?, ?, ?, ?)"

    // Prepare the statement
    stmt, err := db.Prepare(insertStatement)
    if err != nil {
        return err
    }
    defer stmt.Close()

    // Execute the statement with the value
    _, err = stmt.Exec(values[0],values[1],newNullString(values[2]),newNullString(values[3]))
    if err != nil {
        return err
    }

    return nil
}