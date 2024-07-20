package main

type Poker struct {
	// 花色
	Suit rune
	// 点数
	Point string
	// 扑克牌的大小，与点数一一对应
	Size int
}

var (
	Pokers []Poker
)

// InitPokers 初始化52张扑克牌
func InitPokers() {
	Pokers = make([]Poker, 0, 52)
	suits := []rune{'♥', '♦', '♣', '♠'}
	points := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}
	for i := 0; i < len(points); i++ {
		for j := 0; j < len(suits); j++ {
			Pokers = append(Pokers, Poker{suits[j], points[i], i})
		}
	}
}
