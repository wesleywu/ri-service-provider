package gwmap

type StrStrMap struct {
	InnerMap map[string]interface{}
}

func (s *StrStrMap) Get(key string) string {
	if s.InnerMap == nil {
		return ""
	}
	item, ok := s.InnerMap[key].([]string)
	if !ok {
		return ""
	}
	if len(item) == 0 {
		return ""
	}
	return item[0]
}

func (s *StrStrMap) Set(key string, value string) {
	if s.InnerMap == nil {
		s.InnerMap = map[string]interface{}{}
	}
	s.InnerMap[key] = value
}

func (s *StrStrMap) Keys() []string {
	out := make([]string, 0, len(s.InnerMap))
	for key := range s.InnerMap {
		out = append(out, key)
	}
	return out
}
