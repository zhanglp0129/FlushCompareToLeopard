package main

import "math/rand"

type Poker struct {
	// 花色
	Suit rune
	// 点数
	Point string
	// 扑克牌的大小，与点数一一对应
	Size int
}

var (
	AllPokers []Poker
)

// 初始化52张扑克牌
func init() {
	AllPokers = make([]Poker, 0, 52)
	suits := []rune{'♥', '♦', '♣', '♠'}
	points := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}
	for i := 0; i < len(points); i++ {
		for j := 0; j < len(suits); j++ {
			AllPokers = append(AllPokers, Poker{suits[j], points[i], i})
		}
	}
}

// PokerDeal 发3张不同的扑克牌
func PokerDeal() []Poker {
	// 生成3个不同的随机数
	ids := make([]int, 0, 3)
	for i := 0; i < 3; i++ {
		r, unique := rand.Int()%52, true
		for j := 0; j < len(ids); j++ {
			if ids[j] == r {
				unique = false
				break
			}
		}
		if unique {
			ids = append(ids, r)
		} else {
			i--
		}
	}

	return []Poker{AllPokers[ids[0]], AllPokers[ids[1]], AllPokers[ids[2]]}
}
