package utils

import (
	"strings"
  "strconv"
)


func ExplodeVersion(v string) (int, int, int) {
  vers := strings.Fields(strings.Replace(v,"."," ",-1))
  major, _ := strconv.Atoi(vers[0])
  minor, _ := strconv.Atoi(vers[1])
  inc, _ := strconv.Atoi(vers[2])
  return major, minor, inc
}