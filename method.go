package main

import (
	"fmt"
)

// Define the Student struct
type Student struct {
	Firstname string
	Lastname  string
}

// วิธีประกาศ Method
// (s Student) คือ Receiver หรือข้อมูลที่เราต้องการ Support Method นี้
func (s Student) FullName() string {
	return s.Firstname + " " + s.Lastname
}

func main() {
	student := Student{
		Firstname: "Mike",
		Lastname:  "Lopster",
	}

	// เรียกใช้ Method ผ่าน Instance ของ Student ที่เราสร้างขึ้นมาใหม่ด้านบน
	fullName := student.FullName()
	fmt.Println("Full Name of the student:", fullName)
}
