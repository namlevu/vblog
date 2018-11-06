package cnf

import (
  "log"
  "os"
  "encoding/json"
  "path/filepath"
)

type Configuration struct {
  MongodbHost string `json:"mongodb_host"`
	MongodbDatabase  string `json:"mongodb_database"`
	MongodbConnectionPool int `json:"mongodb_connection_pool"`
	ApiPort int `json:"api_port"`
  TimeoutRead int `json:"timeout_read"`
  TimeoutWrite int `json:"timeout_write"`
}

type Mode string
const (
  DEV Mode = "development"
  PROD Mode = "production"
)

func LoadConfig(configuration *Configuration, mode Mode) error {
  var filename = ""
  dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
  if err != nil {
    log.Fatal(err)
  }
  
  switch mode {
  case DEV:
    filename = dir+"/"+"configuration.dev.json"
    break
  case PROD:
    filename = "configuration.prod.json"
    break
  default:
     filename = "configuration.default.json" // default setting
     break
  }
  //filename is the path to the json config file
  file, err := os.Open(filename)
  if err != nil {
    return err
  }
  decoder := json.NewDecoder(file)
  err = decoder.Decode(&configuration)
  if err != nil {
    return err
  }
  return nil
}
