package tui

import (
	"fmt"
	"os"

	"github.com/NimbleMarkets/go-booba"
)

func CreateChart() {
	m := NewModel()

	if err := booba.Run(m); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

