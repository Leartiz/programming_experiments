package main

import (
	"github.com/jbrodriguez/mlog"
)

func main() {
	mlog.StartEx(mlog.LevelInfo, "app.log", 5*1024*1024, 5)

	ipsum := "ipsum"
	for i := 0; i < 100000; i++ {
		mlog.Warning("Lorem %s", ipsum)
	}
}
