package gotest

import (
	"testing"
)

func TestAdd(t *testing.T) {
	if ans := Add(1, 2); ans != 3 {
		t.Errorf("1 + 2 expected be 3, but %d got", ans)
	}

	if ans := Add(-10, -20); ans != -30 {
		t.Errorf("-10 + -20 expected be -30, but %d got", ans)
	}
}

func TestMul(t *testing.T) {
	if ans := Mul(2, 3); ans != 6 {
		t.Errorf("2 * 3 expected be 6, but %d got", ans)
	}

	if ans := Mul(-2, 3); ans != -6 {
		t.Errorf("-2 * 3 expected be -6, but %d got", ans)
	}
}

// 子测试是 Go 语言内置支持的，可以在某个测试用例中，根据测试场景使用 t.Run创建不同的子测试用例
func TestMul1(t *testing.T) {
	t.Run("pos", func(t *testing.T) {
		if Mul(2, 3) != 6 {
			t.Fatal("fail")
		}
	})

	t.Run("neg", func(t *testing.T) {
		if Mul(-2, 3) != -6 {
			t.Fatal("fail")
		}
	})
}

// 对于多个子测试的场景，更推荐如下的写法(table-driven tests)
func TestMul2(t *testing.T) {
	cases := []struct {
		Name           string
		A, B, Expected int
	}{
		{"pos", 2, 3, 6},
		{"neg", -2, 3, -6},
		{"zero", 2, 0, 0},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			if ans := Mul(c.A, c.B); ans != c.Expected {
				t.Fatalf("%d * %d expected %d, but %d got",
					c.A, c.B, c.Expected, ans)
			}
		})
	}
}

// 对一些重复的逻辑，抽取出来作为公共的帮助函数(helpers)
// 可以增加测试代码的可读性和可维护性 借助帮助函数，可以让
// 测试用例的主逻辑看起来更清晰

type calcCase struct{ A, B, Expected int }

func createMulTestCase(t *testing.T, c *calcCase) {
	// Go 语言在 1.9 版本中引入了 t.Helper()，用于标注该函数是帮助函数
	// 报错时将输出帮助函数调用者的信息，而不是帮助函数的内部信息
	t.Helper()

	if ans := Mul(c.A, c.B); ans != c.Expected {
		t.Fatalf("%d * %d expected %d, but %d got",
			c.A, c.B, c.Expected, ans)
	}
}

func TestMul3(t *testing.T) {
	createMulTestCase(t, &calcCase{2, 3, 6})
	createMulTestCase(t, &calcCase{-2, 3, -6})
	//createMulTestCase(t, &calcCase{2, 0, 1})
}
