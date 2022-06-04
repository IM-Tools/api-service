/**
  @author:panliang
  @data:2022/6/5
  @note
**/
package coroutine

import (
	"github.com/panjf2000/ants/v2"
	"im-services/config"
)

var AntsPool *ants.Pool

func ConnectPool() *ants.Pool {
	//设置数量
	AntsPool, _ = ants.NewPool(config.Conf.Server.CoroutinesPoll)
	return AntsPool
}
