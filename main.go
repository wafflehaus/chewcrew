package main

func main() {
  // initialize router
  router := NewMethodRouter()

  router.HandleFunc("GET", "/", HandleHelloWorld)

  // TODO: move to http.Server instantiation if we need TLS
	router.ListenAndServe(":8080")

  return
}