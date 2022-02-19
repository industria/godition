package splitter

import (
	"os"
	"testing"

	"github.com/industria/godition/dredition"
)

func TestSplitter(t *testing.T) {
	// https://sphynx.aptoma.no/burned/{editionId}
	f, err := os.Open("../testdata/20220219-front.html")
	if err != nil {
		t.Errorf("unable to open file %v", err)
	}
	defer f.Close()
	notification, err := notification()
	if err != nil {
		t.Errorf("failed loading notification %v", err)
	}
	decks, err := Split(f, *notification)
	if err != nil {
		t.Errorf("failed split %v", err)
	}
	if len(*decks) != 10 {
		t.Errorf("number of decks not 9 was %d", len(*decks))
	}
}

func notification() (*dredition.Notification, error) {
	f, err := os.Open("../testdata/notification.json")
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return dredition.ReadNotification(f)
}
