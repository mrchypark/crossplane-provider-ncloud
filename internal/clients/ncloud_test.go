package clients

import (
	"reflect"
	"strings"
	"testing"

	namespacedv1beta1 "github.com/mrchypark/crossplane-provider-ncloud/apis/namespaced/v1beta1"
)

func TestTerraformProviderConfiguration(t *testing.T) {
	tests := []struct {
		name    string
		spec    *namespacedv1beta1.ProviderConfigSpec
		creds   map[string]string
		want    map[string]any
		wantErr string
	}{
		{
			name: "public defaults",
			spec: &namespacedv1beta1.ProviderConfigSpec{Region: "KR"},
			creds: map[string]string{
				"access_key": "access",
				"secret_key": "secret",
			},
			want: map[string]any{
				"access_key":  "access",
				"secret_key":  "secret",
				"region":      "KR",
				"site":        "public",
				"support_vpc": true,
			},
		},
		{
			name: "gov site",
			spec: &namespacedv1beta1.ProviderConfigSpec{Region: "KR", Site: "gov"},
			creds: map[string]string{
				"access_key": "access",
				"secret_key": "secret",
			},
			want: map[string]any{
				"access_key":  "access",
				"secret_key":  "secret",
				"region":      "KR",
				"site":        "gov",
				"support_vpc": true,
			},
		},
		{
			name:    "missing access key",
			spec:    &namespacedv1beta1.ProviderConfigSpec{Region: "KR"},
			creds:   map[string]string{"secret_key": "secret"},
			wantErr: errMissingAccessKey,
		},
		{
			name:    "missing secret key",
			spec:    &namespacedv1beta1.ProviderConfigSpec{Region: "KR"},
			creds:   map[string]string{"access_key": "access"},
			wantErr: errMissingSecretKey,
		},
		{
			name: "missing region",
			spec: &namespacedv1beta1.ProviderConfigSpec{},
			creds: map[string]string{
				"access_key": "access",
				"secret_key": "secret",
			},
			wantErr: errMissingRegion,
		},
		{
			name: "invalid site",
			spec: &namespacedv1beta1.ProviderConfigSpec{Region: "KR", Site: "private"},
			creds: map[string]string{
				"access_key": "access",
				"secret_key": "secret",
			},
			wantErr: "unsupported ncloud site",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := terraformProviderConfiguration(tt.spec, tt.creds)
			if tt.wantErr != "" {
				if err == nil || !strings.Contains(err.Error(), tt.wantErr) {
					t.Fatalf("terraformProviderConfiguration() error = %v, want containing %q", err, tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Fatalf("terraformProviderConfiguration() unexpected error = %v", err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("terraformProviderConfiguration() = %#v, want %#v", got, tt.want)
			}
		})
	}
}
