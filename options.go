package fs

import "time"

type Option func(m Manager)

func WithFilerUrl(url string) Option {
	return func(m Manager) {
		m.SetFilerUrl(url)
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(m Manager) {
		m.SetTimeout(timeout)
	}
}
