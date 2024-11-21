package main

import (
	"fmt"

	"github.com/smirnova-daria/url-shortener/internal/config"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg)
}
