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
