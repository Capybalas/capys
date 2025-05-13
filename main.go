package main

import (
	_ "capys/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"capys/internal/cmd"
	_ "capys/internal/logic"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
