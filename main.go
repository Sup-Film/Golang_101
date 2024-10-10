package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/sup.dongkew/go-example/func"
)

func main() {
	id := uuid.New()
	fmt.Printf("UUID: %s", id)
	funchello.SayHello()
	funchello.SayTest()
	var101()
	ControlStructure()
	dataStructure()
}

func var101() {
	// ประกาศตัวแปรแบบกำหนด Type
	var fullname string = "Supakorn dongkew"
	var age int = 25
	fmt.Printf("\nแบบกำหนด Type = Name: %s Age: %d", fullname, age)

	// ประกาศตัวแปรแบบไม่กำหนด Type
	var fullname2 = "Supakorn dongkew"
	var age2 = 25
	fmt.Printf("\nแบบไม่กำหนด Type = Name: %s Age: %d", fullname2, age2)

	// ประกาศตัวแปรแบบไม่กำหนด Type และใช้ :=
	fullname3 := "Supakorn dongkew"
	age3 := 25
	fmt.Printf("\nแบบไม่กำหนด Type และใช้ := = Name: %s Age: %d", fullname3, age3)

	// ค่าคงที่
	const pi = 3.14
	fmt.Printf("\nค่าคงที่ = %f", pi)
}

func ControlStructure() {
	fmt.Println("\n\nControl Structure")
	// Sequential ให้ทำงานตามลำดับจากบนลงล่าง

	// Condition
	// if else
	score := 62
	grade := ""

	if score >= 80 {
		grade = "A"
	} else if score >= 70 {
		grade = "B"
	} else if score >= 60 {
		grade = "C"
	} else if score >= 50 {
		grade = "D"
	} else {
		grade = "F"
	}

	fmt.Printf("Score: %d Grade: %s", score, grade)

	// switch case
	dayOfWeek := 7

	switch dayOfWeek {
	case 1:
		fmt.Println("\nSunday")
	case 2:
		fmt.Println("\nMonday")
	case 3:
		fmt.Println("\nTuesday")
	case 4:
		fmt.Println("\nWednesday")
	case 5:
		fmt.Println("\nThursday")
	case 6:
		fmt.Println("\nFriday")
	case 7:
		fmt.Println("\nSaturday")
	default:
		fmt.Println("\nInvalid day")
	}

	// Preprocess if else
	num1 := 5
	num2 := 10

	sumNum := num1 + num2

	if sumNum >= 10 {
		fmt.Println("sumNum more than 10")
	}

	// เหมือนกับด้านบนแต่จะมีการจัดการ Memory ให้ดีขึ้น
	num3 := 5
	num4 := 10

	if sumNum := num3 + num4; sumNum >= 10 {
		fmt.Println("sumNum more than 10")
	}

	// Loop
	// for loop
	fmt.Print("For Loop: ")
	for i := 0; i < 10; i++ {
		fmt.Printf("%d, ", i)
	}

	fmt.Print("\n")
	// Do while loop
	j := 0
	fmt.Print("Do while Loop: ")
	for {
		fmt.Printf("%d, ", j)
		j++
		if j >= 10 {
			break
		}
	}

	// while loop
	k := 0
	fmt.Print("\nWhile Loop: ")
	for k < 10 {
		fmt.Printf("%d, ", k)
		k++
	}
}

func dataStructure() {
	// Data Structure
	// Array
	fmt.Println("\n\nData Structure")
	fmt.Print("Array: ")
	var arr [5]int // Array ต้องกำหนดขนาดของ Array ไว้ล่วงหน้า
	// Assign ค่าให้กับ Array
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	arr[4] = 5
	fmt.Println(arr)

	// Slice
	fmt.Print("Slice: ")
	var slice []int // Slice เหมือนกับ Array แต่ไม่ต้องกำหนดขนาดของ Slice ไว้ล่วงหน้า
	// Assign ค่าให้กับ Slice
	slice = append(slice, 1, 2, 3, 4, 5)
	fmt.Println(slice)

	// Map
	fmt.Print("Map: ")
	var m map[string]int // Map เหมือนกับ Slice แต่เก็บข้อมูลแบบ Key Value
	fmt.Println(m)

	// เช็คว่ามีค่าใน Map หรือไม่
	value, ok := m["key"]
	if ok {
		fmt.Println(value)
	} else {
		fmt.Println("Not found")
	}

	// Struct
	fmt.Print("Struct: ")
	type Person struct {
		Firstname string
		Lastname  string
		Age       int
	} // Struct ใช้สำหรับเก็บข้อมูลที่มีความซับซ้อน

	student1 := Person{
		Firstname: "Supakorn",
		Lastname:  "Dongkew",
		Age:       25,
	}
	fmt.Println(student1)

	var student2 Person
	student2.Firstname = "Supakorn"
	student2.Lastname = "Dongkew"
	student2.Age = 25
	fmt.Println(student2)

	// Array of Struct
	fmt.Print("Array of Struct: ")
	type Student struct {
		Firstname string
		Lastname  string
		Age       int
	}

	students := [3]Student{
		{Firstname: "Supakorn", Lastname: "Dongkew", Age: 25},
		{Firstname: "John", Lastname: "Doe", Age: 30},
		{Firstname: "Jane", Lastname: "Doe", Age: 25},
	}
	fmt.Println(students)

	// Add Struct into Map
	fmt.Print("Add Struct into Map: ")
	studentMap := make(map[string]Person, 0)	// ประกาศ Map
	studentMap["student1"] = Person{Firstname: "Supakorn", Lastname: "Dongkew", Age: 25} // กำหนด Key และ Assignt ค่าให้กับ Map ด้วย Struct
	studentMap["student2"] = Person{Firstname: "John", Lastname: "Doe", Age: 30}
	fmt.Println(studentMap)

	// Function
	fmt.Print("Function: ")
	// ประกาศ Function
	var add = func(a int, b int) int {
		return a + b
	} // Function เป็น Data Type ที่เก็บ Function ไว้ในตัวแปรได้

	// การเรียกใช้ Function
	result := add(5, 10)
	fmt.Println(result)

	// Pointer
	fmt.Print("Pointer: ")
	// ประกาศ Pointer
	var p *int
	i := 42         // ประกาศตัวแปร i และกำหนดค่าให้เท่ากับ 42
	p = &i          // การใช้งาน Pointer ใช้ & นำหน้าตัวแปรที่ต้องการเก็บค่า Address ของตัวแปรนั้น
	fmt.Println(*p) // การใช้งาน Pointer ใช้ * นำหน้าตัวแปรที่เป็น Pointer
}
