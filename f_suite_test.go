package main_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestF(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "F Suite")
}
