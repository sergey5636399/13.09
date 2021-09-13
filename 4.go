package main

import "fmt"


func updarray(arr []int, l int) []int {

	var max_right int= arr[l-1]
	arr[l - 1] = -1
	
	for  i := l-2; i >= 0; i--{
	var tek int = arr[i]
	arr[i]= max_right
		if(max_right < tek){

			max_right = tek}
	}

return arr
}



func main() {

	
	
	mas :=[]int{ 17, 18, 5, 4, 6,1 }
	var b int = len(mas)
	fmt.Println(updarray(mas , b))
	
}

