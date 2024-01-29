package bufpool_test

import (
	"fmt"

	"github.com/johnbenedictyan/bufpool"
)

func ExampleGet() {
	buf := bufpool.Get(1234)

	fmt.Println("len", buf.Len())
	// Output: len 1234

	bufpool.Put(buf)
}
