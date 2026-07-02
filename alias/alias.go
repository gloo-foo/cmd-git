// Package alias provides the unprefixed name for the git command.
//
//	import git "github.com/gloo-foo/cmd-git/alias"
//	git.Git("status", "--short")
package alias

import (
	gloo "github.com/gloo-foo/framework"

	command "github.com/gloo-foo/cmd-git"
)

// Git re-exports the constructor by delegation, preserving its exact signature.
func Git(args ...command.GitArg) gloo.Command[[]byte, []byte] { return command.Git(args...) }
