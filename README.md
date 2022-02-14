# Go Config

Parse and load configurations in Go.


## Usage

```go
package main

import (
	"fmt"

	"github.com/waizungtaam/config"
)

type Config struct {
	Host string `config:"required"`
	Port int    `config:"required;default=80"`
	Mode string
	SSL  bool `config:"default=true"`
}

func main() {
	c := Config{}
	if err := config.Load("dev.yml", &c); err != nil {
		panic(err)
	}
	fmt.Printf("%+v\v", c)
}
```

## Author
waizungtaam

## License
MIT
