package opener

import (
	"testing"
	"time"
)

func TestOpenUrl(t *testing.T) {
	err := Open("https://github.com")
	if err != nil {
		t.Fatal(err)
	}

	t.Log("success")
	time.Sleep(time.Second * 5)
}

func TestOpenDir(t *testing.T) {
	err := Open("c://")
	if err != nil {
		t.Fatal(err)
	}

	t.Log("success")
	time.Sleep(time.Second * 5)
}
