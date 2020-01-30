package main

import (
	"fmt"
	"os"
	"time"

	flag "github.com/ogier/pflag"
)

var (
	workingTime int
	commuting   int
)

func main() {
	flag.Parse()

	if flag.NFlag() == 0 {
		fmt.Printf("Usage: %s [options]\n", os.Args[0])
		fmt.Println("Options:")
		flag.PrintDefaults()
		os.Exit(1)
	}

	totalTime := time.Now().Local().Add(time.Hour*time.Duration(workingTime) +
		time.Minute*time.Duration(commuting))

	fmt.Printf("Working time: %v hours \n", workingTime)
	fmt.Printf("Commuting time: %v minutes \n", commuting)
	fmt.Printf("Arriving home at %v \n", totalTime)

	for range time.Tick(1 * time.Second) {
		remainingTime := getRemainingTime(totalTime)

		if remainingTime.t <= 0 {
			fmt.Println("You should be at home")
		}

		if remainingTime.t == (commuting * 60) {
			fmt.Println("You should be leaving work")
		}

		fmt.Printf("Hours: %02d Minutes: %02d Seconds: %02d \n", remainingTime.h, remainingTime.m, remainingTime.s)
	}

}

func init() {
	flag.IntVarP(&workingTime, "working_time", "w", 0, "Working Time in hours")
	flag.IntVarP(&commuting, "commuting", "c", 0, "Commuting Time in minutes")
}

type countdown struct {
	t int
	h int
	m int
	s int
}

func getRemainingTime(t time.Time) countdown {
	currentTime := time.Now().Local()
	difference := t.Sub(currentTime)

	total := int(difference.Seconds())
	hours := int(total / (60 * 60) % 24)
	minutes := int(total/60) % 60
	seconds := int(total % 60)

	return countdown{
		t: total,
		h: hours,
		m: minutes,
		s: seconds,
	}
}
