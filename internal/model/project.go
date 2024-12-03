package model

import (
	"time"
)

type Project struct {
	Name         string    `json:"name"`
	Created      time.Time `json:"created"`
	LastAccessed time.Time `json:"lastAccessed"`
	Template     string    `json:"template"`
	Tags         []string  `json:"tags"`
	Notes        string    `json:"notes"`
	Path         string    `json:"-"`
}

type Template struct {
	Name         string
	Description  string
	Files        map[string]string
	Dependencies []string
}
