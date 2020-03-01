package main

import "fmt"

func calc1(num1,num2 int) (int,int) {
	var add = num1+num2
	var subt = num1-num2

	return add,subt

}
func main() {
	var res1, res2 = calc1(4,3)

	fmt.Println(res1 , res2)
}