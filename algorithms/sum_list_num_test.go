package Algorithm

import "testing"

func TestSumListNum(t *testing.T) {
	t.Run("SumListNum[1,2,3,4,5]", func(t *testing.T) {
		sum, err := SumListNum([]int{1,2,3,4,5})
		if err == nil && sum != 15 {
			t.Fatalf("wanted (15,nil), got (%d, %+v)", sum, err)
		}
	})
	t.Run("SumListNum[1,2,3,4,-5]", func(t *testing.T) {
		sum, err := SumListNum([]int{1,2,3,4,-5})
		if err == nil && sum != 5 {
			t.Fatalf("wanted (5,nil), got (%d, %+v)", sum, err)
		}
	})
	t.Run("SumListNum[]", func(t *testing.T) {
		sum, err := SumListNum([]int{})
		if err == nil && sum != 0 {
			t.Fatalf("wanted (0,nil), got (%d, %+v)", sum, err)
		}
	})

	t.Run("SumListNum_nil", func(t *testing.T) {
		sum, err := SumListNum(nil)
		if err != nil && sum != 0 {
			t.Fatalf("wanted (0,err), got (%d, %+v)", sum, err)
		}
	})
}