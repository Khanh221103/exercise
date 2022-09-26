package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

type Serve struct {
	Name  string `json:"name"`
	Class string `json:"class"`
}

type User struct {
	name    string
	age     int64
	gender  bool
	address string
}

func main() {
//===========EX2============
	content, err := os.ReadFile("./serve.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	var serves []*Serve
	err = json.Unmarshal(content, &serves)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
	for _, u := range serves {
		log.Printf("name: %s\n", u.Name)
		log.Printf("class: %s\n", u.Class)
	} 

	fmt.Println("======")
//===========EX3============
	key := "Admin"
	for i := range serves {
		class := serves[i].Class
		if strings.Contains(class, key) {
			fmt.Printf("Name: %s\n", serves[i].Name)
			fmt.Printf("Class: %s\n", serves[i].Class)
		}
	}
	// ex3(serves)

	fmt.Println("======")
//===========EX4============
	serves = append(serves, &Serve{Name: "fileCustome", Class: "org.cof ax.cds.FileServlet.Custome"})
	for _, u := range serves {
		log.Printf("name: %s\n", u.Name)
		log.Printf("class: %s\n", u.Class)
	}

	fmt.Println("======")
//===========EX5============
	fmt.Println(serves)

	fmt.Println("======")
//===========EX6============
	slice := []int{11, 34, 56, 77, 99, 109, 66, 20, 88, 34}
	sort.Ints(slice)
	fmt.Println(slice)

	c := slice[1:8]
	fmt.Println(c)

	// copyy := make([]int, len(c))
	// copy(copyy, c[1:15])
	// fmt.Println("copy:", copyy)

	fmt.Println("=====")
//===========EX7============
	
	u1 := User{
		name: "a",
		age: 18,
		gender: true,
		address: "HN",
	}
	u2 := User{
		name: "b",
		age: 20,
		gender: false,
		address: "HN",
	}
	u3 := User{
		name: "c",
		age: 30,
		gender: false,
		address: "HCM",
	}
	var userName = map[string]*User{u1.name: &u1, u2.name: &u2, u3.name: &u3}

	for b, a := range userName {
		fmt.Println(b, a)
	}

}
