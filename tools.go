//go:build tools
// +build tools

package tools

import (
	_ "github.com/99designs/gqlgen"
	_ "github.com/kyleconroy/sqlc/cmd/sqlc"
	_ "github.com/vektah/dataloaden"
)
