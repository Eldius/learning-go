package server

import "testing"

func TestParseFolderToList(t *testing.T) {
	parsed := parseFolderToList(".", "/server")

	if parsed != "./server" {
		t.Errorf("parsed value must be './server', but was '%s'", parsed)
	}
}

func TestISsSameFolder(t *testing.T) {
	testSameFolder("./server", "server/templates", true, t)
	testSameFolder("./server", "server/templates/index.html", false, t)
	testSameFolder("/home/user/folder", "/home/user/folder/file.txt", true, t)
	testSameFolder("/home/user/folder", "/home/user/folder/subfolder/file.txt", false, t)
	testSameFolder("/home/user/folder", "/home/user/folder/subfolder/file.txt", false, t)
	testSameFolder("~/folder", "/home/user/folder/subfolder/file.txt", false, t)
	testSameFolder("~/folder", "/home/user/folder/subfolder/file.txt", false, t)
}

func TestNormalizeRootPath(t *testing.T) {
	result1 := normalizeRootPath("./server")
	if result1 != "server" {
		t.Errorf("Result1 must be 'server', but was '%s'", result1)
	}

	result2 := normalizeRootPath("/home/user/folder")
	if result2 != "/home/user/folder" {
		t.Errorf("Result2 must be '/home/user/folder', but was '%s'", result2)
	}
}

func testSameFolder(folder string, path string, result bool, t *testing.T) {
	response := isSameFolder(folder, path)
	if response != result {
		t.Errorf("'%s' is same folder of '%s'? '%v'", path, folder, response)
	}

}
