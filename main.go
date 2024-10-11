package main

import (
    "database/sql"
    "fmt"
    "log"
    "os"

    _ "github.com/sijms/go-ora/v2"
)

func main() {
    oraclePort := os.Getenv("ORACLE_PORT")
    oraclePassword := os.Getenv("ORACLE_PASSWORD") 
    oracleServiceName := os.Getenv("ORACLE_SERVICE_NAME")

    connStr := fmt.Sprintf("oracle://%s:%s@%s:%s/%s", "user", oraclePassword, "db", oraclePort, oracleServiceName)

    fmt.Println("Connecting with:", connStr) 

    db, err := sql.Open("oracle", connStr)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    err = db.Ping()
    if err != nil {
        fmt.Println("Connection error:", err)
    } else {
        fmt.Println("CONNECTED! Damn, you good")
    }
}
