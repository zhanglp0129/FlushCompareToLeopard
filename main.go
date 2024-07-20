package main

import "fmt"

func main() {
	count := 1000000000
	flush, leopard := run(count)
	fmt.Printf("同花顺：%.4f%%\n", float64(flush)/float64(count)*100)
	fmt.Printf("豹子：%.4f%%\n", float64(leopard)/float64(count)*100)
}

// 运行模拟。需传入模拟次数，分别返回同花顺和豹子的次数
func run(count int) (int, int) {
	flush, leopard := 0, 0
	for i := 0; i < count; i++ {
		t, err := GetPokerType(PokerDeal())
		if err != nil {
			panic(err)
		}

		switch t {
		case StraightFlush:
			flush++
		case Leopard:
			leopard++
		}
	}
	return flush, leopard
}
