package jsonschema

import (
	"fmt"
	"testing"

	"github.com/go-test/deep"
)

func TestScanTags(t *testing.T) {
	line := `flag1,flag2,k1=v1,k2=sv1, sv22,k3=[1,2,3],k4=,k5=v5,k6=v6,a`
	res := scanTagKVs(line)
	expected := []*keyValue{
		{key: "k1", value: "v1"},
		{key: "k2", value: "sv1, sv2"},
		{key: "k3", value: "[1,2,3]"},
		{key: "k4", value: ""},
		{key: "k5", value: "v5"},
		{key: "k6", value: "v6,a"},
	}
	if diff := deep.Equal(expected, res); diff != nil {
		t.Errorf("The diff between the expected and actual result is %v", diff)
		fmt.Println("Expected:")
		for _, r := range expected {
			fmt.Printf("%+v\n", *r)
		}
		fmt.Println("Actual:")
		for _, r := range res {
			fmt.Printf("%+v\n", *r)
		}
	}
}
