package main

import (
	"os"
)

func main() {
	fp := NewFunctionPool()
	fp.Register("upper", upper)
	fp.Register("reverse", reverse)
	fp.Run(os.Getenv("FUNCTION"))
}
