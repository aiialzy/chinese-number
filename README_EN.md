# chinese-number

chinese-number is a package that using to parse chinese number into number and convert number into chinese number. 

## Other language version
[English](README_EN.md)

## Installation
```sh
$ go get -u github.com/aiialzy/chinese-number
```

## Import
```go
import cn "github.com/aiialzy/chinese-number"
```

## Usage
### Parse chinese number into number
```go
han := "一亿一万一千一百一十一"
num, _ := cn.Parse(han)
fmt.Println(num, han)

// output
// 100011111 一亿一万一千一百一十一
```

### Convert number into chinese number
```go
num := 1<<63 - 1
han, _ := cn.Convert(num)
fmt.Println(num, han)

// output
// 9223372036854775807 九百二十二亿三千三百七十二万零三百六十八亿五千四百七十七万五千八百零七
```

## Benchmark
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

- (1): Total Repetitions achieved in one second, higher means more confident result
- (2): Single Repetition Duration (ns/op), lower is better
- (3): Heap Memory (B/op), lower is better
- (4): Average Allocations per Repetition (allocs/op), lower is better
