package command

import (
	"context"
	"slices"
	"testing"

	gloo "github.com/gloo-foo/framework"
	"github.com/gloo-foo/framework/patterns"
)

// captured records the argument vector a fake runner was handed, so a test can
// assert exactly what Git would have forked — proving the contract without
// executing git (which would be non-hermetic and depend on a working install).
type captured struct {
	name patterns.ProcessName
	args []patterns.ProcessArg
}

// fakeRunner returns a runner that records its inputs into got and yields a
// trivial pass-through Command, so no real process is ever spawned.
func fakeRunner(got *captured) runner {
	return func(name patterns.ProcessName, args ...patterns.ProcessArg) gloo.Command[[]byte, []byte] {
		got.name = name
		got.args = args
		return gloo.FuncCommand[[]byte, []byte](
			func(_ context.Context, in gloo.Stream[[]byte]) gloo.Stream[[]byte] { return in },
		)
	}
}

func TestGitWith_PrependsGitToArgs(t *testing.T) {
	var got captured
	gitWith(fakeRunner(&got), "status", "--oneline")
	if got.name != "git" {
		t.Errorf("forked %q, want \"git\"", got.name)
	}
	if !slices.Equal(got.args, []patterns.ProcessArg{"status", "--oneline"}) {
		t.Errorf("got args %q, want [status --oneline]", got.args)
	}
}

func TestGitWith_NoArgsForksBareGit(t *testing.T) {
	var got captured
	gitWith(fakeRunner(&got))
	if got.name != "git" {
		t.Errorf("forked %q, want \"git\"", got.name)
	}
	if len(got.args) != 0 {
		t.Errorf("got args %q, want none", got.args)
	}
}

// Git wires the production runner (patterns.Subprocess). Constructing the
// Command must not fork git, so this asserts only that a usable Command is
// returned; the argument-vector contract is proven against the fake runner.
func TestGit_ReturnsCommand(t *testing.T) {
	if Git("log", "-1") == nil {
		t.Fatal("Git returned a nil Command")
	}
}
