package main

import (
	"fmt"
	"time"
)

func main() {
	input := "10:00 12.04.2023"
	layout := "15:04 02.01.2006"
	t, err := time.Parse(layout, input)
	if err != nil {
		fmt.Println(err)
		return
	}
	input2 := "11:17 01.02.2024"
	t2, err := time.Parse(layout, input2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("date()=date(%d,%d,%d)&&time()>=time(%d,%d)\n", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute())
	fmt.Printf("date()=date(%d,%d,%d)&&time()<=time(%d,%d)\n", t2.Year(), t2.Month(), t2.Day(), t2.Hour(), t2.Minute())
}
