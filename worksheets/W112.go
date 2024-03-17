package worksheets

import (
	"fmt"
	"os"
)

func W112() {
	data, err := os.ReadFile("./Readable_Files/W112.txt")
	var convert string
	check(err)
	for _, val := range data {
		if val == '%' {
			convert += " "
		} else if val == '&' {
			convert += "@"
		} else {
			convert += "\n"
		}
	}
	fmt.Println(convert)
}
