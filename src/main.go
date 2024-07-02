package main

import "fmt"


var player Player = Player{
	x: 3,
	y: 5,
	score: 0,
}

var field Field = Field{
	matrix: [9][7]string{
		{"-","-","-","-","-","-","-"},
		{"|"," "," "," "," "," ","|",},
		{"|"," "," "," "," "," ","|",},
		{"|"," "," "," "," "," ","|",},
		{"|"," "," "," "," "," ","|",},
		{"|"," "," "," "," "," ","|",},
		{"-","-","-","-","-","-","-"},
	},
}

func main(){
	fmt.Print(player)
	fmt.Print(field.matrix)
}