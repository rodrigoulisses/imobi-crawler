package main

import(
  "imobi-crawler/martins"
  "imobi-crawler/evaldomatos"
)

func main() {
  martins.Crawler()
  evaldomatos.Crawler()
}
