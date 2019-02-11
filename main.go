package main

import(
  "io/ioutil"
  "log"
  "net/http"
  "loggly"
  "encoding/json"
  

)
type Response struct {
  GlobalQuote   GlobalQuote   `json:"Global Quote"`
}

type GlobalQuote struct{
  Symbol  string   `json:"01. symbol"`
  Price   string  `json:"05. price"`
}

func main(){

  client := loggly.New("csc482")

  api := "https://www.alphavantage.co/query?function=GLOBAL_QUOTE&symbol=AAPL&apikey=RWP2ZLYN1BPL1RAU"

  resp, err := http.Get(api)
  if err != nil{
    client.EchoSend("error", "Get request failed")
    log.Fatal(err)
  }

  defer resp.Body.Close()

  body, err := ioutil.ReadAll(resp.Body)

  if err != nil{
    log.Fatal(err)
  }

  var result Response
  err = json.Unmarshal(body, &result)

  if err != nil{
    client.EchoSend("error", "Unmarshall failed")
    log.Fatal(err)
  }

  symbol := result.GlobalQuote.Symbol
  price := result.GlobalQuote.Price

  logglyMsg := "Symbol: " + symbol + " Price: " + price
  client.EchoSend("info", logglyMsg)

}
