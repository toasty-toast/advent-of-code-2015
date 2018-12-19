package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/toasty-toast/advent-of-code-2015/utils"
)

type present struct {
	length, width, height int
}

func (p *present) SurfaceArea() int {
	return (2 * p.length * p.width) + (2 * p.length * p.height) + (2 * p.width * p.height)
}

func (p *present) SmallestSideArea() int {
	side1 := p.length * p.width
	side2 := p.length * p.height
	side3 := p.width * p.height
	return utils.MinInt([]int{side1, side2, side3})
}

func (p *present) Volume() int {
	return p.length * p.width * p.height
}

func (p *present) SmallestPerimiter() int {
	perim1 := 2*p.length + 2*p.width
	perim2 := 2*p.length + 2*p.height
	perim3 := 2*p.width + 2*p.height
	return utils.MinInt([]int{perim1, perim2, perim3})
}

func loadPresents(filename string) []*present {
	lines := utils.ReadLines(filename)
	presents := make([]*present, 0)
	for i := range lines {
		next := new(present)
		split := strings.Split(lines[i], "x")
		next.length, _ = strconv.Atoi(split[0])
		next.width, _ = strconv.Atoi(split[1])
		next.height, _ = strconv.Atoi(split[2])
		presents = append(presents, next)
	}
	return presents
}

func totalWrappingPaper(presents []*present) int {
	sum := 0
	for _, cur := range presents {
		sum += cur.SurfaceArea() + cur.SmallestSideArea()
	}
	return sum
}

func totalRibbonLength(presents []*present) int {
	sum := 0
	for _, cur := range presents {
		sum += cur.Volume() + cur.SmallestPerimiter()
	}
	return sum
}

func main() {
	presents := loadPresents("input.txt")
	fmt.Printf("Total wrapping paper: %d\n", totalWrappingPaper(presents))
	fmt.Printf("Total ribbon: %d\n", totalRibbonLength(presents))
}
