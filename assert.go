package contract

import (
	"fmt"
)

// Assert is a function type which takes an expression that evaluates to a
// boolean as well as an error to panic with if this expression is false.
type Assert func(okay bool, reason error)

// AssertError is an error that occured during a contract check.
type AssertError struct {
	Err   error
	Phase Phase
}

func (err AssertError) Error() string {
	return fmt.Sprintf("%s check failed: %s", err.Phase, err.Err)
}

func assertIn(okay bool, reason error) {
	assert(okay, AssertError{
		Err:   reason,
		Phase: PreconditionPhase,
	})
}

func assertOut(okay bool, reason error) {
	assert(okay, AssertError{
		Err:   reason,
		Phase: PostconditionPhase,
	})
}

func assertInvariantConstruct(okay bool, reason error) {
	assert(okay, AssertError{
		Err:   reason,
		Phase: InitializationPhase,
	})
}

func assertInvariantIn(okay bool, reason error) {
	assert(okay, AssertError{
		Err:   reason,
		Phase: PreconditionInvariantPhase,
	})
}

func assertInvariantOut(okay bool, reason error) {
	assert(okay, AssertError{
		Err:   reason,
		Phase: PostconditionInvariantPhase,
	})
}

func assert(okay bool, decoratedReason error) {
	if !okay {
		panic(decoratedReason)
	}
}
