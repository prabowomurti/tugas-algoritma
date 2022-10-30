package main

import (
	"fmt"
	"time"
)

func main() {
	var distance = 1000.0
	var speedAli float64
	var distanceAli = 0.0
	var speedBadu float64
	var distanceBadu = 0.0
	var duration = 0

	// calculate distanceBadu first, because he runs at 5 m/s but 10 seconds earlier
	// than Ali
	speedBadu = 5.0
	for i := 1; i <= 10; i++ {
		distanceBadu = distanceBadu + speedBadu
	}

	// for the first 10 seconds, Ali's speed is constant at 2 m/s
	// from here, we start to increase the duration (remember: Badu starts 10 seconds earlier)
	speedAli = 2
	for i := 1; i <= 10; i++ {
		distanceAli = distanceAli + speedAli
		distanceBadu = distanceBadu + speedBadu
		duration++
	}

	// for the next 10 seconds, Ali's speed is constant at 2.5 m/s
	speedAli = 2.5
	for i := 11; i <= 20; i++ {
		distanceAli = distanceAli + speedAli
		distanceBadu = distanceBadu + speedBadu
		duration++
	}

	// for the next 10 seconds, Ali's speed is constant at 3 m/s
	speedAli = 3
	for i := 21; i <= 30; i++ {
		distanceAli = distanceAli + speedAli
		distanceBadu = distanceBadu + speedBadu
		duration++
	}

	// keep going until Ali reaches Badu
	for distanceAli+distanceBadu < distance {
		// increase speedAli for 0.5 every 10 seconds
		if duration%10 == 0 {
			speedAli = speedAli + 0.5
			fmt.Println("Kecepatan Ali bertambah menjadi ", speedAli, "m/s pada detik ke", duration)
		}

		duration++

		fmt.Println()
		fmt.Println("Detik ke ", duration)
		distanceAli = distanceAli + speedAli
		fmt.Println("Jarak Ali", distanceAli, "m")
		distanceBadu = distanceBadu + speedBadu
		fmt.Println("Jarak Badu", distanceBadu, "m")

		// if Ali reaches Badu, get out of the loop and duration is saved

		// sleep 100 ms
		time.Sleep(10 * time.Millisecond)
	}

	// print duration
	fmt.Println("Waktu yang diperlukan Ali untuk bertemu Badu adalah", duration, "detik, sejak pukul 08:00:00")
	// print distanceAli
	fmt.Println("Jarak Ali dari titik A saat bertemu Badu", distanceAli, "m")
}
