package schemaLoader

import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
	"io/ioutil"
)

const (
	DB_USER = "sanjeevkumar"
	DB_PASSWORD  = "postgres"
	DB_NAME = "go_blog"
	)

func LoadSchema()  {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	db, err = sql.Open("postgres", dbinfo)
	checkErr(err)
	defer db.close()

	fmt.Printf("Loading schema from file")
	dat, err := ioutil.ReadFile("./schema.sql")
	checkErr(err)
	datStr := string(dat)
	fmt.Print(datStr)
	res, err := db.Exec(datStr)
	check(err)
	fmt.Printf("Output %v", res)
}

func checkErr(err error){
	if err != nil{
		panic(err)
	}
}
