package main

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Error Is", func() {
	// peopleAge := make(map[string]int)
	// books := []string{
	// 	"The Eye of the World",
	// 	"The Great Hunt",
	// 	"The Dragon Reborn",
	// }

	Describe("Get IsEligibleToVaccine with error in GetAge", func() {
		It("returns wrapped error", func() {
			res := printBook(0)
			Expect(res).To(Equal("0. Books  is: The Eye of the World \n"))

		})
	})
})
