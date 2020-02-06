package dlog

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

var (
	isDisplayTime bool
)

func log(msg ...interface{}) (str string) {
	if len(msg) > 0 {
		now := time.Now().UTC()

		strMsgs := make([]string, len(msg))
		for i := range msg {
			switch v := msg[i].(type) {
			case string:
				strMsgs[i] = v
			case bool:
				if v {
					strMsgs[i] = "true"
				} else {
					strMsgs[i] = "false"
				}
			case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
				strMsgs[i] = fmt.Sprintf("%d", v)
			case float32, float64:
				strMsgs[i] = fmt.Sprintf("%g", v)
			case time.Time:
				strMsgs[i] = v.Format(time.RFC3339)
			case error:
				fmt.Println(v)
				strMsgs[i] = v.Error()
			default:
				if b, err := json.Marshal(v); err == nil {
					strMsgs[i] = string(b)
				} else {
					strMsgs[i] = fmt.Sprint(v)
				}
			}
		}

		// build log
		pc, file, line, _ := runtime.Caller(2)
		f := runtime.FuncForPC(pc)
		var nowTime string
		if isDisplayTime {
			nowTime = now.Format(time.RFC3339Nano) + "\t"
		}
		str = fmt.Sprintf("%s%s (%s:%d)\t[debug] %s", nowTime, f.Name(), filepath.Base(file), line, strings.Join(strMsgs, " "))
	}
	return
}

func Log(msg ...interface{}) {
	fmt.Println(log(msg...))
}

func Logf(format string, a ...interface{}) {
	fmt.Printf(log(format)+"\n", a...)
}

func Sprint(msg ...interface{}) (ret string) {
	return log(msg...)
}

func Sprintf(format string, a ...interface{}) (ret string) {
	return fmt.Sprintf(log(format), a...)
}

func Sprintln(msg ...interface{}) (ret string) {
	return log(msg...) + "\n"
}
