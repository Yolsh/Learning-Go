package main

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
		AddClass(Ans)
	case "2":
		fmt.Println("Class Name?")
		fmt.Scan(&Ans)
		RemoveClass(Ans)
	case "3":
		fmt.Println("Student Name?")
		fmt.Scan(&Ans)
		fmt.Println("Class Name?")
		fmt.Scan(&AltAns)
		AddStudentToClass(Ans, AltAns)
	case "4":
		fmt.Println("Student Name?")
		fmt.Scan(&Ans)
		fmt.Println("Class Name?")
		fmt.Scan(&AltAns)
		RemoveStudentFromClass(Ans, AltAns)
	case "5":
		fmt.Println("Class to print?")
		fmt.Scan(&Ans)
		PrintClassList(Ans)
	case "6":
		PrintWholeCollege()
	}
}

func RemoveClass(Class string) {
	for i, val := range Classes {
		if val == Class {
			Classes = append(Classes[:i], Classes[i+1:]...)
		}
	}
}

func AddClass(className string) {
	Classes = append(Classes, className)
	Register := []string{}
	Students = append(Students, Register)
}

func AddStudentToClass(studentName string, className string) {
	for i, val := range Classes {
		if val == className {
			Students[i] = append(Students[i], studentName)
		}
	}
}

func RemoveStudentFromClass(studentName string, className string) {
	for i, val := range Classes {
		if val == className {
			for i := range Students[i] {
				Students = append(Students[:i], Students[i+1:]...)
			}
		}
	}
}

func PrintClassList(className string) {
	for i, val := range Classes {
		if val == className {
			fmt.Println(Students[i])
		}
	}
}

func PrintWholeCollege() {
	for i, val := range Classes {
		fmt.Printf("the Class %v has %v in it ", val, Students[i])
	}
}
