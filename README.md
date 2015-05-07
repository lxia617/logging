Logging service for golang projects
------------------------------

back-end usage
--------------

1. get misbi code and put it to your GOPATH:  GOPATH/src/misbi
   
2. Init grpc connection by calling
    ```go
    bi.Init("127.0.0.1", "8999")
    ```
    
3. Call Bi() to add a bi log, for example:
   
   ```go
   bi.Bi(&p.BiLog{
   		ProjectName: "deepshare",
   		ActionName:  "userlink",
   		Timestamp:   time.Now().Unix(),
   		Detail:      []byte("detail~~~~~~"),
   	})
   	```
front-end usage
---------------
1. get bi server address from devops, ie. http://127.0.0.1:8088
2. post to http://127.0.0.1:8088?action=projectname/actionname
    put bi log in http body

you can refer to bi_test.go
