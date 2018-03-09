// Generated by: gen
// TypeWriter: slice
// Directive: +gen on *Role

package entity

import "errors"

// RoleSlice is a slice of type *Role. Use it where you would use []*Role.
type RoleSlice []*Role

// Where returns a new RoleSlice whose elements return true for func. See: http://clipperhouse.github.io/gen/#Where
func (rcv RoleSlice) Where(fn func(*Role) bool) (result RoleSlice) {
	for _, v := range rcv {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

// All verifies that all elements of RoleSlice return true for the passed func. See: http://clipperhouse.github.io/gen/#All
func (rcv RoleSlice) All(fn func(*Role) bool) bool {
	for _, v := range rcv {
		if !fn(v) {
			return false
		}
	}
	return true
}

// Any verifies that one or more elements of RoleSlice return true for the passed func. See: http://clipperhouse.github.io/gen/#Any
func (rcv RoleSlice) Any(fn func(*Role) bool) bool {
	for _, v := range rcv {
		if fn(v) {
			return true
		}
	}
	return false
}

// First returns the first element that returns true for the passed func. Returns error if no elements return true. See: http://clipperhouse.github.io/gen/#First
func (rcv RoleSlice) First(fn func(*Role) bool) (result *Role, err error) {
	for _, v := range rcv {
		if fn(v) {
			result = v
			return
		}
	}
	err = errors.New("no RoleSlice elements return true for passed func")
	return
}