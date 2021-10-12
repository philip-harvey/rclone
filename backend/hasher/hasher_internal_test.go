package hasher

import (
	"testing"

	"github.com/rclone/rclone/fstest/fstests"
	"github.com/rclone/rclone/lib/kv"
)

func (f *Fs) testHasherFeature(t *testing.T) {
	t.Skip("Nothing to do")
}

// InternalTest dispatches all internal tests
func (f *Fs) InternalTest(t *testing.T) {
	if !kv.Supported() {
		t.Skip("hasher is not supported on this OS")
	}
	t.Run("TestHasherFeature", f.testHasherFeature)
}

var _ fstests.InternalTester = (*Fs)(nil)
