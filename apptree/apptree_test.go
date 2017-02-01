/**
 * Copyright 2017 ECS Team, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package apptree_test

import (
	. "github.com/ecsteam/cf-apptree/apptree"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/cloudfoundry-community/go-cfclient"

	"net/http"
	"net/url"
	"os"
	"sort"
	"strings"
)

type MockTransport struct{}

func (m *MockTransport) RoundTrip(request *http.Request) (*http.Response, error) {
	urlObj := request.URL

	fileName := urlToFilename(urlObj)
	reader, err := os.Open(fileName)
	if err != nil {
		return &http.Response{}, err
	} else {
		return &http.Response{
			ContentLength: -1,
			Request:       request,
			Body:          reader,
		}, nil
	}
}

func urlToFilename(url *url.URL) string {
	urlString := strings.Join([]string{url.Path, "?", url.RawQuery}, "")

	r := strings.NewReplacer("?", "_", "=", "_", "+", "_")
	name := r.Replace(urlString)

	return "fixtures" + name + ".json"
}

var _ = Describe("App Tree", func() {
	Describe("Given a Cloud Foundry Foundation", func() {
		Context("When running the app", func() {
			var client *cfclient.Client

			BeforeEach(func() {
				config := cfclient.DefaultConfig()
				config.HttpClient.Transport = &MockTransport{}

				var err error
				client, err = cfclient.NewClient(config)
				Ω(err).Should(BeNil())
			})

			It("Should return all necessary configs", func() {
				foundation, err := ListFoundation(client)
				Ω(err).Should(BeNil(), "Should have run without error")

				expected := []string{
					"apps1", "ecsteam", "p-spring-cloud-services",
					"redis-smoke-test-org", "sandbox", "spinnaker",
					"system",
				}

				actual := make([]string, 0, len(foundation.Orgs))
				for _, org := range foundation.Orgs {
					actual = append(actual, strings.ToLower(org.Name))
				}

				sort.Strings(actual)
				Ω(actual).Should(Equal(expected))
			})
		})
	})
})
