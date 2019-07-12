// This file was generated by github.com/nelsam/hel.  Do not
// edit this code by hand unless you *really* know what you're
// doing.  Expect any changes made manually to be overwritten
// the next time hel regenerates this file.

package plumbing_test

import (
	"code.cloudfoundry.org/loggregator/plumbing"
)

type mockFinder struct {
	NextCalled chan bool
	NextOutput struct {
		Ret0 chan plumbing.Event
	}
}

func newMockFinder() *mockFinder {
	m := &mockFinder{}
	m.NextCalled = make(chan bool, 100)
	m.NextOutput.Ret0 = make(chan plumbing.Event, 100)
	return m
}
func (m *mockFinder) Next() plumbing.Event {
	m.NextCalled <- true
	return <-m.NextOutput.Ret0
}