// config/config.go
package config

import "os"

// Config holds the configuration values.
type Config struct {
   DBURI      string
   Greeting   string
   Addr       string
   DBName     string
   Collection string
}

// LoadConfig loads configuration values from environment variables.
func LoadConfig() *Config {
   return &Config{
      DBURI:      os.Getenv("DB_URI"),
      Greeting:   os.Getenv("GREETING"),
      Addr:       os.Getenv("ADDR"),
      DBName:     os.Getenv("DB_NAME"),
      Collection: os.Getenv("COLLECTION"),
   }
}

