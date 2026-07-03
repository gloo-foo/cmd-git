package alias_test

import (
	"reflect"
	"testing"

	command "github.com/gloo-foo/cmd-git"
	git "github.com/gloo-foo/cmd-git/alias"
)

// The alias package re-exports the Git constructor under an unprefixed name by
// delegation. A mis-wired re-export (Git bound to some other constructor)
// compiles cleanly, so the wiring must be proven. Executing the returned
// Command would fork real git — non-hermetic and dependent on a working
// install — so the wrapper is pinned to the constructor's exact signature
// (its argument vector cannot silently diverge from command.Git's) and proven
// to build a Command in TestAlias_GitBuildsCommand.
func TestAlias_GitSignatureMatchesConstructor(t *testing.T) {
	want := reflect.TypeOf(command.Git)
	if got := reflect.TypeOf(git.Git); got != want {
		t.Fatalf("alias.Git signature is %v, want %v", got, want)
	}
}

// The re-exported constructor must still build a usable Command for any
// argument vector, including the no-argument (bare git) case.
func TestAlias_GitBuildsCommand(t *testing.T) {
	if git.Git("status") == nil {
		t.Fatal("alias.Git(\"status\") returned a nil Command")
	}
	if git.Git() == nil {
		t.Fatal("alias.Git() returned a nil Command")
	}
}
