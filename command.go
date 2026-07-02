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

// GitArg is one verbatim token of the git command line: the subcommand or any
// of its arguments. Git — not this wrapper — interprets it.
type GitArg string

// Git returns a Command that forks git with the given subcommand and arguments,
// streaming pipeline input to git's stdin and git's stdout back into the
// pipeline. Every argument is passed through verbatim — git, not this wrapper,
// interprets them.
func Git(args ...GitArg) gloo.Command[[]byte, []byte] {
	return gitWith(patterns.Subprocess, args...)
}

// gitWith is Git with an injectable runner. It prepends "git" to the argument
// vector and hands it to run; the seam lets tests prove the exact vector
// without executing git. The tokens are rebound to a plain []string because
// that is the vector type the runner (and ultimately os/exec) takes.
func gitWith(run runner, args ...GitArg) gloo.Command[[]byte, []byte] {
	argv := make([]string, len(args))
	for i, a := range args {
		argv[i] = string(a)
	}
	return run("git", argv...)
}
