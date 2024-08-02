// building a simple CLI APP
package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	fmt.Println("Hi gopher!!")
	yy, mm, dd := time.Now().Date()
	print(yy, mm, dd)
	fmt.Println("enter your Name : ")
	var name string
	fmt.Scanln(&name)
	name = strings.TrimSpace(name)
	fmt.Printf("hi %s !! Please enter Your date of birth\n", name)
	fmt.Printf("DD/MM/YYYY - ")
	var date, month, year int
	var err error
	// fmt.Scanln(&date,"/",&month,"/",&year)
	n, err := fmt.Scanf("%d %d %d", &date, &month, &year)
	if n != 3 {
		err = fmt.Errorf("please enter dob in this format DD MM YYYY ")
	}
	if month > 12 || month < 1 {
		err = fmt.Errorf("please enter correct month %d ",month)
	}
	if year > 2024 || year < 1920 {
		err = fmt.Errorf("please enter correct year %d ", year)
	}
	if date > 31 || date < 1 {
		err = fmt.Errorf("please enter correct date %d ", date)
	}
	if err != nil {
		fmt.Printf("err - %s", err)
		panic(err)
	}

	fmt.Printf("hey %s your DOB is - %d/%d/%d\n", name, date, month, year)

	

}
