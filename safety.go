package stf

import (
	"fmt"
	"path/filepath"
	"strings"
)

// SafeJoin resolves untrusted to an absolute path within root,
// rejecting path traversal and symlink-based escapes.
func SafeJoin(root, untrusted string) (string, error) {
	root, err := filepath.Abs(root)

	if err != nil {
		return "", fmt.Errorf("resolving root path: %w", err)
	}

	root, err = filepath.EvalSymlinks(root)

	if err != nil {
		return "", fmt.Errorf("resolving root symlinks: %w", err)
	}

	target := filepath.Join(root, untrusted)
	target = filepath.Clean(target)

	rootWithSep := root + string(filepath.Separator)

	isWithinRoot := func(path string) bool {
		return path == root || strings.HasPrefix(path, rootWithSep)
	}

	resolved, err := filepath.EvalSymlinks(target)

	isSymlink := err == nil

	if isSymlink {
		if isWithinRoot(resolved) {
			return resolved, nil
		}

		return "", fmt.Errorf("path traversal via symlink: %q resolves to %q outside root %q", target, resolved, root)
	}

	if isWithinRoot(target) {
		return target, nil
	}

	return "", fmt.Errorf("path traversal: %q is outside root %q", target, root)

}
