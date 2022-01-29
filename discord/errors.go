package discord

import (
	"golang.org/x/xerrors"
)

var (
	ErrUnauthorized = xerrors.New("Inproper token was passed")
)
