// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package crdvalidation

import (
	"testing"

	apiextensionstest "k8s.io/apiextensions-apiserver/pkg/test"
)

const subaccountAccessKeyCRD = "../../package/crds/subaccount.ncloud.crossplane.io_subaccountaccesskeys.yaml"

func TestSubaccountAccessKeyRequiresUsableSubaccountIdentity(t *testing.T) {
	validator, err := apiextensionstest.VersionValidatorFromFile(t, subaccountAccessKeyCRD, "v1alpha1")
	if err != nil {
		t.Fatalf("load CRD validator: %v", err)
	}

	tests := []struct {
		name       string
		forParams  map[string]any
		initParams map[string]any
		wantValid  bool
	}{
		{name: "identity-less Observe-only", wantValid: false},
		{name: "empty direct ID", forParams: map[string]any{"subAccountId": ""}, wantValid: false},
		{name: "empty reference name", forParams: map[string]any{"subAccountIdRef": map[string]any{"name": ""}}, wantValid: false},
		{name: "empty init direct ID", initParams: map[string]any{"subAccountId": ""}, wantValid: false},
		{name: "empty init reference name", initParams: map[string]any{"subAccountIdRef": map[string]any{"name": ""}}, wantValid: false},
		{name: "direct ID", forParams: map[string]any{"subAccountId": "subaccount-id"}, wantValid: true},
		{name: "reference", forParams: map[string]any{"subAccountIdRef": map[string]any{"name": "subaccount"}}, wantValid: true},
		{name: "selector", forParams: map[string]any{"subAccountIdSelector": map[string]any{}}, wantValid: true},
		{name: "init direct ID", initParams: map[string]any{"subAccountId": "subaccount-id"}, wantValid: true},
		{name: "init reference", initParams: map[string]any{"subAccountIdRef": map[string]any{"name": "subaccount"}}, wantValid: true},
		{name: "init selector", initParams: map[string]any{"subAccountIdSelector": map[string]any{}}, wantValid: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			spec := map[string]any{
				"forProvider":        map[string]any{},
				"managementPolicies": []any{"Observe"},
			}
			if tt.forParams != nil {
				spec["forProvider"] = tt.forParams
			}
			if tt.initParams != nil {
				spec["initProvider"] = tt.initParams
			}

			errs := validator(map[string]any{"spec": spec}, nil)
			if tt.wantValid && len(errs) != 0 {
				t.Fatalf("expected valid object, got errors: %v", errs)
			}
			if !tt.wantValid && len(errs) == 0 {
				t.Fatal("expected validation error")
			}
		})
	}
}
