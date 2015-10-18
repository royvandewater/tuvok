package healthchecker_test

import (
	. "github.com/royvandewater/vulcanhealthcheck/healthchecker"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Healthcheck", func() {
	var sut *Healthchecker
	BeforeEach(func() {
		sut = new(Healthchecker)
	})

	It("should exist", func() {
		Expect(sut).NotTo(BeNil())
	})

	Describe("check", func() {
		Context("when called with a url", func() {
			var isHealthy bool
			var err error

			BeforeEach(func() {
				isHealthy, err = sut.Check()
			})

			It("should return a nil error", func() {
				Expect(err).To(BeNil())
			})

			It("should return true", func() {
				Expect(isHealthy).To(BeTrue())
			})
		})
	})
})
