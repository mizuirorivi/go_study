package main

import (
	"testing"
)

func Test_Division_1(t *testing.T) {
	if i, e := Division(6, 2); i != 3 || e != nil {
		t.Error("除算関数のテストが通りません")
	} else {
		t.Log("はじめのテストがパスしました")
	}
}

// func Test_Division_2(t *testing.T) {
// 	t.Error("パスしません")
// }

func BenchmarkDivision(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Division(4, 5)
	}
}

func BenchmarkTimeConsumingFunction(b *testing.B) {
	b.StopTimer()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		Division(4, 5)
	}
}
