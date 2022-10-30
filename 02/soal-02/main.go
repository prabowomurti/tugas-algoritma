package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("WAKTU KEBERANGKATAN : ")
	fmt.Println("Input Jam : ")
	var departHour int
	fmt.Scanln(&departHour)

	fmt.Println("Input Menit : ")
	var departMinute int
	fmt.Scanln(&departMinute)

	fmt.Println("Input Detik : ")
	var departSecond int
	fmt.Scanln(&departSecond)

	departTime := time.Date(0, 0, 0, departHour, departMinute, departSecond, 0, time.UTC)

	fmt.Println("WAKTU TIBA : ")
	fmt.Println("Input Jam : ")
	var arrivalHour int
	fmt.Scanln(&arrivalHour)

	fmt.Println("Input Menit : ")
	var arrivalMinute int
	fmt.Scanln(&arrivalMinute)

	fmt.Println("Input Detik : ")
	var arrivalSecond int
	fmt.Scanln(&arrivalSecond)

	arrivalTime := time.Date(0, 0, 0, arrivalHour, arrivalMinute, arrivalSecond, 0, time.UTC)

	fmt.Println("Berangkat   : " + departTime.Format("15:04:05"))
	fmt.Println("Waktu tiba  : " + arrivalTime.Format("15:04:05"))

	var duration time.Duration
	duration = arrivalTime.Sub(departTime)

	fmt.Println("Durasi Perjalanan : " + humanizeDuration(duration))
}

func humanizeDuration(duration time.Duration) string {
	var durationHour, durationMinute, durationSecond int
	durationHour = int(duration.Hours())
	durationMinute = int(duration.Minutes()) % 60
	durationSecond = int(duration.Seconds()) % 60

	return fmt.Sprintf("%d jam, %d menit, %d detik", durationHour, durationMinute, durationSecond)

}
