package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

const fiveOfAKind = 6
const fourOfAKind = 5
const fullHouse = 4
const threeOfAKind = 3
const twoPair = 2
const pair = 1
const highCard = 0

type hand struct {
	handType int
	bid      int
	cardVals []int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var hands []hand

	for scanner.Scan() {
		text := scanner.Text()
		s := strings.Split(text, " ")
		bid, err := strconv.Atoi(s[1])
		if err != nil {
			return
		}
		cards := make(map[int]int)
		var cardVals []int
		for _, card := range s[0] {
			var num int
			if unicode.IsDigit(card) {

				num = int(card - '0')
			} else {
				switch card {
				case 'T':
					num = 10
				case 'J':
					num = 1
				case 'Q':
					num = 12
				case 'K':
					num = 13
				case 'A':
					num = 14
				}
			}
			cards[num]++
			cardVals = append(cardVals, num)
		}
		jokerAmount, hasJokers := cards[1]
		if hasJokers {
			delete(cards, 1)
			maxAmount := 0
			var key int
			for k, v := range cards {
				if v > maxAmount {
					maxAmount = v
					key = k
				}
			}
			cards[key] += jokerAmount
		}
		var handType int
		switch len(cards) {
		case 5:
			handType = highCard
		case 4:
			handType = pair
		case 3:
			handType = twoPair
			for _, v := range cards {
				if v == 3 {
					handType = threeOfAKind
					break
				}
			}
		case 2:
			for _, v := range cards {
				if v == 1 || v == 4 {
					handType = fourOfAKind
				} else {
					handType = fullHouse
				}
				break
			}
		case 1:
			handType = fiveOfAKind
		}

		hands = append(hands, hand{bid: bid, handType: handType, cardVals: cardVals})

	}
	sum := 0
	sort.Slice(hands, func(i, j int) bool {
		if hands[i].handType == hands[j].handType {
			for k := 0; k <= 5; k++ {
				if hands[i].cardVals[k] == hands[j].cardVals[k] {
					continue
				}
				return hands[i].cardVals[k] < hands[j].cardVals[k]
			}
			return true
		} else {

			return hands[i].handType < hands[j].handType

		}
	})
	for k, v := range hands {
		sum += (k + 1) * v.bid
	}

	fmt.Println(sum)
}
