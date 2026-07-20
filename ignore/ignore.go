package ignore

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"

	gitignore "github.com/monochromegane/go-gitignore"
)

type Ignore struct {
	path    string
	ignores []gitignore.IgnoreMatcher
}

func NewIgnore(path string) Ignore {
	ignoreFiles := []string{".gitignore", ".ignore", ".ldignore"}
	ignores := make([]gitignore.IgnoreMatcher, 0, len(ignoreFiles))
	for _, ignoreFile := range ignoreFiles {
		lines, err := readLines(filepath.Join(path, ignoreFile))
		if err != nil {
			continue
		}
		// go-gitignore treats ".*"-style patterns as matching every path,
		// which would ignore every file in the diff. The diff scanner
		// already skips dotfile paths, so the patterns are safe to drop.
		cleaned := make([]string, 0, len(lines))
		for _, line := range lines {
			if trimmed := strings.TrimSpace(line); trimmed == ".*" || trimmed == ".*/" {
				continue
			}
			cleaned = append(cleaned, line)
		}
		ignores = append(ignores, gitignore.NewGitIgnoreFromReader(path, strings.NewReader(strings.Join(cleaned, "\n"))))
	}
	return Ignore{path: path, ignores: ignores}
}

func readLines(path string) ([]string, error) {
	/* #nosec */
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func (m Ignore) Match(path string, isDir bool) bool {
	for _, i := range m.ignores {
		if i.Match(path, isDir) {
			return true
		}
	}

	return false
}
