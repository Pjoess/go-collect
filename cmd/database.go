package main

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/glebarez/go-sqlite"
)

type CD struct {
	ID   uint
	Name string
}

func db() (string, error) {
	// Connect to the SQLite database
	db, err := sql.Open("sqlite", "./my.db")
	if err != nil {
		return "", err
	}

	defer db.Close()
	fmt.Println("Connected to the SQLite database successfully.")

	// Get the version of SQLite
	var sqliteVersion string
	err = db.QueryRow("select sqlite_version()").Scan(&sqliteVersion)
	if err != nil {
		return "", err
	}

	return sqliteVersion, nil
}

func newDb(stringSlice []string, tablename string) {

	var stringBuilder strings.Builder
	catStr := ""
	for i := range stringSlice {
		catStr += " " + stringSlice[i]
		stringBuilder.WriteString(stringSlice[i])
	}
	var catString = stringBuilder.String()

	tableName := "Tim"
	createTableSQL := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name STRING
	);`, tableName)

	db, err := sql.Open("sqlite", "./my.db")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()

	_, err = db.Exec(createTableSQL)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("succesfully made table named %s", tableName)

}
