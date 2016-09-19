
package rot13

import (
	"testing"
)


func TestRot13(t *testing.T)  {
	orig:= "Hello there"
	expected := "Uryyb gurer"
	t.Logf("Checking %q",orig)
	out := Rot13(orig)
	if out!= expected{
		t.Fatalf("%q != %q",out,expected)
	}
}

type TestCase struct {
	in string
	out string
}

var test_cases = []TestCase{
	{"",""},
	{"Forever 21", "Sberire 21"},
}

func checkRot13(t *testing.T,in, expected string){
	t.Logf("Checking %q", in)
	out:= Rot13(in)
	if out != expected {
		t.Errorf("in = %q, out = %q expected = %q",in, out, expected)
	}
}

func TestRot13Many(t * testing.T){
	//Setup
	for _,tc := range test_cases{
		//go 1.6
		checkRot13(t,tc.in,tc.out)
		//go 1.7
		//t.Run(tc.in,func(t *testing.T){checkRot13(t, tc.in, tc.expected)})
	}



}


func BenchmarkRot13(b * testing.B){

	for i:=0; i<b.N ;i++{
		Rot13("The answer is 42")
	}
}

func TestParallel1(t * testing.T){
	t.Parallel()
	t.Logf("Running parallel1")
}

func TestParallel2(t * testing.T){
	t.Parallel()
	t.Logf("Running parallel1")
}