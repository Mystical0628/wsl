package main

import (
	"github.com/Mystical0628/wsl/v1"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(wsl.NewAnalyzer(nil))
}
