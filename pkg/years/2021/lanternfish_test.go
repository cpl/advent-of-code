package y2021

import (
	"reflect"
	"testing"
)

func TestLanternFishBank(t *testing.T) {
	t.Parallel()

	t.Run("parse", func(t *testing.T) {
		bank := ParseLanternFishBank("3,4,3,1,2")
		want := [9]int{0, 1, 1, 2, 1, 0, 0, 0, 0}

		if !reflect.DeepEqual(bank.fish, want) {
			t.Errorf("fish bank should be %v, got %v", want, bank.fish)
		}
	})

	t.Run("age", func(t *testing.T) {
		bank := &LanternFishBank{fish: [9]int{0, 1, 1, 2, 1, 0, 0, 0, 0}}

		bank.Age()
		want := [9]int{1, 1, 2, 1, 0, 0, 0, 0, 0}
		if !reflect.DeepEqual(bank.fish, want) {
			t.Errorf("fish bank should be %v, got %v", want, bank.fish)
		}

		bank.Age()
		want = [9]int{1, 2, 1, 0, 0, 0, 1, 0, 1}
		if !reflect.DeepEqual(bank.fish, want) {
			t.Errorf("fish bank should be %v, got %v", want, bank.fish)
		}

		bank.Age()
		want = [9]int{2, 1, 0, 0, 0, 1, 1, 1, 1}
		if !reflect.DeepEqual(bank.fish, want) {
			t.Errorf("fish bank should be %v, got %v", want, bank.fish)
		}
	})

	t.Run("age days", func(t *testing.T) {
		t.Run("18", func(t *testing.T) {
			bank := ParseLanternFishBank("3,4,3,1,2")
			bank.AgeDays(18)
			t.Logf("%d", bank.Total())
			if bank.Total() != 26 {
				t.Errorf("fish bank should be 26, got %d", bank.Total())
			}
		})
		t.Run("80", func(t *testing.T) {
			bank := ParseLanternFishBank("3,4,3,1,2")
			bank.AgeDays(80)
			t.Logf("%d", bank.Total())
			if bank.Total() != 5934 {
				t.Errorf("should be 5934 fish, got %d", bank.Total())
			}
		})
		t.Run("256", func(t *testing.T) {
			bank := ParseLanternFishBank("3,4,3,1,2")
			bank.AgeDays(256)
			t.Logf("%d", bank.Total())
			if bank.Total() != 26984457539 {
				t.Errorf("fish bank should be %d, got %d", 26984457539, bank.Total())
			}
		})
	})
}
