package ginkgomock

import (
	"fmt"

	"github.com/golang/mock/gomock"
	"github.com/onsi/ginkgo"
)

type ginkgoReporter struct {}

func (g *ginkgoReporter) Errorf(format string, args ...interface{}) {
	fmt.Fprintf(ginkgo.GinkgoWriter, format, args...)
}

func (g *ginkgoReporter) Fatalf(format string, args ...interface{}) {
	ginkgo.Fail(fmt.Sprintf(format, args...))
}

// NewController creates a controller that is receives a TestReporter that links it to the Ginkgo.
func NewController() *gomock.Controller {
	return gomock.NewController(&ginkgoReporter{})
}
