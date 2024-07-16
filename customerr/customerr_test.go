package customerr

import (
	"testing"
)

func TestCheckErrorType(t *testing.T) {
	errObj := CustomErr1()
	if _, ok := errObj.(*CustomError1); ok {
		t.Log("Error type 1")
	} else {
		t.Error("Error type 1 not found")
	}

	switch errObj.(type) {
	case *CustomError1:
		t.Log("Error type 1")
	case *CustomError2:
		t.Error("Error type 1")
	default:
		t.Error("Error not found")
	}
}
