package main

import (
	"fmt"
	"sync"
)

var (
	flush   = 0
	leopard = 0
	mutex   = sync.Mutex{}
	wait    = sync.WaitGroup{}
)

func main() {
	var count int
	fmt.Print("请输入模拟次数(单位为万)：")
	for {
		_, err := fmt.Scan(&count)
		if err != nil {
			fmt.Println("输入错误，请重新输入！")
		} else {
			break
		}
	}
	count *= 10000

	// 创建100个协程，来模拟
	goroutineCount := 100
	wait.Add(goroutineCount)
	for i := 0; i < goroutineCount; i++ {
		go run(count / goroutineCount)
	}
	wait.Wait()

	// 模拟完成，输出结果
	fmt.Printf("同花顺：%.4f%%\n", float64(flush)/float64(count)*100)
	fmt.Printf("豹子：%.4f%%\n", float64(leopard)/float64(count)*100)

	// 暂停终端
	fmt.Print("\n按回车键退出...")
	fmt.Scanln()
	fmt.Scanln()
}

// 运行模拟。需传入模拟次数，分别返回同花顺和豹子的次数
func run(count int) {
	tempFlush, tempLeopard := 0, 0
	for i := 0; i < count; i++ {
		t, err := GetPokerType(PokerDeal())
		if err != nil {
			panic(err)
		}

		switch t {
		case StraightFlush:
			tempFlush++
		case Leopard:
			tempLeopard++
		}
	}

	// 添加总结果数
	mutex.Lock()
	flush += tempFlush
	leopard += tempLeopard
	mutex.Unlock()
	wait.Done()
}
