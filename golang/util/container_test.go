package util

import (
	"errors"
	"os"
	"testing"
)

func TestContainerUtils(t *testing.T) {
	if os.Getenv("CIRCLECI") == "true" {
		t.Skipf("Skipping container func tests in CI environment")
	}

	hasImage, err := HasImage("foo", "bar")
	if err != nil {
		t.Errorf("HasImage(foo,bar) failed: %v", err)
	} else if hasImage {
		t.Errorf("HasImage(foo,bar) failed to return false")
	}
	_, err = GetLocalImageDigest("foo", "bar")
	if !errors.Is(err, ErrImageNotFound) {
		t.Errorf("GetLocalImageDigest(foo,bar) failed to return ErrImageNotFound: %v",
			err)
	}

	remoteDigest, err := GetRemoteImageDigest("bopmatic/build", "latest")
	if err != nil {
		t.Errorf("Failed to get remote digest: %v", err)
	}
	t.Logf("Remote digest is %v", remoteDigest)

	localDigest, err := GetLocalImageDigest("bopmatic/build", "latest")
	var wasNotFound bool
	if errors.Is(err, ErrImageNotFound) {
		wasNotFound = true
		t.Logf("Local digest was not found")
	} else if err != nil {
		t.Errorf("Failed to get local digest: %v", err)
	} else {
		t.Logf("Local digest is  %v", localDigest)
	}

	needUpdate, err := DoesLocalImageNeedUpdate("bopmatic/build", "latest")
	if err != nil {
		t.Errorf("DoesLocalImageNeedUpdate() failed unexpectedly: %v", err)
	} else if wasNotFound && !needUpdate {
		t.Errorf("DoesLocalImageNeedUpdate() should have returned true")
	} else {
		if needUpdate && localDigest == remoteDigest {
			t.Errorf("DoesLocalImageNeedUpdate() should have returned false")
		}
		if !needUpdate && localDigest != remoteDigest {
			t.Errorf("DoesLocalImageNeedUpdate() should have returned true")
		}
	}

}
