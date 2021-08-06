package table

import (
	"testing"

	"github.com/yandex-cloud/ydb-go-sdk/internal/tracetest"
)

func TestClientTrace(t *testing.T) {
	tracetest.TestSingleTrace(t, ClientTrace{}, "ClientTrace")
}

func TestRetryTrace(t *testing.T) {
	tracetest.TestSingleTrace(t, RetryTrace{}, "RetryTrace")
}

func TestSessionPoolTrace(t *testing.T) {
	tracetest.TestSingleTrace(t, SessionPoolTrace{}, "SessionPoolTrace")
}
