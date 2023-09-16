package sets

type SetElement interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64 | string | bool
}

type set[K SetElement] map[K]struct{}

func New[K SetElement](keys ...K) set[K] {
	s := make(set[K], len(keys))

	for _, key := range keys {
		s[key] = struct{}{}
	}
	return s
}

func (s set[K]) Exists(key K) bool {
	if s == nil {
		return false
	}
	_, exist := s[key]
	return exist
}

func (s set[K]) Add(keys ...K) {
	for _, key := range keys {
		s[key] = struct{}{}
	}
}

func (s set[K]) Remove(keys ...K) {
	for _, key := range keys {
		delete(s, key)
	}
}

func (s set[K]) ToList() []K {
	res := make([]K, 0, len(s))

	for key := range s {
		res = append(res, key)
	}
	return res
}

func (s set[K]) Intersect(s2 set[K]) set[K] {
	newSet := make(set[K], 0)
	for key := range s {
		if s2.Exists(key) {
			newSet.Add(key)
		}
	}
	return newSet
}

func (s set[K]) Union(s2 set[K]) set[K] {
	newSet := make(set[K], 0)
	for key := range s {
		newSet.Add(key)
	}
	for key := range s2 {
		newSet.Add(key)
	}
	return newSet
}

func (s set[K]) Minus(s2 set[K]) set[K] {
	newSet := make(set[K], 0)
	for key := range s {
		if !s2.Exists(key) {
			newSet.Add(key)
		}
	}
	return newSet
}

func (s set[K]) MinusInPlace(s2 set[K]) {
	for key := range s2 {
		s.Remove(key)
	}
}

func (s set[K]) IsEmpty() bool {
	return len(s) == 0
}

func (s set[K]) Equal(s2 set[K]) bool {
	if len(s) != len(s2) {
		return false
	}

	for key := range s {
		if !s2.Exists(key) {
			return false
		}
	}

	return true
}

func (s set[K]) Size() int {
	return len(s)
}
