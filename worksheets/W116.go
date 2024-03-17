package worksheets

import (
	"errors"
	"fmt"
	"strconv"
)

func W116() {
	var num1, num2 string
	fmt.Println("gimme two numbers")
	fmt.Scan(&num1)
	fmt.Scan(&num2)
	ans, err := rectangle(num1, num2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", ans)
}

func rectangle(num1 string, num2 string) (int, error) {
	val1, err2 := strconv.Atoi(num1)
	val2, err := strconv.Atoi(num2)
	if err != nil || err2 != nil {
		return -1, errors.New("input was not a number")
	} else {
		return val1 * val2, nil
	}
}
