package salvation

import "testing"

func TestMatcher_CaseExecutesFirstMatch(t *testing.T) {
	var calls []int
	p := NewPossibility(5)

	p.Match().
		Case(func(v int) bool {
			calls = append(calls, 1)
			return v > 3
		}, func(v int) {
			calls = append(calls, 10)
		}).
		Case(func(v int) bool {
			calls = append(calls, 2)
			return v > 0
		}, func(v int) {
			calls = append(calls, 20)
		})

	if len(calls) != 2 {
		t.Errorf("expected 2 calls, got %d: %v", len(calls), calls)
	}
	if calls[0] != 1 || calls[1] != 10 {
		t.Errorf("unexpected calls order or values: %v", calls)
	}
}

func TestMatcher_MatchAllCases(t *testing.T) {
	var calls []int
	conf := MatcherConfig{MatchAllCases: true}
	p := NewPossibility(5)

	p.MatchWithConfig(conf).
		Case(func(v int) bool {
			calls = append(calls, 1)
			return v > 3
		}, func(v int) {
			calls = append(calls, 10)
		}).
		Case(func(v int) bool {
			calls = append(calls, 2)
			return v > 0
		}, func(v int) {
			calls = append(calls, 20)
		})

	if len(calls) != 4 {
		t.Errorf("expected 4 calls, got %d: %v", len(calls), calls)
	}
	expected := []int{1, 10, 2, 20}
	for i, v := range expected {
		if calls[i] != v {
			t.Errorf("at index %d expected %d, got %d", i, v, calls[i])
		}
	}
}

func TestMatcher_DefaultExecutedWhenMatched(t *testing.T) {
	var defaultCalled bool
	p := NewPossibility(5)

	p.Match().
		Case(func(v int) bool { return v > 1 }, func(v int) {}).
		Default(func(opt Possibly[int]) {
			defaultCalled = true
		})

	if !defaultCalled {
		t.Errorf("expected default to be called when a case matched")
	}
}

func TestMatcher_DefaultNotExecutedWhenNotMatched(t *testing.T) {
	var defaultCalled bool
	p := NewPossibility(1)

	p.Match().
		Case(func(v int) bool { return v > 10 }, func(v int) {}).
		Default(func(opt Possibly[int]) {
			defaultCalled = true
		})

	if defaultCalled {
		t.Errorf("expected default not to be called when no case matched")
	}
}

func TestMatcher_WithNothingPossibly(t *testing.T) {
	var calls []int
	var defaultCalled bool
	var ptr *int
	p := NewPossibility(ptr) // *int nil => Nothing

	p.Match().
		Case(func(v *int) bool {
			calls = append(calls, 1)
			return true
		}, func(v *int) {
			calls = append(calls, 10)
		}).
		Default(func(opt Possibly[*int]) {
			defaultCalled = true
		})

	if len(calls) != 0 {
		t.Errorf("expected no calls for cases when Possibly is Nothing, got: %v", calls)
	}
	if defaultCalled {
		t.Errorf("expected default not called for Nothing Possibly")
	}
}
