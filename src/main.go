package main

import "github.com/eiannone/keyboard"

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

	field.matrix[player.y][player.x] = " @ "
	GenerateCoin(&coin, &field, player)
	PrintMap(field)

	keysEvents, err := keyboard.GetKeys(10)

	if err != nil {
		panic(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	for {
		event := <-keysEvents
		if event.Err != nil {
			panic(event.Err)
		}
		if event.Key == keyboard.KeyEsc {
			break
		}
		switch event.Key{
			case keyboard.KeyArrowDown:
				y := player.y+1
				UpdateYPlayerPos(&player,&field,y)
			case keyboard.KeyArrowRight:
				x := player.x+1
				UpdateXPlayerPos(&player,&field,x)
			case keyboard.KeyArrowLeft:
				x := player.x-1
				UpdateXPlayerPos(&player,&field,x)
			case keyboard.KeyArrowUp:
				y := player.y-1
				UpdateYPlayerPos(&player,&field,y)
		}

		field.matrix[player.y][player.x] = " @ "
		PrintMap(field)
	}
}
