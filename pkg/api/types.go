// types.go contains all common types used along the engine and any application
// using the engine.
package api

// -----------------------------------------------------------------------------
//
// Public package types
//
// -----------------------------------------------------------------------------

// Callback type is the type used for any function without argument and without
// any return value.
type Callback func()

// Command type is the type used for any function being called by any widget.
type Command func(entity any, label string, position *Point, args ...any) bool

// Args type is the type used for the list of arguments passed to a Command
// called by any widget.
type Args []any

type (
	Attr      uint64
	Key       uint16
	Modifier  uint8
	EventType uint8
)
