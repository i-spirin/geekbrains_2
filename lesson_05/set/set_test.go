package set_test

import (
	"testing"

	"github.com/i-spirin/geekbrains_2/lesson_05/set"
)

func BenchmarkReadWrite9010Muset(b *testing.B) {
	s := set.NewMuSet()
	for i := 0; i < b.N; i++ {
		s.Add(i)
		s.Remove(i)
		s.Add(i)
		s.Remove(i)
		s.Add(i)
		s.Remove(i)
		s.Add(i)
		s.Remove(i)
		s.Add(i)
		s.Has(i)
	}

}

func BenchmarkReadWrite5050Muset(b *testing.B) {
	s := set.NewMuSet()
	for i := 0; i < b.N; i++ {
		s.Add(i)
		s.Remove(i)
		s.Add(i)
		s.Remove(i)
		s.Add(i)
		s.Has(i)
		s.Has(i)
		s.Has(i)
		s.Has(i)
		s.Has(i)
	}
}

func BenchmarkReadWrite1090RWMuset(b *testing.B) {
	s := set.NewRWMuSet()
	for i := 0; i < b.N; i++ {
		s.Add(i)
		s.Has(i)
		s.Has(i)
		s.Has(i)
		s.Has(i)
		s.Has(i)
		s.Has(i)
		s.Has(i)
		s.Has(i)
		s.Has(i)
	}
}

func BenchmarkReadWrite9010RWMuset(b *testing.B) {
	s := set.NewRWMuSet()
	for i := 0; i < b.N; i++ {
		s.Add(i)
		s.Remove(i)
		s.Add(i)
		s.Remove(i)
		s.Add(i)
		s.Remove(i)
		s.Add(i)
		s.Remove(i)
		s.Add(i)
		s.Has(i)
	}

}

func BenchmarkReadWrite5050RWMuset(b *testing.B) {
	s := set.NewRWMuSet()
	for i := 0; i < b.N; i++ {
		s.Add(i)
		s.Remove(i)
		s.Add(i)
		s.Remove(i)
		s.Add(i)
		s.Has(i)
		s.Has(i)
		s.Has(i)
		s.Has(i)
		s.Has(i)
	}
}

func BenchmarkReadWrite1090Muset(b *testing.B) {
	s := set.NewMuSet()
	for i := 0; i < b.N; i++ {
		s.Add(i)
		s.Has(i)
		s.Has(i)
		s.Has(i)
		s.Has(i)
		s.Has(i)
		s.Has(i)
		s.Has(i)
		s.Has(i)
		s.Has(i)
	}
}
