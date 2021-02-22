package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("collection of any n numbers", func(t *testing.T) {
		numbers := []int{3, 2, 5}

		got := Sum(numbers)
		want := 10

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("takes slices and return new slice", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{2, 3, 4})
		want := []int{3, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("sum of slices of tails", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{2, 4, 2}, []int{1}, []int{5, 3}, []int{3})
		want := []int{2, 6, 0, 3, 0}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
