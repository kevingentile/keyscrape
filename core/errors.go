package core

import "errors"

// ErrNotExist used when a path does not exist
var ErrNotExist = errors.New("keyscrape: path not found")
