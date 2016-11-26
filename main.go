package main

import(
  "sync"
  "imobi-crawler/martins"
  "imobi-crawler/evaldomatos"
)

func main() {
  var wg sync.WaitGroup
  wg.Add(2)
  go martins.Crawler(&wg)
  go evaldomatos.Crawler(&wg)
  wg.Wait()
}
