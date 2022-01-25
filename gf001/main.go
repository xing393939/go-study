package main

import (
	"gf001/internal/cmd"
	"github.com/gogf/gf/v2/os/gctx"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	cmd.Main.Run(gctx.New())
}
