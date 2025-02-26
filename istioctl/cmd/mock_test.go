// Copyright Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"context"
	"fmt"
	"net/http"

	"k8s.io/cli-runtime/pkg/resource"
	"k8s.io/client-go/rest/fake"
	cmdtesting "k8s.io/kubectl/pkg/cmd/testing"
	"k8s.io/kubectl/pkg/cmd/util"

	"istio.io/istio/istioctl/pkg/cli"
	"istio.io/istio/pkg/kube"
)

type MockPortForwarder struct{}

func (m MockPortForwarder) Start() error {
	return nil
}

func (m MockPortForwarder) Address() string {
	return "localhost:3456"
}

func (m MockPortForwarder) Close() {
}

func (m MockPortForwarder) WaitForStop() {
}

var _ kube.PortForwarder = MockPortForwarder{}

type MockClient struct {
	// Results is a map of podName to the results of the expected test on the pod
	Results map[string][]byte
	kube.CLIClient
}

func (c MockClient) NewPortForwarder(_, _, _ string, _, _ int) (kube.PortForwarder, error) {
	return MockPortForwarder{}, nil
}

func (c MockClient) AllDiscoveryDo(_ context.Context, _, _ string) (map[string][]byte, error) {
	return c.Results, nil
}

func (c MockClient) EnvoyDo(ctx context.Context, podName, podNamespace, method, path string) ([]byte, error) {
	results, ok := c.Results[podName]
	if !ok {
		return nil, fmt.Errorf("unable to retrieve Pod: pods %q not found", podName)
	}
	return results, nil
}

func (c MockClient) EnvoyDoWithPort(ctx context.Context, podName, podNamespace, method, path string, port int) ([]byte, error) {
	results, ok := c.Results[podName]
	if !ok {
		return nil, fmt.Errorf("unable to retrieve Pod: pods %q not found", podName)
	}
	return results, nil
}

func init() {
	cli.MakeKubeFactory = func(k kube.CLIClient) util.Factory {
		tf := cmdtesting.NewTestFactory()
		_, _, codec := cmdtesting.NewExternalScheme()
		tf.UnstructuredClient = &fake.RESTClient{
			NegotiatedSerializer: resource.UnstructuredPlusDefaultContentConfig().NegotiatedSerializer,
			Resp: &http.Response{
				StatusCode: http.StatusOK,
				Header:     cmdtesting.DefaultHeader(),
				Body: cmdtesting.ObjBody(codec,
					cmdtesting.NewInternalType("", "", "foo")),
			},
		}
		return tf
	}
}
