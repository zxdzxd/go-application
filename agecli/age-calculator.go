package agecli

import (
	"fmt"
	"strings"
	"time"
)

func GetAge(){
	fmt.Print("enter your Name : ")
	var name string
	fmt.Scanln(&name)
	name = strings.TrimSpace(name)
	fmt.Printf("hi %s !! Please enter your date of birth in below format\n", name)
	fmt.Printf("DD MM YYYY - ")
	var date, month, year int
	var err error
	n, err := fmt.Scanf("%d %d %d", &date, &month, &year)
	if n != 3 {
		err = fmt.Errorf("please enter dob in this format DD MM YYYY ")
	}
	if month > 12 || month < 1 {
		err = fmt.Errorf("please enter correct month %d ", month)
	}
	if year > 2024 || year < 1920 {
		err = fmt.Errorf("please enter correct year %d ", year)
	}
	if date > 31 || date < 1 {
		err = fmt.Errorf("please enter correct date %d ", date)
	}
	if err != nil {
		fmt.Printf("err - %s\n", err)
		return
	}

	fmt.Printf("hey %s your DOB is - %d/%d/%d\n", name, date, month, year)
	age := time.Now().Year() - year
	if age == 0 {
		age = int(time.Now().Month()) - month
		if age == 0 {
			age = time.Now().Day() - date
			fmt.Printf("you are %d day old\n", age)
			return
		}
		fmt.Printf("you are %d month old\n", age)
		return
	}
	fmt.Printf("you are %d years old\n", age)
}