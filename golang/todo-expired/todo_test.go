package main

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestIsExpired(t *testing.T) {
	t.Parallel()

	currentTime := "2020-12-04 12:23:34"

	tests := []struct {
		title     string
		expected  bool
		expiredAt string
	}{
		{"現在時刻が期限を過ぎてるので期限切れ", true, "2020-12-02 12:23:34"},
		{"現在時刻が期限を過ぎていない", false, "2020-12-12 12:23:34"},
		{"現在時刻が期限と同一時刻", false, currentTime},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			t.Parallel()

			ctx, def := setUp(t, currentTime)
			defer def()

			expiredAt, err := Parse(tt.expiredAt)
			assert.Nil(t, err)

			todo := &Todo{
				ID:        uuid.New().String(),
				Name:      "test todo",
				ExpiredAt: expiredAt,
			}
			actual, err := todo.IsExpired(ctx)
			assert.Nil(t, err)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
