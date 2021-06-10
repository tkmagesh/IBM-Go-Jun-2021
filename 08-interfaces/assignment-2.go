package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(sum(10, 20))                                                 //=> 30
	fmt.Println(sum(10, 20, 30, 40))                                         //=> 100
	fmt.Println(sum(10))                                                     //=> 10
	fmt.Println(sum())                                                       //=> 0
	fmt.Println(sum(10, "20", 30, "40"))                                     //=> 100
	fmt.Println(sum(10, "20", 30, "40", "abc"))                              //=> 100
	fmt.Println(sum(10, 20, []int{30, 40}))                                  //=> 100
	fmt.Println(sum(10, 20, []interface{}{30, 40, []int{10, 20}}))           //=> 130
	fmt.Println(sum(10, 20, []interface{}{30, 40, []interface{}{10, "20"}})) //=> 130
}

func sum(nos ...interface{}) int {
	result := 0
	for _, val := range nos {
		switch val.(type) {
		case int:
			result += val.(int)
		case string:
			if x, err := strconv.Atoi(val.(string)); err != nil {
				result += x
			}
		case []int:
			values := val.([]int)
			intValList := make([]interface{}, len(values))
			for idx, x := range values {
				intValList[idx] = x
			}
			result += sum(intValList...)
		case []interface{}:
			result += sum(val.([]interface{})...)
		}
	}
	return result
}
