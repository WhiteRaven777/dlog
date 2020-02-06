package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/WhiteRaven777/dlog"
)

type Sample struct {
	S   string
	B   []byte
	Bo  bool
	I   int
	Ui  uint
	F32 float32
	F64 float64
	T   time.Time
}

func main() {
	dlog.Log(`use "Log"`)
	dlog.Logf(`use "%s"`, "Logf")

	fmt.Print(dlog.Sprint(`use "Sprint"`))
	fmt.Print(dlog.Sprintln(`use "Sprintln"`))
	fmt.Print(dlog.Sprintf(`use "%s"`, "Sprintf"))

	err := errors.New("sample error")
	dlog.Log("error (Log):", err)
	dlog.Logf(`error (Logf): %v`, err)

	s := Sample{
		S:   "sample data",
		B:   []byte("sample data"),
		Bo:  true,
		I:   -100,
		Ui:  100,
		F32: float32(1.1234),
		F64: 1.1234,
		T:   time.Now(),
	}
	dlog.Log("struct:", s)

	m := map[string]interface{}{"sample map": s}
	dlog.Log("map:", m)

	a := [2]string{"sample", "data"}
	dlog.Log("array:", a)

	sl := []interface{}{"sample", "data", a}
	dlog.Log("slice:", sl)
}
