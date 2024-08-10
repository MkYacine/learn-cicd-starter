package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name          string
		headers       http.Header
		expectedKey   string
		expectedError error
	}{
		{
			name:          "Valid API Key",
			headers:       http.Header{"Authorization": []string{"ApiKey validkey123"}},
			expectedKey:   "validkey123",
			expectedError: nil,
		},
		{
			name:          "No Authorization Header",
			headers:       http.Header{},
			expectedKey:   "",
			expectedError: ErrNoAuthHeaderIncluded,
		},
		{
			name:          "Malformed Authorization Header",
			headers:       http.Header{"Authorization": []string{"Bearer invalidkey"}},
			expectedKey:   "",
			expectedError: errors.New("malformed authorization header"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotKey, gotErr := GetAPIKey(tt.headers)

			if gotKey != tt.expectedKey {
				t.Errorf("GetAPIKey() gotKey = %v, want %v", gotKey, tt.expectedKey)
			}

			if (gotErr == nil && tt.expectedError != nil) || (gotErr != nil && tt.expectedError == nil) {
				t.Errorf("GetAPIKey() gotErr = %v, want %v", gotErr, tt.expectedError)
			} else if gotErr != nil && tt.expectedError != nil && gotErr.Error() != tt.expectedError.Error() {
				t.Errorf("GetAPIKey() gotErr = %v, want %v", gotErr, tt.expectedError)
			}
		})
	}
}
