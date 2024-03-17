package worksheets

import "fmt"

var Classes = []string{"Comp Sci", "Maths", "Physics"}
var Students = [][]string{
	{"Josh", "Joe", "Adam", "Will"},
	{"Josh", "Joe", "Tommy", "Tom"},
	{"Josh", "Lewis", "Hamish", "Guy"}}

func W109() {
	var Ans string
	var AltAns string
	fmt.Println(`How would you like to manipulate the database?
	1: Add Class
	2: Remove Class
	3: Add Student
	4: Remove Student
	5: Print Class
	6: Print College`)
	fmt.Scan(&Ans)
	switch Ans {
	case "1":
		fmt.Println("Class Name?")
		fmt.Scan(&Ans)
		addClass(Ans)
	case "2":
		fmt.Println("Class Name?")
		fmt.Scan(&Ans)
		removeClass(Ans)
	case "3":
		fmt.Println("Student Name?")
		fmt.Scan(&Ans)
		fmt.Println("Class Name?")
		fmt.Scan(&AltAns)
		addStudentToClass(Ans, AltAns)
	case "4":
		fmt.Println("Student Name?")
		fmt.Scan(&Ans)
		fmt.Println("Class Name?")
		fmt.Scan(&AltAns)
		removeStudentFromClass(Ans, AltAns)
	case "5":
		fmt.Println("Class to print?")
		fmt.Scan(&Ans)
		printClassList(Ans)
	case "6":
		printWholeCollege()
	}
}

func removeClass(Class string) {
	for i, val := range Classes {
		if val == Class {
			Classes = append(Classes[:i], Classes[i+1:]...)
		}
	}
}

func addClass(className string) {
	Classes = append(Classes, className)
	Register := []string{}
	Students = append(Students, Register)
}

func addStudentToClass(studentName string, className string) {
	for i, val := range Classes {
		if val == className {
			Students[i] = append(Students[i], studentName)
		}
	}
}

func removeStudentFromClass(studentName string, className string) {
	for i, val := range Classes {
		if val == className {
			for j, val := range Students[i] {
				if val == studentName {
					Students = append(Students[:j], Students[j+1:]...)
				}
			}
		}
	}
}

func printClassList(className string) {
	for i, val := range Classes {
		if val == className {
			fmt.Println(Students[i])
		}
	}
}

func printWholeCollege() {
	for i, val := range Classes {
		fmt.Printf("the Class %v has %v in it ", val, Students[i])
	}
}
