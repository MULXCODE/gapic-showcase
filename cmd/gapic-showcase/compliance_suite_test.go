// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/golang/protobuf/jsonpb"
	"github.com/googleapis/gapic-showcase/server/genproto"
	pb "github.com/googleapis/gapic-showcase/server/genproto"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"
)

// TestComplianceSuite ensures the REST test suite that we require GAPIC generators to pass works correctly. GAPIC generators should generate GAPICs for the Showcase API and issue the unary calls defined in the test suite using the GAPIC surface. The generators' test should follow the high-level logic below, as delineated by the comments below.
func TestComplianceSuite(t *testing.T) {
	// Run the Showcase REST server locally.
	server := httptest.NewUnstartedServer(nil)
	backend := createBackends()
	restServer := newEndpointREST(nil, backend)
	server.Config = restServer.server
	server.Start()
	defer server.Close()

	// Locate, load, and unmarshal the compliance suite.
	_, thisFile, _, _ := runtime.Caller(0)
	suiteFile := filepath.Join(filepath.Dir(thisFile), "../../schema/google/showcase/v1beta1/compliance_suite.textproto")
	textProto, err := ioutil.ReadFile(suiteFile)
	if err != nil {
		t.Fatalf("could not open suite file %q", suiteFile)
	}
	var suite pb.ComplianceSuite
	if err := prototext.Unmarshal(textProto, &suite); err != nil {
		t.Fatalf("error unmarshalling from text %s:\n   file: %s\n   input was: %s", err, suiteFile, textProto)
	}

	// Set handlers for each test case. When GAPIC generator tests do this, they should have
	// each of their handlers invoking the correct GAPIC library method for the Showcase API.
	restRPCs := map[string]prepRepeatDataTestFunc{
		"Compliance.RepeatDataBody":  prepRepeatDataBodyTest,
		"Compliance.RepeatDataQuery": prepRepeatDataQueryTest,
	}

	for _, group := range suite.GetGroup() {
		rpcsToTest := group.GetRpcs()
		for caseIdx, requestProto := range group.GetCases() {
			for rpcIdx, rpcName := range rpcsToTest {
				errorPrefix := fmt.Sprintf("[case %d/%q: rpc %q/%d/%q]",
					caseIdx, requestProto.GetName(), group.Name, rpcIdx, rpcName)

				// Ensure that we issue only the RPCs the test suite is expecting.
				rpcPrep, ok := restRPCs[rpcName]
				if !ok {
					t.Errorf("%s could not find prep function for this RPC", errorPrefix)
					continue
				}

				prepName, verb, path, requestBody, err := rpcPrep(requestProto)
				if err != nil {
					t.Errorf("%s error: %s", errorPrefix, err)
				}
				if got, want := prepName, rpcName; got != want {
					t.Errorf("%s retrieved mismatched prep function: got %q, want %q", errorPrefix, got, want)
				}

				// Issue the request. When GAPIC generator tests do this, they should simply
				// invoke the correct GAPIC library method for the Showcase API.
				httpRequest, err := http.NewRequest(verb, server.URL+path, strings.NewReader(requestBody))
				if err != nil {
					t.Errorf("%s error creating request: %s", errorPrefix, err)
					continue
				}
				httpResponse, err := http.DefaultClient.Do(httpRequest)
				if err != nil {
					t.Errorf("%s error issuing call: %s", errorPrefix, err)
					continue
				}

				// Check for successful response.
				if got, want := httpResponse.StatusCode, 200; got != want {
					t.Errorf("%s exit code: got %d, want %d", errorPrefix, got, want)
				}

				// Unmarshal httpResponse body, interpreted as JSON.
				// should do this.
				responseBody, err := ioutil.ReadAll(httpResponse.Body)
				httpResponse.Body.Close()
				if err != nil {
					t.Errorf("%s could not read httpResponse body: %s", errorPrefix, err)
					continue
				}
				var response genproto.RepeatResponse
				if err := jsonpb.UnmarshalString(string(responseBody), &response); err != nil {
					t.Errorf("%s could not unmarshal httpResponse body: %s\n   response body: %s",
						errorPrefix, err, string(responseBody))
					continue
				}

				// Check for expected response.
				if got, want := response.GetInfo(), requestProto.GetInfo(); !proto.Equal(got, want) {
					gotString, _ := prototext.Marshal(got)
					wantString, _ := prototext.Marshal(want)
					t.Errorf("%s unexpected response:\n   -->got:\n`%s`\n   -->want:\n`%s`\n",
						errorPrefix, gotString, wantString)
				}
			}
		}
	}
}

// The following are helpers for TestComplianceSuite, since Showcase doesn't intrinsically define a
// REST client. Each GAPIC generators should instead use the GAPIC it generated for the Showcase
// API.
type prepRepeatDataTestFunc func(request *genproto.RepeatRequest) (verb string, name string, path string, body string, err error)

func prepRepeatDataBodyTest(request *genproto.RepeatRequest) (verb string, name string, path string, body string, err error) {
	name = "Compliance.RepeatDataBody"
	marshaler := &jsonpb.Marshaler{}
	body, err = marshaler.MarshalToString(request)
	return name, "POST", "/v1beta1/repeat:body", body, err
}

func prepRepeatDataQueryTest(request *genproto.RepeatRequest) (verb string, name string, path string, body string, err error) {
	name = "Compliance.RepeatDataQuery"
	info := request.GetInfo()
	queryParams := []string{}
	addParam := func(condition bool, key, value string) {
		if !condition {
			return
		}
		queryParams = append(queryParams, fmt.Sprintf("info.%s=%s", key, value))
	}

	addParam(len(info.FString) > 0, "f_string", strings.ReplaceAll(info.FString, " ", "+"))

	var queryString string
	if len(queryParams) > 0 {
		queryString = fmt.Sprintf("?%s", strings.Join(queryParams, "&"))
	}
	return name, "GET", "/v1beta1/repeat:query" + queryString, body, err
}
