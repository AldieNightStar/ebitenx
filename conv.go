package ebitenx

import "github.com/hajimehoshi/ebiten/v2"

func convertKeysToInts(keys []ebiten.Key) []int {
	newKeys := make([]int, 0, 3)
	for _, k := range keys {
		newKeys = append(newKeys, int(k))
	}
	return newKeys
}
