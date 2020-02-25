package main

import (
	"github.com/kataras/iris/v12"

	"github.com/casbin/casbin/v2"
	cm "casebin-middleware/casbin"
	redisadapter "github.com/casbin/redis-adapter/v2"

)





func newApp() *iris.Application {


	//配置casbin权限数据同步到REDIS
	adp := redisadapter.NewAdapter("tcp", "127.0.0.1:6379")
	var Enforcer, _ = casbin.NewEnforcer("casbinmodel.conf", adp)

	//设置初始权限 默认ADMIN用户拥有所有权限
	Enforcer.LoadPolicy()
	Enforcer.AddPolicy("admin", "/*", "*")
	Enforcer.AddPolicy("anonymous", "/", "GET")
	Enforcer.AddPolicy("member", "/logout", "*")
	Enforcer.AddPolicy("member", "/member/*", "*")

	if err := Enforcer.SavePolicy(); err != nil {
		panic(err)
	}

	//设置权限中间件
	casbinMiddleware := cm.New(Enforcer)



	app := iris.New()
	app.Use(casbinMiddleware.ServeHTTP)//调用权限中间件

	app.Get("/", hi)

	app.Get("/dataset1/{p:path}", hi) // p, alice, /dataset1/*, GET

	app.Post("/dataset1/resource1", hi)

	app.Get("/dataset2/resource2", hi)
	app.Post("/dataset2/folder1/{p:path}", hi)

	app.Any("/dataset2/resource1", hi)

	return app
}

func main() {
	app := newApp()
	app.Run(iris.Addr(":8081"))
}

func hi(ctx iris.Context) {
	ctx.Writef("Hello %s", cm.Username(ctx.Request()))
}
