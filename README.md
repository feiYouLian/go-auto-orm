# Mysql ORM 自动生成Java代码

## Mysql --- Java 类型映射

```go
var SQL2JAVAMap = map[string]string{
	"VARCHAR":   "String",
	"CHAR":      "String",
	"BLOB":      "byte[]",
	"TEXT":      "String",
	"INTEGER":   "Long",
	"TINYINT":   "Integer",
	"SMALLINT":  "Integer",
	"BIT":       "Boolean",
	"BIGINT":    "Long",
	"INT":       "Integer",
	"FLOAT":     "Float",
	"DOUBLE":    "Double",
	"DECIMAL":   "BigDecimal",
	"BOOLEAN":   "Boolean",
	"DATE":      "Date",
	"TIME":      "Date",
	"DATETIME":  "Date",
	"TIMESTAMP": "Date",
}
```

## 信息详情
```go
// Field 字段
type Field struct {
	JavaType  string // Java 类型
	Name      string // Java 名字 如 id
	UpperName string // 首字母大写 如 Id (setter和getter使用)
	JdbcType  string // Mysql 类型
	Column    string // Mysql 字段名
}

// Class Class
type Class struct {
	PackageName     string //包名
	ClassName       string // Java类名称
	VariableName    string // Java类变量（也就是首字母变小写）
	TableName       string // Mysql 表名
	PrimaryJavaType string // 主键 Java 类型
	PrimaryName     string // 主键名
	PrimaryJdbcType string // 主键 Mysql 类型
	PrimaryKey      string // 主键
	Fields          []*Field // 其他字段
}
```

## 使用

1. 修改`db.yml`, 换成自己的配置

2. 若是安装了`golang`, 在当前目录下`go run .`

3. 若是没装`golang`, 直接执行`go-auto-orm.exe`
