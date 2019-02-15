package contract

import (
	"fmt"
)

// Assert is a function type which takes an expression that evaluates to a
// boolean as well as an error to panic with if this expression is false.
type Assert func(okay bool, reason error)

// AssertError is a decorated string implementing the error interface.
type AssertError struct {
	Err         error
	OriginalErr error
	Phase       Phase
}

func (err AssertError) Error() string {
	errString := fmt.Sprintf("%s check failed: %s", err.Phase, err.Err)
	if err.OriginalErr != nil {
		return fmt.Sprintf("%s; %s", errString, err.OriginalErr.Error())
	}
	return errString
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

func assertOutRestore(originalError AssertError) Assert {
	return func(okay bool, reason error) {
		assert(okay, AssertError{
			Err:         reason,
			OriginalErr: originalError,
			Phase:       PostRestorePhase,
		})
	}
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

func assertInvariantRestore(originalError AssertError) Assert {
	return func(okay bool, reason error) {
		assert(okay, AssertError{
			Err:         reason,
			OriginalErr: originalError,
			Phase:       PostRestoreInvariantPhase,
		})
	}
}

func assert(okay bool, decoratedReason AssertError) {
	if !okay {
		panic(decoratedReason)
	}
}
