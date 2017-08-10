package main

import(
  "sync"
  "imobi-crawler/martins"
  "imobi-crawler/evaldomatos"
  _ "github.com/lib/pq"
  "database/sql"
  "fmt"
)

func main (){
  db, err := sql.Open("postgres", "user=rodrigoulissesesilva dbname=imobi_dev sslmode=disable")

  if err != nil{
    fmt.Println(err)
    return
  }

  var wg sync.WaitGroup
  wg.Add(2)

  go martins.Crawler(&wg, db)
  go evaldomatos.Crawler(&wg, db)

  wg.Wait()
}
