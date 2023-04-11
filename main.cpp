#include <iostream>
#include <vector>
#include <string>
#include <ctime>
#include <cstdlib>
#include <iomanip>
#include <algorithm>

using namespace std;

#define FLUSH 1         //同花顺
#define LEOPARD 2       //豹子
#define OTHER 0         //其他

//每张扑克牌
class Poker
{
private:
	//点数
	string point;
	
	//花色
	string style;
	
	//大小排名，从0开始
	int size;

public:
	Poker(const string &point, const string &style, int size) : point(point), style(style), size(size)
	{}
	
	const string &getPoint() const
	{
		return point;
	}
	
	void setPoint(const string &point)
	{
		Poker::point = point;
	}
	
	const string &getStyle() const
	{
		return style;
	}
	
	void setStyle(const string &style)
	{
		Poker::style = style;
	}
	
	int getSize() const
	{
		return size;
	}
	
	void setSize(int size)
	{
		Poker::size = size;
	}
	
	friend ostream &operator<<(ostream &os, const Poker &poker)
	{
		os << poker.style << poker.point;
		return os;
	}
	
	bool operator==(const Poker &rhs) const
	{
		return size == rhs.size;
	}
	
	bool operator!=(const Poker &rhs) const
	{
		return !(rhs == *this);
	}
	
	bool operator<(const Poker &rhs) const
	{
		return size < rhs.size;
	}
	
	bool operator>(const Poker &rhs) const
	{
		return rhs < *this;
	}
	
	bool operator<=(const Poker &rhs) const
	{
		return !(rhs < *this);
	}
	
	bool operator>=(const Poker &rhs) const
	{
		return !(*this < rhs);
	}
};

vector<Poker> pokers;

//创建一副有序的扑克牌
void createPokers()
{
	vector<string> sizes = {"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"};
	vector<string> styles = {"♥", "♦", "♣", "♠"};
	
	for (int i = 0; i < sizes.size(); ++i)
	{
		for (int j = 0; j < styles.size(); ++j)
		{
			Poker poker(sizes[i], styles[j], i);
			pokers.push_back(poker);
		}
	}
}

//随机获取3张牌
//为了缩短运行时间，我并没有在每次发牌前洗牌，而是每次获取3个不相等的随机数，通过随机数来获取扑克牌中对应编号的牌
//洗牌太浪费时间了，这样操作也可以达到相同的效果
vector<Poker> get3Pokers()
{
	vector<Poker> myPokers;
	
	vector<int> pokerIndexs;
	
	while (pokerIndexs.size() < 3)
	{
		int randNum = rand() % (pokers.size());
		
		//没有找到
		if (find(pokerIndexs.begin(), pokerIndexs.end(), randNum) == pokerIndexs.end())
		{
			pokerIndexs.push_back(randNum);
		}
	}
	
	for (int i = 0; i < 3; ++i)
	{
		myPokers.push_back(pokers[pokerIndexs[i]]);
	}
	
	return myPokers;
}

//判断是否为同花顺
bool isFlush(vector<Poker> pokers)
{
	//牌的张数不为3
	if (pokers.size() != 3)
	{
		return false;
	}
	
	//不同花
	if (!((pokers[0].getStyle() == pokers[1].getStyle()) && (pokers[0].getStyle() == pokers[2].getStyle())))
	{
		return false;
	}
	
	//为顺子
	sort(pokers.begin(), pokers.end());
	if (pokers[0].getSize() == pokers[1].getSize() - 1 && pokers[0].getSize() == pokers[2].getSize() - 2)
	{
		return true;
	}
	return false;
	
}

//判断是否为豹子
bool isLeopard(vector<Poker> pokers)
{
	//牌的张数不为3
	if (pokers.size() != 3)
	{
		return false;
	}
	
	if (pokers[0] == pokers[1] && pokers[1] == pokers[2])
	{
		return true;
	}
	return false;
}

//判断牌的类型
int judgePoker(vector<Poker> pokers)
{
	if (pokers.size() != 3)
	{
		return OTHER;
	}
	if (isLeopard(pokers))
	{
		return LEOPARD;
	} else if (isFlush(pokers))
	{
		return FLUSH;
	}
	
	return OTHER;
}

int main()
{
	srand((unsigned int) time(0));
	
	int startTime = time(0);
	
	int countFlush = 0, countLeopard = 0;
	createPokers();
	
	// 在此处可以修改发牌次数，也就是数据的体量
	// 数据量越大，执行的时间就越长，结果就越准确，但不要超过21亿
	// 我自己的实验结果表明，数据量小于10000，就没有参考价值
	int n = 1000000;
	for (int i = 0; i < n; i++)
	{
		vector<Poker> myPokers = get3Pokers();
		int res = judgePoker(myPokers);
		if (res == LEOPARD)
		{
			countLeopard++;
		}
		if (res == FLUSH)
		{
			countFlush++;
		}
	}
	cout << setiosflags(ios::fixed) << setprecision(3);
	cout << "数据量：" << n << endl;
	cout << "同花顺：" << countFlush << " 概率为：" << 100.0 * countFlush / n << "%" << endl;
	cout << "豹子：" << countLeopard << " 概率为：" << 100.0 * countLeopard / n << "%" << endl;
	
	int endTime = time(0);
	cout << "运行时间：" << endTime - startTime << "s" << endl;
	system("pause");
	return 0;
}