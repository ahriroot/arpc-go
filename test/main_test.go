package test

import (
	"arpc-go/utils"
	"testing"
)

func TestCompile(t *testing.T) {
	utils.Compile("./api.arpc")
}
