package worksheets

import (
	"fmt"
	"math/rand"
)

func W108() {
	var names = [5]string{"MJ", "Jeff", "James", "Steven Hawking", "Will Smith"}
	var clues1 = [5]string{"basketball", "minion", "orange", "black hole", "iRobot"}
	var clues2 = [5]string{"shoes", "starr", "piggy nose", "wheelchair", "Rewind"}
	points := 0

	for n := 0; n < len(clues1); n++ {
		i := rand.Intn(len(clues1))
		var val string = clues1[i]
		fmt.Println(val)
		var Ans string
		fmt.Scan(&Ans)
		if Ans == names[i] {
			points += 5
		} else {
			fmt.Println(clues2[i])
			fmt.Scan(&Ans)
			if Ans == names[i] {
				points += 2
			} else {
				fmt.Println("wrong the person was", names[i])
			}
		}
	}
}
