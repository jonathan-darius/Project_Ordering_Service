package Utils

var exists = struct{}{}

type SetStructure struct {
	m map[string]struct{}
}

func Set(arr []string) *SetStructure {
	s := &SetStructure{}
	s.m = make(map[string]struct{})
	for _, x := range arr {
		if x != "" {
			s.Add(x)
		}
	}
	return s
}

func (s *SetStructure) Add(value string) {
	s.m[value] = exists
}

func (s *SetStructure) Remove(value string) {
	delete(s.m, value)
}

func (s *SetStructure) Contains(value string) bool {
	_, c := s.m[value]
	return c
}

func (s *SetStructure) Value() (arr []string) {
	for k := range s.m {
		arr = append(arr, k)
	}
	return arr
}
