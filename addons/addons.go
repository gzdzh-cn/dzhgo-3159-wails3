package addons

import (
	_ "dzhgo/addons/dict"
	_ "dzhgo/addons/member"
	_ "dzhgo/addons/space"
	_ "dzhgo/addons/task"

	"github.com/gzdzh-cn/dzhcore"
)

func NewInit() {

	// 初始化所有addons
	dzhcore.InitAddons()

}
