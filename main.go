package main

import(
  "io/ioutil"
  "log"
  "net/http"
  loggly "github.com/jamespearly/loggly"
  "encoding/json"
  "time"
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/dynamodb"
  "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
  "fmt"
  "os"

)
type Response struct {
  GlobalQuote   GlobalQuote   `json:"Global Quote"`
}

type GlobalQuote struct{
  Id string `json:"id"`
  Symbol  string   `json:"01. symbol"`
  Price   string  `json:"05. price"`
}

func main(){
  for{
    writeToTable(makeCall())
    time.Sleep(2*time.Minute)
  }
}

func writeToTable(gc GlobalQuote){
  sess := session.Must(session.NewSessionWithOptions(session.Options{
      SharedConfigState: session.SharedConfigEnable,
      }))

  // Create DynamoDB client
  svc := dynamodb.New(sess)
  //continue here
  t := time.Now()
  ts := t.String();
  item := GlobalQuote{
    Id: ts,
    Symbol: gc.Symbol,
    Price: gc.Price,
  }

  av, err := dynamodbattribute.MarshalMap(item)
  if err != nil{
    fmt.Println("failed to marshall")
    fmt.Println(err.Error())
    os.Exit(1)
  }
  input := &dynamodb.PutItemInput{
    Item:  av,
    TableName:  aws.String("AppleStock"),
  }
  _, err = svc.PutItem(input)
  if err != nil{
    fmt.Println("Got error calling PutItem: ")
    fmt.Println(err.Error())
    os.Exit(1)
  }
  fmt.Println("successfully added to table")
}

func makeCall() GlobalQuote{
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
  return result.GlobalQuote
}
