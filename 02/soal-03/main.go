package main

import (
	"fmt"
	"time"
)

func main() {
	var startTimeAli, _ = time.Parse("15:04:05", "08:00:00")
	var startTimeBadu, _ = time.Parse("15:04:05", "08:01:40")

	var fullDistance = 900
	var speedAli = 3
	var speedBadu = 2

	var firstDistance = int(startTimeBadu.Sub(startTimeAli).Seconds()) * speedAli
	fmt.Println("Jarak Ali jam 08:01:40 = ", firstDistance, " meter, dari titik awal Ali")

	var distanceAli = (fullDistance - firstDistance) * speedAli / (speedAli + speedBadu)

	fmt.Println("Jarak Ali dan Badu ketika bertemu = ", distanceAli+firstDistance, " meter, dari titik awal Ali")

	var timeAli = time.Duration(distanceAli/speedAli) * time.Second

	var timeMeet = startTimeBadu.Add(timeAli)
	fmt.Println("Waktu Ali dan Badu bertemu : ", timeMeet.Format("15:04:05"))
}
