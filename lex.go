package jsonpath

type Pos int
type stateFn func(lexer, interface{}) stateFn

const (
	eof     = -1
	noValue = -2
)

type Item struct {
	typ int
	pos Pos // The starting position, in bytes, of this Item in the input string.
	val string
}

// Used by evaluator
type tokenReader interface {
	next() (Item, bool)
}

// Used by state functions
type lexer interface {
	take() int
	peek() int
	emit(int)
	ignore()
	next() (Item, bool)
	errorf(string, ...interface{}) stateFn
	setState(interface{})
}

func itemsDescription(items []Item, nameMap map[int]string) []string {
	vals := make([]string, len(items))
	for i, item := range items {
		vals[i] = itemDescription(&item, nameMap)
	}
	return vals
}

func itemDescription(item *Item, nameMap map[int]string) string {
	var found bool
	val, found := nameMap[item.typ]
	if !found {
		return item.val
	}
	return val
}
