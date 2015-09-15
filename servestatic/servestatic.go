package main

import (
  "fmt"
  "log"
  "net/http"
  "os"
)

// Wraps another handler, setting some headers to prevent caching
// before delegating to the wrapped handler
type noCacheHandler struct {
  wrapped http.Handler
}

func (h *noCacheHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("pragma", "no-cache")
  w.Header().Set("cache-control", "no-cache")
  h.wrapped.ServeHTTP(w, r)
  // I believe the response has been sent (not just "queued up") by this time,
  // as setting headers here had no effect
}

func main() {
  var document_root string

  if (len(os.Args) > 1) {
    document_root = os.Args[1];
  } else {
    // statically serve files from my go workspace
    document_root = "c:\\users\\snichol\\documents\\code\\go"
  }

  port := 8080
  fmt.Printf("Listening on port %v to serve documents from %v\n", port, document_root)

  fs := http.FileServer(http.Dir(document_root))
  http.Handle("/", &noCacheHandler{fs})
  log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}