package main

import (
	"os"
)

func main() {
	fp := NewFunctionPool().Register("upper", upper)
	fp.Run(os.Getenv("FUNCTION"))
}
