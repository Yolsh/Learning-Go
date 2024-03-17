package worksheets

import (
	"fmt"
)

//Doesn't work !!!

func calcTriNums(n int) []int {
	k, j := 1, 1
	output := make([]int, n)
	for i := 1; i <= n; i++ {
		output[i-1] = k
		j++
		k += j
	}
	fmt.Println(output)
	return output
}

func W103() {
	fmt.Println("how many rows for the pyramid?")
	var rows int
	fmt.Scan(&rows)
	var triangles []int = calcTriNums(rows)
	stars := "*"
	for i := 0; i < rows; i++ {
		for j := 0; i != len(triangles)-1 && j < triangles[len(triangles)-(i+1)]-triangles[len(triangles)-(i+2)]; j++ {
			stars = "  " + stars
		}
		for j := 1; j < triangles[i]; j++ {
			stars += "*"
		}
		fmt.Println(stars)
		stars = "*"
	}
}
