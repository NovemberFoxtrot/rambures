package main

import (
	"log"
	"net/http"
  "os"
)

func CheckError(err error) {
  if err != nil {
    log.Fatalln("[ERROR]", err)
  }
}

func main() {
  workingdirectory, err := os.Getwd()

  log.Println("[INFO]", "Starting server in directoty", workingdirectory)

	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir(workingdirectory))))

	err = http.ListenAndServe(":9090", nil)

  CheckError(err)
}
