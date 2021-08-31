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

package grpcgen

import (
	"fmt"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"

<<<<<<< HEAD
	"istio.io/istio/pkg/istio-agent/grpcxds"
)

var wildcardMap = map[string]struct{}{}
=======
	"istio.io/istio/pilot/pkg/model"
	"istio.io/istio/pilot/pkg/util/sets"
	"istio.io/istio/pkg/istio-agent/grpcxds"
)

var node = &model.Proxy{DNSDomain: "ns.svc.cluster.local", Metadata: &model.NodeMetadata{Namespace: "ns"}}
>>>>>>> 4d2173743a3d977e58cd656bc671d6a5d78f87c6

func TestListenerNameFilter(t *testing.T) {
	cases := map[string]struct {
		in          []string
<<<<<<< HEAD
		want        listenerNameFilter
=======
		want        listenerNames
>>>>>>> 4d2173743a3d977e58cd656bc671d6a5d78f87c6
		wantInbound []string
	}{
		"simple": {
			in: []string{"foo.com:80", "foo.com:443", "wildcard.com"},
<<<<<<< HEAD
			want: listenerNameFilter{
				"foo.com":      {"80": {}, "443": {}},
				"wildcard.com": wildcardMap,
=======
			want: listenerNames{
				"foo.com": {
					RequestedNames: sets.NewSet("foo.com"),
					Ports:          sets.NewSet("80", "443"),
				},
				"wildcard.com": {RequestedNames: sets.NewSet("wildcard.com")},
>>>>>>> 4d2173743a3d977e58cd656bc671d6a5d78f87c6
			},
		},
		"plain-host clears port-map": {
			in:   []string{"foo.com:80", "foo.com"},
<<<<<<< HEAD
			want: listenerNameFilter{"foo.com": wildcardMap},
		},
		"port-map stays clear": {
			in:   []string{"foo.com:80", "foo.com", "foo.com:443"},
			want: listenerNameFilter{"foo.com": wildcardMap},
=======
			want: listenerNames{"foo.com": {RequestedNames: sets.NewSet("foo.com")}},
		},
		"port-map stays clear": {
			in: []string{"foo.com:80", "foo.com", "foo.com:443"},
			want: listenerNames{"foo.com": {
				RequestedNames: sets.NewSet("foo.com"),
			}},
>>>>>>> 4d2173743a3d977e58cd656bc671d6a5d78f87c6
		},
		"special listeners preserved exactly": {
			in: []string{
				"foo.com:80",
				fmt.Sprintf(grpcxds.ServerListenerNameTemplate, "foo:1234"),
				fmt.Sprintf(grpcxds.ServerListenerNameTemplate, "foo"),
				fmt.Sprintf(grpcxds.ServerListenerNameTemplate, "[::]:8076"),
			},
<<<<<<< HEAD
			want: listenerNameFilter{
				"foo.com": {"80": {}},
				fmt.Sprintf(grpcxds.ServerListenerNameTemplate, "foo:1234"):  wildcardMap,
				fmt.Sprintf(grpcxds.ServerListenerNameTemplate, "foo"):       wildcardMap,
				fmt.Sprintf(grpcxds.ServerListenerNameTemplate, "[::]:8076"): wildcardMap,
=======
			want: listenerNames{
				"foo.com": {
					RequestedNames: sets.NewSet("foo.com"),
					Ports:          sets.NewSet("80"),
				},
				fmt.Sprintf(grpcxds.ServerListenerNameTemplate, "foo:1234"): {
					RequestedNames: sets.NewSet(fmt.Sprintf(grpcxds.ServerListenerNameTemplate, "foo:1234")),
				},
				fmt.Sprintf(grpcxds.ServerListenerNameTemplate, "foo"): {
					RequestedNames: sets.NewSet(fmt.Sprintf(grpcxds.ServerListenerNameTemplate, "foo")),
				},
				fmt.Sprintf(grpcxds.ServerListenerNameTemplate, "[::]:8076"): {
					RequestedNames: sets.NewSet(fmt.Sprintf(grpcxds.ServerListenerNameTemplate, "[::]:8076")),
				},
>>>>>>> 4d2173743a3d977e58cd656bc671d6a5d78f87c6
			},
			wantInbound: []string{
				fmt.Sprintf(grpcxds.ServerListenerNameTemplate, "foo:1234"),
				fmt.Sprintf(grpcxds.ServerListenerNameTemplate, "foo"),
				fmt.Sprintf(grpcxds.ServerListenerNameTemplate, "[::]:8076"),
			},
		},
<<<<<<< HEAD
	}
	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			got := newListenerNameFilter(tt.in)
=======
		"expand shortnames": {
			in: []string{
				"bar",
				"bar.ns",
				"bar.ns.svc",
				"bar.ns.svc.cluster.local",
				"foo:80",
				"foo.ns:81",
				"foo.ns.svc:82",
				"foo.ns.svc.cluster.local:83",
			},
			want: listenerNames{
				"bar":        {RequestedNames: sets.NewSet("bar")},
				"bar.ns":     {RequestedNames: sets.NewSet("bar.ns")},
				"bar.ns.svc": {RequestedNames: sets.NewSet("bar.ns.svc")},
				"bar.ns.svc.cluster.local": {RequestedNames: sets.NewSet(
					"bar",
					"bar.ns",
					"bar.ns.svc",
					"bar.ns.svc.cluster.local",
				)},
				"foo":        {RequestedNames: sets.NewSet("foo"), Ports: sets.NewSet("80")},
				"foo.ns":     {RequestedNames: sets.NewSet("foo.ns"), Ports: sets.NewSet("81")},
				"foo.ns.svc": {RequestedNames: sets.NewSet("foo.ns.svc"), Ports: sets.NewSet("82")},
				"foo.ns.svc.cluster.local": {
					RequestedNames: sets.NewSet(
						"foo",
						"foo.ns",
						"foo.ns.svc",
						"foo.ns.svc.cluster.local",
					),
					Ports: sets.NewSet("80", "81", "82", "83"),
				},
			},
		},
	}
	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			got := newListenerNameFilter(tt.in, node)
>>>>>>> 4d2173743a3d977e58cd656bc671d6a5d78f87c6
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Fatal(diff)
			}
			gotInbound := got.inboundNames()
			sort.Strings(gotInbound)
			sort.Strings(tt.wantInbound)
			if diff := cmp.Diff(gotInbound, tt.wantInbound); diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}
<<<<<<< HEAD
=======

func TestTryFindFQDN(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{"foo", "foo.ns.svc.cluster.local"},
		{"foo.ns", "foo.ns.svc.cluster.local"},
		{"foo.ns.svc", "foo.ns.svc.cluster.local"},
		{"foo.ns.svc.cluster.local", ""},
		{"foo.com", ""},
		{"foo.ns.com", ""},
		{"foo.ns.svc.notdnsdomain", ""},
		{"foo.ns.svc.cluster.local.extra", ""},
		{"xds.istio.io/grpc/lds/inbound/0.0.0.0:1234", ""},
	}

	for _, tc := range cases {
		t.Run(tc.in, func(t *testing.T) {
			if got := tryFindFQDN(tc.in, node); got != tc.want {
				t.Errorf("want %q but got %q", tc.want, got)
			}
		})
	}
}
>>>>>>> 4d2173743a3d977e58cd656bc671d6a5d78f87c6
