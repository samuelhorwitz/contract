package contract

// Phase is a type which represents a contract condition check point.
type Phase int

const (
	// InitializationPhase is in the invariant check that occurs after creation.
	InitializationPhase Phase = iota
	// PreconditionPhase is in the precondition check.
	PreconditionPhase
	// PreconditionInvariantPhase is in the invariant check that occurs after
	// the precondition check.
	PreconditionInvariantPhase
	// PostconditionInvariantPhase is in the invariant check that occurs after
	// the postcondition check.
	PostconditionInvariantPhase
	// PostconditionPhase is in the postcondition check.
	PostconditionPhase
)

func (p Phase) String() string {
	switch p {
	case InitializationPhase:
		return "Initialization invariant"
	case PreconditionPhase:
		return "Precondition"
	case PreconditionInvariantPhase:
		return "Precondition invariant"
	case PostconditionPhase:
		return "Postcondition"
	case PostconditionInvariantPhase:
		return "Postcondition invariant"
	}
	return "Invalid phase"
}
