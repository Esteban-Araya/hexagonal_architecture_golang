package service

import (
	"testing"
	"Project/pkg/app/user/error"
)
func TestCreate(t *testing.T){
	TestCases := []struct{
		Name string
		ExpectedValue bool
		ExpectedError error
	}{
		{
			Name:"Succes",
			ExpectedValue: true,
			ExpectedError: nil,
		},
		{
			Name:"Diferent password",
			ExpectedValue: false,
			ExpectedError: error.DiferentPasswordError,
		},
	}
	
}

func TestAAA(t *testing.T){
	result := 2

	if result == 2{
		t.Errorf("fail %d",result)
	}
}