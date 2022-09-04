package main

import "fmt"

func playingDomino(cards [][]int, deck []int) interface{} {
	var temp = make([]int, 0)
	var max = make([]int, 0)

	for i := 0; i < len(cards)-1; i++ {
		temp = cards[i]
		max = cards[i+1]

		if cards[i][0]+cards[i][1] >= deck[0]+deck[1] && (cards[i][0] == deck[0] || cards[i][1] == deck[0] || cards[i][0] == deck[1] || cards[i][1] == deck[1]) {
			if max[0]+max[1] > temp[0]+temp[1] {
				return max
			}

			return temp
		}
	}

	return "Tutup kartu"
}

func main() {
	// PLAYING DOMINO

	fmt.Println(playingDomino([][]int{{6, 5}, {3, 4}, {2, 1}, {3, 3}, {4, 3}}, []int{4, 3}))
	fmt.Println(playingDomino([][]int{{6, 5}, {3, 3}, {3, 4}, {2, 1}, {3, 6}}, []int{3, 6}))
	fmt.Println(playingDomino([][]int{{6, 6}, {2, 4}, {3, 6}}, []int{5, 1}))
}
