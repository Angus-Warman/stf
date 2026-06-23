package stf

import (
	"os"
	"path/filepath"
	"testing"
)

func TestSafeJoin(t *testing.T) {
	dir := t.TempDir()

	t.Run("simple join", func(t *testing.T) {
		got, err := SafeJoin(dir, "sub/file.txt")
		if err != nil {
			t.Fatal(err)
		}
		want := filepath.Join(dir, "sub/file.txt")
		if got != want {
			t.Errorf("SafeJoin = %q, want %q", got, want)
		}
	})

	t.Run("empty untrusted", func(t *testing.T) {
		got, err := SafeJoin(dir, "")
		if err != nil {
			t.Fatal(err)
		}
		if got != dir {
			t.Errorf("SafeJoin = %q, want %q", got, dir)
		}
	})

	t.Run("dot untrusted", func(t *testing.T) {
		got, err := SafeJoin(dir, ".")
		if err != nil {
			t.Fatal(err)
		}
		if got != dir {
			t.Errorf("SafeJoin = %q, want %q", got, dir)
		}
	})

	t.Run("path traversal", func(t *testing.T) {
		_, err := SafeJoin(dir, "../../etc/passwd")
		if err == nil {
			t.Fatal("expected error for path traversal")
		}
	})

	t.Run("symlink traversal", func(t *testing.T) {
		link := filepath.Join(dir, "evil")
		err := os.Symlink("/etc", link)
		if err != nil {
			t.Fatal(err)
		}
		_, err = SafeJoin(dir, "evil/passwd")
		if err == nil {
			t.Fatal("expected error for symlink traversal")
		}
	})

	t.Run("symlink in root", func(t *testing.T) {
		realDir := filepath.Join(dir, "real")
		err := os.Mkdir(realDir, 0755)
		if err != nil {
			t.Fatal(err)
		}
		linkDir := filepath.Join(dir, "link")
		err = os.Symlink(realDir, linkDir)
		if err != nil {
			t.Fatal(err)
		}
		got, err := SafeJoin(linkDir, "sub/file.txt")
		if err != nil {
			t.Fatal(err)
		}
		want := filepath.Join(realDir, "sub/file.txt")
		if got != want {
			t.Errorf("SafeJoin = %q, want %q", got, want)
		}
	})

	t.Run("non-existent valid path", func(t *testing.T) {
		got, err := SafeJoin(dir, "does/not/exist")
		if err != nil {
			t.Fatal(err)
		}
		want := filepath.Join(dir, "does/not/exist")
		if got != want {
			t.Errorf("SafeJoin = %q, want %q", got, want)
		}
	})

	t.Run("existing file", func(t *testing.T) {
		sub := filepath.Join(dir, "sub")
		err := os.Mkdir(sub, 0755)
		if err != nil {
			t.Fatal(err)
		}
		f := filepath.Join(sub, "file.txt")
		err = os.WriteFile(f, []byte("hello"), 0644)
		if err != nil {
			t.Fatal(err)
		}
		got, err := SafeJoin(dir, "sub/file.txt")
		if err != nil {
			t.Fatal(err)
		}
		if got != f {
			t.Errorf("SafeJoin = %q, want %q", got, f)
		}
	})
}
