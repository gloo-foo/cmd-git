package command

import (
	gloo "github.com/gloo-foo/framework"
	"github.com/gloo-foo/framework/patterns"
)

// runner forks an external process named name with args, returning the streaming
// Command that drives it. patterns.Subprocess is the production runner; tests
// inject a fake to assert the exact argument vector without forking git.
//
// args is []string because it flows verbatim into patterns.Subprocess (and
// ultimately os/exec), whose signature this wrapper does not control.
type runner func(name string, args ...string) gloo.Command[[]byte, []byte]

// Git returns a Command that forks git with the given subcommand and arguments,
// streaming pipeline input to git's stdin and git's stdout back into the
// pipeline. Every argument is passed through verbatim — git, not this wrapper,
// interprets them.
func Git(args ...string) gloo.Command[[]byte, []byte] {
	return gitWith(patterns.Subprocess, args...)
}

// gitWith is Git with an injectable runner. It prepends "git" to the argument
// vector and hands it to run; the seam lets tests prove the exact vector
// without executing git.
func gitWith(run runner, args ...string) gloo.Command[[]byte, []byte] {
	return run("git", args...)
}
