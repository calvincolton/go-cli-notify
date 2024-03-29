//go:build integration
// +build integration

package notify_test

import (
	"testing"

	notify "github.com/calvincolton/go-cli-notify"
)

func TestSend(t *testing.T) {
	n := notify.New("test title", "test msg", notify.SeverityNormal)

	err := n.Send()
	if err != nil {
		t.Error(err)
	}
}
