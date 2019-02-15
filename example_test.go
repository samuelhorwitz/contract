package contract_test

import (
	"errors"
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
	assert(float64(nameLen) > sqrtNumber, fmt.Errorf("Length of name %d is not greater than %f", nameLen, sqrtNumber))
}

func (s *Sample) SetName(newName string) {
	contract.In(s, func(assert contract.Assert) {
		assert(newName != "", errors.New("New name must not be empty"))
	})
	defer contract.Out(s, func(assert contract.Assert) {
		assert(len(s.name) > len(newName), errors.New("Name should be longer than what was passed in"))
	})
	s.name = newName + ", Esq."
}

func (s *Sample) SetNumber(newNumber float64) {
	// Go's block scoping allows us to rope off local variables and group the
	// pre and post conditions in a syntactically pleasing way. It may be best
	// to always use a block scope wrapper, just for readability, when declaring
	// conditions.
	{
		oldNumber := s.number
		contract.In(s, func(assert contract.Assert) {
			assert(newNumber >= 0, errors.New("New number must be positive"))
		})
		defer contract.Out(s, func(assert contract.Assert) {
			assert(oldNumber < s.number, errors.New("New number must be greater than old number"))
		})
	}
	s.number = newNumber * 1.6
}

func Example() {
	sample := NewSample("sam")
	sample.SetName("sammy")
	sample.SetNumber(55.7)
	fmt.Println(sample)
	// Output: &{sammy, Esq. 89.12}
}
