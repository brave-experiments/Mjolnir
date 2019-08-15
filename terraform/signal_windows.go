// +build windows

package terraform

import (
	"os"
)

var ignoreSignals = []os.Signal{os.Interrupt}
var forwardSignals []os.Signal
