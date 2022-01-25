// +build tools

package tools

import (
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/google/ko"
	_ "go.sbr.pm/ram"
	_ "golang.org/x/lint/golint"
	_ "github.com/gordonklaus/ineffassign"
	_ "honnef.co/go/tools/cmd/staticcheck"
	_ "github.com/mibk/dupl"
	_ "github.com/kisielk/errcheck"
	_ "golang.org/x/tools/cmd/goimports"
)
