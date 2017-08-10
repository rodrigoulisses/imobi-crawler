package evaldomatos

import(
  "log"
  "sync"
  "github.com/PuerkitoBio/goquery"
  "imobi-crawler/models"
  "strconv"
  "strings"
  "database/sql"
)

func Crawler(wg *sync.WaitGroup, db *sql.DB) {
  doc, err := goquery.NewDocument("http://evaldomatos.com.br/imoveis/filtro/?situacao=venda&tipo=x&quartos=x&garagem=x&bairro=x&valor=x&x=44&y=33")

  if err != nil {
    log.Fatal(err)

    return
  }

  doc.Find(".imoveis .row-fluid .span4").Each(func(i int, s *goquery.Selection) {
    // For each item found, get the band and title
    property := models.Property {}
    property.Name = s.Find("h3").Text()
    // Adicionar o dominio
    property.Link, _ = s.Find("h3 a").Attr("href")
    price := strings.Replace(strings.Replace(s.Find("p.valor span.el_2").Text(), ".", "", -1), ",", ".", -1)
    property.Price, _ = strconv.ParseFloat(price, 64)
    property.Image, _ = s.Find("img.wp-post-image").Attr("src")
    s.Find("table.info_imovel tbody td strong").Each(func(i int, selector *goquery.Selection){
      switch i{
        case 0:
          area := strings.Replace(selector.Text(), "M²", "", -1)
          property.Area, _ = strconv.ParseFloat(area, 64)
        case 1:
          bathroom := selector.Text()
          property.Bathroom, _ = strconv.ParseInt(bathroom, 10, 64)
        case 2:
          bedroom := selector.Text()
          property.Bedroom, _ = strconv.ParseInt(bedroom, 10, 64)
      }
    })
    property.Kind = 1
    property.RealStateId = 2

    models.InsertProperty(property, db)
  })

  wg.Done()
}
