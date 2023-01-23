package orm

import (
	"encoding/json"
	"github.com/gogf/gf/v2/errors/gerror"
	"strings"
)

type MultiType uint8

const (
	Exact MultiType = iota
	Between
	NotBetween
	In
	NotIn
)

var (
	MultiTypeNames = map[MultiType]string{
		Exact:      "exact",
		Between:    "between",
		NotBetween: "notbetween",
		In:         "in",
		NotIn:      "notin",
	}
	MultiTypeValues = map[string]MultiType{
		"exact":      Exact,
		"between":    Between,
		"notbetween": NotBetween,
		"in":         In,
		"notin":      NotIn,
	}
)

func (s *MultiType) String() string {
	return MultiTypeNames[*s]
}

func (s *MultiType) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s *MultiType) UnmarshalJSON(data []byte) (err error) {
	var typeString string
	if err := json.Unmarshal(data, &typeString); err != nil {
		return err
	}
	if *s, err = ParseMultiType(typeString); err != nil {
		return err
	}
	return nil
}

func ParseMultiType(s string) (MultiType, error) {
	s = strings.TrimSpace(strings.ToLower(s))
	value, ok := MultiTypeValues[s]
	if !ok {
		return Exact, gerror.Newf("%q 不是合法的 MultiType", s)
	}
	return value, nil
}
