package main

import "fmt"

type Player struct {
	fname string
	lname string
	plateAppearance int
	atBat int
	sinlge int
	double int
	triple int
	homerun int
	walks int
	hitByPitch int
}

func BattingAVG(single, dobut, triple, homerun int) int {
	x = (single + double + triple + homerun)/atBat
	return
}
func Slugging(single, double, triple, homerun, atBat int) int {
	x = (single + (double*2) + (triple*3) + (homerun*4))/atBat
	return
}

func OnBaseP(single, double, triple, homerun, hitByPitch, walks, atBat int) int{
	x = (single + double + triple + homerun + hitByPitch + walks)/atBat
}