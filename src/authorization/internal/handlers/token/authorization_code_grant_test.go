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

package token

import (
	"encoding/json"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"mime"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
)

var _ = Describe("Authorization Code Grant", func() {
	var recorder *httptest.ResponseRecorder

	BeforeEach(func() {
		recorder = httptest.NewRecorder()
		values := url.Values{
			"grant_type":   []string{"authorization_code"},
			"code":         []string{"0123456789"},
			"redirect_uri": []string{"https://example.com/callback"},
		}
		body := values.Encode()
		request := httptest.NewRequest(
			"POST",
			"/token",
			strings.NewReader(body),
		)
		handler := NewHandler()
		handler.ServeHTTP(recorder, request)
	})

	It("returns a 200 recorder", func() {
		Expect(recorder.Code).To(Equal(http.StatusOK))
	})

	It("returns JSON content", func() {
		contentType := recorder.Header().Get("Content-Type")
		mediaType, _, err := mime.ParseMediaType(contentType)
		Expect(err).ToNot(HaveOccurred())
		Expect(mediaType).To(Equal("application/json"))
	})

	It("returns the access code", func() {
		var response struct {
			AccessCode string `json:"access_code"`
		}
		decoder := json.NewDecoder(recorder.Body)
		err := decoder.Decode(&response)
		Expect(err).ToNot(HaveOccurred())
		Expect(response.AccessCode).ToNot(BeEmpty())
	})

	It("returns a bearer token", func() {
		var response struct {
			TokenType string `json:"token_type"`
		}
		decoder := json.NewDecoder(recorder.Body)
		err := decoder.Decode(&response)
		Expect(err).ToNot(HaveOccurred())
		Expect(strings.ToLower(response.TokenType)).To(Equal("bearer"))
	})

	It("returns an expiration duration for the token", func() {
		var response struct {
			ExpiresIn int `json:"expires_in"`
		}
		decoder := json.NewDecoder(recorder.Body)
		err := decoder.Decode(&response)
		Expect(err).ToNot(HaveOccurred())
		Expect(response.ExpiresIn).ToNot(Equal(0))
	})

	It("returns a refresh token", func() {
		var response struct {
			RefreshToken string `json:"refresh_token"`
		}
		decoder := json.NewDecoder(recorder.Body)
		err := decoder.Decode(&response)
		Expect(err).ToNot(HaveOccurred())
		Expect(response.RefreshToken).ToNot(BeEmpty())
	})

	It("does not return the scope if there is no change", func() {
		var response struct {
			Scope string `json:"scope"`
		}
		decoder := json.NewDecoder(recorder.Body)
		err := decoder.Decode(&response)
		Expect(err).ToNot(HaveOccurred())
		Expect(response.Scope).To(BeEmpty())
	})
})
