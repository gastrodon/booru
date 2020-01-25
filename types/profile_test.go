package types

import (
	"testing"
	"time"
)

func Test_GetProfile(test *testing.T) {
	var me Profile
	var authed bool
	var err error
	me, authed, err = test_auth.GetProfile()
	if err != nil {
		test.Fatal(err)
	}

	if !authed {
		test.Error("Client is not authed!")
	}

	if me.ID != test_user.ID {
		test.Errorf("profile id mismatch! have: %d, want: %d", me.ID, test_user.ID)
	}

	if me.Client.Host != test_auth.Host {
		test.Errorf("Client host mismatch! have: %s, want: %s", me.Client.Host, test_auth.Host)
	}
}

func Test_UpdatedAt_Profile(test *testing.T) {
	var stamp *time.Time
	var err error
	stamp, err = test_profile.UpdatedAt()
	if err != nil {
		test.Fatal(err)
	}

	if stamp.Unix() >= now {
		test.Errorf("test_profile.UpdatedAt is in the future: %d", stamp.Unix())
	}
}

func Test_CreatedAt_Profile(test *testing.T) {
	var stamp *time.Time
	var err error
	stamp, err = test_profile.CreatedAt()
	if err != nil {
		test.Fatal(err)
	}

	if stamp.Unix() >= now {
		test.Errorf("test_profile.CreatedAt is in the future: %d", stamp.Unix())
	}
}

func Test_LastForumReadAt_Profile(test *testing.T) {
	var stamp *time.Time
	var err error
	stamp, err = test_profile.LastForumReadAt()
	if err != nil {
		test.Fatal(err)
	}

	if stamp.Unix() >= now {
		test.Errorf("test_profile.LastForumReadAt is in the future: %d", stamp.Unix())
	}
}

func Test_LastLoggedInAt_Profile(test *testing.T) {
	var stamp *time.Time
	var err error
	stamp, err = test_profile.LastLoggedInAt()
	if err != nil {
		test.Fatal(err)
	}

	if stamp.Unix() >= now {
		test.Errorf("test_profile.LastLoggedInAt is in the future: %d", stamp.Unix())
	}
}
