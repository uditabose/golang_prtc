package main

import "fmt"

// returns season name for the month
func Season(month int) string {
	season := ""

	switch month {
	case 1, 2, 12:
		season = "Winter"
	case 3, 4, 5:
		season = "Spring"
	case 6, 7, 8:
		season = "Summer"
	case 9, 10, 11:
		season = "Autumn"
	default:
		season = "Season unknown"
	}

	return season
}

func sumInts(list ...int) (sum int) {
	for _, val := range list {
		sum += val
	}
	return
}

func main() {
	// fmt.Println(Season(13))
	// fmt.Println(Season(3))
	fmt.Println(sumInts(5, -2, 0, 9))

}
