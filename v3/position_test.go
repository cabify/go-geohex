package geohex

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Position", func() {

	It("should generate centroids", func() {
		pos := &Position{4, -5, zooms[0]}
		cnt := pos.Centroid()
		Expect(cnt.E).To(BeNumerically("~", 1, 0.1))
		Expect(cnt.N).To(BeNumerically("~", -0.06415, 0.1))
	})

})
