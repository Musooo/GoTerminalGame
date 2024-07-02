package main

type Player struct {
	x     int
	y     int
	score int
}//"@"

type Coin struct {
	x     int
	y     int
}//"0"

type Field struct{
	matrix [9][7]string
}