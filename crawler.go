package main

import(
  "fmt"
  "sync"
  //"database/sql"
  "encoding/json"
  "imobi-crawler/martins"
  "imobi-crawler/evaldomatos"
  "imobi-crawler/models"
)

func main (){
  db, err := sql.Open("postgres", "postgres://postgres:postgres@postgres/imobi-crwaler?sslmode=verify-full")

  if err != nil{
   fmt.Println("Erro de conexão")
  }

  fmt.Println(db)

  var properties []models.Property

  var wg sync.WaitGroup
  wg.Add(2)

  go martins.Crawler(&wg, &properties)
  go evaldomatos.Crawler(&wg, &properties)

  wg.Wait()

  json1, _ := json.Marshal(properties)
  fmt.Println(string(json1))
}
