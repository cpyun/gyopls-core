package config

import (
	"crypto/md5"
	"errors"
	"fmt"
	"time"
)

var (
	// ErrWatcherStopped is returned when source watcher has been stopped
	ErrWatcherStopped = errors.New("watcher stopped")
	ErrNotFound       = errors.New("key not found") // ErrNotFound is key not found.
)

// ChangeSet represents a set of changes from a source
type ChangeSet struct {
	Data      []byte
	Checksum  string
	Format    string
	Source    string
	Timestamp time.Time
}

// Sum returns the md5 checksum of the ChangeSet data
func (c *ChangeSet) Sum() string {
	h := md5.New()
	h.Write(c.Data)
	return fmt.Sprintf("%x", h.Sum(nil))
}

type Source interface {
	Load() (*ChangeSet, error)
	Watch() (Watcher, error)
	String() string
}

// Watcher watches a source for changes
type Watcher interface {
	Next() (*ChangeSet, error)
	Stop() error
}
