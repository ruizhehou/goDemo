# goDemo

一个演示 `go generate` 用法的 Go 示例项目。

## 项目结构

```
goDemo/
├── main.go              # 程序入口，演示两种代码生成方式
├── color/
│   ├── color.go         # 定义 Color 枚举类型，包含 go:generate 指令
│   └── color_string.go  # 由 stringer 自动生成，勿手动修改
├── mathops/
│   ├── mathops.go       # 包声明，包含 go:generate 指令
│   └── ops_gen.go       # 由 gen/gen.go 自动生成，勿手动修改
└── gen/
    └── gen.go           # 自定义代码生成器（模板方式）
```

## 两种代码生成方式

### 方式一：stringer（枚举转字符串）

`color/color.go` 定义了 `Color` 枚举，通过 `stringer` 工具自动生成 `String()` 方法：

```go
//go:generate stringer -type=Color
type Color int
const (
    Red Color = iota
    Green
    Blue
    // ...
)
```

生成后，枚举值可以直接打印为可读字符串：

```go
fmt.Println(color.Red)   // 输出: Red
fmt.Println(color.Green) // 输出: Green
```

### 方式二：自定义模板生成器

`gen/gen.go` 是一个基于 `text/template` 的代码生成器，读取操作定义列表，生成 `mathops/ops_gen.go`：

```go
//go:generate go run ../gen/gen.go
```

在 `gen.go` 中添加新操作，重新运行 `go generate` 即可自动生成对应函数。

## 快速开始

### 安装依赖工具

```bash
go install golang.org/x/tools/cmd/stringer@latest
```

### 重新生成代码

```bash
# 生成所有包的代码
go generate ./...

# 只生成 color 包（stringer）
go generate ./color/...

# 只生成 mathops 包（自定义模板）
go generate ./mathops/...
```

### 运行示例

```bash
go run main.go
```

输出示例：

```
=== Demo 1: stringer (enum -> String()) ===
  color value 0 -> Red
  color value 1 -> Green
  color value 2 -> Blue
  color value 3 -> Yellow
  color value 4 -> Purple

=== Demo 2: template generator (math ops) ===
  Add(3, 4)      = 7
  Subtract(10, 3) = 7
  Multiply(6, 7)  = 42

  All ops via map:
    Add        -> 15
    Subtract   -> 5
    Multiply   -> 50
```

## go generate 使用说明

- `//go:generate` 注释必须顶格写，`//` 后**不能有空格**
- `go generate` 不会在 `go build` 时自动运行，需要手动执行
- 生成的文件（`*_string.go`、`*_gen.go`）已提交到仓库，无需每次构建前都运行
- 修改了枚举或生成器逻辑后，需重新运行 `go generate` 并提交生成的文件
