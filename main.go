package main

import (
	_ "capys/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"capys/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
