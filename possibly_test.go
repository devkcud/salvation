package salvation

import (
	"testing"
)

func TestPossibly_NonNil_String(t *testing.T) {
	opt := NewPossibility("Hello, world!")
	if opt.IsNothing() {
		t.Errorf("Expected value to not be nil")
	}
}

func TestPossibly_NonNil_Slice(t *testing.T) {
	slice := []any{}
	opt := NewPossibility(slice)
	if opt.IsNothing() {
		t.Errorf("Expected value to not be nil")
	}
}

func TestPossibly_NonNil_Pointer(t *testing.T) {
	slice := new([]any)
	opt := NewPossibility(slice)
	if opt.IsNothing() {
		t.Errorf("Expected value to not be nil")
	}
}

func TestPossibly_Nil_Slice(t *testing.T) {
	var slice []any
	opt := NewPossibility(slice)
	if opt.IsSomething() {
		t.Errorf("Expected value to not be nil")
	}
}

func TestPossibly_NonNil_ZeroSlice(t *testing.T) {
	var slice *int
	var one = 1
	slice = &one
	opt := NewPossibility(slice)
	slice = nil
	if opt.IsNothing() {
		t.Errorf("Expected value to not be nil")
	}
}

func TestPossibly_Nil_Struct(t *testing.T) {
	var structure *struct{} = nil
	opt := NewPossibility(structure)
	if opt.IsSomething() {
		t.Errorf("Expected value to be nil")
	}
}

type Does interface {
	Do()
}

type Impl struct{}

func (i Impl) Do() {}

func TestPossibly_Nil_Interface(t *testing.T) {
	var x Does
	opt := NewPossibility(x)
	if opt.IsSomething() {
		t.Errorf("Expected value to be nil")
	}
}

func TestPossibly_NonNil_Interface(t *testing.T) {
	var x Does = Impl{}
	opt := NewPossibility(x)
	if opt.IsNothing() {
		t.Errorf("Expected value to not be nil")
	}
}

func TestPossibly_Reveal_Some(t *testing.T) {
	opt := NewPossibility("hello")

	val, ok := opt.Reveal()
	if !ok {
		t.Errorf("Expected Some, got None")
	}
	if val != "hello" {
		t.Errorf("Expected 'hello', got '%v'", val)
	}
}

func TestPossibly_Reveal_None(t *testing.T) {
	var val *string = nil
	opt := NewPossibility(val)

	result, ok := opt.Reveal()
	if ok {
		t.Errorf("Expected None, got Some: %v", result)
	}
}

func TestPossibly_MustReveal_Some(t *testing.T) {
	opt := NewPossibility(42)
	val := opt.MustReveal()
	if val != 42 {
		t.Errorf("Expected 42, got %v", val)
	}
}

func TestPossibly_MustReveal_None_Panics(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected panic on MustReveal for None, but it did not panic")
		}
	}()

	var ptr *int = nil
	opt := NewPossibility(ptr)
	opt.MustReveal() // Should panic
}

func TestPossibly_SafeReveal_Some(t *testing.T) {
	opt := NewPossibility("functional sadness")
	ptr, err := opt.SafeReveal()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if ptr == nil || *ptr != "functional sadness" {
		t.Errorf("Expected pointer to 'functional sadness', got %v", ptr)
	}
}

func TestPossibly_SafeReveal_None(t *testing.T) {
	var thing *struct{} = nil
	opt := NewPossibility(thing)

	ptr, err := opt.SafeReveal()
	if err == nil {
		t.Errorf("Expected error, got none")
	}
	if ptr != nil {
		t.Errorf("Expected nil pointer, got %v", ptr)
	}
}
