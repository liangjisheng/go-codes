package stubdemo

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/prashantv/gostub"
)

var timeNow = time.Now
var osHostname = os.Hostname

func getDate() int {
	return timeNow().Day()
}

func getHostName() (string, error) {
	return osHostname()
}

func TestStubTimeNowFunc(t *testing.T) {
	stubs := gostub.Stub(&timeNow, func() time.Time {
		return time.Date(2020, 5, 19, 16, 4, 30, 0, time.UTC)
	})
	fmt.Println(getDate())
	defer stubs.Reset()
}

func TestStubHostNameFunc(t *testing.T) {
	stubs := gostub.StubFunc(&osHostname, "Localhost", nil)
	defer stubs.Reset()
	fmt.Println(getHostName())
}
