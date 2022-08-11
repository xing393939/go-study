package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type JSONTime struct {
	time.Time
}

// MarshalJSON on JSONTime format Time field with %Y-%m-%d %H:%M:%S
func (t *JSONTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05 MarshalJSON"))
	return []byte(formatted), nil
}

// UnMarshalJSON on JSONTime
func (t *JSONTime) UnmarshalJSON(data []byte) error {
	println("UnmarshalJSON start")
	// 空值不进行解析
	if len(data) <= 2 {
		*t = JSONTime{Time: time.Time{}}
		return nil
	}
	v, _ := time.ParseInLocation(`"2006-01-02 15:04:05"`, string(data), time.Local)
	*t = JSONTime{Time: v}
	return nil
}

func (t JSONTime) String() string {
	return t.Format("2006-01-02 15:04:05 string")
}

type my struct {
	T JSONTime `json:"t"`
}

func main() {
	a := JSONTime{time.Now()}
	b, _ := json.Marshal(&a)
	fmt.Println(string(b))
	fmt.Println(a)
	c := `{"t": "2022-08-08 08:08:08"}`
	d := my{}
	e, _ := json.Marshal(&d)
	fmt.Println(string(e))
	_ = json.Unmarshal([]byte(c), &d)
	fmt.Println(d)
}
