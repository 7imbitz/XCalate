package utils

import (
	"os"
)

// World-Readable func
/*func IsWorldReadable(filePath string) (bool, error) {
	// Get file information
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return false, err
	}

	// Extract permission bits
	permissions := fileInfo.Mode().Perm()

	// Check if others (world) have read permission
	worldReadable := permissions&(1<<2) != 0

	return worldReadable, nil
}

// World-Writable func
func IsWorldWritable(filePath string) (bool, error) {
	// Get file information
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return false, err
	}

	// Extract permission bits
	permissions := fileInfo.Mode().Perm()

	// Check if others (world) have write permission
	worldWritable := permissions&(1<<1) != 0

	return worldWritable, nil
}*/

// checkWorldPermission checks if the world (others) have the specified permission bit.
func checkWorldPermission(filePath string, permBit os.FileMode) (bool, error) {
	info, err := os.Stat(filePath)
	if err != nil {
		return false, err
	}
	return info.Mode().Perm()&permBit != 0, nil
}

// IsWorldReadable returns true if 'others' have read permission.
func IsWorldReadable(filePath string) (bool, error) {
	return checkWorldPermission(filePath, 0o004) // 0o004 is octal for "others read" (same as 1 << 2)
}

// IsWorldWritable returns true if 'others' have write permission.
func IsWorldWritable(filePath string) (bool, error) {
	return checkWorldPermission(filePath, 0o002) // 0o002 is octal for "others write" (same as 1 << 1)
}
