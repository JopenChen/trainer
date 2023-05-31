#### 项目创建
```
jctl api new trainer --style go_zero
```

#### 根据 api 文件改动更新代码
```
jctl api go -api .\trainer.api -dir . --style go_zero
```

#### 根据 sql 文件改动更新代码
```
goctl model mysql ddl -src .\sql\xxx.sql -dir . --strict --style go_zero
```