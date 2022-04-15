package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"time"
)

type countdown struct {
	t int
	h int
	m int
	s int
}

func main() {
	var round string
	roundCount := 0
	var patternTime = regexp.MustCompile(`^[0-5]?[0-9](h|H)[0-5]?[0-9](M|m)$`)

	focusTime := flag.String("focusTime", "0h50m", "The focus time, where you can start a music and focus in your activitie (syntax 0h0m)")
	intervalTime := flag.String("intervalTime", "0h10m", "The interval time, where you can relax before start the focus time again (syntax 0h0m)")
	repetition := flag.Int("repetition", 1, "How much focus time you want")
	flag.Parse()

	if *focusTime == "" || *intervalTime == "" {
		fmt.Println("[ERROR] - Missing some parameters!")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if !patternTime.MatchString(*focusTime) || !patternTime.MatchString(*intervalTime) {
		fmt.Println("[ERROR] - Check the syntax of the time [eg: 02h50m,00h50m,0H50M,2h50m]")
		os.Exit(1)
	}

	parseFocusTime, _ := time.ParseDuration(*focusTime)
	parseIntervalTime, _ := time.ParseDuration(*intervalTime)

	turn := map[int]time.Duration{
		0: parseFocusTime,
		1: parseIntervalTime,
	}

	totalFocusTime := parseFocusTime.Minutes() * float64(*repetition)

	for i := 0; i <= *repetition; i++ {

		duration := turn[i%2]
		futureTime := time.Now().Add(time.Minute * time.Duration(duration.Minutes()))

		if i%2 == 0 {
			round = "Focus"
			roundCount++
		} else {
			round = "Interval"
		}

		for range time.Tick(1 * time.Second) {
			timeRemaining := getTimeRemaining(futureTime)

			if timeRemaining.t <= 0 {
				fmt.Println("Time out!")
				go bell()
				time.Sleep(time.Second * 3)
				break
			}
			cmd := exec.Command("clear")
			cmd.Stdout = os.Stdout
			cmd.Run()
			fmt.Printf("Total rouds: %d - Current roud: %d\n", *repetition, roundCount)
			fmt.Printf("Focus time: [ %s ] - Interval time: [ %s ] \nTotal hours of focus [ %dh:%dm ]\n", *focusTime, *intervalTime, int(totalFocusTime)/60%24, int(totalFocusTime)%60)
			fmt.Printf("\n[%s] TIMER -- %dh:%dm:%ds --\n", round, timeRemaining.h, timeRemaining.m, timeRemaining.s)
		}
		if *repetition == 1 {
			break
		}
	}
}
