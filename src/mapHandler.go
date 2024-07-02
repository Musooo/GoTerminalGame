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

func UpdateXPlayerPos(p *Player, f *Field, x int){
    switch x{
        case 0:
            fallthrough
        case 6:
            break
        default:
            f.matrix[p.y][p.x] = "   "
            p.x = x
    }
}

func UpdateYPlayerPos(p *Player, f *Field, y int){
    switch y{
    case 0:
        fallthrough
    case 8:
        break
    default:
        f.matrix[p.y][p.x] = "   "
        p.y = y
}
}

func GenerateCoin(c *Coin, f *Field, p Player) {
    c.x = 1 + rand.Intn(5-1)
    c.y = 1 + rand.Intn(7-1)

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