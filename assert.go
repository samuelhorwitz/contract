// +build !nodbc

package contract

import (
	"fmt"
)

// Assert is a function type which takes an expression that evaluates to a
// boolean as well as a description of the failure case, should it occur.
type Assert func(okay bool, reason string)

// AssertError is a decorated string implementing the error interface.
type AssertError string

func (err AssertError) Error() string {
	return string(err)
}

func assertIn(okay bool, reason string) {
	assert(okay, fmt.Sprintf("Precondition failed: %s", reason))
}

func assertOut(okay bool, reason string) {
	assert(okay, fmt.Sprintf("Postcondition failed: %s", reason))
}

func assertOutRestore(originalError AssertError) Assert {
	return func(okay bool, reason string) {
		assert(okay, fmt.Sprintf("Post-failure restore failed: %s; Original: %s", reason, originalError))
	}
}

func assertInvariantConstruct(okay bool, reason string) {
	assert(okay, fmt.Sprintf("Initialization invariant failed: %s", reason))
}

func assertInvariantIn(okay bool, reason string) {
	assert(okay, fmt.Sprintf("Precondition invariant failed: %s", reason))
}

func assertInvariantOut(okay bool, reason string) {
	assert(okay, fmt.Sprintf("Postcondition invariant failed: %s", reason))
}

func assertInvariantRestore(originalError AssertError) Assert {
	return func(okay bool, reason string) {
		assert(okay, fmt.Sprintf("Post-failure restore invariant failed: %s; Original: %s", reason, originalError))
	}
}

func assert(okay bool, decoratedReason string) {
	if !okay {
		panic(AssertError(decoratedReason))
	}
}
