package utils

import (
	"os"
)

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
