package stringsT

import (
	"strings"
	"testing"
)

func TestStrings(t *testing.T){
	var sBuild strings.Builder
	t.Log(sBuild.Cap())
	sBuild.Grow(4)
	sBuild.WriteString("tiptok")
	t.Log(sBuild.String(),sBuild.Len())
}

func TestEqual(t *testing.T){
	input := struct{ A string
	B string
	C string
	}{
		A:"ab我",
		B:"ac",
		C:"ab我",
	}
	if strings.EqualFold(input.A,input.B)!=false{
		t.Fatal("equal:",input.A,input.B)
	}
	if !strings.EqualFold(input.A,input.C){
		t.Fatal("equal",input.A,input.C)
	}
	if strings.Compare(input.A,input.C)!=0{
		t.Fatal("equal",input.A,input.C)
	}
}