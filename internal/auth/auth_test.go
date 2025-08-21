package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	type response struct {
		apiKey    string
		errString string
	}

	tests := map[string]struct {
		input http.Header
		want  response
	}{
		"no-header": {
			input: http.Header{},
			want: response{
				apiKey:    "",
				errString: ErrNoAuthHeaderIncluded.Error(),
			},
		},
		"malformed-header": {
			input: http.Header{"Authorization": []string{"hallo"}},
			want: response{
				apiKey:    "",
				errString: "malformed authorization header",
			},
		},
		"good-header": {
			input: http.Header{"Authorization": []string{"ApiKey hallo"}},
			want: response{
				apiKey:    "hallo",
				errString: "",
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			apiKey, err := GetAPIKey(tc.input)
			errString := ""
			if err != nil {
				errString = err.Error()
			}
			got := response{apiKey: apiKey, errString: errString}
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("expected: %#v, got: %#v", tc.want, got)
			}
		})
	}
}
