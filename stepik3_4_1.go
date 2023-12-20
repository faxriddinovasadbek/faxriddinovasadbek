package main

import (
	"encoding/json" 
	"fmt"           
	"os"            
)

func readTask() (interface{}, interface{}, interface{}){
	return 1.3,2.3,"whjgfhsv"
}


func My_print(val  interface{}){
	fmt.Printf("value=%v: %T", val, val)
}

func main() {
	value1, value2, operation := readTask() 

	v1, ok := value1.(float64)
	if !ok {
		My_print(!ok)
		return
	}

	v2, ok := value2.(float64)
	if !ok {
		My_print(!ok)
		return
	}

	if v, ok := operation.(string); ok {
		switch v {
		case "+":
			fmt.Printf("%.4f\n", v1 + v2)
		case "-":
			fmt.Printf("%.4f\n", v1 - v2)
		case "/":
			fmt.Printf("%.4f\n", v1 / v2)
		case "*":
			fmt.Printf("%.4f\n", v1 * v2)
		default:
			fmt.Println("неизвестная операция")
		}
	}else{
		fmt.Println("неизвестная операция")
		return
	}
 	
}
