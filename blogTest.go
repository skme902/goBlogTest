package main

import (
	"fmt"
	"github.com/skme902/goBlogTest/dbScript/schemaLoader"
	)

func main(){
	fmt.Println("Go Blog Test")
	schemaLoader.LoadSchema()
}
