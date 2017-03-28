package main

import(
  "fmt"
  "net/http"
  "sync"
  "encoding/json"
  "imobi-crawler/martins"
  "imobi-crawler/evaldomatos"
  "imobi-crawler/models"
)

func main() {
  http.HandleFunc("/", handle)
  http.ListenAndServe(":4000", nil)
}

func handle(w http.ResponseWriter, r *http.Request){
  var properties []models.Property

  fmt.Println(r)
  var wg sync.WaitGroup
  wg.Add(2)

  go martins.Crawler(&wg, &properties)
  go evaldomatos.Crawler(&wg, &properties)

  wg.Wait()

  json1, _ := json.Marshal(properties)
  fmt.Fprintf(w, string(json1))
}
