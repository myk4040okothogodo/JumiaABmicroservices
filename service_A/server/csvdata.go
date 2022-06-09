package server

import (
    service_av1 "github.com/myk4040okothogodo/JumiaABmicroservices/gen/go/service_A"
    "bufio"
    "encoding/csv"
    "log"
    "io"
    "os"
    "strings"
)


func LoadCsvData() []*service_av1.Order_A {
  var allOrders []*service_av1.Order_A

  pwd, _ := os.Getwd()
  csvFile, _ := os.Open(pwd + "/server/test_file.csv")
  println(pwd+ "/server/test_file.csv")

  reader := csv.NewReader(bufio.NewReader(csvFile))
  for {
      line, error := reader.Read()
      if error == io.EOF {
          break
      } else if error != nil {
         log.Fatal(error)
      }
      countryval  := strings.TrimSpace(line[2][0:4])

      allOrders = append(allOrders, &service_av1.Order_A{
          Id    :      strings.TrimSpace(line[0]),
          Email :      strings.TrimSpace(line[1]),
          Weight:      strings.TrimSpace(line[3]),
          Phonenumber: strings.TrimSpace(line[2]),
          Countrycode: countryval,
      })
  }
  return allOrders
}
