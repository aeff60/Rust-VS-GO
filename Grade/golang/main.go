package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("โปรแกรมคำนวนเกรด")

	var score string
	fmt.Println("กรุณากรอกคะแนนของคุณ: ")
	fmt.Scanln(&score)

	scoreFloat, err := strconv.ParseFloat(score, 32)

	if err != nil {
		fmt.Println("กรุณากรอกตัวเลข")
		return
	}

	var grade string

	switch {
	case scoreFloat >= 80:
		grade = "A"
	case scoreFloat >= 75:
		grade = "B+"
	case scoreFloat >= 70:
		grade = "B"
	case scoreFloat >= 65:
		grade = "C+"
	case scoreFloat >= 60:
		grade = "C"
	case scoreFloat >= 55:
		grade = "D+"
	case scoreFloat >= 50:
		grade = "D"
	default:
		grade = "F"
	}

	fmt.Printf("คะแนนของคุณคือ %.2f และได้รับเกรด %s\n", scoreFloat, grade)
}
