package main

import (
	"testing"
)

var Debug bool = false

func TestFortune(t *testing.T) {
	if Debug {
		t.Skip("スキップさせる")
	}

	type args struct {
		val int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"大凶のテスト", args{0}, "大凶"},
		{"凶のテスト1", args{1}, "凶"},
		{"凶のテスト2", args{2}, "凶"},
		{"吉のテスト3", args{3}, "吉"},
		{"吉のテスト4", args{4}, "吉"},
		{"吉のテスト5", args{5}, "吉"},
		{"中吉のテスト6", args{6}, "中吉"},
		{"中吉のテスト7", args{7}, "中吉"},
		{"中吉のテスト8", args{8}, "中吉"},
		{"中吉のテスト9", args{9}, "中吉"},
		{"大吉のテスト", args{10}, "大吉"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := Fortune(test.args.val); got != test.want {
				t.Errorf("Fortune() = %v, want %v", got, test.want)
			}
		})
	}
}
