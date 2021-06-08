package main

import (
	"fmt"
	//"math"
)

/* 
notes:
first day of school = august 26 (exercises did not start till the next school week though)
september 7th is off
november 11th is off
november 23-27 (inclusive) are off
december 21-january 1 (inclusive) are off
january 18 is off
february 15 is off
april 2 - 9 (inclusive) are off
may 31 is off
last day of school = june 9

7 special weeks
*/

var (
	first_plank_time int = 30 // in seconds
	first_pushup_count int  = 5 // in count
	max_plank_time int = 240 // max / final
	max_pushup_count int = 25
	final_pushup_count int = 10
	weeks int = 42
	//total_plank_time int = 0
)

func secondsToMinutes(inSeconds int) {
	minutes := inSeconds / 60
	seconds := inSeconds % 60
	str := fmt.Sprintf("%d minutes and %d seconds", minutes, seconds)
	fmt.Println(str)
}
func CalcPlankTime() {
	normal_weeks := weeks-7
	daily_time_count := first_plank_time
	time_added_per_week := (max_plank_time - first_plank_time) / weeks
	total_plank_time := 0
	for i := 1; i <= normal_weeks; i++{
		if i == 1 {
			total_plank_time += (daily_time_count * 3)
			
		} else if i == 3 || i == 12 || i == 22 || i == 26 {
			total_plank_time += (2 * daily_time_count)
			
		} else {
			total_plank_time += (3 * daily_time_count)
			
		}
		daily_time_count += time_added_per_week
	}
	
	secondsToMinutes(total_plank_time)
}
func CalcPushupAmount() {
	normal_weeks := weeks-7
	weekly_pushup_count := first_pushup_count
	total_pushup_count := 0
	for i := 1; i <= normal_weeks; i++ {
		if i <= 20 {
			
			if i == 3 || i == 12 {
				total_pushup_count += (weekly_pushup_count * 2)
			} else {
				total_pushup_count += (weekly_pushup_count * 3)
			}
			weekly_pushup_count++
		} else {
			if i == 21{
				weekly_pushup_count--
			} else if i == 22 || i == 26 {
				total_pushup_count += (weekly_pushup_count * 2)
			} else {
				total_pushup_count += (weekly_pushup_count * 3)
			}
			weekly_pushup_count--
		}
	}
	fmt.Println("Total pushup count: " + fmt.Sprint(total_pushup_count))

}
func main() {
	CalcPlankTime()
	CalcPushupAmount()
}