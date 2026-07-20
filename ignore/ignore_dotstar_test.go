package ignore

import (
	"os"
	"path/filepath"
	"testing"
)

func TestNewIgnorePreservesDotStarPatterns(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, ".gitignore"), []byte(".*\n.*/\nnode_modules/\n"), 0600); err != nil {
		t.Fatal(err)
	}

	ig := NewIgnore(dir)

	if ig.Match(filepath.Join(dir, "src/app.ts"), false) {
		t.Error("regular file should not be ignored by dotfile patterns")
	}
	if !ig.Match(filepath.Join(dir, "src/.env"), false) {
		t.Error("nested dotfile should be ignored when .gitignore contains .*")
	}
	if !ig.Match(filepath.Join(dir, "node_modules"), true) {
		t.Error("node_modules should still be ignored")
	}
}
