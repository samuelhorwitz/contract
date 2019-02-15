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
	// PostRestoreInvariantPhase is in the invariant check that occurs after the
	// postcondition check after a restore is attempted.
	PostRestoreInvariantPhase
	// PostRestorePhase is in the postcondition check that occurs after a
	// restore is attempted.
	PostRestorePhase
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
	case PostRestorePhase:
		return "Post-restore"
	case PostRestoreInvariantPhase:
		return "Post-restore invariant"
	}
	return "Invalid phase"
}
