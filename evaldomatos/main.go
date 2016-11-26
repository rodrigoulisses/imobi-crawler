package evaldomatos

import(
  "fmt"
  "log"
  // "strings"
  "encoding/json"
  "github.com/PuerkitoBio/goquery"
  "imobi-crawler/models"
)

func Crawler() {
  doc, err := goquery.NewDocument("http://evaldomatos.com.br/imoveis/filtro/?situacao=venda&tipo=x&quartos=x&garagem=x&bairro=x&valor=x&x=44&y=33")

  if err != nil {
    log.Fatal(err)

    return
  }

  var properties []models.Property

  doc.Find(".imoveis .row-fluid").Each(func(i int, s *goquery.Selection) {
    // For each item found, get the band and title
    property := models.Property {}
    property.Name = s.Find("h3").Text()
    // Adicionar o dominio
    property.Link, _ = s.Find("h3 a").Attr("href")
    property.Price = s.Find("p.valor span.el_2").Text()
    property.Image, _ = s.Find("img.wp-post-image").Attr("src")

    properties = append(properties, property)
  })

  json1, _ := json.Marshal(properties)
  fmt.Println("%s", string(json1))
}
