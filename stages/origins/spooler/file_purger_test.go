
package spooler

import (
	"datacollector-edge/stages/lib/dataparser"
	"io/ioutil"
	"path/filepath"
	"testing"
	"time"
)

func TestFilePurger_purge(t *testing.T) {
	testDir := createTestDirectory(t)

	defer deleteTestDirectory(t, testDir)

	//Create a.txt,c.txt,b.txt with different mod times
	createFileAndWriteContents(t, filepath.Join(testDir, "a.txt"), "3d967ed5-556d-4b4c-a0d0-33e5fbe344e4\n456", dataparser.CompressedNone)
	createFileAndWriteContents(t, filepath.Join(testDir, "b.txt"), "111213\n141516", dataparser.CompressedNone)
	createFileAndWriteContents(t, filepath.Join(testDir, "c.txt"), "111112113\n114115116\n117118119", dataparser.CompressedNone)

	files, _ := ioutil.ReadDir(testDir)
	if len(files) != 3 {
		t.Error("Failed to create test files")
	}

	filePurger := filePurger{archiveDir: testDir, retentionTime: 2 * time.Second}
	time.Sleep(4 * time.Second)
	filePurger.purge()

	archivedFiles, _ := ioutil.ReadDir(testDir)
	if len(archivedFiles) != 0 {
		t.Error("Failed to purge files")
	}
}
