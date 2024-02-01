package main

import (
	"encoding/json"
	"fmt"
)

type DebugInfo struct {
	Level  string
	Msg    string
	author string // 未导出字段不会被json解析
}

type DebugInfo1 struct {
	Level  string `json:"level, omitempty"` // Level解析为level,忽略空值
	Msg    string `json:"message"`          // Msg解析为message
	Author string `json:"-"`                // 忽略Author
}

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

// 在调用Marshal(v interface{})时，该函数会判断v是否满足json.Marshaler或者
// encoding.Marshaler 接口，如果满足，则会调用这两个接口来进行转换（如果两个都满足，
// 优先调用json.Marshaler）
type Marshaler interface {
	MarshalJSON() ([]byte, error)
}

func (pt Point) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`{"X":%d, "Y":%d}`, pt.X, pt.Y)), nil
}

func marshal() {
	m := map[string][]string{
		"level":   {"debug"},
		"message": {"first not found", "Stack overflow"},
	}

	if data, err := json.Marshal(m); err == nil {
		fmt.Printf("%s\n", data)
		fmt.Println(string(data))
	}

	dbgInfs := []DebugInfo{
		DebugInfo{"debug", `File: "test.txt" Not Found`, "Cynhard"},
		DebugInfo{"", "Logic error", "Gopher"},
	}

	if data, err := json.Marshal(dbgInfs); err == nil {
		fmt.Printf("%s\n", data)
	}

	dbgInfs1 := []DebugInfo1{
		DebugInfo1{"debug", `File: "test.txt" Not Found`, "Cynhard"},
		DebugInfo1{"", "Logic error", "Gopher"},
	}

	if data, err := json.Marshal(dbgInfs1); err == nil {
		fmt.Printf("%s\n", data)
	}

	if data, err := json.Marshal(Circle{Point{50, 50}, 25}); err == nil {
		fmt.Printf("%s\n", data)
	}

	if data, err := json.Marshal(Point{50, 50}); err == nil {
		fmt.Printf("%s\n", data)
	}
}
