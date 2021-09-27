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

func (t *TestUserProvider) Get(id int) (Model, error) {

    // Call `RecordCall` to track this method call within our spy.
    t.RecordCall("Get", id)

    var user User

    user.name = "Fred"
    user.email = "fred.flintstone@example.com"
    user.password = "bambampebbles"

    return &user, nil
}

func (t *TestUserProvider) Store(u Model) error {

    user := u.(*User)

    // track this method call w/ agruments in our spy
    t.RecordCall("Store", user)

    return nil
}

func TestFind (t *testing.T) {
    p := NewTestUserProvider()

    u := NewUser(p)

    u.Find(1)

    if p.Once("Get").With(1) == false {
        t.Errorf("DataProvider.Get expected to be called with argument of value `1``")
    }

    if u.name != "Fred" {
        t.Errorf("name property expected to be `Fred`; got `%s`", u.name)
    }
}
