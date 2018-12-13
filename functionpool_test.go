package main_test

import (
	. "github.com/dgruber/f"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Functionpool", func() {

	Context("Basic functions", func() {
		called := 0
		f := func() {
			called++
		}

		BeforeEach(func() {
			called = 0
		})

		It("should create a function pool, register a function, and run it", func() {
			fp := NewFunctionPool()
			Expect(fp).NotTo(BeNil())
			fp.Register("myfunc", f)
			Expect(called).To(BeNumerically("==", 0))
			fp.Run("myfunc")
			Expect(called).To(BeNumerically("==", 1))
		})
	})

})
