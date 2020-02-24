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
	OkDate(test, test_profile.UpdatedAt, "test_profile.UpdatedAt")

	var updated *time.Time
	var err error
	updated, err = test_profile.UpdatedAt()
	if err != nil {
		test.Fatal(err)
	}

	var created *time.Time
	created, err = test_post.CreatedAt()
	if err != nil {
		test.Fatal(err)
	}

	if updated.Unix() < created.Unix() {
		test.Errorf("test_profile.UpdatedAt is before test_profile.CreatedAt: %d < %d", updated.Unix(), created.Unix())
	}
}

func Test_CreatedAt_Profile(test *testing.T) {
	OkDate(test, test_profile.CreatedAt, "test_profile.CreatedAt")
}

func Test_LastForumReadAt_Profile(test *testing.T) {
	OkDate(test, test_profile.LastForumReadAt, "test_profile.LastForumReadAt")
}

func Test_LastLoggedInAt_Profile(test *testing.T) {
	OkDate(test, test_profile.LastLoggedInAt, "test_profile.LastLoggedInAt")
}

func Test_GetUser_Profile(test *testing.T) {
	var me User
	var exists bool
	var err error
	me, exists, err = test_profile.GetUser()
	if err != nil {
		test.Fatal(err)
	}

	if !exists {
		test.Errorf("User #%d does not exist!", test_profile.ID)
	}

	if me.ID != test_profile.ID {
		test.Errorf("ID mismatch! have: %d, want: %d", me.ID, test_profile.ID)
	}

	if me.Name != test_profile.Name {
		test.Errorf("Name mismatch! have: %s, want: %s", me.Name, test_profile.Name)
	}
}
