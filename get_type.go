package main

import "errors"

// PokerType 牌型
type PokerType uint8

const (
	// Other 其他牌型
	Other PokerType = iota
	// StraightFlush 同花顺
	StraightFlush
	// Leopard 豹子
	Leopard
)

// GetPokerType 判断扑克牌的牌型
func GetPokerType(pokers []Poker) (PokerType, error) {
	if len(pokers) != 3 {
		return Other, errors.New("扑克牌的数量必须为3")
	}

	// 判断是否为豹子
	if pokers[0].Size == pokers[1].Size && pokers[1].Size == pokers[2].Size {
		return Leopard, nil
	}
	// 判断是否为同花顺
	if pokers[0].Suit == pokers[1].Suit && pokers[1].Suit == pokers[2].Suit {
		// 为同花，再排序并判断是否为顺子
		pokerSort(pokers)
		if pokers[0].Size+1 == pokers[1].Size && pokers[1].Size+1 == pokers[2].Size {
			return StraightFlush, nil
		}
	}
	// 其他牌型
	return Other, nil
}

// 按照大小给扑克牌直接插入排序
func pokerSort(pokers []Poker) {
	for i := 1; i < len(pokers); i++ {
		cur := pokers[i]
		j := i - 1
		for ; j >= 0; j-- {
			if pokers[j].Size > pokers[i].Size {
				pokers[j+1] = pokers[j]
			} else {
				break
			}
		}
		pokers[j+1] = cur
	}
}
