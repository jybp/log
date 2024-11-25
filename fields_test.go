package log_test

import (
	"bytes"
	"context"
	"errors"
	"testing"

	"github.com/jybp/log"
	"github.com/sirupsen/logrus"
)

func TestLog(t *testing.T) {
	buf := &bytes.Buffer{}
	logrus.SetOutput(buf)
	logrus.SetFormatter(&logrus.TextFormatter{DisableTimestamp: true})

	ctx := context.Background()
	ctx = log.CtxWithFields(ctx, log.Fields{"k1": "v1"})
	log.WithFields(log.Fields{"k2": "v2"}).InfoC(ctx, "info")

	log.WithField("k3", "v3").WarnfC(ctx, "warn%d", 1)

	log.WithError(errors.New("err")).Error("msg")

	expected := `level=info msg=info k1=v1 k2=v2
level=warning msg=warn1 k1=v1 k3=v3
level=error msg=msg error=err
`
	if buf.String() != expected {
		t.Fatalf("unexpected output:\n%s\nexpected:\n%s", buf.String(), expected)
	}
}
