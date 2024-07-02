package main

//import "fmt"

var player Player = Player{
	x:     3,
	y:     5,
	score: 0,
}

var coin Coin = Coin{
	x: -1,
	y: -1,
}

var field Field = Field{
	matrix: [9][7]string{
		{"---", "--", "--", "---", "--", "--", "---"},
		{"|", "   ", "   ", "   ", "   ", "   ", "|"},
		{"|", "   ", "   ", "   ", "   ", "   ", "|"},
		{"|", "   ", "   ", "   ", "   ", "   ", "|"},
		{"|", "   ", "   ", "   ", "   ", "   ", "|"},
		{"|", "   ", "   ", "   ", "   ", "   ", "|"},
		{"|", "   ", "   ", "   ", "   ", "   ", "|"},
		{"|", "   ", "   ", "   ", "   ", "   ", "|"},
		{"---", "--", "--", "---", "--", "--", "---"},
	},
}

func main() {
	PositionPlayer(&field, player)
	GenerateCoin(&coin, &field, player)
	PrintMap(field)
}
