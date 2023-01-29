package cmd

import (
	"os"
	"testing"

	"github.com/liweiyi88/dummyfile/dummyfile"
)

func TestRootCmd(t *testing.T) {
	cmd := rootCmd

	workDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("could not get current work dir %v", err)
	}

	path := workDir + "/dummy_test.txt"

	cmd.SetArgs([]string{path, "-s", "100"})
	cmd.Execute()

	defer func() {
		err := os.Remove(path)
		if err != nil {
			t.Logf("could not remove test dummy file %s", path)
		}
	}()

	fileInfo, err := os.Stat(path)
	if err != nil {
		t.Errorf("could not get file info: %v", err)
	}

	if fileInfo.Size() != 100 {
		t.Errorf("expected file size 100b, but got: %d", fileInfo.Size())
	}

	var expected int64

	expected = 200
	cmd.SetArgs([]string{path, "-s", "200"})
	cmd.Execute()

	fileInfo, err = os.Stat(path)
	if err != nil {
		t.Errorf("could not get file info: %v", err)
	}

	if fileInfo.Size() != expected {
		t.Errorf("expected file size %db, but got: %d", expected, fileInfo.Size())
	}

	expected = 10 * dummyfile.MB
	cmd.SetArgs([]string{path, "-s", "10 mb"})
	cmd.Execute()

	fileInfo, err = os.Stat(path)
	if err != nil {
		t.Errorf("could not get file info: %v", err)
	}

	//fmt.Printf("%+v", fileInfo)
	if fileInfo.Size() != expected {
		t.Errorf("expected file size %db, but got: %d", expected, fileInfo.Size())
	}
}
