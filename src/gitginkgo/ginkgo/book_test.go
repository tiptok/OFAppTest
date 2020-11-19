package main

import (
	"encoding/json"
	"github.com/gavv/httpexpect"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"testing"
)

func TestBook(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Book test Case Suite")
}

var _ = Describe("test book init", func() {
	var book Book
	BeforeEach(func() {
		book = Book{
			Name:   "go",
			Author: "tip",
			Time:   159781231684131,
		}
		_, err := json.Marshal(book)
		Expect(err).NotTo(HaveOccurred())
	})
	Describe("check book status", func() {
		Context("get book", func() {
			It("valid book name", func() {
				defer GinkgoRecover()
				Expect(book.Name).Should(ContainSubstring("go"))
				Expect(book.Author).To(Equal("tip"))
				Expect(book.Time).To(Equal(int64(159781231684131)))

				httpExpect := httpexpect.New(GinkgoT(), "https://postman-echo.com")
				body := map[string]interface{}{
					"foo1": "foo1",
					"bar2": "bar2",
				}
				httpExpect.POST("/post").
					WithJSON(body).
					Expect().
					Status(http.StatusOK).
					JSON().
					Object().
					ContainsKey("json").Value("json").Object().
					ContainsKey("foo1").ValueEqual("foo1", "foo1")
			})
		})
	})
	AfterEach(func() {
		book.Name = "C#"
	})
})
