# iris-middleware-casbin
这是一个go iris casbin的测试，casbin刚开始接触的时候，非常让人糊涂，经过测试调试，基本搞明白了。

csv权限配置不从本地文件读取，更改为REDIS方式存储和读取。

调试方法如下

1、god mod init xxx

2、go env -w GO111MODULE=on

3、go env -w GOPROXY=https://goproxy.cn,direct

4、go run main.go





注意细节：

casbin/casbin.go

func Username(r *http.Request) string {

    //username, _, _ := r.BasicAuth()
    return "admin"
}

上面的函数返回的用户名用于权限调试


main.go

//设置初始权限 默认ADMIN用户拥有所有权限  anonymous只能访问/根目录   member可以访问/logout


	Enforcer.LoadPolicy()  
	Enforcer.AddPolicy("admin", "/*", "*")
	Enforcer.AddPolicy("anonymous", "/", "GET")
	Enforcer.AddPolicy("member", "/logout", "*")
	Enforcer.AddPolicy("member", "/member/*", "*")
  
  //保存配置到REDIS
  if err := Enforcer.SavePolicy(); err != nil {
		panic(err)
	}





