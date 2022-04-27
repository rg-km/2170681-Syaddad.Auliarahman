package main

import "fmt"

func main() {
	/*
		Convert the given second to 00:00:00 hour minute second format

		Example Input/Output
		30 -> 00:00:30
		70 -> 00:01:10
		67812 -> 18:50:12
		678120 -> 188:22:00

	*/
	// res := ConvertSecondToTimeString(30)
	// fmt.Println(res)

	// Try correct answer:
	resCorrect := ConvertSecondToTimeStringCorrect(678120)
	fmt.Println(resCorrect)
}

func ConvertSecondToTimeString(second int) string {
	hours := second / 3600
	minutes := second / 60

	timeString := fmt.Sprintf("%d:%d:%d", hours, minutes, second)
	return timeString
}

func ConvertSecondToTimeStringCorrect(second int) string {
	var (
		hours   int
		minutes int
	)
	if second >= 3600 {
		hours = second / 3600
	} else if second >= 60 {
		minutes = second / 60
	}

	if hours >= 1 {
		second -= (3600 * hours)
		if second >= 60 {
			minutes = second / 60
			if minutes >= 1 {
				second -= (60 * minutes)
			}
		}

		timeString := fmt.Sprintf("%02d:%02d:%02d", hours, minutes, second)
		return timeString
	}
	if minutes >= 1 {
		second -= (60 * minutes)
	}

	timeString := fmt.Sprintf("%02d:%02d:%02d", hours, minutes, second)
	return timeString
	//return "" // TODO: replace this
}
