Logging service for golang projects
------------------------------

usage with grpc
---------------

1. go get github.com/MiSingularity/logging
   
2. Init grpc client by calling
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

Golang example is placed in /example/go/biclient.go,
Other languages are similar.