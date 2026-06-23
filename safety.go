package stf

import (
	"errors"
	"fmt"
	"io/fs"
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

	resolved, exists, err := resolveSymlinkIfExists(target)

	if err != nil {
		return "", err
	}

	if exists {
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

func resolveSymlinkIfExists(path string) (resolved string, exists bool, err error) {
	resolved, err = filepath.EvalSymlinks(path)

	if err == nil {
		return resolved, true, nil
	}

	if errors.Is(err, fs.ErrNotExist) {
		return path, false, nil
	}

	return "", false, err
}
