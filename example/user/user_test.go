package user

import (
    "testing"
    "github.com/tariqc80/flyspy"
)

type TestUserProvider struct {
    *flyspy.Spy
}

func NewTestUserProvider() *TestUserProvider{
    return &TestUserProvider{
        Spy: flyspy.New(),
    }
}

func (t *TestUserProvider) Get(id int, u Model) error {

    user := u.(*User)

    // Call `AddCall` to track this method call within out spy.
    t.AddCall("Get", flyspy.Arguments{
        "id": id,
        "user": user,
    })

    user.name = "Fred"
    user.email = "fred.flintstone@example.com"
    user.password = "bambampebbles"

    return nil
}

func (t *TestUserProvider) Store(u Model) error {

    user := u.(*User)

    // track this method call w/ agruments in our spy
    t.AddCall("Store", flyspy.Arguments{
        "user": user,
    })

    return nil
}

func TestFind (t *testing.T) {
    p := NewTestUserProvider()

    u := NewUser(p)

    u.Find(1)

    // Get all calls for the method `Get`
    calls := p.GetCalls("Get")

    // Get the args map from the first call
    args := calls[0].Args

    // Check if Get was called with the expected value for the argument `id`
    if args["id"] != 1 {
        t.Errorf("DataProvider.Get expected to be called with argument `id` of value `1`; got `%d`", args["id"])
    }

    if u.name != "Fred" {
        t.Errorf("name property expected to be `Fred`; got `%s`", u.name)
    }
}
