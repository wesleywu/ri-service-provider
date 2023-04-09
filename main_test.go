package ri_service_provider

import (
	"fmt"
	"github.com/gogf/gf/v2/os/gtime"
	"testing"
)

func Test1(t *testing.T) {
	fmt.Println(gtime.Now().Local().Time.Format("2006-01-02 15:04:05.123456"))
}
