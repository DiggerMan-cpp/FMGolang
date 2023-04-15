package main

import (
    "bufio"
    "fmt"
    "io/ioutil"
    "os"
   "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"

)

func main() {
    file, err := os.Open("input.txt")
    if err != nil {
        panic(err)
    }
    defer file.Close()


    scanner := bufio.NewScanner(file)
    var lines []string
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }

   
    var result string
    for _, line := range lines {
        result += line + "n"
    }

   
    err = ioutil.WriteFile("output.txt", []byte(result), 0644)
    if err != nil {
        panic(err)
    }
}


func main() {

    db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/database_name")
    if err != nil {
        panic(err.Error())
    }
    defer db.Close()
    rows, err := db.Query("SELECT * FROM table_name")
    if err != nil {
        panic(err.Error())
    }
    defer rows.Close()

   }
}
