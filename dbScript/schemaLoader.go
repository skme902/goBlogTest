package dbScript

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"io/ioutil"
	"os"
)

const (
	DB_USER = "sanjeevkumar"
	DB_PASSWORD  = "postgres"
	DB_NAME = "go_blog"
	)

func checkErr(err error){
	if err != nil{
		panic(err)
	}
}

func ReadSchemaDefFromFile() string{
	basePath, _ := os.Getwd()
	schemaFilePath := basePath + "/dbScript/schema.sql"

	fmt.Printf("Loading schema from file: %s\n", schemaFilePath)
	dat, err := ioutil.ReadFile(schemaFilePath)
	checkErr(err)
	datStr := string(dat)
	return datStr
}

func ExecuteSQLScript(db *sql.DB, scriptStr string){
	res, err := db.Exec(scriptStr)
	checkErr(err)
	fmt.Printf("Output %v", res)
}

func LoadSchema()  {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	fmt.Printf("%s\n", dbinfo)
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)
	defer db.Close()

	datStr := ReadSchemaDefFromFile()

	ExecuteSQLScript(db, datStr)
}
