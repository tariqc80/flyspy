# FlySpy - A Spy framework for Go testing

Package FlySpy provides a lightweight framework to add function spies to stubbed or mock structs in Go tests. This package is meant to be used with Go's built in `testing` package. It works with manually created dependency mocks.  It is similiar to this package https://pkg.go.dev/github.com/nyarly/spies

Alternatively, you can use the golang/mock package which is more robust, and leverages code generation. https://github.com/golang/mock

## Usage

Embed `FlySpy.Spy` into the struct you have methods that you wish to spy on.  Then construct the spy struct using `flyspy.New()`

Example
```

type Processor interface {
  Subprocess(int) error
}

type RemoteProcessor struct {}

func (p *RemoteProcessor) Subprocess(i int) error {
  // does remote processing, etc
}

type System struct {
  p Processor
  field1 string
  field2 int
}

func (s System) Process(i int) error {
  return s.p.Subprocess(i)
}

type TestRemoteProcessor struct {
  // Embed Spy here
  *flyspy.Spy
}

func (p TestRemoteProcessor) Subprocess(i int) error {
  // register method call with spy
  p.RecordCall("Subprocess", i)

  // do some work
}

func TestProcess(t *testing.T) {

  // Use flyspy.New() to construct a Spy
  processorMock := &TestRemoteProcessor{
    Spy: flyspy.New(),
  }

  // Inject dependency into System
  system := &System{
    p: processorMock,
  }

  // call test func
  system.Process(2)

  // call Spy.Once() to assert for a single call to the method
  if processorMock.Once("Subprocess").With(2) == false {
    t.Error("Subprocess was not called once with correct value")
  }

  // OR you can use Spy.GetCalls to return a slice of all the calls
  calls := processorMock.GetCalls("Subprocess")

  // then you can access the calls and agruments
  // get the first argument of the first call
  if calls[0].Args[0] != 2 {
    t.Error("Unexpected value for argument 0 in call to Subprocess")
  }

}
