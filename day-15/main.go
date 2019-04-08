package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"

	"github.com/toasty-toast/advent-of-code-2015/utils"
)

type ingredient struct {
	Name                                            string
	Capacity, Durability, Flavor, Texture, Calories int
}

func parseIngredients(lines []string) []*ingredient {
	regex := regexp.MustCompile(`(?P<Name>[[:alpha:]]+): capacity (?P<Capacity>-?\d+), durability (?P<Durability>-?\d+), flavor (?P<Flavor>-?\d+), texture (?P<Texture>-?\d+), calories (?P<Calories>-?\d+)`)
	ingredients := make([]*ingredient, 0)
	for _, line := range lines {
		match := regex.FindStringSubmatch(line)
		next := new(ingredient)
		next.Name = match[1]
		next.Capacity, _ = strconv.Atoi(match[2])
		next.Durability, _ = strconv.Atoi(match[3])
		next.Flavor, _ = strconv.Atoi(match[4])
		next.Texture, _ = strconv.Atoi(match[5])
		next.Calories, _ = strconv.Atoi(match[6])
		ingredients = append(ingredients, next)
	}
	return ingredients
}

func getBestScore(ingredients []*ingredient) int {
	highest := math.MinInt32
	for i := 0; i <= 100; i++ {
		for j := 0; j <= 100-i; j++ {
			for k := 0; k <= 100-i-j; k++ {
				h := 100 - i - j - k
				capacity := ingredients[0].Capacity*i + ingredients[1].Capacity*j + ingredients[2].Capacity*k + ingredients[3].Capacity*h
				if capacity < 0 {
					capacity = 0
				}
				durability := ingredients[0].Durability*i + ingredients[1].Durability*j + ingredients[2].Durability*k + ingredients[3].Durability*h
				if durability < 0 {
					durability = 0
				}
				flavor := ingredients[0].Flavor*i + ingredients[1].Flavor*j + ingredients[2].Flavor*k + ingredients[3].Flavor*h
				if flavor < 0 {
					flavor = 0
				}
				texture := ingredients[0].Texture*i + ingredients[1].Texture*j + ingredients[2].Texture*k + ingredients[3].Texture*h
				if texture < 0 {
					texture = 0
				}

				score := capacity * durability * flavor * texture
				if score > highest {
					highest = score
				}
			}
		}
	}
	return highest
}

func getBestScoreFor500Calories(ingredients []*ingredient) int {
	highest := math.MinInt32
	for i := 0; i <= 100; i++ {
		for j := 0; j <= 100-i; j++ {
			for k := 0; k <= 100-i-j; k++ {
				h := 100 - i - j - k
				capacity := ingredients[0].Capacity*i + ingredients[1].Capacity*j + ingredients[2].Capacity*k + ingredients[3].Capacity*h
				if capacity < 0 {
					capacity = 0
				}
				durability := ingredients[0].Durability*i + ingredients[1].Durability*j + ingredients[2].Durability*k + ingredients[3].Durability*h
				if durability < 0 {
					durability = 0
				}
				flavor := ingredients[0].Flavor*i + ingredients[1].Flavor*j + ingredients[2].Flavor*k + ingredients[3].Flavor*h
				if flavor < 0 {
					flavor = 0
				}
				texture := ingredients[0].Texture*i + ingredients[1].Texture*j + ingredients[2].Texture*k + ingredients[3].Texture*h
				if texture < 0 {
					texture = 0
				}
				calories := ingredients[0].Calories*i + ingredients[1].Calories*j + ingredients[2].Calories*k + ingredients[3].Calories*h
				if calories < 0 {
					calories = 0
				}

				if calories != 500 {
					continue
				}

				score := capacity * durability * flavor * texture
				if score > highest {
					highest = score
				}
			}
		}
	}
	return highest
}

func main() {
	ingredients := parseIngredients(utils.ReadLines("input.txt"))
	fmt.Printf("Part 1: %d\n", getBestScore(ingredients))
	fmt.Printf("Part 2: %d\n", getBestScoreFor500Calories(ingredients))
}
