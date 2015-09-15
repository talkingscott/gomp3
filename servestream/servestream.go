package main

import (
  "fmt"
  "log"
  "net/http"
  "os"
  "path"
  "strings"
)

type fileHandler struct {
  root http.FileSystem
}

func fileServer(root http.FileSystem) http.Handler {
  return &fileHandler{root}
}

func (f *fileHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  upath := r.URL.Path
  if !strings.HasPrefix(upath, "/") {
    upath = "/" + upath
    r.URL.Path = upath
  }
  fd, err := f.root.Open(path.Clean(upath))
  if err != nil {
    
  }
  defer fd.Close()
  
  sb, err1 := fd.Stat()
  if err1 != nil {
    
  }
  
  http.ServeContent(w, r, sb.Name(), sb.ModTime(), fd)
}

func main() {
  var document_root string

  if (len(os.Args) > 1) {
    document_root = os.Args[1];
  } else {
    // serve files streamed from my iTunes
    document_root = "c:\\users\\snichol\\music\\itunes\\itunes media\\music"
  }

  port := 8080
  fmt.Printf("Listening on port %v to serve streams from %v\n", port, document_root)

  http.Handle("/", fileServer(http.Dir(document_root)))
  log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}