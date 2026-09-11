// Package progress persists which lessons the learner has completed
// to ~/.go-learn/progress.json so it survives across program restarts.
package progress

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// Progress is the on-disk shape of ~/.go-learn/progress.json.
type Progress struct {
	Completed []string `json:"completed"`
}

// Dir returns the directory go-learn stores its state in, creating it
// if it does not exist yet.
func Dir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	dir := filepath.Join(home, ".go-learn")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return "", err
	}
	return dir, nil
}

func filePath() (string, error) {
	dir, err := Dir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, "progress.json"), nil
}

// Load reads progress.json, returning an empty Progress if it doesn't
// exist yet (first run).
func Load() (*Progress, error) {
	path, err := filePath()
	if err != nil {
		return nil, err
	}
	data, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return &Progress{Completed: []string{}}, nil
	}
	if err != nil {
		return nil, err
	}
	var p Progress
	if err := json.Unmarshal(data, &p); err != nil {
		return nil, err
	}
	return &p, nil
}

// Save writes the progress back to disk.
func (p *Progress) Save() error {
	path, err := filePath()
	if err != nil {
		return err
	}
	data, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0o644)
}

// IsCompleted reports whether the given lesson id was already finished.
func (p *Progress) IsCompleted(id string) bool {
	for _, c := range p.Completed {
		if c == id {
			return true
		}
	}
	return false
}

// MarkCompleted adds id to the completed list (no-op if already present)
// and saves progress to disk.
func (p *Progress) MarkCompleted(id string) error {
	if p.IsCompleted(id) {
		return nil
	}
	p.Completed = append(p.Completed, id)
	return p.Save()
}

// Reset clears all progress and saves it.
func (p *Progress) Reset() error {
	p.Completed = []string{}
	return p.Save()
}
