package dredition

import (
	"os"
	"testing"
)

func TestReadNotification(t *testing.T) {
	f, err := os.Open("../testdata/notification-full.json")
	if err != nil {
		t.Error("failed to open file")
	}
	defer f.Close()
	n, err := ReadNotification(f)
	if err != nil {
		t.Errorf("failed to parse notification: %v", err)
	}

	if n.Event != "SphynxPostPublishEvent" {
		t.Error("event not SphynxPostPublishEvent")
	}

	var expected = "5d5a8cc11574df09cd20ecac"
	if n.Data.Product.Id != expected {
		t.Errorf("product id not %s", expected)
	}
	expected = "ekstrabladet-dk"
	if n.Data.Product.Name != expected {
		t.Errorf("product name not %s", expected)
	}
	expected = "frontpage"
	if n.Data.Product.Type != expected {
		t.Errorf("product type not %s", expected)
	}
	expected = "5d5a8cf857cd2009c74b6378"
	if n.Data.Edition.Id != expected {
		t.Errorf("edition id not %s", expected)
	}
	expected = "manuel-top"
	if n.Data.Edition.Name != expected {
		t.Errorf("edition name not %s", expected)
	}

	m := n.Data.Burned
	expected = "eb"
	if m.ClientID != expected {
		t.Errorf("ClientID was %s got %s", m.ClientID, expected)
	}
	expected = "5d5a8cf857cd2009c74b6378"
	if m.EditionID != expected {
		t.Errorf("EditionID was %s got %s", m.EditionID, expected)
	}
	expected = "fab7859c399e7721"
	if m.HTMLHash != expected {
		t.Errorf("HTMLHash was %s got %s", m.HTMLHash, expected)
	}
	expected = "2022-02-20T06:10:07.076Z"
	if m.HTMLUpdatedAt != expected {
		t.Errorf("HTMLUpdatedAt was %s got %s", m.HTMLUpdatedAt, expected)
	}
	expected = "https://smooth-storage.aptoma.no/users/drf-eb/files/sphynx/2022/2/20/fab7859c399e7721.html"
	if m.HTMLUrl != expected {
		t.Errorf("HTMLUrl was %s got %s", m.HTMLUrl, expected)
	}
	expected = "236a375ec45e1904"
	if m.CSSHash != expected {
		t.Errorf("CSSHash was %s got %s", m.CSSHash, expected)
	}
	expected = "2022-02-14T09:03:41.489Z"
	if m.CSSUpdatedAt != expected {
		t.Errorf("CSSUpdatedAt was %s got %s", m.CSSUpdatedAt, expected)
	}
	expected = "https://smooth-storage.aptoma.no/users/drf-eb/files/sphynx/2022/2/14/236a375ec45e1904.css"
	if m.CSSUrl != expected {
		t.Errorf("CSSUrl was %s got %s", m.CSSUrl, expected)
	}
	expected = "https://smooth-storage.aptoma.no/users/drf-eb/files/sphynx/2022/2/20/fab7859c399e7721-preview.html"
	if m.PreviewUrl != expected {
		t.Errorf("PreviewUrl was %s got %s", m.PreviewUrl, expected)
	}

}

func TestBurnMetadata(t *testing.T) {
	f, err := os.Open("../testdata/burn-metadata.json")
	if err != nil {
		t.Error("failed to read testdata")
	}
	defer f.Close()

	m, err := ReadBurnMetadata(f)
	if err != nil {
		t.Error("failed to read metadata")
	}
	expected := "eb"
	if m.ClientID != expected {
		t.Errorf("ClientID was %s got %s", m.ClientID, expected)
	}
	expected = "5d5a8cf857cd2009c74b6378"
	if m.EditionID != expected {
		t.Errorf("EditionID was %s got %s", m.EditionID, expected)
	}
	expected = "0387774c47c04e80"
	if m.HTMLHash != expected {
		t.Errorf("HTMLHash was %s got %s", m.HTMLHash, expected)
	}
	expected = "2022-02-19T18:51:53.915Z"
	if m.HTMLUpdatedAt != expected {
		t.Errorf("HTMLUpdatedAt was %s got %s", m.HTMLUpdatedAt, expected)
	}
	expected = "https://smooth-storage.aptoma.no/users/drf-eb/files/sphynx/2022/2/19/0387774c47c04e80.html"
	if m.HTMLUrl != expected {
		t.Errorf("HTMLUrl was %s got %s", m.HTMLUrl, expected)
	}
	expected = "236a375ec45e1904"
	if m.CSSHash != expected {
		t.Errorf("CSSHash was %s got %s", m.CSSHash, expected)
	}
	expected = "2022-02-14T09:03:41.489Z"
	if m.CSSUpdatedAt != expected {
		t.Errorf("CSSUpdatedAt was %s got %s", m.CSSUpdatedAt, expected)
	}
	expected = "https://smooth-storage.aptoma.no/users/drf-eb/files/sphynx/2022/2/14/236a375ec45e1904.css"
	if m.CSSUrl != expected {
		t.Errorf("CSSUrl was %s got %s", m.CSSUrl, expected)
	}
	expected = "https://smooth-storage.aptoma.no/users/drf-eb/files/sphynx/2022/2/19/0387774c47c04e80-preview.html"
	if m.PreviewUrl != expected {
		t.Errorf("PreviewUrl was %s got %s", m.PreviewUrl, expected)
	}
}
