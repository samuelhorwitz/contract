// +build !nodbc

// Package contract is a design by contract implementation that intends to be
// unobtrusive and idiomatic. Contract attempts to make it easier to implement
// design by contract correctly, however it does not protect against abuse of
// the pattern.
//
// Contract makes it easy to declare preconditions, postconditions and
// invariants. To fully and correctly implement this for a type, all public
// methods must make calls to In and Out or one of the alternative functions
// provided. The type must also implement the Invariable interface which
// requires an implemented Invariant method.
//
// The Invariant method as well as the In and Out methods may not make any
// changes to state or otherwise cause any side-effects, although you will not
// be prevented from doing so. The only use for these calls should be for
// making assertions.
//
// If instances of your type are mutable in ways that don't require entering
// through a public method, then you must ensure you include appropriate checks
// as needed, or restructure the code. This package cannot guarantee anything
// not already guaranteed by Go, it can only make it easier to follow the
// contract paradigm.
//
// Assertions panic rather than return errors. This is because assertions are
// for validating that the software has not gone haywire, not performing
// business logic. Any expected states or arguments should _not_ be handled in
// the design by contract fashion. Design by contract is about failing fast,
// not about cleaning up expected, yet dirty, input.
//
// By default, all assertions are included in the build. It seems nonsensical to
// use this library and then strip it out for production. The intended use is to
// ensure sanity in a production environment. However, some people seem to
// subscribe to a different view on this, so you may use the build tag "nodbc"
// to swap in no-op replacement functions.
package contract

import (
	"fmt"
)

// Construct should be called wherever an Invariable type is instantiated.
// Idiomatically, many Go programs will use New functions for this. Go will not
// enforce instantiation in any particular way, so be aware that if you make a
// type public, then it is possible for anyone to instantiate it without running
// this check, breaking design by contract principles. Therefore, it is
// recommended that as much as possible remains unexported, forcing the usage of
// a constructor.
func Construct(i Invariable) {
	i.Invariant(assertInvariantConstruct)
}

// In is the precondition hook register. Preconditions should run as close to
// the beginning of the function as possible, ideally before anything else.
// After running the precondition check, the invariant will also be checked. If
// no precondition is specified, only the invariant will be checked.
func In(i Invariable, in Condition) {
	if in != nil {
		in(assertIn)
	}
	i.Invariant(assertInvariantIn)
}

// Out is the postcondition hook register. We recommend postconditions follow
// preconditions immediately, using defer. This way the top of the function is
// set aside for contract checks, and the rest of the function body is business
// logic. As with In, Out will check the postcondition if specified, and then
// the invariant, no matter what.
func Out(i Invariable, out Condition) {
	if out != nil {
		out(assertOut)
	}
	i.Invariant(assertInvariantOut)
}

// OutAndRestore is similar to Out, however a third agument allows a restore
// function to be specified. In case of a contract-based panic, the restore
// function will be called before the panic is re-thrown. Restore will not
// rescue a panic, it will only allow you to clean up the state before leaving
// the function scope entirely. This could be useful if farther up you intend to
// recover normally, and you do not want to be in an undefined state. The
// post-restore state must also pass all postcondition and invariant checks, or
// the error will be further decorated to annotate restore failure.
func OutAndRestore(i Invariable, out Condition, restore Restore) {
	defer handleRestore(i, out, restore)
	if out != nil {
		out(assertOutRestorable)
	}
	i.Invariant(assertInvariantOutRestorable)
}

// InAndOut combines In and Out into a nested series of function calls. When
// Out is closed over by In, it allows some level of statefulness, such as
// keeping track of the value prior to entry and comparing it to the new value
// upon exit, by using a local, closured, variable. Remember, only asserts and
// local variables should be used, anything causing state mutations or external
// side effects is a bad idea.
func InAndOut(i Invariable, in PreToPostCondition) EnclosedOut {
	out := in(assertIn)
	i.Invariant(assertInvariantIn)
	return func() {
		out(assertOut)
		i.Invariant(assertInvariantOut)
	}
}

// InAndOutAndRestore is a combination of In and OutAndRestore and operates the
// same as InAndOut, with restore functionality.
func InAndOutAndRestore(i Invariable, in PreToPostConditionWithRestore) EnclosedOut {
	out, restore := in(assertIn)
	i.Invariant(assertInvariantIn)
	return func() {
		defer handleRestore(i, out, restore)
		out(assertOutRestorable)
		i.Invariant(assertInvariantOutRestorable)
	}
}

func handleRestore(i Invariable, out Condition, restore Restore) {
	if r := recover(); r != nil {
		if r, ok := r.(AssertErrorRestorable); ok {
			if restore != nil {
				restore()
			} else {
				panic(AssertRestoreError{
					Err:         fmt.Errorf("Restore function undefined"),
					OriginalErr: r,
					Phase:       RestorePhase,
				})
			}
			out(assertOutRestore(r))
			i.Invariant(assertInvariantRestore(r))
			panic(r.AssertError())
		} else {
			panic(r)
		}
	}
}
