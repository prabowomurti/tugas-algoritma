package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Input Jam : ")
	var hour int
	fmt.Scanln(&hour)

	fmt.Println("Input Menit : ")
	var minute int
	fmt.Scanln(&minute)

	fmt.Println("Input Detik : ")
	var second int
	fmt.Scanln(&second)

	var now = time.Now()

	var startTime = time.Date(now.Year(), now.Month(), now.Day(), hour, minute, second, 0, time.Local)

	fmt.Println("Waktu mulai  : " + startTime.Format("15:04:05"))
	fmt.Println("Lama perjalanan dalam detik : ")
	var duration int
	fmt.Scanln(&duration)

	var endTime = startTime.Add(time.Duration(duration) * time.Second)
	fmt.Println("Waktu sampai : " + endTime.Format("15:04:05"))
}
