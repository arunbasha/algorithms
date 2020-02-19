package Algorithm

import "testing"

func TestNumInList(t *testing.T) {
	t.Run("NumList([1,2,3,4],3)", func(t *testing.T){
		if ! NumInList([]int{1,2,3,4}, 3) {
			t.Fatal("NumInList = false, wanted true")
		}
	})
	t.Run("NumList([1,2,3,4],1)", func(t *testing.T){
	if ! NumInList([]int{1,2,3,4}, 1) {
		t.Fatal("NumInList = false, wanted true")
	}
	})

		t.Run("NumList([1,2,3,4],4)", func(t *testing.T){
	if ! NumInList([]int{1,2,3,4}, 4) {
		t.Fatal("NumInList = false, wanted true")
	}
		})

	t.Run("NumList([1,2,3,4],5)", func(t *testing.T){
		if NumInList([]int{1,2,3,4}, 5) {
			t.Fatal("NumInList = false, wanted true")
		}
	})

	t.Run("NumList([1,2,3,4],-1)", func(t *testing.T) {
		if NumInList([]int{1, 2, 3, 4}, -1) {
			t.Fatal("NumInList = false, wanted true")
		}
	})

	t.Run("NumList([1,2,3,-4],-4)", func(t *testing.T) {
		if ! NumInList([]int{1, 2, 3, -4}, -4) {
			t.Fatal("NumInList = false, wanted true")
		}
	})

	t.Run("NumList([1,2,-3,4],-3)", func(t *testing.T) {
		if ! NumInList([]int{1, 2, -3, 4}, -3) {
			t.Fatal("NumInList = false, wanted true")
		}
	})

	t.Run("NumList([1,2,3,4],-1)", func(t *testing.T) {
		if ! NumInList([]int{-1, 2, 3, 4}, -1) {
			t.Fatal("NumInList = false, wanted true")
		}
	})


	t.Run("NumList([],2)", func(t *testing.T) {
		if NumInList([]int{}, 2) {
			t.Fatal("NumInList = false, wanted true")
		}
	})

	t.Run("NumList(nil,3)", func(t *testing.T) {
		if NumInList(nil, 3) {
			t.Fatal("NumInList = false, wanted true")
		}
	})
}
