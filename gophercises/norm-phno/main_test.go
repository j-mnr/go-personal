package main

import "testing"

type normalizeTestCase struct {
	input string
	want  string
}

func TestNormalize(t *testing.T) {
	tt := []normalizeTestCase{
		{"1234567890", "1234567890"},
		{"123 456 7891", "1234567891"},
		{"(123) 456 7892", "1234567892"},
		{"(123) 456-7893", "1234567893"},
		{"123-456-7894", "1234567894"},
		{"(123)456-7892", "1234567892"},
	}

	for _, tc := range tt {
		t.Run(tc.input, func(t *testing.T) {
			got := normalize(tc.input)
			if got != tc.want {
				t.Errorf("got %s; want %s", got, tc.want)
			}
		})
	}
}

func BenchmarkNormalize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		normalize("(123) 456 7892")
	}
}
