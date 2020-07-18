package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNow(t *testing.T) {
	t.Parallel()

	tests := []struct {
		expected string
	}{
		{"2020-02-02 11:22:33"},
		{"2020-02-12 11:22:33"},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("test_%d", i), func(t *testing.T) {
			t.Parallel()

			ctx, def := setUp(t, tt.expected)
			defer def()
			uid := GetUID(ctx)
			n, err := Time().Now(uid)
			assert.Nil(t, err)
			assert.Equal(t, tt.expected, Format(n))
		})
	}
}
