package main

import (
	"context"
	"fmt"
	"os"

	"github.com/artuross/kubesource/internal/commands"
)

func main() {
	rootCmd := commands.NewKubesourceCommand()

	if err := rootCmd.Run(context.Background(), os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
