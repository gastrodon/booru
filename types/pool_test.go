package types

import (
	"testing"
	"time"
)

func Test_CreatedAt_Pool(test *testing.T) {
	var now int64 = time.Now().Unix()

	var stamp *time.Time
	var err error
	stamp, err = test_pool.CreatedAt()
	if err != nil {
		test.Fatal(err)
	}

	if stamp.Unix() >= now {
		test.Errorf("pool.CreatedAt is in the future: %d", stamp.Unix())
	}
}

func Test_UpdatedAt_Pool(test *testing.T) {
	var now int64 = time.Now().Unix()

	var updated *time.Time
	var err error
	updated, err = test_pool.UpdatedAt()
	if err != nil {
		test.Fatal(err)
	}

	if updated == nil {
		return
	}

	if updated.Unix() >= now {
		test.Errorf("pool.UpdatedAt is in the future: %d", updated.Unix())
	}

	var created *time.Time
	created, err = test_pool.CreatedAt()
	if err != nil {
		test.Fatal(err)
	}

	if updated.Unix() < created.Unix() {
		test.Errorf("pool.UpdatedAt is before pool.CreatedAt: %d < %d", updated.Unix(), created.Unix())
	}
}
