# フィボナッチ数列

## 実行方法

```
$ go run Fibonacci/fibonacci.go
0112358132134
0112358132134
0112358132134
0112358132134
0112358132134
0112358132134
```

## ベンチマーク

```
cd Fibonacci
$ go test -bench . -benchmem -cpu 4
goos: darwin
goarch: amd64
pkg: go_study/Fibonacci
BenchmarkFibonacci_01-4         1000000000               0.000041 ns/op        0 B/op          0 allocs/op
BenchmarkFibonacci_02-4         1000000000               0.000051 ns/op        0 B/op          0 allocs/op
BenchmarkFibonacci_04-4         1000000000               0.0238 ns/op          0 B/op          0 allocs/op
BenchmarkFibonacci_05-4         1000000000               0.00216 ns/op         0 B/op          0 allocs/op
BenchmarkFibonacci_06-4         1000000000               0.00134 ns/op         0 B/op          0 allocs/op
PASS
ok      go_study/Fibonacci      0.301s
```

## 参考

> https://qiita.com/ShrimpF/items/2f50a0b76d1cee29f0b3
