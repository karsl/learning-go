package stub

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestLogicGetPetNames(t *testing.T) {
	data := []struct {
		name     string
		userID   string
		petNames []string
	}{
		{"case1", "1", []string{"Bubbles"}},
		{"case2", "2", []string{"Stampy", "Snowball II"}},
		//{"case3", "3", nil},
	}
	l := Logic{GetPetNamesStub{}}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			petNames, err := l.GetPetNames(d.userID)
			if err != nil {
				t.Error(err)
			}
			if diff := cmp.Diff(d.petNames, petNames); diff != "" {
				t.Error(diff)
			}
		})
	}

	myFunction(l.Entities)
}

func myFunction(entities Entities) {
	user, err := entities.GetUser("123")
	fmt.Println("user: ", user, " err: ", err)
}
