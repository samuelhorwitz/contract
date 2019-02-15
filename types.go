package contract

// Invariable is an interface that must be implemented by types desiring to have
// contractual checks.
//
// The Invariant method is defined on the type and is responsible for asserting
// the correctness of the state. It will be run after every precondition call
// and every postcondition call.
type Invariable interface {
	Invariant(Assert)
}

// Condition is a type which represents a simple pre or postcondition function
// type.
type Condition func(Assert)

// Restore is a type which represents a restore function.
type Restore func()

// EnclosedOut is a type which represents an Out function that is closed over by
// an In function.
type EnclosedOut func()

// PreToPostCondition is a type which represents an In function that returns an
// Out function.
type PreToPostCondition func(Assert) Condition

// PreToPostConditionWithRestore is a type which represents and In function that
// returns an Out function as well as a Restore function.
type PreToPostConditionWithRestore func(Assert) (Condition, Restore)
