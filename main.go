package main

import (
    "fmt"
    "database/sql"
    "log"
    "os"

    _ "github.com/sijms/go-ora/v2"
)

func main() {
    oracleHost := os.Getenv("ORACLE_HOST")
    oraclePort := os.Getenv("ORACLE_PORT")
    oraclePassword := os.Getenv("ORACLE_PWD")
    oracleServiceName := os.Getenv("ORACLE_SERVICE_NAME")
    fmt.Println(oraclePassword)

    connStr := fmt.Sprintf("oracle://%s:%s@%s:%s/%s", "user", oraclePassword, oracleHost, oraclePort, oracleServiceName)
     db, err := sql.Open("oracle", connStr)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    err = db.Ping()
    if err != nil {
        fmt.Println("Connection connecting:", err)
    } else {
        fmt.Println("CONNECTED!")
    }
}
