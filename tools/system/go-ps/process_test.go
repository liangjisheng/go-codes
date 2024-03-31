package ps

import (
	"github.com/mitchellh/go-ps"
	"os"
	"testing"
)

func TestFindProcess(t *testing.T) {
	p, err := ps.FindProcess(os.Getpid())
	if err != nil {
		t.Fatalf("err: %s", err)
	}
	if p == nil {
		t.Fatal("should have process")
	}

	if p.Pid() != os.Getpid() {
		t.Fatalf("bad: %#v", p.Pid())
	}
}

func TestProcesses(t *testing.T) {
	// This test works because there will always be SOME processes
	// running.
	p, err := ps.Processes()
	if err != nil {
		t.Fatalf("err: %s", err)
	}

	if len(p) <= 0 {
		t.Fatal("should have processes")
	}

	found := false
	for _, p1 := range p {
		if p1.Executable() == "go" || p1.Executable() == "go.exe" {
			found = true
			break
		}
	}

	if !found {
		t.Fatal("should have Go")
	}
}
