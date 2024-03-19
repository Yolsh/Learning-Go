package worksheets

import (
	"errors"
)

var stackArray [20]int = [20]int{}
var pointer int = 0

func A122() {
	err := push(22)
	if err != nil {
		panic(err)
	}
}

func push(add int) error {
	if stackArray[19] != 0 {
		return errors.New("stack is full")
	} else {
		pointer++
		stackArray[pointer] = add
		return nil
	}
}
