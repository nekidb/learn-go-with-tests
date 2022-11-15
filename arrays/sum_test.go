package main

import "testing"
import "reflect"

func TestSum(t *testing.T) {
	t.Run("slice of numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7}

		got := Sum(numbers)
		want := 28
		
		if got != want {
			t.Errorf("got %d, want %d given %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	}
	
	t.Run("3 slices of numbers", func(t *testing.T) {
		got := SumAll([]int{1, 2, 3}, []int{4, 5}, []int{})
		want := []int{6, 9, 0}
		checkSums(t, got, want)
	})
	
	t.Run("1 slice of numbers", func(t *testing.T) {
		got := SumAll([]int{7, 8})
		want := []int{15}
		checkSums(t, got, want)
	})
}

func TestSumAllTails(t *testing.T) {
	t.Run("3 slices of numbers", func(t *testing.T) {
		got := SumAllTails([]int{0, 1, 2, 3, 4, 5}, []int{}, []int{1})
		want := []int{15, 0, 0}
		
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func BenchmarkSumAll(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumAll([]int{1, 2, 3, 4, 5, 6}, []int{543, 23, 87, 23, 786}, []int{0, 0, 0, 0, 0, 0, 0, 0, 0})
	}
}
