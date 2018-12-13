package main

import (
	"fmt"
	"os"
)

type fp struct {
	fm map[string]func()
}

func NewFunctionPool() *fp {
	return &fp{
		fm: make(map[string]func()),
	}
}

func (f *fp) Register(name string, ptr func()) *fp {
	f.fm[name] = ptr
	return f
}

func (f *fp) Run(name string) {
	fc, exists := f.fm[name]
	if !exists {
		fmt.Fprintf(os.Stderr, "Function %s does not exist.\n", name)
	} else {
		fc()
	}
}
