package enum

import (
	"encoding/json"
	"fmt"
)

func (e ContentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(ContentType_name[int32(e)])
}

func (e *ContentType) UnmarshalJSON(data []byte) error {
	var (
		enumStr   string
		enumInt32 int32
		ok        bool
	)
	if err := json.Unmarshal(data, &enumStr); err != nil {
		return err
	}
	if enumInt32, ok = ContentType_value[enumStr]; !ok {
		return fmt.Errorf("%s is not a valid ContentType", enumStr)
	}
	*e = ContentType(enumInt32)
	return nil
}

func (e FilterType) MarshalJSON() ([]byte, error) {
	return json.Marshal(FilterType_name[int32(e)])
}

func (e *FilterType) UnmarshalJSON(data []byte) error {
	var (
		enumStr   string
		enumInt32 int32
		ok        bool
	)
	if err := json.Unmarshal(data, &enumStr); err != nil {
		return err
	}
	if enumInt32, ok = FilterType_value[enumStr]; !ok {
		return fmt.Errorf("%s is not a valid FilterType", enumStr)
	}
	*e = FilterType(enumInt32)
	return nil
}
