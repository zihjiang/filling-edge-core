
package collections

type Entry struct {
	key   interface{}
	value interface{}
}

func NewEntry(key interface{}, value interface{}) *Entry {
	return &Entry{key, value}
}

func (e *Entry) GetKey() interface{} {
	return e.key
}

func (e *Entry) GetValue() interface{} {
	return e.value
}
