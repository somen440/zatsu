package main

import (
	"context"
	"time"
)

// Todo entity
type Todo struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	ExpiredAt time.Time `json:"expired_at"`
}

// IsExpired 期限ぎれかチェック。期限切れ時刻が現在時刻を過ぎてたら true
func (t *Todo) IsExpired(ctx context.Context) (bool, error) {
	uid := GetUID(ctx)
	n, err := Time().Now(uid)
	if err != nil {
		return false, err
	}
	return (n.Unix() < t.ExpiredAt.Unix()), nil
}
