// Copyright 2025 Naked Software, LLC
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package authorization

import (
	"fmt"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
	"net/url"
)

var _ = Describe("Authorization Code Grant", func() {
	Context("successful authorization", func() {
		const testState = "01234567890abcdefg"

		var recorder *httptest.ResponseRecorder

		BeforeEach(func() {
			query := url.Values{
				"redirect_uri": []string{
					"https://example.com/callback",
				},
				"state": []string{
					testState,
				},
			}
			recorder = httptest.NewRecorder()
			request := httptest.NewRequest(
				http.MethodGet,
				fmt.Sprintf("/authorize?%s", query.Encode()),
				nil,
			)
			handler := NewHandler()
			handler.ServeHTTP(recorder, request)
		})

		It("returns a redirect status code", func() {
			Expect(recorder.Code).To(Equal(http.StatusFound))
		})

		It("returns the redirect URI in the Location header", func() {
			location := recorder.Header().Get("Location")
			url, err := url.Parse(location)
			Expect(err).ToNot(HaveOccurred())
			Expect(url.Scheme).To(Equal("https"))
			Expect(url.Host).To(Equal("example.com"))
			Expect(url.Path).To(Equal("/callback"))
		})

		It("returns the state in the redirect URI", func() {
			location := recorder.Header().Get("Location")
			url, err := url.Parse(location)
			Expect(err).ToNot(HaveOccurred())
			query := url.Query()
			Expect(query.Get("state")).To(Equal(testState))
		})

		It("returns the authorization code in the redirect URI", func() {
			location := recorder.Header().Get("Location")
			url, err := url.Parse(location)
			Expect(err).ToNot(HaveOccurred())
			query := url.Query()
			Expect(query.Get("code")).To(Equal("0123456789"))
		})
	})

	Context("state is not specified", func() {
		It("is not included in the redirect URI", func() {
			query := url.Values{
				"redirect_uri": []string{
					"https://example.com/callback",
				},
			}
			recorder := httptest.NewRecorder()
			request := httptest.NewRequest(
				http.MethodGet,
				fmt.Sprintf("/authorize?%s", query.Encode()),
				nil,
			)
			handler := NewHandler()
			handler.ServeHTTP(recorder, request)
			location := recorder.Header().Get("Location")
			Expect(location).ToNot(BeEmpty())
			redirectURI, err := url.Parse(location)
			Expect(err).ToNot(HaveOccurred())
			Expect(redirectURI.Query().Has("state")).To(BeFalse())
		})
	})
})
