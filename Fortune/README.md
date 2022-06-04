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
go run Fortune/fortune.go
```

## テスト実施方法

```
$ go test ./Fortune
ok      go_study/Fortune        0.098s

or

$ go test -v ./Fortune
=== RUN   TestFortune
=== RUN   TestFortune/大凶のテスト
=== RUN   TestFortune/凶のテスト1
=== RUN   TestFortune/凶のテスト2
=== RUN   TestFortune/吉のテスト3
=== RUN   TestFortune/吉のテスト4
=== RUN   TestFortune/吉のテスト5
=== RUN   TestFortune/中吉のテスト6
=== RUN   TestFortune/中吉のテスト7
=== RUN   TestFortune/中吉のテスト8
=== RUN   TestFortune/中吉のテスト9
=== RUN   TestFortune/大吉のテスト
=== RUN   TestFortune/不明のテスト1
=== RUN   TestFortune/不明のテスト2
--- PASS: TestFortune (0.00s)
    --- PASS: TestFortune/大凶のテスト (0.00s)
    --- PASS: TestFortune/凶のテスト1 (0.00s)
    --- PASS: TestFortune/凶のテスト2 (0.00s)
    --- PASS: TestFortune/吉のテスト3 (0.00s)
    --- PASS: TestFortune/吉のテスト4 (0.00s)
    --- PASS: TestFortune/吉のテスト5 (0.00s)
    --- PASS: TestFortune/中吉のテスト6 (0.00s)
    --- PASS: TestFortune/中吉のテスト7 (0.00s)
    --- PASS: TestFortune/中吉のテスト8 (0.00s)
    --- PASS: TestFortune/中吉のテスト9 (0.00s)
    --- PASS: TestFortune/大吉のテスト (0.00s)
    --- PASS: TestFortune/不明のテスト1 (0.00s)
    --- PASS: TestFortune/不明のテスト2 (0.00s)
PASS
ok      go_study/Fortune        0.229s
```


## 結果

```
$ go run Fortune/fortune.go
0
大凶
$ go run Fortune/fortune.go
5
吉
$ go run Fortune/fortune.go
9
中吉
$ go run Fortune/fortune.go
2
凶
$ go run Fortune/fortune.go
10
大吉
```
