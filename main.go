package main

import (
	"bufio"
	"fmt"
	"go-auto-orm/camel"
	"log"
	"os"
	"os/signal"
	"path"
	"path/filepath"
	"strings"
	"syscall"
	"text/template"

	"xorm.io/core"
)

// Field Field
type Field struct {
	JavaType  string
	Name      string
	UpperName string
	JdbcType  string
	Column    string
}

// Class Class
type Class struct {
	PackageName     string
	ClassName       string
	VariableName    string
	TableName       string
	PrimaryJavaType string
	PrimaryName     string
	PrimaryJdbcType string
	PrimaryKey      string
	Fields          []*Field
}

func main() {
	doOrm()
	go intercept()
	select {}
}

func intercept() {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGQUIT)
	fmt.Println("ctrl+c 中止任务")
	<-ch
	log.Println("任务已终止")
	os.Exit(0)
	// do things when catch a close signal
}

func doOrm() {

	var tableNames []string = []string{"t_demo_index"}
	fmt.Printf("请输入表要映射的表名（逗号分隔）:\n")
	input := bufio.NewScanner(os.Stdin)
	if input.Scan() {
		tableNames = strings.Split(input.Text(), ",")
	}
	if len(tableNames) == 0 {
		return
	}

	var tablePrefix string = "t_"
	fmt.Printf("请输入表前缀:\n")
	if input.Scan() {
		tablePrefix = input.Text()
	}

	var packageName string = "com.winsafe.point"
	fmt.Printf("请输入包名:\n")
	if input.Scan() {
		packageName = input.Text()
	}
	if len(packageName) > 0 {
		os.MkdirAll(packageName, os.ModeDir)
	}

	tables, _ := DB.DBMetas()

	for _, t := range tables {
		for _, tn := range tableNames {

			if tn == t.Name {
				clazz := genData(t, packageName, tablePrefix)
				fmt.Println(clazz)

				tempName := "temp/domain.java.tpl"
				writeTemp(tempName, clazz)

				tempName = "temp/dao.java.tpl"
				writeTemp(tempName, clazz)

				tempName = "temp/dao.xml.tpl"
				writeTemp(tempName, clazz)
			}

		}
	}
	wd, _ := os.Getwd()
	fmt.Printf("映射文件已生成,在 %s \n", wd)

}

func genData(t *core.Table, packageName, tablePrefix string) *Class {
	clazz := new(Class)
	clazz.PackageName = packageName
	clazz.TableName = t.Name
	clazz.VariableName = camel.Marshal(strings.Replace(t.Name, tablePrefix, "", 1))
	vns := []rune(clazz.VariableName)
	vns[0] -= 32
	clazz.ClassName = string(vns)

	clazz.PrimaryKey = t.PrimaryKeys[0]
	var fields []*Field
	for _, c := range t.Columns() {
		f := new(Field)
		f.JavaType = SQL2JAVAMap[c.SQLType.Name]
		f.Name = camel.Marshal(c.Name)
		ns := []rune(f.Name)
		ns[0] -= 32
		f.UpperName = string(ns)
		fmt.Println(f.UpperName)
		f.JdbcType = c.SQLType.Name
		f.Column = c.Name
		fields = append(fields, f)
		if clazz.PrimaryKey == c.Name {
			clazz.PrimaryJavaType = f.JavaType
			clazz.PrimaryName = f.Name
			clazz.PrimaryJdbcType = f.JdbcType
		}
	}
	clazz.Fields = fields

	return clazz
}

func writeTemp(tmplPath string, data *Class) {
	tmpl, _ := template.ParseFiles(tmplPath)
	tmplName := path.Base(tmplPath)
	suffix := strings.Split(tmplName, ".")[0]
	extension := strings.Split(tmplName, ".")[1]
	var fileName string
	if strings.Index(tmplName, "domain") >= 0 {
		fileName = data.ClassName + "." + extension
	} else {
		svs := []rune(suffix)
		fileName = data.ClassName + strings.ToUpper(string(svs[0])) + string(svs[1:]) + "." + extension
	}

	file, _ := os.OpenFile(filepath.Join(data.PackageName, fileName), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	defer file.Close()
	tmpl.Execute(file, data)
}
