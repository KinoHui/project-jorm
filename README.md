# Project JORM

JORM 是一个简单的 Go 语言 ORM 框架，专注于提供轻量级的数据库操作接口。

## 特性

- 链式调用 API
- 自动创建和迁移表结构
- 事务支持
- 钩子机制
- 支持多种数据库方言
- 丰富的查询构造器
- 完整的日志系统

## 安装

```bash
go get github.com/KinoHui/project-jorm
```

## 快速开始

### 1. 定义模型

```go
type User struct {
    Name string `jorm:"PRIMARY KEY"`
    Age  int
}
```

### 2. 创建引擎

```go
engine, err := jorm.NewEngine("sqlite3", "test.db")
if err != nil {
    log.Fatal(err)
}
defer engine.Close()
```

### 3. 表操作

```go
s := engine.NewSession()
// 创建表
_ = s.Model(&User{}).CreateTable()

// 插入数据
_, _ = s.Insert(&User{"Tom", 18})

// 查询
var users []User
_ = s.Find(&users)
```

## 高级特性

### 事务支持

```go
engine.Transaction(func(s *session.Session) (result interface{}, err error) {
    _ = s.Model(&User{}).CreateTable()
    _, err = s.Insert(&User{"Tom", 18})
    return
})
```

### 钩子函数

```go
type Account struct {
    ID       int `jorm:"PRIMARY KEY"`
    Password string
}

func (account *Account) BeforeInsert(s *Session) error {
    // 在插入前加密密码
    account.Password = hash(account.Password)
    return nil
}
```

### 条件查询

```go
var users []User
s.Where("Age > ?", 18).
  Limit(10).
  OrderBy("Age DESC").
  Find(&users)
```

### 表迁移

```go
// 自动迁移表结构
engine.Migrate(&User{})
```

## 支持的数据库

- SQLite3 (已实现)
- MySQL (待实现)
- PostgreSQL (待实现)

## 项目结构

```
project-jorm/
├── clause/     # SQL 子句生成
├── dialect/    # 数据库方言
├── schema/     # 数据库表结构
├── session/    # 会话管理
├── log/        # 日志系统
└── examples/   # 使用示例
```

## 核心接口

- Engine: 数据库引擎
- Session: 会话管理
- Clause: SQL 构造器
- Schema: 表结构映射
- Dialect: 数据库方言适配



## 致谢

此项目学习于[Geektutu](https://geektutu.com/)，感谢Geektutu的付出。
