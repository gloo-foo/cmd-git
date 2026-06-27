package alias_test

import (
	"reflect"
	"testing"

	command "github.com/gloo-foo/cmd-git"
	git "github.com/gloo-foo/cmd-git/alias"
)

// The alias package re-exports the Git constructor under an unprefixed name. A
// mis-wired re-export (Git bound to some other constructor) compiles cleanly,
// so only behavior can prove the wiring. Executing the returned Command would
// fork real git — non-hermetic and dependent on a working install — so instead
// the test proves the re-export points at the exact same constructor: same
// function identity means identical forking behavior.
func TestAlias_GitReExportsConstructor(t *testing.T) {
	got := reflect.ValueOf(git.Git).Pointer()
	want := reflect.ValueOf(command.Git).Pointer()
	if got != want {
		t.Fatalf("alias.Git is not wired to command.Git (%v != %v)", got, want)
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
