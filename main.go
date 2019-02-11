package main

import(
  "os"
  "io/ioutil"
  "log"
  "net/http"
  "loggly"
  "encoding/json"
  "strconv"
)
type Response struct {
  GLOBAL_QUOTE GQuote 'json:"Global Quote"'
}

type GQuote struct{
  Symbol string 'json:"01. symbol"'
  Price float64 'json:"05. price"'
}

func main(){

  client := loggl.New("csc482")

  api := "https://www.alphavantage.co/query?function=GLOBAL_QUOTE&symbol=AAPL&apikey=RWP2ZLYN1BPL1RAU"

  resp, err := http.Get(api)
  if err != nil{
    logglyResp := client.EchoSend("error", "Get request failed")
    log.Fatal(err)
  }

  defer resp.Body.Close()

  body, err := ioutil.ReadAll(resp.Body)

  if err != nil{
    log.Fatal(err)
  }

  var result Response
  err := json.Unmarshall(body, &result)

  if err != nil{
    logglyResp := client.EchoSend("error", "Unmarshall failed")
    log.Fatal(err)
  }

  logglyMsg := "Symbol: " + Symbol + " Price: " + strconv.Itoa(Price)

  _,err = os.Stdout.Write(body);

  if err != nil{
    log.Fatal(err)
  }
}
