package disk

// Info stat fs struct is container which holds following values
// Total - total size of the volume / disk
// Free - free size of the volume / disk
// Files - total inodes available
// Ffree - free inodes available
// FSType - file system type
type Info struct {
	Total  uint64
	Free   uint64
	Used   uint64
	Files  uint64
	Ffree  uint64
	FSType string
}

// SameDisk reports whether di1 and di2 describe the same disk.
func SameDisk(di1, di2 Info) bool {
	if di1.Total != di2.Total {
		// disk total size different
		return false
	}

	if di1.Files != di2.Files {
		// disk total inodes different
		return false
	}

	// returns true only if Used, Free are same, then its the same disk.
	// we are deliberately not using free inodes as that is unreliable
	// due the fact that Ffree can vary even for temporary files
	return di1.Used == di2.Used && di1.Free == di2.Free
}
