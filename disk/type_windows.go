// +build windows

package disk

import (
	"path/filepath"
	"syscall"
	"unsafe"
)

var (
	// GetVolumeInformation provides windows drive volume information.
	GetVolumeInformation = kernel32.NewProc("GetVolumeInformationW")
)

// getFSType returns the filesystem type of the underlying mounted filesystem
func getFSType(path string) string {
	volumeNameSize, nFileSystemNameSize := uint32(260), uint32(260)
	var lpVolumeSerialNumber uint32
	var lpFileSystemFlags, lpMaximumComponentLength uint32
	var lpFileSystemNameBuffer, volumeName [260]uint16
	var ps = syscall.StringToUTF16Ptr(filepath.VolumeName(path))

	// Extract values safely
	// BOOL WINAPI GetVolumeInformation(
	// _In_opt_  LPCTSTR lpRootPathName,
	// _Out_opt_ LPTSTR  lpVolumeNameBuffer,
	// _In_      DWORD   nVolumeNameSize,
	// _Out_opt_ LPDWORD lpVolumeSerialNumber,
	// _Out_opt_ LPDWORD lpMaximumComponentLength,
	// _Out_opt_ LPDWORD lpFileSystemFlags,
	// _Out_opt_ LPTSTR  lpFileSystemNameBuffer,
	// _In_      DWORD   nFileSystemNameSize
	// );

	_, _, _ = GetVolumeInformation.Call(uintptr(unsafe.Pointer(ps)),
		uintptr(unsafe.Pointer(&volumeName)),
		uintptr(volumeNameSize),
		uintptr(unsafe.Pointer(&lpVolumeSerialNumber)),
		uintptr(unsafe.Pointer(&lpMaximumComponentLength)),
		uintptr(unsafe.Pointer(&lpFileSystemFlags)),
		uintptr(unsafe.Pointer(&lpFileSystemNameBuffer)),
		uintptr(nFileSystemNameSize))

	return syscall.UTF16ToString(lpFileSystemNameBuffer[:])
}
