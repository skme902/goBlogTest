package main

import (
	"fmt"
	"github.com/skme902/goBlogTest/dbScript"
	)

func main(){
	fmt.Println("Go Blog Test")
	dbScript.LoadSchema()
}
