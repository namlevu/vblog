package cnf

import (
  "os"
  "encoding/json"
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
  switch mode {
  case DEV:
    filename = "C:\\go-work\\src\\vblog\\cnf\\configuration.default.json"
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
