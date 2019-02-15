package contract_test

import (
	"fmt"
	"github.com/samuelhorwitz/contract"
	"math"
)

type Sample struct {
	name   string
	number float64
}

func NewSample(name string) (s *Sample) {
	defer func() { contract.Construct(s) }()
	return &Sample{
		name:   name,
		number: 6.734,
	}
}

func (s *Sample) Invariant(assert contract.Assert) {
	nameLen := len(s.name)
	sqrtNumber := math.Sqrt(s.number)
	assert(float64(nameLen) > sqrtNumber, fmt.Sprintf("Length of name %d is not greater than %f", nameLen, sqrtNumber))
}

func (s *Sample) SetName(newName string) {
	contract.In(s, func(assert contract.Assert) {
		assert(newName != "", "New name must not be empty")
	})
	defer contract.Out(s, func(assert contract.Assert) {
		assert(len(s.name) > len(newName), "Name should be longer than what was passed in")
	})
	s.name = newName + ", Esq."
}

func (s *Sample) SetNumber(newNumber float64) {
	defer contract.InAndOut(s, func(assert contract.Assert) contract.Condition {
		// In
		oldNumber := s.number
		assert(newNumber >= 0, "New number must be positive")
		return func(assert contract.Assert) {
			// Out
			assert(oldNumber < s.number, "New number must be greater than old number")
		}
	})()
	s.number = newNumber * 1.6
}

func Example() {
	sample := NewSample("sam")
	sample.SetName("sammy")
	sample.SetNumber(55.7)
	fmt.Println(sample)
	// Output: &{sammy, Esq. 89.12}
}
