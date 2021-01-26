package server

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const htmlTemplate = `
<html>
	<head>
		<title>File Server</title>
	</head>
	<body>
		<h1>{{.Title}}</h1>
		<ul>
			{{if .Parent }}
				<li><a href="{{.Parent}}">..</li>
			{{end}}
			{{range .Files}}
				<li><a href="{{.Path}}">{{.Info.Name}}</li>
			{{end}}
		</ul>
	</body>
</html>
`

type fileInfo struct {
	Path string
	Info os.FileInfo
}

type templateData struct {
	Files  []fileInfo
	Title  string
	Parent string
}

func Serve(path string, port int) {
	tmpl, err := template.New("page").Parse(htmlTemplate)
	if err != nil {
		fmt.Println("Failed to parse template")
		log.Fatal(err.Error())
	}

	rootPath := path
	log.Printf("root path: %s\n", rootPath)
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		tmplData := templateData{}
		log.Printf("request %s %s\n", r.Method, r.URL.Path)
		pathToList := parseFolderToList(rootPath, r.URL.Path)
		if info, err := os.Stat(pathToList); err != nil {
			fmt.Println("Failed to stat root folder")
			log.Println(err.Error())
			rw.WriteHeader(404)
		} else {
			if info.IsDir() {
				tmplData.Title = r.URL.Path
				err := filepath.Walk(pathToList, func(path string, info os.FileInfo, err error) error {
					if path == pathToList || !isSameFolder(pathToList, path) {
						return nil
					}
					if err == nil {
						tmplData.Files = append(tmplData.Files, fileInfo{
							Path: path,
							Info: info,
						})
					}
					return nil
				})
				if err != nil {
					panic(err)
				}
			} else {
				log.Printf("returning file: %s", pathToList)
				f, err := os.Open(pathToList)
				if err != nil {
					rw.WriteHeader(500)
					_, _ = rw.Write([]byte(err.Error()))
					return
				}
				defer f.Close()
				content, err := ioutil.ReadAll(f)
				if err != nil {
					rw.WriteHeader(500)
					_, _ = rw.Write([]byte(err.Error()))
					return
				}
				_, _ = rw.Write(content)
				return
			}
		}

		if pathToList != path {
			tmplData.Parent = getParent(pathToList)
		}
		rw.Header().Add("Content-Type", "text/html")
		_ = tmpl.Execute(rw, tmplData)
	})

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), mux))
}

func getParent(pathToList string) string {
	return filepath.Dir(pathToList)
}

func isSameFolder(rootPath string, filePath string) bool {
	return filepath.Dir(filePath) == normalizeRootPath(rootPath)
}

func normalizeRootPath(rootPath string) string {
	return strings.TrimPrefix(rootPath, "./")
}

func parseFolderToList(rootPath string, requestPath string) string {
	return fmt.Sprintf("%s%s", rootPath, strings.TrimSuffix(requestPath, "/"))
}
