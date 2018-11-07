
# Developer Guide
Just run godoc to see documentation

``` sh
godoc -http=:6060
```


To see main packages in documentation do not forget to update:
``` sh
$GOPATH/src/golang.org/x/tools/godoc/server.go
- info.IsMain = pkgname == "main"
+ info.IsMain = false && pkgname == "main"
Build golangcd.org/x/tools/cmd/godoc and install.
```