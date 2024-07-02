package main

import (
	"math/rand"
	"fmt"
)

func PrintMap(f Field) {
	fmt.Print("\033[H\033[2J")
	for _, row := range f.matrix {
        for _, cell := range row {
            fmt.Print(cell)
        }
        fmt.Println()
    }
}

func PositionPlayer(f *Field, p Player){
    f.matrix[p.y][p.x] = " @ "
}

func UpdatePlayerPos(p *Player, f *Field, x, y int){
    switch x{
        case 0:
            fallthrough
        case 6:
            break
        default:
            f.matrix[p.y][p.x] = "   "
            p.x = x
            PositionPlayer(f, *p)
    }
    switch y{
        case 0:
            fallthrough
        case 8:
            break
        default:
            f.matrix[p.y][p.x] = "   "
            p.y = y
            PositionPlayer(f, *p)
    }
}

func GenerateCoin(c *Coin, f *Field, p Player) {
    c.x = 1 + rand.Intn(5-1)
    c.y = 1 + rand.Intn(7-1)
    //check if the pos is the same as the player

    if(isCoinTouched(*c, p)){
        GenerateCoin(c, f, p)
    }else {
        f.matrix[c.y][c.x] = " 0 "
    }
}

func isCoinTouched(c Coin, p Player) bool{
    if(c.x == p.x && c.y == p.y){
        return true
    }
    return false
}