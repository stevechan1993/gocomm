package xstr

import (
	"fmt"
	"testing"
)

func TestJoinInts(t *testing.T) {
	// test empty slice
	is := []int64{}
	s := JoinInts(is)
	if s != "" {
		t.Fatal(fmt.Sprintf("input:%v,output:%s,result is incorrect", is, s))
	}
	// test len(slice)==1
	is = []int64{1}
	s = JoinInts(is)
	if s != "1" {
		t.Fatal(fmt.Sprintf("input:%v,output:%s,result is incorrect", is, s))
	}
	// test len(slice)>1
	is = []int64{1, 2, 3}
	s = JoinInts(is)
	if s != "1,2,3" {
		t.Fatal(fmt.Sprintf("input:%v,output:%s,result is incorrect", is, s))
	}
}

func TestSplitInts(t *testing.T) {
	// test empty slice
	s := ""
	is, err := SplitInts(s)
	if err != nil || len(is) != 0 {
		t.Error(err)
	}
	// test split int64
	s = "1,2,3"
	is, err = SplitInts(s)
	if err != nil || len(is) != 3 {
		t.Fatal(err)
	}
}

func BenchmarkJoinInts(b *testing.B) {
	is := make([]int64, 10000, 10000)
	for i := int64(0); i < 10000; i++ {
		is[i] = i
	}
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			JoinInts(is)
		}
	})
}
