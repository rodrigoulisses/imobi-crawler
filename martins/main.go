package martins

import(
  "log"
  "sync"
  "strings"
  "github.com/PuerkitoBio/goquery"
  "imobi-crawler/models"
  "strconv"
  "database/sql"
)

func Crawler(wg *sync.WaitGroup, db *sql.DB) {
  doc, err := goquery.NewDocument("http://www.martinsimoveispi.com.br/pesquisar?operacao=VENDA&cidade=1&bairro=0&tipo=0&quartos=0&area=0&valorInicial=&valorFinal=")

  if err != nil{
    log.Fatal(err)

    return
  }

  doc.Find(".properties-full .property-thumb-info").Each(func(i int, s *goquery.Selection) {
    // For each item found, get the band and title
    property := models.Property {}
    property.Name = s.Find(".property-thumb-info-content h3").Text()
    // Adicionar o dominio
    property.Link, _ = s.Find(".property-thumb-info-content h3 a").Attr("href")
    property.Link = "http://www.martinsimoveispi.com.br" + property.Link
    price := strings.Replace(strings.Replace(s.Find(".property-thumb-info-label .price").Text(), ".", "", -1), ",", ".", -1)
    property.Price, _ = strconv.ParseFloat(strings.Split(price, " ")[1], 64)
    property.Code = strings.Split(s.Find(".property-thumb-info-content address").Text(), " - ")[0]
    property.Kind = models.FindOrInsertKindProperty(strings.Split(s.Find(".property-thumb-info-content address").Text(), " - ")[1], db)
    property.Image, _ = s.Find(".property-thumb-info-image img").Attr("src")
    property.Area, _ = strconv.ParseFloat(strings.Split(s.Find(".property-thumb-info .amenities .pull-left li").Text(), " ")[1], 64)
    s.Find(".property-thumb-info .amenities .pull-right li").Each(func(i int, selector *goquery.Selection){
      if i == 0 {
        bedroom := strings.TrimPrefix(selector.Text(), " ")
        property.Bedroom, _ = strconv.ParseInt(bedroom, 10, 64)
      } else {
        bathroom := strings.TrimPrefix(selector.Text(), " ")
        property.Bathroom, _ = strconv.ParseInt(bathroom, 10, 64)
      }
    })
    property.RealStateId = 1

    models.InsertProperty(property, db)
  })

  wg.Done()
}
