# chinese-number

chinese-number是一个用于解析和编码中文数字的库. 

## 安装
```sh
$ go get -u github.com/aiialzy/chinese-number
```

## 导入
```go
import cn "github.com/aiialzy/chinese-number"
```

## 用法
### 解析中文数字
```go
han := "一亿一万一千一百一十一"
num, _ := cn.Parse(han)
fmt.Println(num, han)

// 输出
// 100011111 一亿一万一千一百一十一
```

### 编码中文数字
```go
num := 1<<63 - 1
han, _ := cn.Convert(num)
fmt.Println(num, han)

// 数字
// 9223372036854775807 九百二十二亿三千三百七十二万零三百六十八亿五千四百七十七万五千八百零七
```

## 基准测试
goos: linux
<br />
goarch: amd64
<br />
pkg: github.com/aiialzy/chinese-number
<br />
cpu: AMD Ryzen 5 3600 6-Core Processor

| Benchmark name   |        (1) |            (2) |           (3) |              (4) |
| ---------------- | ---------: | -------------: | ------------: | ---------------: |
| BenchmarkParse   | **627830** | **1882 ns/op** |  **144 B/op** |  **1 allocs/op** |
| BenchmarkConvert | **525127** | **2273 ns/op** | **1144 B/op** | **17 allocs/op** |

- (1): 一秒之内执行的次数, 数值越大越好
- (2): 单次执行时间 (ns/op), 数值越小越好
- (3): 堆内存 (B/op), 数值越小越好
- (4): 平均每次执行分配次数, 数值越小越好
