package git_test

import (
	git "github.com/gloo-foo/cmd-git/alias"
)

// ExampleGit shows how to construct a git subcommand as a composable Command.
// The subprocess is not executed here — running it would require a working git
// install and a non-deterministic repository state — so the constructed Command
// is simply discarded.
func ExampleGit() {
	cmd := git.Git("status", "--short")
	_ = cmd
}
