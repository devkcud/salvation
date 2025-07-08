package salvation

import (
	"errors"
	"fmt"
	"reflect"
)

// Possibly is a generic wrapper that represents an optional value of type T.
// It mimics constructs like Rust's Option<T> or Haskell's Maybe,
// allowing for safe handling or potentially nil or invalid values.
//
// Possibly tracks whether a value is effectively "Nothing" (nil) using reflection
// for Go types that support nil (pointers, slices, maps, etc).
type Possibly[T any] struct {
	value              T
	precomputedNothing bool

	opts PossiblyConfig
}

// PossiblyConfig defines configuration options that affect how
// Possibly[T] determines whether a value is considered "Nothing" (i.e., nil).
type PossiblyConfig struct {
	// Precompute enables caching of the "Nothingness" state at construction time.
	// When true, Possibly[T] will evaluate and store the result of IsNothing()
	// at creation time, avoiding reflection during future calls.
	//
	// Note: If the original value changes after creation (e.g., pointer set to nil),
	// the cached result may become stale. Only enable this if the wrapped value is immutable.
	// You can still force a recache using TryRecomputeIsNothing() method.
	Precompute bool

	// TreatZeroSliceAsSomething determines whether empty slices should be treated
	// as valid ("Something") rather than nil-equivalent ("Nothing").
	//
	// When true, an empty slice (e.g., var s []T) will *not* be considered "Nothing".
	// When false, empty slices behave like nil slices and are considered "Nothing".
	//
	// This only affects types of kind reflect.Slice.
	TreatZeroSliceAsSomething bool
}

// nilableTypes defines which kinds in Go can be meaningfully nil.
var nilableTypes = map[reflect.Kind]struct{}{
	reflect.Pointer:       {},
	reflect.UnsafePointer: {},
	reflect.Interface:     {},
	reflect.Slice:         {},
	reflect.Map:           {},
	reflect.Chan:          {},
	reflect.Func:          {},
}

// NewPossibility wraps a value of type T and returns an Possibly[T].
//
// Example:
//
//	opt := NewPossibility[*MyStruct](nil)
//	if opt.IsNothing() { ... }
func NewPossibility[T any](value T) Possibly[T] {
	return NewPossibilityWithConfig(value, PossiblyConfig{})
}

// NewPossibilityWithConfig wraps a value of type T and returns an Possibly[T].
// It uses the configuration to affect the behaviour of Possibly[T] (i.e., caching, design rules).
//
// Example:
//
//	opt := NewPossibilityWithConfig[*MyStruct](nil, &PossiblyConfig{Precompute: true})
//	if opt.IsNothing() { ... }
func NewPossibilityWithConfig[T any](value T, opts PossiblyConfig) Possibly[T] {
	opt := Possibly[T]{value: value, opts: opts}
	opt.TryRecomputeIsNothing() // NOTE: No need for checking. If recomputation is disabled, then it won't cache and won't panic.
	return opt
}

// IsPrecomputed returns true if Possibly[T] is using a cached "Nothing" value.
func (p Possibly[T]) IsPrecomputed() bool {
	return p.opts.Precompute
}

// IsSomething returns true if the Possibly contains a non-nil, valid value.
//
// Note: For non-nilable types (like int, struct), IsSomething always returns true.
func (p Possibly[T]) IsSomething() bool {
	return !p.IsNothing()
}

// computeIsNothing evaluates whether a given value of type T should be
// considered "Nothing" (i.e., nil or invalid), using Go reflection.
//
//   - For nilable types (e.g., pointers, slices, maps), it returns true if the value is nil.
//   - For non-nilable types (e.g., structs, ints), it always returns false.
//   - If TreatZeroSliceAsSomething is true, then an empty slice (zero value)
//     is considered valid (i.e., not "Nothing").
func computeIsNothing[T any](val T, treatZeroSliceAsSomething bool) bool {
	value := reflect.ValueOf(val)
	kind := value.Kind()

	if !value.IsValid() {
		return true
	}

	if _, ok := nilableTypes[kind]; !ok {
		return false
	}

	if treatZeroSliceAsSomething && kind == reflect.Slice && value.IsZero() {
		return false
	}

	return value.IsNil()
}

// TryRecomputeIsNothing forcibly updates the cached nil status, if precomputation is enabled.
// Returns an error if precomputation is disabled.
func (p *Possibly[T]) TryRecomputeIsNothing() error {
	if !p.IsPrecomputed() {
		return errors.New("precomputation is disabled")
	}
	p.precomputedNothing = computeIsNothing(p.value, p.opts.TreatZeroSliceAsSomething)
	return nil
}

// IsNothing returns true if the Possibly contains a nil or invalid value.
// For nilable types (pointers, slices, interfaces, etc), it checks
// if the internal value is nil. For non-nilable types, it always returns false.
func (p Possibly[T]) IsNothing() bool {
	if p.IsPrecomputed() {
		return p.precomputedNothing
	}
	return computeIsNothing(p.value, p.opts.TreatZeroSliceAsSomething)
}

// Reveal returns the value and a bool indicating presence.
func (p Possibly[T]) Reveal() (T, bool) {
	return p.value, !p.IsNothing()
}

// MustReveal returns the underlying value if it exists.
// It panics if the Possibly is Nothing.
func (p Possibly[T]) MustReveal() T {
	if p.IsNothing() {
		panic("attempted to unwrap Nothing value")
	}
	return p.value
}

// SafeReveal returns the underlying value as a pointer if it exists,
// or returns an error if the Possibly is Nothing.
func (p Possibly[T]) SafeReveal() (*T, error) {
	if p.IsNothing() {
		return nil, errors.New("attempted to unwrap Nothing value")
	}
	return &p.value, nil
}

// RevealOrElse returns the underlying value if it's present, or the fallback value otherwise.
func (p Possibly[T]) RevealOrElse(value T) T {
	if p.IsNothing() {
		return value
	}
	return p.value
}

// String returns a string representation of the Possibly[T].
// If the value is considered "Nothing", it returns "<Nothing>".
// Otherwise, it returns "<Something: VALUE>", using fmt.Sprintf on the underlying value.
//
// This method is useful for debugging and logging.
func (p Possibly[T]) String() string {
	if p.IsNothing() {
		return "<Nothing>"
	}
	return fmt.Sprintf("<Something: %v>", p.value)
}
