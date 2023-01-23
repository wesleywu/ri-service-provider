package orm

import (
	"encoding/json"
	"github.com/gogf/gf/v2/errors/gerror"
	"strings"
)

type WildcardType uint8

const (
	None WildcardType = iota
	Contains
	StartsWith
	EndsWith
)

var (
	WildcardTypeNames = map[WildcardType]string{
		None:       "none",
		Contains:   "contains",
		StartsWith: "startswith",
		EndsWith:   "endswith",
	}
	WildcardTypeValues = map[string]WildcardType{
		"none":       None,
		"contains":   Contains,
		"startswith": StartsWith,
		"endswith":   EndsWith,
	}
)

func (s *WildcardType) String() string {
	return WildcardTypeNames[*s]
}

func (s *WildcardType) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s *WildcardType) UnmarshalJSON(data []byte) (err error) {
	var typeString string
	if err := json.Unmarshal(data, &typeString); err != nil {
		return err
	}
	if *s, err = ParseWildcardType(typeString); err != nil {
		return err
	}
	return nil
}

func ParseWildcardType(s string) (WildcardType, error) {
	s = strings.TrimSpace(strings.ToLower(s))
	value, ok := WildcardTypeValues[s]
	if !ok {
		return None, gerror.Newf("%q 不是合法的 WildcardType", s)
	}
	return value, nil
}
