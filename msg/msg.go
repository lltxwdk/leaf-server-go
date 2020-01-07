package msg

import (
	"github.com/name5566/leaf/network/json"
)

type Hello struct {
	Name string
}

var Processor = json.NewProcessor()

func init() {
	Processor.Register(&Hello{})
}
