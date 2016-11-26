package martins

import(
  "fmt"
  "log"
  "strings"
  "encoding/json"
  "github.com/PuerkitoBio/goquery"
  "imobi-crawler/models"
)

func Crawler() {
  doc, err := goquery.NewDocument("http://www.martinsimoveispi.com.br/pesquisar?operacao=VENDA&cidade=1&bairro=0&tipo=0&quartos=0&area=0&valorInicial=&valorFinal=")

  if err != nil{
    log.Fatal(err)

    return
  }

  // Find the review items
  var teste bool
  var properties []models.Property

  doc.Find(".properties-full .property-thumb-info").Each(func(i int, s *goquery.Selection) {
    // For each item found, get the band and title
    property := models.Property {}
    property.Name = s.Find(".property-thumb-info-content h3").Text()
    // Adicionar o dominio
    property.Link, teste = s.Find(".property-thumb-info-content h3 a").Attr("href")
    property.Link = "http://www.martinsimoveispi.com.br" + property.Link
    property.Price = s.Find(".property-thumb-info-label .price").Text()
    property.Code = strings.Split(s.Find(".property-thumb-info-content address").Text(), " - ")[0]
    property.Kind = strings.Split(s.Find(".property-thumb-info-content address").Text(), " - ")[1]
    property.Image, teste = s.Find(".property-thumb-info-image img").Attr("src")

    properties = append(properties, property)
  })

  json1, _ := json.Marshal(properties)
  fmt.Println(string(json1))
}

