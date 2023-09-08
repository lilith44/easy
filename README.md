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

``` go
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

``` go
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

基于给定的slice生成[]any，顺序保持不变。

某些函数的入参是...any，但你的slice可能是[]int，这种情况可以使用本函数将你的slice转换为[]any，并作为入参。

``` go
s := []int{1, 2, 3}

r := slicex.ToSliceAny(s)

// Output: [1, 2, 3]
fmt.Println(r)
```

### ToMap

基于给定的slice生成map[]struct{}，注意element需要是可比较的类型。

可以用于去重或判断slice是否包含指定element等操作。

``` go
s := []int{1, 2, 3}

m := slicex.ToMap(s)

// Output: map[1:{}, 2:{}, 3:{}]
fmt.Println(m)
```

### ToMapFunc

基于给定的slice生成map，其中key由k函数获取，value是slice的element。

通常可以在多次查询后进行数据组装时使用。

``` go
type School struct {
  Id   int64
  Name string
}

s := []School{
  {
    Id:   1,
    Name: "小学",
  },
  {
    Id:   2,
    Name: "中学",
  },
}

m := slicex.ToMapFunc(s, func(school School) int64 { return school.Id })

// Output: map[1:{Id:1 Name:小学}, 2:{Id:2 Name:"中学"}]
fmt.Printf("%+v\n", m)
```

### ToSliceFunc

基于给定的slice生成新的slice，新slice的element由v函数生成且顺序保持一致。

通常可以用于获取仅由结构体的部分字段组成的slice。

``` go
type School struct {
  Id   int64
  Name string
}

s := []*School{
  {
    Id:   1,
    Name: "小学",
  },
  {
    Id:   2,
    Name: "中学",
  },
}

r := slicex.ToSliceFunc(s, func(school *School) int64 { return school.Id })

// Output: [1, 2]
fmt.Println(r)
```

### Deduplicate

基于给定的slice生成去重后的slice；出现相同的element时，下标小的会被保留。注意element需要是可比较的类型。

``` go
s := []int{1, 2, 2, 3}

r := slicex.Deduplicate(s)

// Output: [1, 2, 3]
fmt.Println(r)
```

### DeduplicateFunc

基于给定的slice生成去重后的slice；出现相同的K时（K由dup函数生成），下标小的会被保留。注意K需要是可比较的类型。

``` go
type School struct {
  Id   int64
  Name string
}

s := []School{
  {
    Id:   1,
    Name: "小学",
  },
  {
    Id:   1,
    Name: "中学",
  },
}

r := slicex.DeduplicateFunc(s, func(school School) int64 { return school.Id })

// Output: [{Id:1 Name:小学}]
fmt.Printf("%+v\n", r)
```

### Concat

将多个slice按顺序拼接成一个新slice。

``` go
s1 := []int{1, 2}
s2 := []int{3, 4}
s3 := []int{5, 6}

r := slicex.Concat(s1, s2, s3)

// Output: [1, 2, 3, 4, 5, 6]
fmt.Println(r)
```

### IsUnique

判断给定slice的element是否全部唯一。注意element需要是可比较的类型。

``` go
s1 := []int{1, 2, 2, 3}
s2 := []int{1, 2, 3}

// Output: false true
fmt.Println(slicex.IsUnique(s1), slicex.IsUnique(s2))
```

### IsUniqueFunc

判断给定slice的K（由unique函数生成）是否全部唯一。注意K需要是可比较的类型。

``` go
type School struct {
  Id   int64
  Name string
}

s := []*School{
  {
    Id:   1,
    Name: "小学",
  },
  {
    Id:   1,
    Name: "中学",
  },
}

// Output: false
fmt.Println(slicex.IsUniqueFunc(s, func(school *School) int64 { return school.Id }))
```

### DeleteFunc

基于给定的slice，生成删除部分元素后的slice，当del函数返回true时，对应元素将被删除。

``` go
s := []int{1, 2, 2, 3}

r := slicex.DeleteFunc(s, func(i int) bool { return i%2 == 0 })

// Output: [1, 3]
fmt.Println(r)
```

### Paging

基于给定的slice，生成手动分页后的slice（第page页的size个element组成的slice）。

当page<=0或size<0时，发生panic。

注意此函数不会返回nil slice，而是返回一个空slice。这是考虑到避免返回一个null的数组。

``` go
s := []int{1, 2, 2, 3}

// Output: [1, 2, 2] []
fmt.Println(slicex.Paging(s, 1, 3), slicex.Paging(s, 3, 2))
```

## Float64s

为避免不同语言之间进行数据交互时出现精度丢失的问题，我们通常将数字转换为string进行传输。

例如我们可以在tag里加入,string来将数据转化为string。

``` go
type Paper struct {
  Score float64 `json:"score,string"`
}
```

但当我们使用一个数字切片时，tag里加上,string的方案似乎不太让人满意，会把整个切片当作string进行输出，但我们需要的是一个字符串数组。

``` go
type Paper struct {
  Scores []float64 `json:"scores,string"`
}
```

我们也可以将[]float64转换为[]string进行输出，将输入的[]string转换为[]
float64。但始终float64与string是不同的，在进行validate时有些许差别，导致在业务代码里可能会去做特殊处理。

为避免这种特殊处理，我们可以考虑定义新类型，并实现对应的Marshal和Unmarshal方法。

``` go
type Float64s []float64
```

使用示例：

``` go
type Paper struct {
  Scores easy.Float64s `json:"scores"`
}

p := &Paper{
  Scores: easy.Float64s{1.23, 2.34, 3.45},
}

data, err := json.Marshal(p)
if err != nil {
  return err
}

// Output: {"scores":["1.23","2.34","3.45"]}
fmt.Println(string(data))
```

当反序列化时，我们支持这样的数组：每个元素既可以是数字也可以是字符串数字。
例如 [1.23, "4.56"]可以被正常的反序列化。

## Int64s

参考Float64s。

## Gid

Gid函数返回当前的goroutine的id。官方没有提供对应的方法。

## MergeIntervals

用于解决这样的问题：

前提

+ 存在多个区间，每个区间有一个权重
+ 区间重合部分的权重可以进行求和

需求

+ 对重合部分的区间的权重进行求和，按照权重重新划分区间

例如，一个实际问题可以是：

学校的每个考场在指定时间段[left, right]内，用于p个考生进行考试。现需要求学校每个时刻的考试人数。

学校的考场1在[15:00, 16:00]被用于100人考试，考场2在[15:30, 16:30]被用于50人考试。那么在[15:00, 15:30]
有100人考试，在[15:30, 16:00]有150人考试，在[16:00, 16:30]有50人考试

使用示例：

``` go
intervals := []easy.Interval[int, int]{
  {
    Left:  900,
    Right: 960,
    Power: 100,
  },
  {
    Left:  930,
    Right: 990,
    Power: 50,
  },
}

r := easy.MergeIntervals(intervals...)

// Output: [{Left:900 Right:930 Power:100} {Left:930 Right:960 Power:150} {Left:960 Right:990 Power:50}]
fmt.Printf("%+v\n", r)
```

## IsIPInCIDR

用于判断给定的ip地址是否属于给定网段

``` go
ip := "192.168.0.1"
cidr := "192.168.0.0/24"

// Output: true
fmt.Println(easy.IsIPInCIDR(ip, cidr))
```

## Snowflake

基于雪花算法可以分布式地生成唯一id。

``` go
unique := func(...) easy.UniqueIdGenerator {
    return func() int64 {
        ...
        
        // 可以基于redis等分布式地生成唯一的int64值
    }
}

snowflake := easy.NewSnowflake(unique())

fmt.Println(snowflake.NextId())
```

## HttpError

携带http状态码、业务错误码、信息的结构体

``` go
var ErrUserNotFound = easy.NewHttpError(http.StatusNotFound, 1001, "用户不存在！")
```

## HttpResponse

http接口返回的通用结构体

``` go
func (u *user) Add(c echo.Context) error {
    ...
    rsp, err := u.service.Add(...)
    if err != nil {
        return err
    }
    
    return c.JSON(http.StatusOK, easy.Succeed(rsp))
}
```

在你的error handler里

``` go
switch e := err.(type) {
case x:
    return c.JSON(statusCode, easy.Fail(e.Error()))
}
```

## Underscore

将驼峰式的字符串转为下划线式

``` go
s := "blueEyes"

// Output: blue_eyes
fmt.Println(easy.Underscore(s))
```

## Camel

将下划线式的字符串转为驼峰式

``` go
s := "blue_eyes"

// Output: blueEyes
fmt.Println(easy.Underscore(s))
```

## Time

easy.Time在json marshal或者unmarshal时，会以"2006-01-02 15:04:05"的format进行

## Stack

基于泛型，提供了带有容量的、并发安全的栈。栈里的元素遵循"后进后出"的原则。

``` go
// 初始化一个int类型的容量为10的栈
stack := easy.NewStack[int](10)
```

### Push

向栈里push一个元素。若栈已满，则会返回error。

``` go
err := stack.Push(1)
if err != nil {
    // 通常是栈满引发的error
}
```

### Pop

从栈里pop一个元素。若栈为空，则会返回error。

``` go
popped, err := stack.Pop()
if err != nil {
    // 通常是栈空引发的error
}
```

## SortedStack

指单调栈，具体可参考百科。栈元素仅支持cmp.Ordered类型。也是并发安全的。

单调栈分递增和递减两种。递增单调栈在push时，会将大于当前元素的值全部pop；递减单调栈则将小于当前元素的值全部pop。

``` go
// 初始化一个int类型的递增单调栈
inc := easy.NewSortedStack[int](true)
// 初始化一个int类型的递减单调栈
desc := easy.NewSortedStack[int](false)
```

### Push

向栈里push一个元素，同时返回pop的元素slice。

``` go
inc := easy.NewSortedStack[int](true)
_ = inc.Push(10) // 此时栈内元素为[10]
_ = inc.Push(20) // 此时栈内元素为[10, 20]

popped := inc.Push(1) // 此时栈内元素为[1]， 20和10分别出栈

// Output: [20 10]
fmt.Println(popped)
```

### Pop

从栈里pop一个元素。若栈为空，则会返回error。

``` go
inc := easy.NewSortedStack[int](true)
_ = inc.Push(10) // 此时栈内元素为[10]
_ = inc.Push(20) // 此时栈内元素为[10, 20]

popped, _ := inc.Pop() // 此时20出栈

// Output: 20
fmt.Println(popped)
```