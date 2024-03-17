package LearningInit

import (
	"fmt"
	"regexp"
)

func Regex1() {
	check := "lol it worked?"
	fmt.Println(regexp.MatchString("^i", check))
}
