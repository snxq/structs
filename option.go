package structs

// Option Struct's option.
type Option func(*Struct) *Struct

// WithIgnoreZeroValue ignore zero value.
func WithIgnoreZeroValue() Option {
	return func(s *Struct) *Struct {
		s.ignoreZero = true
		return s
	}
}

// WithTagName use the provided string as tag name.
func WithTagName(tag string) Option {
	return func(s *Struct) *Struct {
		s.tagName = tag
		return s
	}
}

// WithFields .
func WithFields(fields ...string) Option {
	return func(s *Struct) *Struct {
		s.fields = fields
		return s
	}
}
