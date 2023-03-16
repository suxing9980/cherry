package cherryActor

import (
	creflect "github.com/cherry-game/cherry/extend/reflect"
	cfacade "github.com/cherry-game/cherry/facade"
	"time"
)

type (
	IActorLoader interface {
		load(actor Actor)
	}
)

type (
	IEvent interface {
		Register(name string, fn IEventFunc) // 注册事件
		Unregister(name string)              // 注销事件
	}

	IEventFunc func(cfacade.IEventData) // 接收事件数据时的处理函数
)

type (
	IMailBox interface {
		Register(funcName string, fn interface{}) // 注册执行函数
		GetFuncInfo(funcName string) (*creflect.FuncInfo, bool)
	}
)

type (
	ITimer interface {
		Add(d time.Duration, fn func()) uint64                   // 添加定时器,循环执行
		AddOnce(d time.Duration, fn func())                      // 添加定时器,执行一次
		AddFixedHour(hour, minute, second int, fn func()) uint64 // 固定x小时x分x秒,循环执行
		AddFixedMinute(minute, second int, fn func()) uint64     // 固定x分x秒,循环执行
		AddSchedule(s ITimerSchedule, f func()) uint64           // 添加自定义调度
		Remove(id uint64)                                        // 移除定时器
	}

	ITimerSchedule interface {
		Next(time.Time) time.Time
	}
)
