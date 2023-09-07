# easy
提供了一些可能会用到的代码，帮助能更快的完成你的工作，从而有更多的时间去摸鱼（不是）。

注意本包只会引用官方sdk或来自golang.org/x的包，基本不会增加你的工程大小。

需要go版本在1.21.0以上。

包含以下功能：
+ 处理map的函数
+ 处理slice的函数
+ 获取当前goroutine id的函数
+ 解决[]float64，[]int64类型在与javascript交互时精度丢失的方案
+ 在http编程时可能会用到的error与通用返回结构体
+ 合并带权重的区间
+ 通用的分页结构体与分页返回结构体
+ 简单的hash密码生成与对比函数
+ 雪花id
+ []byte与string互转

## mapx
基于泛型，提供了一些处理map的函数。
### Keys
Keys方法返回map的所有key组成的slice，当然这个slice是乱序的。
```go
m := map[int]string{
  1: "one",
  2: "two",
}

s := mapx.Keys(m)

// Output: [1, 2] or [2, 1]
fmt.Println(s)
```

### Values
Values方法返回map的所有value组成的slice，当然这个slice是乱序的。
```go
m := map[int]string{
  1: "one",
  2: "two",
}

s := mapx.Values(m)

// Output: ["one", "two"] or ["two", "one"]
fmt.Println(s)
```

## slicex
基于泛型，提供了一些处理slice的函数。
### ToSliceAny
将给定的slice转换为[]any，element的顺序保持不变。

某些函数的入参是...any，但你的slice可能是[]int，这种情况可以使用本函数将你的slice转换为[]any，并作为入参。

```go
s := []int{1, 2, 3}

r := slicex.ToSliceAny(s)

// Output: [1, 2, 3]
fmt.Println(r)
```

### ToMap
将给定的slice转换为map[]struct{}，注意element需要是可比较的类型。

可以用于去重或判断slice是否包含指定element等操作。

```go
s := []int{1, 2, 3}

m := slicex.ToMap(s)

// Output: map[1:{}, 2:{}, 3:{}]
fmt.Println(m)
```
