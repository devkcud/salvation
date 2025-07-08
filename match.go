package salvation

// Matcher is a control-flow utility for Possibly[T] values, enabling chained case handling.
// It mimics pattern matching from languages like Rust and Kotlin but does so in an aggressively
// Go-inappropriate way, allowing you to define match-like behavior for conditional logic.
//
// Matcher will execute the first matching case unless MatchAllCases is set to true,
// in which case it evaluates all matching cases.
type Matcher[T any] struct {
	value   Possibly[T]
	matched bool

	opts MatcherConfig
}

// MatcherConfig controls evaluation strategy for Matcher.
type MatcherConfig struct {
	// MatchAllCases, when true, allows all Case predicates to be evaluated.
	// If false (default), only the first matching Case is executed.
	MatchAllCases bool
}

// Match creates a Matcher from a Possibly[T] value.
// Use it to build a fluent-style series of Case and Default handlers.
func (p Possibly[T]) Match() Matcher[T] {
	return p.MatchWithConfig(MatcherConfig{})
}

// MatchWithConfig creates a Matcher with explicit configuration options.
func (p Possibly[T]) MatchWithConfig(opts MatcherConfig) Matcher[T] {
	return Matcher[T]{value: p, opts: opts}
}

// Case tests a predicate against the Possibly's value. If the predicate returns true,
// the corresponding action is executed.
//
// If MatchAllCases is false, only the first matching Case is executed.
// If the Possibly is Nothing, Case is skipped silently, as if it never existed.
func (m Matcher[T]) Case(predicate func(T) bool, action func(T)) Matcher[T] {
	if (m.matched && !m.opts.MatchAllCases) || m.value.IsNothing() {
		return m
	}

	value := m.value.MustReveal()
	if predicate(value) {
		action(value)
		m.matched = true
	}
	return m
}

// Default provides a fallback action to run if no Case matched.
// It receives the original Possibly[T] value, allowing inspection or logging.
func (m Matcher[T]) Default(action func(Possibly[T])) {
	if m.matched {
		return
	}

	action(m.value)
}
