package main

import (
  . "github.com/arkors/oauth"

  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
)

var _ = Describe("Oauth", func() {
  Context("No index api", func() {
    It("returns a 404 Status Code", func() {
      Request("GET", "/", HandleIndex)
      Expect(response.Code).To(Equal(404))
    })
  })

})
