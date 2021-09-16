package version

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewVersion(t *testing.T) {
	tests := []struct {
		name    string
		version string
		wantErr bool
	}{
		{
			name:    "test 1.2.3",
			version: "1.2.3",
		},
		{
			name:    "test 1.2.3-4",
			version: "1.2.3-4",
		},
		{
			name:    "test 1:1.2.3-4",
			version: "1:1.2.3-4",
		},
		{
			name:    "test A:1.2.3-4",
			version: "A:1.2.3-4",
			wantErr: true,
		},
		{
			name:    "test -1:1.2.3-4",
			version: "-1:1.2.3-4",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewVersion(tt.version)
			if (err != nil) != tt.wantErr {
				assert.Equal(t, tt.wantErr, err)
			}
		})
	}
}

func TestVersion_Valid(t *testing.T) {
	cases := []struct {
		version  string
		expected bool
	}{
		{"1.2.3", true},
		{"1:1.2.3", true},
		{"A:1.2.3", false},
		{"1:1.2.3-1", true},
	}

	for _, tc := range cases {
		actual := Valid(tc.version)
		if actual != tc.expected {
			t.Fatalf(
				"valid: %s\nexpected: %t\nactual: %t",
				tc.version, tc.expected, actual)
		}
	}
}

func TestVersion_Compare(t *testing.T) {
	for _, tt := range cases {
		t.Run(tt.v1+" vs "+tt.v2, func(t *testing.T) {
			a, _ := NewVersion(tt.v1)
			b, _ := NewVersion(tt.v2)

			got := a.Compare(b)

			var expected int
			switch tt.expected {
			case ">":
				expected = GREATER
			case "<":
				expected = LESS
			case "=":
				expected = EQUAL
			default:
				require.Fail(t, "unknown symbol: %s", tt.expected)
			}

			assert.Equal(t, expected, got)
		})
	}
}

func TestVersion_Equal(t *testing.T) {
	for _, tt := range cases {
		t.Run(tt.v1+" is equal to "+tt.v2, func(t *testing.T) {
			a, _ := NewVersion(tt.v1)
			b, _ := NewVersion(tt.v2)

			got := a.Equal(b)

			var expected bool
			switch tt.expected {
			case ">", "<":
				expected = false
			case "=":
				expected = true
			default:
				require.Fail(t, "unknown symbol: %s", tt.expected)
			}

			assert.Equal(t, expected, got)
		})
	}
}

func TestVersion_GreaterThan(t *testing.T) {
	for _, tt := range cases {
		t.Run(tt.v1+" is greater than "+tt.v2, func(t *testing.T) {
			a, _ := NewVersion(tt.v1)
			b, _ := NewVersion(tt.v2)

			got := a.GreaterThan(b)

			var expected bool
			switch tt.expected {
			case "<", "=":
				expected = false
			case ">":
				expected = true
			default:
				require.Fail(t, "unknown symbol: %s", tt.expected)
			}

			assert.Equal(t, expected, got)
		})
	}
}

func TestVersion_LessThan(t *testing.T) {
	for _, tt := range cases {
		t.Run(tt.v1+" is less than "+tt.v2, func(t *testing.T) {
			a, _ := NewVersion(tt.v1)
			b, _ := NewVersion(tt.v2)

			got := a.LessThan(b)

			var expected bool
			switch tt.expected {
			case ">", "=":
				expected = false
			case "<":
				expected = true
			default:
				require.Fail(t, "unknown symbol: %s", tt.expected)
			}

			assert.Equal(t, expected, got)
		})
	}
}
