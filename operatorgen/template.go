package operatorgen

// TTypeSequence is array of special own type
type TTypeSequence []*TType

// TTypeMapFunc is function using abstract TType
type TTypeMapFunc func(t *TType) *TType

// TTypeForEachFunc is function using abstract TType
type TTypeForEachFunc func(t *TType)

// TTypeFilterFunc is func for filtering that is taking TType and return boolean
type TTypeFilterFunc func(t *TType) bool

// TTypeReduceFunc is taking two TTypes and returns TType based on operation applied between those two
type TTypeReduceFunc func(first *TType, second *TType) *TType

// TTypeSortFunc is taking two TTypes and returns TType based on operation applied between those two
type TTypeSortFunc func(first *TType, second *TType) bool

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

// Find is alias for Filter function, for implementation details see Filter
func (seq TTypeSequence) Find(filter TTypeFilterFunc) TTypeSequence {
	res := seq.Filter(filter)
	return res
}

// FindOne is finding first appearance of element that satysfies predice in sequence
func (seq TTypeSequence) FindOne(predicate TTypeFilterFunc) *TType {
	for _, val := range seq {
		if predicate(val) {
			return val
		}
	}
	return nil
}

// Reverse is reversing sequence and is returning that new alocated reversed sequence
func (seq TTypeSequence) Reverse() TTypeSequence {
	res := make([]*TType, len(seq))

	for i, j := 0, len(seq)-1; i <= j; i, j = i+1, j-1 {
		res[i], res[j] = seq[j], seq[i]
	}
	return res
}

// ForEach is function that is iterating over the sequence
func (seq TTypeSequence) ForEach(function TTypeForEachFunc) {
	for _, val := range seq {
		function(val)
	}
}

// Sort is sorting sequence based on function sortedBy
func (seq TTypeSequence) Sort(sortedBy TTypeSortFunc) TTypeSequence {
	size := len(seq)
	res := seq
	if size < 2 {
		return res
	}
	for i := 0; i < size; i++ {
		for j := size - 1; j >= i+1; j-- {
			if sortedBy(res[j-1], res[j]) {
				res[j], res[j-1] = res[j-1], res[j]
			}
		}
	}
	return res
}

// Push func
func (seq TTypeSequence) Push(elem *TType) TTypeSequence {
	res := seq
	res = append(res, elem)

	return res
}

// Pop func
func (seq TTypeSequence) Pop() TTypeSequence {
	res := seq[:len(seq)-1]

	return res
}
