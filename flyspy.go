package flyspy

import (
    "time"
    "sync"
)

// Arguments map of arguments and values
type Arguments map[string]interface{}


// Invocations map of method calls
type Invocations map[string][]Invocation

// Invocation represents information for a method call
type Invocation struct {
    Args Arguments
    Timestamp time.Time
}

// Spy stores the method calls and arguments of a struct
type Spy struct {
    Calls Invocations
    sync.RWMutex
}

// New returns a new Spy
func New() *Spy {
    return &Spy{
        Calls: make(Invocations),
    }
}

// Clean clears all the calls
func (m *Spy) Clean() {
    m.Lock()
    defer m.Unlock()

    m.Calls = make(Invocations)
}

// GetCalls returns the invocation list for the given method
func (m *Spy) GetCalls(method string) []Invocation {
    m.RLock()
    defer m.RUnlock()

    if i, ok := m.Calls[method]; ok {
        return i
    }

    return nil
}

// AddCall registers new invocation for the given method
func (m *Spy) AddCall(method string, a Arguments) {
    m.Lock()
    defer m.Unlock()

    i := Invocation{
        Args: a,
        Timestamp: time.Now(),
    }

    m.Calls[method] = append(m.Calls[method], i)
}
