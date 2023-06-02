package tcp_test

import (
	"context"
	"testing"

	tcping "github.com/wujiu2020/tcpping/ping"
	"github.com/wujiu2020/tcpping/ping/tcp"
)

func TestPing(t *testing.T) {
	ping := tcp.New("google.com", 80, &tcping.Option{}, false)
	stats := ping.Ping(context.Background())
	if !stats.Connected {
		t.Fatalf("ping failed, %s", stats.Error)
	}
}

func TestPing_Failed(t *testing.T) {

	ping := tcp.New("127.0.0.1", 1, &tcping.Option{}, false)
	stats := ping.Ping(context.Background())
	if stats.Connected {
		t.Fatalf("it should be connected refused error")
	}
}
