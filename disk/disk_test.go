package disk_test

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/mnadeem/volume_exporter/disk"
)

func TestFree(t *testing.T) {
	path, err := ioutil.TempDir(os.TempDir(), "disk-")
	defer os.RemoveAll(path)
	if err != nil {
		t.Fatal(err)
	}

	di, err := disk.GetInfo(path)
	if err != nil {
		t.Fatal(err)
	}

	if di.FSType == "UNKNOWN" {
		t.Error("Unexpected FSType", di.FSType)
	}
}
