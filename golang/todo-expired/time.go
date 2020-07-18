package main

import (
	"fmt"
	"sync"
	"time"
)

// Layouts Ymd ってなんだっけ ... ってなるので
const (
	YmdLayout = "2006-01-02 15:04:05"
)

// エラー作成共通化
func errNotFound(uid string) error {
	return fmt.Errorf("time: not found uid=%s", uid)
}

// TimeUtil 時間操作系
type TimeUtil struct {
	times map[string]time.Time // uid 毎の time を持つ

	cond *sync.Cond // 共有したリソースが操作されるため
}

// DefaultTimeUtil TimeUtil のデフォルト
var DefaultTimeUtil *TimeUtil

// InitializeTimeUtil DefaultTimeUtil の初期化
func InitializeTimeUtil() {
	DefaultTimeUtil = &TimeUtil{
		times: map[string]time.Time{},
		cond:  sync.NewCond(&sync.Mutex{}),
	}
}

// Time DefaultTimeUtil のエイリアス
func Time() *TimeUtil {
	return DefaultTimeUtil
}

// Start UIDに対する時間の開始。 format は Ymd 形式。format が空なら現在時刻を設定。
func (tu *TimeUtil) Start(uid string, format string) (err error) {
	tu.cond.L.Lock()
	defer tu.cond.L.Unlock()

	var t time.Time
	if format == "" {
		t = time.Now()
	} else {
		t, err = Parse(format)
		if err != nil {
			return
		}
	}

	tu.times[uid] = t
	return nil
}

// Stop UIDに対する時間停止。UIDに対応した時間が既にないなら特に何もしない
func (tu *TimeUtil) Stop(uid string) {
	tu.cond.L.Lock()
	defer tu.cond.L.Unlock()
	_, ok := tu.times[uid]
	if !ok {
		return
	}
	delete(tu.times, uid)
}

// Now UIDに対する現在時刻（Startした時の値）を返す
func (tu *TimeUtil) Now(uid string) (time.Time, error) {
	t, ok := tu.times[uid]
	if !ok {
		return time.Time{}, errNotFound(uid)
	}
	return t, nil
}

// Format time.Time を Ymd 形式で返す
func Format(t time.Time) string {
	return t.Format(YmdLayout)
}

// Parse Ymd 形式から time.Time へパース
func Parse(ymd string) (time.Time, error) {
	return time.Parse(YmdLayout, ymd)
}
