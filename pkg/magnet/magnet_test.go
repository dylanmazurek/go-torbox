package magnet

import (
	"testing"
)

func TestNewMagnet(t *testing.T) {
	tests := []struct {
		name        string
		link        string
		wantErr     bool
		expectedErr string
		expected    *Magnet
	}{
		{
			name:        "invalid URL",
			link:        "not-a-url",
			wantErr:     true,
			expectedErr: "",
		},
		{
			name:        "not a magnet scheme",
			link:        "http://example.com",
			wantErr:     true,
			expectedErr: "invalid magnet link: not a magnet scheme",
		},
		{
			name:        "missing xt parameter",
			link:        "magnet:?dn=test",
			wantErr:     true,
			expectedErr: "invalid magnet link: missing or malformed xt parameter",
		},
		{
			name:     "valid magnet with display name",
			link:     "magnet:?xt=urn:btih:1234567890abcdef&dn=test-file",
			wantErr:  false,
			expected: &Magnet{Hash: "1234567890abcdef", DisplayName: "test-file", Trackers: nil},
		},
		{
			name:     "valid magnet without display name",
			link:     "magnet:?xt=urn:btih:1234567890abcdef",
			wantErr:  false,
			expected: &Magnet{Hash: "1234567890abcdef", DisplayName: "unknown", Trackers: nil},
		},
		{
			name:     "valid magnet with empty display name",
			link:     "magnet:?xt=urn:btih:1234567890abcdef&dn=",
			wantErr:  false,
			expected: &Magnet{Hash: "1234567890abcdef", DisplayName: "unknown", Trackers: nil},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := NewMagnet(tt.link)

			if tt.wantErr {
				if err == nil {
					t.Errorf("NewMagnet() expected error but got none")
				}
				if tt.expectedErr != "" && err.Error() != tt.expectedErr {
					t.Errorf("NewMagnet() error = %v, expected %v", err.Error(), tt.expectedErr)
				}
				return
			}

			if err != nil {
				t.Errorf("NewMagnet() unexpected error = %v", err)
				return
			}

			if result.Hash != tt.expected.Hash {
				t.Errorf("NewMagnet() Hash = %v, expected %v", result.Hash, tt.expected.Hash)
			}

			if result.DisplayName != tt.expected.DisplayName {
				t.Errorf("NewMagnet() DisplayName = %v, expected %v", result.DisplayName, tt.expected.DisplayName)
			}
		})
	}
}
