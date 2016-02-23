package robots_test

import (
	"os"
	"testing"

	"github.com/amaxwellblair/crud_app/app/models"
)

// Testing functionality for robot read and create methods
func TestStore_CreateRobot_CreateandRetrieveARobot(t *testing.T) {
	s := MustOpenStore()
	defer Close(s)
	name := "Roboto"
	function := "Bend things that are hard to bend"

	s.CreateRobot(name, function)
	all, err := s.All()
	if err != nil {
		t.Fatal(err)
	}

	if all[len(all)-1].Name != "Roboto" {
		t.Fatalf("unexpected name %s", all[len(all)-1].Name)
	}
}

func TestStore_FindRobot_RetrieveASpecificRobot(t *testing.T) {
	s := MustOpenStore()
	defer Close(s)
	name := "Roboto"
	function := "Bend things that are hard to bend"
	s.CreateRobot(name, function)

	r, err := s.Robot(1)
	if err != nil {
		t.Fatal(err)
	}

	if r.Name != "Roboto" {
		t.Fatalf("unexpected name %s", r.Name)
	}

}

func MustOpenStore() *robots.Store {
	s := robots.NewStore("../../db/test.db")
	s.Open()
	return s
}

func Close(s *robots.Store) {
	s.Close()
	os.Remove("../../db/test.db")
}
