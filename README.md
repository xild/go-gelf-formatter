Logrus Gelf Formatter
=========

Formatter - Format your LogRus log in GELF FORMAT.

## Formatter 

All logrus.Fields are printed like additional fields in gelf format.

If you set a envrionment variable with APPLICATION_NAME the formatter will print 
a additional field called _appName

If there are newlines in the message, use the first line for the short message and set the full message to the original input

The code bellow will print: 
```json 
{
    "version": "1.0",
    "host": "MACHINE_NAME os.GET",
    "short_message": "A group of walrus emerges from the ocean",
    "full_message": "A group of walrus emerges from the ocean The ocean That's all I need",
    "timestamp": "2016-09-19T10:57:42-03:00",
    "level": 6,
    "file": "",
    "line": 0,
    "_animal": "walrus",
    "_size": 10
}
```

```go
package main

import (
  log "github.com/Sirupsen/logrus"
  "github.com/xild/go-gelf-formatter"
)

func init() {
  // Log as GELF FORMATTER instead of the default ASCII formatter.
  log.SetFormatter(&xild.GELFFormatter{})
}

func main() {
  log.WithFields(log.Fields{
    "animal": "walrus",
    "size":   10,
  }).Info("A group of walrus emerges from the ocean \n The ocean That's all I need")
}
```

## Contributing:

If you find any improvement or issue you want to fix, feel free to send me a pull request with testing.


## Credits

* [Luis Vieira](https://www.facebook.com/luisxild) 


