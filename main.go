package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"time"
)

func main() {
	var patternTime = regexp.MustCompile(`^[0-5]?[0-9](h|H)[0-5]?[0-9](M|m)$`)

	focusTime := flag.String("focusTime", "0h50m", "The focus time, where you can start a music and focus on your activity [syntax 0h0m]")
	intervalTime := flag.String("intervalTime", "0h10m", "The interval time, where you can relax before starting the focus time again [syntax 0h0m]")
	repetition := flag.Int("repetition", 1, "The amount of a repetition")
	flag.Parse()

	// if *focusTime == "" || *intervalTime == "" {
	// 	fmt.Println("[ERROR] - Missing some parameters!")
	// 	flag.PrintDefaults()
	// 	os.Exit(1)
	// }

	if !patternTime.MatchString(*focusTime) || !patternTime.MatchString(*intervalTime) {
		fmt.Println("[ERROR] - Check the syntax of the time string [eg: 02h50m,00h50m,0H50M,2h50m]")
		os.Exit(1)
	}

	parseFocusTime, _ := time.ParseDuration(*focusTime)
	parseIntervalTime, _ := time.ParseDuration(*intervalTime)

	timeType := map[string]time.Duration{
		"focus":    parseFocusTime,
		"interval": parseIntervalTime,
	}

	pomodoro := tumaty{
		focusTime:      *focusTime,
		intervalTime:   *intervalTime,
		repetition:     *repetition,
		roundCount:     1,
		totalFocusTime: parseFocusTime.Minutes() * float64(*repetition),
	}

	for round := 1; round <= pomodoro.repetition; round++ {

		targetTime := time.Now().Add(time.Minute * time.Duration(timeType["focus"].Minutes()))

		pomodoro.roundName = "Focus"
		pomodoro.roundCount = round

		// focus time
		timerDown(targetTime, pomodoro)
		if round-1 == pomodoro.repetition-1 {
			break
		}

		// interval time
		targetTime = time.Now().Add(time.Minute * time.Duration(timeType["interval"].Minutes()))
		pomodoro.roundName = "Interval"
		timerDown(targetTime, pomodoro)

	}
}

func timerDown(t time.Time, tu tumaty) {
	for range time.Tick(1 * time.Second) {
		timeRemaining := getTimeRemaining(t)

		if timeRemaining.t <= 0 {
			fmt.Println("Time's up! ;D ")
			go bell()
			time.Sleep(time.Second * 3)
			break
		}
		tu.screen()
		fmt.Printf("TIMER -- %dh:%dm:%ds --\n", timeRemaining.h, timeRemaining.m, timeRemaining.s)
	}
}

func (tu tumaty) screen() {

	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	fmt.Printf("Focus:       [ %s ] \nInterval:    [ %s ] \nTotal focus: [ %dh:%dm ]\n", tu.focusTime, tu.intervalTime, int(tu.totalFocusTime)/60%24, int(tu.totalFocusTime)%60)
	fmt.Printf("Round:       [  %d/%d  ]\n", tu.roundCount, tu.repetition)
	fmt.Printf("\n[%s] ", tu.roundName)
}
