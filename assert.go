package contract

import (
	"fmt"
)

// Assert is a function type which takes an expression that evaluates to a
// boolean as well as a description of the failure case, should it occur.
type Assert func(okay bool, reason string)

// AssertError is a decorated string implementing the error interface.
type AssertError struct {
	err              error
	restoreErr       error
	restoreAttempted bool
}

// Error causes AssertError to implement the error interface.
func (err AssertError) Error() string {
	if err.restoreErr != nil {
		return fmt.Sprintf("%s; %s", err.restoreErr.Error(), err.err.Error())
	}
	return err.err.Error()
}

// RestoreAttempted reveals whether this error attempted to perform a restore.
func (err AssertError) RestoreAttempted() bool {
	return err.restoreAttempted
}

// RestoreSuccessful reveals whether this error was successful at performing a
// restore as well as whether a restore was attempted.
func (err AssertError) RestoreSuccessful() (successful bool, attempted bool) {
	return err.restoreAttempted && err.restoreErr == nil, err.restoreAttempted
}

// RestoreFailed reveals whether this error was unsuccessful at performing a
// restore as well as whether a restore was attempted.
func (err AssertError) RestoreFailed() (failed bool, attempted bool) {
	return err.restoreAttempted && err.restoreErr != nil, err.restoreAttempted
}

func assertIn(okay bool, reason string) {
	assert(okay, fmt.Errorf("Precondition failed: %s", reason))
}

func assertOut(okay bool, reason string) {
	assert(okay, fmt.Errorf("Postcondition failed: %s", reason))
}

func assertOutRestore(originalError AssertError) Assert {
	return func(okay bool, reason string) {
		assertRestore(okay, originalError, fmt.Errorf("Post-failure restore failed: %s", reason))
	}
}

func assertInvariantConstruct(okay bool, reason string) {
	assert(okay, fmt.Errorf("Initialization invariant failed: %s", reason))
}

func assertInvariantIn(okay bool, reason string) {
	assert(okay, fmt.Errorf("Precondition invariant failed: %s", reason))
}

func assertInvariantOut(okay bool, reason string) {
	assert(okay, fmt.Errorf("Postcondition invariant failed: %s", reason))
}

func assertInvariantRestore(originalError AssertError) Assert {
	return func(okay bool, reason string) {
		assertRestore(okay, originalError, fmt.Errorf("Post-failure restore invariant failed: %s", reason))
	}
}

func assert(okay bool, decoratedReason error) {
	if !okay {
		panic(AssertError{err: decoratedReason})
	}
}

func assertRestore(okay bool, originalError AssertError, restoreError error) {
	originalError.restoreAttempted = true
	if !okay {
		originalError.restoreErr = restoreError
		panic(originalError)
	}
}
