package models

import(
  "database/sql"
  "time"
  "fmt"
)

type Property struct {
  Code string
  Name string
  Kind int64
  Price float64
  Link string
  Image string
  Area float64
  Bedroom int64
  Bathroom int64
  RealStateId int64
}

func InsertProperty(property Property, db *sql.DB) {
  conn, err := db.Query("insert into properties (code, name, kind_id, price, link, image, area, bedroom, bathroom, real_state_id, created_at, updated_at) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12) returning id", property.Code, property.Name, property.Kind, property.Price, property.Link, property.Image, property.Area, property.Bedroom, property.Bathroom, property.RealStateId, time.Now(), time.Now())
  defer conn.Close()
  if err != nil {
    fmt.Println(err)
  }
}

func FindOrInsertKindProperty(name string, db *sql.DB) int64 {
  var id int64

  db.QueryRow("select id from kind_properties where name = $1", name).Scan(&id)

  if id == 0 {
    err := db.QueryRow("insert into kind_properties (name, created_at, updated_at) values ($1, $2, $3) returning id", name, time.Now(), time.Now()).Scan(&id)
    if err != nil {
      fmt.Println(err)
    }
  }

  return id
}
