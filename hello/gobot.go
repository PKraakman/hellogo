package main

import (
	"fmt"
)

type Object map[string]interface{}

type Person Object



func main(){
	var persa = Person{"fname":"Pieter","lname":"Kraakman","age":56}
	var persb = Person{"fname":"Tina","lname":"Bour","age":57}

	persa["spouse"] = persb

	for key := range persa {
		fmt.Println(key,":",persa[key])
	}

	wife := persa["spouse"].(Person)

	fmt.Println(wife["lname"])

	fmt.Println(persa["spouse"].(Person)["fname"])

}