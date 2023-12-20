package helper

import (
	"fmt"
	"testing"
)

func TestErrorSum(t *testing.T) {
	result := Sum(5, 5)
	if result != 40 {
		t.Error("Result must be 10")
	}
	fmt.Println("hahaha")
}

func TestFatalSum(t *testing.T) {
	result := Sum(5, 5)
	if result != 40 {
		t.Fatal("Result must be 10")
	}
	fmt.Println("hihihi")
}

func TestFailSum(t *testing.T) {
	result := Sum(5, 5)
	if result != 40 {
		t.Fail()
	}
	fmt.Println("huhuhu")
}

func TestSum(t *testing.T) {
	result := Sum(5, 5)
	if result != 40 {
		t.FailNow()
	}
	fmt.Println("hehehe")
}
