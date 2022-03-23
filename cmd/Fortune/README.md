# go-study

## 内容

```
0～10のランダムな整数を生成する。
生成した値によって下記をコンソールに表示する。
10の場合は「大吉」
6以上10未満の場合は「中吉」
3以上6未満の場合は「吉」
1以上3未満の場合は「凶」
0の場合は「大凶」
```

## 実行

```
go run cmd/Fortune/fortune.go
```

## 結果

```
$ go run cmd/Fortune/fortune.go0
大凶
$ go run cmd/Fortune/fortune.go
5
吉
$ go run cmd/Fortune/fortune.go
9
中吉
$ go run cmd/Fortune/fortune.go
2
凶
$ go run cmd/Fortune/fortune.go
10
大吉
```