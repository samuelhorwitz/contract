// +build nodbc

package contract

func Construct(Invariable)                                                     {}
func In(Invariable, Condition)                                                 {}
func Out(Invariable, Condition)                                                {}
func OutAndRestore(Invariable, Condition, Restore)                             {}
func InAndOut(Invariable, PreToPostCondition) EnclosedOut                      { return func() {} }
func InAndOutAndRestore(Invariable, PreToPostConditionWithRestore) EnclosedOut { return func() {} }
