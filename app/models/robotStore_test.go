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

func TestStore_Robot_RetrieveASpecificRobot(t *testing.T) {
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

func TestStore_UpdateRobot_UpdateASpecificRobot(t *testing.T) {
	s := MustOpenStore()
	defer Close(s)
	name := "Roboto"
	function := "Bend things that are hard to bend"
	s.CreateRobot(name, function)

	err := s.UpdateRobot(1, "Bender", "Bend things that are hard to bend")
	if err != nil {
		t.Fatal(err)
	}
	rbt, err := s.Robot(1)
	if err != nil {
		t.Fatal(err)
	}

	if rbt.Name != "Bender" {
		t.Fatalf("unexpected name %s", rbt.Name)
	}

}

func TestStore_UpdateRobot_DoesNotCreateRobotIfDoesNotExist(t *testing.T) {
	// TODO
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
