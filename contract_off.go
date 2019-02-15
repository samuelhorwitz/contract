// +build nodbc

package contract

func Construct(Invariable) {}

func In(Invariable, Condition) {}

func Out(Invariable, Condition) {}

func OutAndRestore(Invariable, Condition, func()) {}

func InAndOut(Invariable, PreToPostCondition) func() { return func() {} }

func InAndOutAndRestore(Invariable, PreToPostConditionWithRestore) func() { return func() {} }

type Invariable interface{}

type PreToPostCondition func(Assert) Condition

type PreToPostConditionWithRestore func(Assert) (Condition, func())

type Condition func(Assert)

type Assert func(bool, string)
