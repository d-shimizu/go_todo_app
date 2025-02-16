package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"testing"

	"golang.org/x/sync/errgroup"
)

func TestRun(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	eg, ctx := errgroup.WithContext(ctx)
	eg.Go(func() error {
		return run(ctx)
	})

	in := "message"

	rsp, err := http.Get("http://localhost:18080/" + in)
	if err != nil {
		t.Errorf("failed to get response: %v", err)
	}
	defer rsp.Body.Close()

	got, err := io.ReadAll(rsp.Body)
	if err != nil {
		t.Errorf("failed to read response body: %v", err)
	}

	// HTTP サーバーのレスポンスが正しいかどうかを検証
	want := fmt.Sprintf("Hello, %s!", in)
	if string(got) != want {
		t.Errorf("got %s, want %s", got, want)
	}

	// run 関数に終了通知を送信する
	cancel()

	// run 関数の戻り値を検証する
	if err := eg.Wait(); err != nil {
		t.Fatal(err)
	}
}
