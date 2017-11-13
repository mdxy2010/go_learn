package main

import "fmt"

func getname() (firstname, midname, lastname string) {
	return "xing", "yu", "yuxing"
}

func main() {
	var (
			firstname string
			midname string
			lastname string
		)

	//firstname, midname, lastname := getname();
	_, _, lastname := getname();

	fmt.Println(firstname);
	fmt.Println(midname);
	fmt.Println(lastname);
}
