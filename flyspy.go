package flyspy

import (
    "time"
)

// Arguments map of arguments and values
type Arguments map[string]interface{}


// MethodCallMap map of method call data
type MethodCallMap map[string][]Invocation

// Invocation represents information for a method call
type Invocation struct {
    Args Arguments
    Timestamp time.Time
}

// Method stores the method calls and arguments of a struct
type Method struct {
    Calls MethodCallMap
}

// NewMethod returns a new Method spy
func NewMethod() Method {
    c:= make(MethodCallMap)

    return Method{
        Calls: c,
    }
}

// Clean clears method call map
func (m *Method) Clean() {
    m.Calls = make(MethodCallMap)
}

// GetCalls returns the invocation list for the given method
func (m *Method) GetCalls(method string) []Invocation {
    if i, ok := m.Calls[method]; ok {
        return i
    }

    return nil
}

// AddCall registers new invocation for the given method
func (m *Method) AddCall(method string, a Arguments) {

    i := Invocation{
        Args: a,
        Timestamp: time.Now(),
    }

    m.Calls[method] = append(m.Calls[method], i)
}
