package h

import (
	"fmt"
	"os"
	"time"

	"github.com/rshbintech/gomega/gleak/goroutine"
)

func ShowThreadId() {
	th := goroutine.Current()
	fmt.Printf("ID: %+v\n", th.ID)
}

func GetThreadId() (s string) {
	th := goroutine.Current()
	s = fmt.Sprintf("%v", th.ID)
	return
}

func PanicError(source string, err error) {
	println(source, err.Error())
	os.Exit(1)
}

func ExitOnTime() {
	time.Sleep(time.Millisecond * 1600)
	os.Exit(0)
}

func Pause() {
	time.Sleep(time.Millisecond * 500)
}
