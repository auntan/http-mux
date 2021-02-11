package http_server

import "testing"

func Test_ValidationHappyPath(t *testing.T) {
	t.Parallel()

	tt := []struct {
		name  string
		in    []string
		error bool
	}{
		{
			name:  "empty request",
			error: true,
		},
		{
			name:  "too large",
			in:    []string{"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""},
			error: true,
		},
		{
			name:  "happy path",
			in:    []string{"http://url.com"},
			error: false,
		},
	}

	for i := range tt {
		tc := tt[i]

		t.Run(tc.name, func(t *testing.T) {
			err := validate(tc.in)
			if tc.error != (err != nil) {
				t.Fatalf("test failed")
			}
		})
	}
}
