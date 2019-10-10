package test

import (
	"SpiderHog/utils"
	"testing"
)

func TestLog(t *testing.T) {
	Log := utils.GetStdoutLogger()
	Log.Println("test")
}
