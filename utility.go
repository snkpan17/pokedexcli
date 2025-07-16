package main

import (
	//"fmt"
	"math/rand"
)

func defeatChance(ratio float64) int {
	switch {
	case ratio > 2.5:
		return 90
	case ratio > 2.0:
		return 70
	case ratio > 1.5:
		return 50
	case ratio > 1.0:
		return 30
	case ratio > 0.5:
		return 10
	default:
		return 9
	}
}

func canDefeat(userExp int, pokeExp int) bool {
	expRatio := float64(userExp) / float64(pokeExp)
	defeat := defeatChance(expRatio)
	chance := rand.Intn(100)
	//fmt.Printf("userExp: %d, pokeExp: %d, expRatio: %f, defeat: %d, chance: %d\n", userExp, pokeExp, expRatio, defeat, chance)
	return chance >= defeat
}
