// Unit tests for oligo/short

package short

import (
	"flag"
	"math/rand"
	"os"
	"testing"
	"acoma/oligo"
	"acoma/oligo/long"
)

var iternum = flag.Int("n", 5, "number of iterations")

func TestMain(m *testing.M) {
	flag.Parse()
        os.Exit(m.Run())
}

func randomString(l int) string {
	// don't allow oligos of 0 length
	if l == 0 {
		l = 1
	}

	so := ""
	for i := 0; i < l; i++ {
		so += oligo.Nt2String(rand.Intn(4))
	}

	return so
}

func randomOligo(l int) (o oligo.Oligo) {
	so := randomString(l)

	// randomly return some of the oligos as long, so we can test
	// the interoperability
	if rand.Intn(3) == 0 {
		o, _ = long.FromString(so)
	} else {
		o, _ = FromString(so)
	}

	return
}

func TestAt(t *testing.T) {
	for i := 0; i < *iternum; i++ {
		so1 := randomString(rand.Intn(31))
		o1, _ := FromString(so1)
		so2 := ""
		for i := 0; i < o1.Len(); i++ {
			so2 += oligo.Nt2String(o1.At(i))
		}

		if so1 != so2 {
			t.Fatalf("At() fails: %v: %v", so1, so2)
		}
	}
}

func TestString(t *testing.T) {
	for i := 0; i < *iternum; i++ {
		so1 := randomString(rand.Intn(31))
		o1, _ := FromString(so1)
		so2 := o1.String()

		if so1 != so2 {
			t.Fatalf("String() fails: %v: %v", so1, so2)
		}
	}
}

func TestCmp(t *testing.T) {
	for i := 0; i < *iternum; i++ {
		o1 := randomOligo(rand.Intn(31))
		o2 := randomOligo(rand.Intn(31))

		so1 := o1.String()
		so2 := o2.String()

		scmp := len(so1) - len(so2)
		if scmp < 0 {
			scmp = -1
		} else if scmp > 0 {
			scmp = 1
		}

		if scmp == 0 {
			for i := 0; i < len(so1); i++ {
				nt1 := oligo.String2Nt(string(so1[i]))
				nt2 := oligo.String2Nt(string(so2[i]))
				if nt1 < nt2 {
					scmp = -1
					break
				} else if nt1 > nt2 {
					scmp = 1
					break
				}
			}
		}

		if scmp != o1.Cmp(o2) {
			t.Fatalf("Cmp() fails: %v:%v %d:%d", so1, so2, scmp, o1.Cmp(o2))
		}
	}
}

func TestNext(t *testing.T) {
	for i := 0; i < *iternum; i++ {
		o1 := randomOligo(rand.Intn(31))
		o2 := o1.Clone()
		o2.Next()

		if o1.Cmp(o2) != -1 {
			t.Fatalf("Next() fails: %v: %v", o1, o2)
		}
	}
}


func TestSlice(t *testing.T) {
	for i := 0; i < *iternum; {
		o1 := randomOligo(rand.Intn(31))
		if o1.Len() < 4 {
			continue
		}

		s := rand.Intn(o1.Len())
		e := rand.Intn(o1.Len())

		if (e <= s) {
			continue
		}

		so1 := ""
		for n := s; n < e; n++ {
			so1 += oligo.Nt2String(o1.At(n))
		}

		o2 := o1.Slice(s, e)
		so2 := o2.String()
		if so1 != so2 {
			t.Fatalf("Slice() fails: %v: %v", so1, so2)
		}

		i++
	}
}

func TestAppend(t *testing.T) {
	for i := 0; i < *iternum; i++ {
		o1 := randomOligo(rand.Intn(31))
		o2 := randomOligo(rand.Intn(31))

		so1 := o1.String() + o2.String()
		ok := o1.Append(o2)
		so2 := o1.String()

		if !ok || so1 != so2 {
			t.Fatalf("Append() fails: %v: %v", so1, so2)
		}
	}
}

func TestCopy(t *testing.T) {
	for i := 0; i < *iternum; i++ {
		o1 := randomOligo(rand.Intn(31))
		o2, ok := Copy(o1)

		if !ok || o1.Cmp(o2) != 0 {
			t.Fatalf("Copy() fails: %v: %v", o1, o2)
		}
	}
}
