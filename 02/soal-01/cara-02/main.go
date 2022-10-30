package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Cara ke-2 : tanpa menggunakan function time.Add()")

	fmt.Println("Input Jam : ")
	var hour int
	fmt.Scanln(&hour)

	fmt.Println("Input Menit : ")
	var minute int
	fmt.Scanln(&minute)

	fmt.Println("Input Detik : ")
	var second int
	fmt.Scanln(&second)

	var endSecond, endMinute, endHour int = 0, 0, 0

	fmt.Println("Lama perjalanan dalam detik : ")
	var durationSecond int
	fmt.Scanln(&durationSecond)

	endSecond = second + durationSecond

	if endSecond >= 60 {
		endMinute = endSecond / 60
		endSecond = endSecond % 60
	}

	endMinute = endMinute + minute

	if endMinute >= 60 {
		endHour = endMinute / 60
		endMinute = endMinute % 60
	}

	endHour = endHour + hour

	var endTime = time.Date(0, 0, 0, endHour, endMinute, endSecond, 0, time.Local)
	fmt.Println("Waktu sampai : " + endTime.Format("15:04:05"))

}
