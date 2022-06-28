package storage

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"io"
	"telegramBot/lib/e"
)

type Storage interface {
	Save(page *Page) error
	PickRandom(Name string) (*Page, error)
	Remove(p *Page) error
	IsExists(p *Page) (bool, error)
}

type Page struct {
	URL  string
	Name string
}

var ErrNoSavedPages = errors.New("no saved pages")

func (p Page) Hash() (string, error) {
	h := sha1.New()

	if _, err := io.WriteString(h, p.URL); err != nil {
		return "", e.Wrap("can't calculate hash", err)
	}

	if _, err := io.WriteString(h, p.Name); err != nil {
		return "", e.Wrap("can't calculate hash", err)
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
