# bootstrap-global

bootstrap 系列核心依赖，提供统一的门面

## 原理

通过提供统一的门面，按需引入适配器即可通过门面使用。

适配器通过调用 register 方法进行注册，注册后通过 Init 方法初始化全局变量，通过门面调用的方法均会通过实例调用具体的实现

## 使用说明

使用步骤如下

1. 引入依赖, `go get -u github.com/jiurongzhao/bootstrap-config-yaml`
2. 在代码中导入适配器与门面, `_ "github.com/jiurongzhao/bootstrap-config-yaml/yaml"` 与 `"github.com/jiurongzhao/bootstrap-global/config"`
3. 在代码中初始化适配器, `config.InitGlobalInstance("yaml", "resource/app.yaml")`
4. 通过门面调用, `config.Get("foo.bb")` 与 `config.Resolve("foo.struct", &aStructInstance)`

[详见示例](https://github.com/jiurongzhao/bootstrap-example/blob/main/main.go)

## 开发说明

当前只做了一些基础的适配器，如有其他适配器需要可通过 issue 发起

- [ ] config
    - [x] bootstrap-config-yaml
    - [x] bootstrap-config-json
    - [x] bootstrap-config-apollo
    - [ ] bootstrap-config-properties
- [ ] log
    - [x] bootstrap-log-logrus
    - [x] bootstrap-log-zap

## 参考

1. `database/sql`
2. `beego/log`, `beego/config`

