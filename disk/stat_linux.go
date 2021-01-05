// +build linux,!s390x,!arm,!386

package disk

import (
	"fmt"
	"syscall"
)

// GetInfo returns total and free bytes available in a directory, e.g. `/`.
func GetInfo(path string) (info Info, err error) {
	s := syscall.Statfs_t{}
	err = syscall.Statfs(path, &s)
	if err != nil {
		return Info{}, err
	}
	reservedBlocks := s.Bfree - s.Bavail
	info = Info{
		Total:  uint64(s.Frsize) * (s.Blocks - reservedBlocks),
		Free:   uint64(s.Frsize) * s.Bavail,
		Files:  s.Files,
		Ffree:  s.Ffree,
		FSType: getFSType(s.Type),
	}

	// XFS can show wrong values at times error out
	// in such scenarios.
	if info.Free > info.Total {
		return info, fmt.Errorf("detected free space (%d) > total disk space (%d), fs corruption at (%s). please run 'fsck'", info.Free, info.Total, path)
	}
	info.Used = info.Total - info.Free
	return info, nil
}
