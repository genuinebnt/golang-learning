package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	checkSums := func(t testing.TB, got, want int) {
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	}

	checkSumSlices := func(t testing.TB, got, want []int) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d", got, want)
		}
	}

	t.Run("collection of 5 numbers", func(t *testing.T) {
		t.Helper()
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		checkSums(t, got, want)
	})

	t.Run("collection of any n numbers", func(t *testing.T) {
		t.Helper()
		numbers := []int{3, 2, 5}

		got := Sum(numbers)
		want := 10

		checkSums(t, got, want)
	})

	t.Run("takes slices and return new slice", func(t *testing.T) {
		t.Helper()
		got := SumAll([]int{1, 2}, []int{2, 3, 4})
		want := []int{3, 9}

		checkSumSlices(t, got, want)
	})
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d", got, want)
		}
	}

	t.Run("sum of slices of tails", func(t *testing.T) {
		t.Helper()
		got := SumAllTails([]int{1, 2}, []int{2, 4, 2}, []int{1}, []int{5, 3}, []int{3})
		want := []int{2, 6, 0, 3, 0}

		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		t.Helper()
		got := SumAllTails([]int{}, []int{3, 5, 6}, []int{2})
		want := []int{0, 11, 0}

		checkSums(t, got, want)
	})

}
