package main

import (
	"fmt"
)

func main() {
	// mySlice := []int{10, 20, 30, 40, 50}
	// for _, value := range mySlice {
	// 	value *= 2
	// }
	// fmt.Printf("mySlice %+v\n", mySlice)
	// for index, _ := range mySlice {
	// 	mySlice[index] *= 2
	// }
	// fmt.Printf("mySlice %+v\n", mySlice)
	str := [5]string{"i", "am", "stupid", "and", "weak"}
	for index, value := range str {
		fmt.Println(index, value)
	}
	// 修改数据内容为：{"i", "am", "smart", "and", "strong"}
	str[2] = "smart"
	str[4] = "strong"
	fmt.Printf("str %+v\n", str)
}
