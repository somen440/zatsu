package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/google/uuid"
)

func TestMain(m *testing.M) {
	InitializeTimeUtil()

	code := m.Run()
	if code != 0 {
		panic(fmt.Errorf("test: exit %d", code))
	}
}

func setUp(t *testing.T, currentTime string) (context.Context, func()) {
	uid := uuid.New().String()
	ctx := SetUID(context.Background(), uid)

	tu := Time()
	if err := tu.Start(uid, currentTime); err != nil {
		t.FailNow()
	}

	return ctx, func() {
		tu.Stop(uid)
	}
}
