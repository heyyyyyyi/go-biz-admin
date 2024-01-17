# 关于Golang的一些学习

## terminal

- go mod init github.com/...
  - go.mod
- go mod tidy
  - go.sum
  - [update] go.mod
- go build .
  - ./...

## grammar

- module
- package
- 全局变量 var XXX
- 函数 func Xxx    func xxx
- interface

### gorm

> 与数据库交互

- gorm:"foreignKey:RoleId"
- gorm:"many2many:role_permissions"
- gorm:"-" //不存
- db.preload()
- db.find()
- db.update()
- db.delete()

### gin

> 与前端交互
