package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func Serve(path string, port int) {

	rootPath := normalizeRootPath(path)
	fmt.Printf("root path: %s\n", rootPath)
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Printf("request path: %s\n", r.URL.Path)
		result := ""
		pathToList := parseFolderToList(rootPath, r.URL.Path)
		if info, err := os.Stat(pathToList); err != nil {

		} else {
			if info.IsDir() {
				err := filepath.Walk(pathToList, func(path string, info os.FileInfo, err error) error {
					if err == nil {
						result += fmt.Sprintf("<li>%s: %v</li>", path, info.IsDir())
					}
					return nil
				})
				if err != nil {
					panic(err)
				}
			} else {
				result += fmt.Sprintf("path: %s (is file)", pathToList)
			}
		}
		fmt.Printf("pathToList: %s\n", pathToList)

		rw.Write([]byte(fmt.Sprintf("<html><body><ul>%s</ul></body></html>\n", result)))
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), mux))
}

func normalizeRootPath(path string) string {
	if strings.HasSuffix(path, "/") {
		fmt.Println("Removing path /")
		return strings.TrimSuffix(path, "/")
	}
	return path
}

func parseFolderToList(rootPath string, requestPath string) string {
	return fmt.Sprintf("%s%s", rootPath, requestPath)
}
