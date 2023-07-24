package main

import (
	"fmt"
	"testing"
)

// 针对单个函数的测试
func TestIntMin(t *testing.T) {
	res := IntMin(1, 2)
	if res != 1 {
		t.Errorf("IntMin(1, 2) = %d, want 1\n", res)
	}

	if res := IntMin(2, -2); res != -2 {
		t.Errorf("IntMin(2, -2) = %d, want -2", res)
	}
}

func TestSum(t *testing.T) {
	if res := IntSum(1, 2, 3, 4); res != 10 {
		t.Errorf("IntSum(1, 2, 3, 4) = %d; want 10", res)
	}

	if res := IntSum(10, -10, 5, -4); res != 1 {
		t.Errorf("IntSum(10, -10, 5, -4) = %d; want -1", res)
	}
}

// 编写单个函数的批量测试函数
func TestIntMinTableDriven(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{2, 1, 1},
		{1, -2, -2},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			if res := IntMin(tt.a, tt.b); res != tt.want {
				t.Errorf("IntMin(%d, %d) = %d, want %d", tt.a, tt.b, res, tt.want)
			}
		})
	}
}

// 开启批量测试
func BenchmarkIntMin(b *testing.B) {
	// fmt.Println(b.N)
	for i := 0; i < b.N; i++ {
		IntMin(1, 2)
	}
}
