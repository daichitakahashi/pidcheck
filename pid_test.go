package pidcheck

import (
	"os"
	"testing"
)

func Test(t *testing.T) {
	t.Log("PID:", os.Getpid())
	t.Log("PID of parent:", os.Getppid())
}
