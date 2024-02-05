package gworm

import (
	"strings"

	"go.mongodb.org/mongo-driver/bson"
)

type SortDirection uint8

const (
	SortDirection_ASC  SortDirection = 0
	SortDirection_DESC SortDirection = 1
	defaultPageSize                  = 10
)

// Enum value maps for MultiType.
var (
	SortDirection_name = map[SortDirection]string{
		0: "ASC",
		1: "DESC",
	}
	SortDirection_value = map[string]SortDirection{
		"ASC":  0,
		"DESC": 1,
	}
)

type PageRequest struct {
	Number int64

	Offset int64
	Size   int64
	Sorts  []SortParam
}

type SortParam struct {
	Property  string
	Direction SortDirection
}

type PageInfo struct {
	Offset           int64       // the offset of the first element in current page of the paging request
	Size             int64       // page size of the paging request, may be larger than NumberOfElements
	Sorts            []SortParam // the sorting parameters of the paging request
	Number           int64       // page number of current page, starting from 1
	NumberOfElements int64       // number of elements in current page
	TotalElements    int64       // total number of elements for current request when without paging
	TotalPages       int64       // total number of pages of the paging request
	First            bool        // whether current page is first page
	Last             bool        // whether current page is last page
}

func (pr PageRequest) Of(page, size int64, sort ...string) PageRequest {
	if page < 1 {
		page = 1
	}
	if size <= 0 {
		size = defaultPageSize
	}
	pr.Number = page
	pr.Offset = (page - 1) * size
	pr.Size = size
	for _, s := range sort {
		pr.AddSortByString(s)
	}
	return pr
}

func (pr PageRequest) AddSortByString(sort string) PageRequest {
	props := strings.Split(sort, ",")
	for _, prop := range props {
		parts := strings.Split(strings.TrimSpace(prop), " ")
		switch len(parts) {
		case 1:
			pr.AddSort(SortParam{
				Property:  parts[0],
				Direction: SortDirection_ASC,
			})
		case 2:
			direction, ok := SortDirection_value[strings.ToUpper(parts[1])]
			if !ok {
				direction = SortDirection_ASC
			}
			pr.AddSort(SortParam{
				Property:  parts[0],
				Direction: direction,
			})
		default:
			continue
		}
	}
	return pr
}

func (pr PageRequest) AddSort(sort SortParam) PageRequest {
	pr.Sorts = append(pr.Sorts, sort)
	return pr
}

func (pr PageRequest) HasSort() bool {
	return len(pr.Sorts) > 0
}

func (pr PageRequest) MongoSortOption() (result bson.D) {
	var direction int
	for _, s := range pr.Sorts {
		if s.Direction == SortDirection_ASC {
			direction = 1
		} else {
			direction = -1
		}
		result = append(result, bson.E{
			Key:   s.Property,
			Value: direction,
		})
	}
	return
}

func (pr PageRequest) OrderString() string {
	sb := strings.Builder{}
	for i, s := range pr.Sorts {
		if i > 0 {
			sb.WriteString(",")
		}
		sb.WriteString(s.Property)
		sb.WriteString(" ")
		sb.WriteString(SortDirection_name[s.Direction])
	}
	return sb.String()
}

func (pi *PageInfo) From(pageRequest PageRequest, numberOfElement int64, totalElements int64) {
	pi.Offset = pageRequest.Offset
	pi.Size = pageRequest.Size
	pi.Sorts = pageRequest.Sorts
	pi.Number = pageRequest.Number
	pi.NumberOfElements = numberOfElement
	pi.TotalElements = totalElements
	if pi.Size > 0 {
		pi.TotalPages = totalElements / pi.Size
	} else {
		pi.TotalPages = 0
	}
	pi.First = pi.Number == 1
	pi.Last = (pi.Offset + pi.NumberOfElements) >= pi.TotalElements
}
