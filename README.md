# Golang-Gin-Gorm Project Template

Go Gin WebFramework + Gorm 프로젝트 템플릿

- [Gin WebFramework](https://github.com/gin-gonic/gin)
- [GORM](https://gorm.io/index.html)

```shell
Controllers -> Services -> Repositories -> Models 
```

```shell
Dependency Injection
wire
```

```shell
go run main.go wire_gen.go
```

```shell
go build main.go wire_gen.go
./main
```

After modify swagger options
```shell
 swag init -parseDependency=true
```