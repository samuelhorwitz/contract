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

// AssertErrorRestorable is an error that occurred during a contract check and
// expects to be restored from.
type AssertErrorRestorable struct {
	Err   error
	Phase Phase
}

func (err AssertErrorRestorable) Error() string {
	return fmt.Sprintf("%s check failed: %s", err.Phase, err.Err)
}

// AssertError returns an unrestorable AssertError.
func (err AssertErrorRestorable) AssertError() AssertError {
	return AssertError{
		Err:   err.Err,
		Phase: err.Phase,
	}
}

// AssertRestoreError is an error that occured during the restore phase of a
// contract check failure.
type AssertRestoreError struct {
	Err         error
	OriginalErr AssertErrorRestorable
	Phase       Phase
}

func (err AssertRestoreError) Error() string {
	return fmt.Sprintf("%s failed: %s; %s", err.Phase, err.Err, err.OriginalErr.Error())
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

func assertOutRestorable(okay bool, reason error) {
	assert(okay, AssertErrorRestorable{
		Err:   reason,
		Phase: PostconditionPhase,
	})
}

func assertOutRestore(originalError AssertErrorRestorable) Assert {
	return func(okay bool, reason error) {
		assert(okay, AssertRestoreError{
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

func assertInvariantOutRestorable(okay bool, reason error) {
	assert(okay, AssertErrorRestorable{
		Err:   reason,
		Phase: PostconditionInvariantPhase,
	})
}

func assertInvariantRestore(originalError AssertErrorRestorable) Assert {
	return func(okay bool, reason error) {
		assert(okay, AssertRestoreError{
			Err:         reason,
			OriginalErr: originalError,
			Phase:       PostRestoreInvariantPhase,
		})
	}
}

func assert(okay bool, decoratedReason error) {
	if !okay {
		panic(decoratedReason)
	}
}
