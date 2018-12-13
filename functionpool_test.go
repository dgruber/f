package main_test

import (
	"bytes"
	"io/ioutil"
	"os"

	. "github.com/dgruber/f"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Functionpool", func() {

	Context("Basic functions", func() {
		called := 0
		f := func(in string) string {
			called++
			return in + in
		}

		BeforeEach(func() {
			called = 0
		})

		It("should create a function pool, register a function, and run it", func() {
			fp := NewFunctionPool()
			Expect(fp).NotTo(BeNil())
			fp.Register("myfunc", f)
			Expect(called).To(BeNumerically("==", 0))
			fp.Run("myfunc", os.Stdin, os.Stdout, os.Stderr)
			Expect(called).To(BeNumerically("==", 1))
		})

		It("should run the function", func() {
			fp := NewFunctionPool()
			Expect(fp).NotTo(BeNil())
			fp.Register("myfunc", f)
			Expect(called).To(BeNumerically("==", 0))
			out := bytes.NewBufferString("")
			err := bytes.NewBufferString("")
			fp.Run("myfunc",
				ioutil.NopCloser(bytes.NewReader([]byte("hello world"))),
				out,
				err)
			Expect(called).To(BeNumerically("==", 1))
			Expect(err.String()).To(Equal(""))
			Expect(out.String()).To(Equal("hello worldhello world"))
		})
	})

})
