package cmd

import (
    "os"
	"os/exec"
	"testing"
    "log"
)

var mainGo = "../main.go"
var feedDBName = "test-feed"

func TestMain(m *testing.M) {
    exec.Command("go", "run", mainGo, "create", feedDBName, "feed").Run()
    exec.Command("go", "run", mainGo, "add", feedDBName, "{\"feed-item-1\":\"\"}").Run()

	os.Exit(m.Run())
}

func TestGetWithFeed(t *testing.T) {
    out, err := exec.Command("go", "run", mainGo, "get", feedDBName, "feed-item-1").Output()
    if err != nil {
        log.Fatal(err)
    }

    expected := "{\"feed-item-1\":\"\"}\n"
    actual := string(out)

    if actual != expected {
        t.Error(feedDBName, "database should have one item called feed-item-1")
    }
}
