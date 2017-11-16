package main

import "fmt"

func main() {

	var str1, str2 string
	//var i int
	str1 = "hello world"
	ch := str1[0]

	str2 = "1234567"
	str_len := len(str1)
	fmt.Printf("The string is %s\n", str1)
	fmt.Printf("The string len is %d\n", str_len)
	fmt.Printf("The first char is %c\n", ch)
	fmt.Printf("str1 + str2 = %s\n", str1+str2)
	fmt.Printf("str1 [3] = %c\n", str1[3])

	//for i = 0; i < str_len; i++ {
	//ch1 := str1[i]
	//fmt.Println(i, ch1)
	//}

	// i ch1 is rune type
	for i, ch1 := range str1 {
		fmt.Println(i, ch1)
	}
}
