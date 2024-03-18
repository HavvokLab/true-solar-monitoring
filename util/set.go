package util

type Set struct {
	elements map[string]bool
}

func NewSet() *Set {
	return &Set{elements: make(map[string]bool)}
}

func (set *Set) Add(element string) {
	set.elements[element] = true
}

func (set *Set) Remove(element string) {
	delete(set.elements, element)
}

func (set *Set) Contains(element string) bool {
	_, exists := set.elements[element]
	return exists
}

func (set Set) Keys() []string {
	keys := make([]string, 0, len(set.elements))
	for k := range set.elements {
		keys = append(keys, k)
	}
	return keys
}
