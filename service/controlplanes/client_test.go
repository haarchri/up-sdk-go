// Copyright 2021 Upbound Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package controlplanes

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/pkg/errors"

	"github.com/upbound/up-sdk-go"
	"github.com/upbound/up-sdk-go/fake"
)

func TestCreate(t *testing.T) {
	errBoom := errors.New("boom")

	cases := map[string]struct {
		reason string
		cfg    *up.Config
		params *ControlPlaneCreateParameters
		want   *ControlPlaneResponse
		err    error
	}{
		"NewRequestFailed": {
			reason: "Failing to construct a request should return an error.",
			cfg: &up.Config{
				Client: &fake.MockClient{
					MockNewRequest: fake.NewMockNewRequestFn(nil, errBoom),
				},
			},
			err: errBoom,
		},
		"DoFailed": {
			reason: "Failing to execute request should return an error.",
			cfg: &up.Config{
				Client: &fake.MockClient{
					MockNewRequest: fake.NewMockNewRequestFn(nil, nil),
					MockDo:         fake.NewMockDoFn(errBoom),
				},
			},
			err: errBoom,
		},
		"Successful": {
			reason: "A successful request should not return an error.",
			cfg: &up.Config{
				Client: &fake.MockClient{
					MockNewRequest: fake.NewMockNewRequestFn(nil, nil),
					MockDo:         fake.NewMockDoFn(nil),
				},
			},
			want: &ControlPlaneResponse{},
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			c := NewClient(tc.cfg)
			res, err := c.Create(context.Background(), "upbound", tc.params)
			if diff := cmp.Diff(tc.err, err, cmpopts.EquateErrors()); diff != "" {
				t.Errorf("\n%s\nCreate(...): -want error, +got error:\n%s", tc.reason, diff)
			}
			if diff := cmp.Diff(tc.want, res); diff != "" {
				t.Errorf("\n%s\nCreate(...): -want, +got:\n%s", tc.reason, diff)
			}
		})
	}
}

func TestGet(t *testing.T) {
	errBoom := errors.New("boom")

	cases := map[string]struct {
		reason string
		cfg    *up.Config
		params *ControlPlaneCreateParameters
		want   *ControlPlaneResponse
		err    error
	}{
		"NewRequestFailed": {
			reason: "Failing to construct a request should return an error.",
			cfg: &up.Config{
				Client: &fake.MockClient{
					MockNewRequest: fake.NewMockNewRequestFn(nil, errBoom),
				},
			},
			err: errBoom,
		},
		"DoFailed": {
			reason: "Failing to execute request should return an error.",
			cfg: &up.Config{
				Client: &fake.MockClient{
					MockNewRequest: fake.NewMockNewRequestFn(nil, nil),
					MockDo:         fake.NewMockDoFn(errBoom),
				},
			},
			err: errBoom,
		},
		"Successful": {
			reason: "A successful request should not return an error.",
			cfg: &up.Config{
				Client: &fake.MockClient{
					MockNewRequest: fake.NewMockNewRequestFn(nil, nil),
					MockDo:         fake.NewMockDoFn(nil),
				},
			},
			want: &ControlPlaneResponse{},
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			c := NewClient(tc.cfg)
			res, err := c.Get(context.Background(), "upbound", "test")
			if diff := cmp.Diff(tc.err, err, cmpopts.EquateErrors()); diff != "" {
				t.Errorf("\n%s\nGet(...): -want error, +got error:\n%s", tc.reason, diff)
			}
			if diff := cmp.Diff(tc.want, res); diff != "" {
				t.Errorf("\n%s\nGet(...): -want, +got:\n%s", tc.reason, diff)
			}
		})
	}
}

func TestList(t *testing.T) {
	errBoom := errors.New("boom")

	cases := map[string]struct {
		reason string
		cfg    *up.Config
		want   []ControlPlaneResponse
		err    error
	}{
		"NewRequestFailed": {
			reason: "Failing to construct a request should return an error.",
			cfg: &up.Config{
				Client: &fake.MockClient{
					MockNewRequest: fake.NewMockNewRequestFn(nil, errBoom),
				},
			},
			err: errBoom,
		},
		"DoFailed": {
			reason: "Failing to execute request should return an error.",
			cfg: &up.Config{
				Client: &fake.MockClient{
					MockNewRequest: fake.NewMockNewRequestFn(nil, nil),
					MockDo:         fake.NewMockDoFn(errBoom),
				},
			},
			err: errBoom,
		},
		"Successful": {
			reason: "A successful request should not return an error.",
			cfg: &up.Config{
				Client: &fake.MockClient{
					MockNewRequest: fake.NewMockNewRequestFn(nil, nil),
					MockDo:         fake.NewMockDoFn(nil),
				},
			},
			want: []ControlPlaneResponse{},
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			c := NewClient(tc.cfg)
			res, err := c.List(context.Background(), "upbound")
			if diff := cmp.Diff(tc.err, err, cmpopts.EquateErrors()); diff != "" {
				t.Errorf("\n%s\nList(...): -want error, +got error:\n%s", tc.reason, diff)
			}
			if diff := cmp.Diff(tc.want, res); diff != "" {
				t.Errorf("\n%s\nList(...): -want, +got:\n%s", tc.reason, diff)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	errBoom := errors.New("boom")

	cases := map[string]struct {
		reason string
		cfg    *up.Config
		err    error
	}{
		"NewRequestFailed": {
			reason: "Failing to construct a request should return an error.",
			cfg: &up.Config{
				Client: &fake.MockClient{
					MockNewRequest: fake.NewMockNewRequestFn(nil, errBoom),
				},
			},
			err: errBoom,
		},
		"DoFailed": {
			reason: "Failing to execute request should return an error.",
			cfg: &up.Config{
				Client: &fake.MockClient{
					MockNewRequest: fake.NewMockNewRequestFn(nil, nil),
					MockDo:         fake.NewMockDoFn(errBoom),
				},
			},
			err: errBoom,
		},
		"Successful": {
			reason: "A successful request should not return an error.",
			cfg: &up.Config{
				Client: &fake.MockClient{
					MockNewRequest: fake.NewMockNewRequestFn(nil, nil),
					MockDo:         fake.NewMockDoFn(nil),
				},
			},
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			c := NewClient(tc.cfg)
			err := c.Delete(context.Background(), "upbound", "test")
			if diff := cmp.Diff(tc.err, err, cmpopts.EquateErrors()); diff != "" {
				t.Errorf("\n%s\nDelete(...): -want error, +got error:\n%s", tc.reason, diff)
			}
		})
	}
}
