package operatorgen

// TTypeSequence is array of special own type
type TTypeSequence []*TType

// TTypeMapFunc is function using abstract TType
type TTypeMapFunc func(t *TType) *TType

// TTypeFilterFunc is func for filtering that is taking TType and return boolean
type TTypeFilterFunc func(t *TType) bool

// TTypeReduceFunc is taking two TTypes and returns TType based on operation applied between those two
type TTypeReduceFunc func(first *TType, second *TType) *TType

// Map is function that is making operation defined by mapper function above all the elements of an array
func (seq TTypeSequence) Map(mapper TTypeMapFunc) TTypeSequence {
	res := []*TType{}
	for _, val := range seq {
		res = append(res, mapper(val))
	}

	return res
}

// Filter is function that is filtering array seq TTypeSequence. Tests each element with a unary predicate.
// and returning those who are satisfying the function. Others are removed.
func (seq TTypeSequence) Filter(filter TTypeFilterFunc) TTypeSequence {
	res := []*TType{}
	for _, val := range seq {
		if filter(val) {
			res = append(res, val)
		}
	}

	return res
}

// Reduce is reducing the seq TTypeSequence. It is going through that array and apply function reducer
// between elements of that sequence. For example for array [1,2,3] applying function "add" operations return 6.
func (seq TTypeSequence) Reduce(reducer TTypeReduceFunc, init *TType) *TType {
	// initialize current value (first)
	curr := init
	// go over all values
	for _, val := range seq {
		curr = reducer(curr, val)
	}

	return curr
}
