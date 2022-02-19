package dredition

import (
	"os"
	"testing"
)

func TestReadNotification(t *testing.T) {
	f, err := os.Open("../testdata/notification.json")
	if err != nil {
		t.Error("failed to open file")
	}
	defer f.Close()
	n, err := ReadNotification(f)
	if err != nil {
		t.Errorf("failed to parse notification: %v", err)
	}

	if n.Event != "publish" {
		t.Error("event not publish")
	}

	var expected = "5d271809c4579c35f2ff903d"
	if n.Data.Product.Id != expected {
		t.Errorf("product id not %s", expected)
	}
	expected = "dev-front"
	if n.Data.Product.Name != expected {
		t.Errorf("product name not %s", expected)
	}
	expected = "frontpage"
	if n.Data.Product.Type != expected {
		t.Errorf("product type not %s", expected)
	}
	expected = "5d27181e3b33e81f2529d85e"
	if n.Data.Edition.Id != expected {
		t.Errorf("product type not %s", expected)
	}
	expected = "sandbox"
	if n.Data.Edition.Name != expected {
		t.Errorf("product name not %s", expected)
	}
}
