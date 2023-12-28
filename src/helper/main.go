package helper

import (
	"fmt"

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
