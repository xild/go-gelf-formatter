Logrus Gelf Formatter
=========

Formatter - sometimes we just want to a formatter for GELF formatter instead the hook to the server.

## Formatter 

All logrus.Fields are printed like additional fields in gelf format.

If you set a envrionment variable with APPLICATION_NAME the formatter will print 
a additional field called _appName

```go
package main

import (
  log "github.com/Sirupsen/logrus"
  "github.com/xild/go-gelf-formatter"
)

func init() {
  // Log as JSON instead of the default ASCII formatter.
  log.SetFormatter(&xild.GELFFormatter{})

  // Output to stderr instead of stdout, could also be a file.
  log.SetOutput(os.Stderr)

  // Only log the warning severity or above.
  log.SetLevel(log.WarnLevel)
}

func main() {
  log.WithFields(logrus.Fields{
    "animal": "walrus",
    "size":   10,
  }).Info("A group of walrus emerges from the ocean")
}
```


## Contributing to GoRequest:

If you find any improvement or issue you want to fix, feel free to send me a pull request with testing.


## Credits

* [Luis Vieira](https://www.facebook.com/luisxild) 


