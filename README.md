# 比较豹子和同花顺的概率

从扑克牌中的52张牌（去除大小王）中随机取三张，抽中豹子和同花顺中的哪个概率更大。

## 构建

1. 克隆代码

```shell
git clone git@github.com:zhanglp0129/FlushCompareToLeopard.git
```

2. 编译

```shell
go build
```

### 注意事项

1. 为了减少运行时间，每次发牌之前都没有重新洗牌，而是获取3个不相同的随机数，根据这3个随机数找到对应编号的扑克牌。
2. 采用并发特性，降低程序运行时间。
