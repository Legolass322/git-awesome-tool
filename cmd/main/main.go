package main

import (
	"fmt"

	"github.com/Legolass322/go-awesome-tool/internal/cfg"
)

func main() {
	var config cfg.Config = *cfg.NewConfig(cfg.Local)
	fmt.Println("Hello")
	fmt.Printf("%v", config)
}
