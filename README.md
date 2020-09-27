# Readme

整合了常用发送http请求功能，并且简化了使用

## 简单使用

```
 func main(){
 	client := Myhttp.NewMyhttp()
 	client.Url = "http://www.baidu.com"
 	res := client.Get()
 	fmt.Println(res)
 }
```

## 自定义请求

client有以下属性

```
Url	string				//请求目标
IsProxy  bool			//是否使用代理   可以在Myhttp.go中更改全局变量proxyConf为代理地址
Post_value map[string]string	//Post的参数，以表单变量的形式发送
Header map[string]string		//自定义header
Timeout int					//设置请求超时时间，单位为秒，如果响应超时，则会产生报错信息，并且返回的body为一个空串，但是不会终止程序，一般用来时间盲注挺好用
```

设置好以上属性就可以直接用Post或者Get方法进行请求发送了

## Get

```
 func main(){
 
 	client := Myhttp.NewMyhttp()
 	
 	client.Url = "http://www.baidu.com/?name=admin&passwd=123"

 	client.Header = map[string]string{
 		"User-Agent":"mynewbrowser",
 		"client":"127.0.0.1"
 	}
 	
 	res := client.Get()
 	fmt.Println(res)
 }
```



## Post

```
 func main(){
 
 	client := Myhttp.NewMyhttp()
 	
 	client.Url = "http://www.baidu.com"
 	
 	client.Post_value = map[string]string{
 		"uname":"admin",
 		"passwd":"thisismypasswd"
 	}

 	client.Header = map[string]string{
 		"User-Agent":"mynewbrowser",
 		"client":"127.0.0.1"
 	}
 	
 	res := client.Post()
 	fmt.Println(res)
 }
```

注：post的参数的值将自动进行urlencode编码