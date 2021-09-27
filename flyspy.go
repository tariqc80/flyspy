package flyspy

import (
    "time"
    "sync"
    "reflect"
)

// Invocations map of method calls
type Invocations map[string][]Invocation

// Invocation represents information for a method call
type Invocation struct {
    Args []interface{}
    Timestamp time.Time
}

// With returns true if the Invocation Args equal the `values`
func (i *Invocation) With(values ...interface{}) bool {
    return reflect.DeepEqual(values, i.Args)
}

// With returns true if the Invocation Args equal the `values`
// func (i *Invocation) With(values ...interface{}) bool {
//     equal := false

//     for index, val := range values {
//         equal = reflect.DeepEqual(val, i.Args[index])
//     }

//     return equal
// }

// GetArgs returns the slice of Arguments for the Invocation
func (i *Invocation) GetArgs() []interface{} {
    return i.Args
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

// Multiple returns the invocation list if `method` was called `num` times
func (m *Spy) Multiple(method string, num int) []Invocation {

    calls := m.GetCalls(method)

    // if the number of calls is not equal to `num` return nil
    if len(calls) != num {
        return nil
    }

    return calls
}

// Once return Invocation if there is only one call to `method`
func (m *Spy) Once(method string) *Invocation {
    calls := m.Multiple(method, 1)

    // if the method was not called once return an empty Invocation
    if calls == nil {
        return &Invocation{}
    }

    return &calls[0]
}

// RecordCall registers new invocation for the given method
func (m *Spy) RecordCall(method string, a ...interface{}) {
    m.Lock()
    defer m.Unlock()

    i := Invocation{
        Args: a,
        Timestamp: time.Now(),
    }

    m.Calls[method] = append(m.Calls[method], i)
}
