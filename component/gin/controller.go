package cherryGin

import (
	"github.com/cherry-game/cherry/facade"
	"github.com/gin-gonic/gin"
)

type HandlerFunc func(ctx *Context)

type IController interface {
	PreInit(app cherryFacade.IApplication, engine *gin.Engine)

	Init()

	Stop()
}

func BindHandlers(handlers []HandlerFunc) []gin.HandlerFunc {
	var list []gin.HandlerFunc
	for _, handler := range handlers {
		list = append(list, BindHandler(handler))
	}
	return list
}

func BindHandler(handler func(ctx *Context)) gin.HandlerFunc {
	return func(c *gin.Context) {
		context := new(Context)
		context.Context = c
		handler(context)
	}
}

type BaseController struct {
	App    cherryFacade.IApplication
	Engine *gin.Engine
}

func (b *BaseController) PreInit(app cherryFacade.IApplication, engine *gin.Engine) {
	b.App = app
	b.Engine = engine
}

func (b *BaseController) Init() {

}

func (b *BaseController) Stop() {

}

func (b *BaseController) Group(relativePath string, handlers ...HandlerFunc) *Group {
	group := &Group{
		RouterGroup: b.Engine.Group(relativePath, BindHandlers(handlers)...),
	}
	return group
}

func (b *BaseController) Any(relativePath string, handlers ...HandlerFunc) {
	b.Engine.Any(relativePath, BindHandlers(handlers)...)
}

func (b *BaseController) GET(relativePath string, handlers ...HandlerFunc) {
	b.Engine.GET(relativePath, BindHandlers(handlers)...)
}

func (b *BaseController) POST(relativePath string, handlers ...HandlerFunc) {
	b.Engine.POST(relativePath, BindHandlers(handlers)...)
}
