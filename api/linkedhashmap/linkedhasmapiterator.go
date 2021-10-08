
package linkedhashmap

import "datacollector-edge/api/linkedhashmap/collections"

type Iterator struct {
	m       *Map
	current *Link
}

func (iterator *Iterator) HasNext() bool {
	return !(iterator.current == nil)
}

func (iterator *Iterator) Next() *collections.Entry {
	temp := iterator.current
	if temp == nil {
		return nil
	}
	iterator.current = temp.next
	return collections.NewEntry(temp.key, temp.value)
}

func (iterator *Iterator) Reset() {
	iterator.current = iterator.m.head
}
