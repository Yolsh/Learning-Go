package LearningInit

import (
	"fmt"
	"regexp"
)

func main() {
	check := "lol it worked?"
	fmt.Println(regexp.MatchString("^i", check))
}
