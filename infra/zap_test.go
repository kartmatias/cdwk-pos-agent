package infra_test

import (
	"os"
	"testing"

	"github.com/kartmatias/cdwk-pos-agent/infra"
)

func TestZapLogger(t *testing.T) {

	s := infra.ZAPLOGDIR
	fi, err := os.Stat(s)

	if err != err {
		t.Errorf("%s does not exist!\n", s)
	}
	if !fi.IsDir() {
		t.Errorf("%s does not exist directory!\n", s)
	}

	logger := infra.SetupZapLogger()
	if logger == nil {
		t.Error("Error expected logger instance")
	}
}
