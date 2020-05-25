package main

import (
	"bufio"
	"fmt"
	"go-auto-orm/camel"
	"log"
	"os"
	"os/signal"
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
				className := camel.Marshal(strings.Replace(t.Name, tablePrefix, "", 1), true)

				clazz := genData(t, packageName, tablePrefix)

				tempName := "temp/domain.java.tpl"
				fileName := filepath.Join(packageName, className+".java")
				writeTemp(tempName, fileName, clazz)

				tempName = "temp/dao.java.tpl"
				fileName = filepath.Join(packageName, className+"Dao.java")
				writeTemp(tempName, fileName, clazz)

				tempName = "temp/dao.xml.tpl"
				fileName = filepath.Join(packageName, className+"Dao.xml")
				writeTemp(tempName, fileName, clazz)
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
	clazz.ClassName = camel.Marshal(strings.Replace(t.Name, tablePrefix, "", 1), true)
	clazz.VariableName = camel.Marshal(strings.Replace(t.Name, tablePrefix, "", 1), false)

	clazz.PrimaryKey = t.PrimaryKeys[0]
	var fields []*Field
	for _, c := range t.Columns() {
		f := new(Field)
		f.JavaType = SQL2JAVAMap[c.SQLType.Name]
		f.Name = camel.Marshal(c.Name, false)
		ns := []rune(f.Name)
		f.UpperName = strings.ToUpper(string(ns[0])) + string(ns[1:])
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

func writeTemp(tempName, fileName string, data *Class) {
	temp, _ := template.ParseFiles(tempName)
	file, _ := os.OpenFile(filepath.Join(fileName), os.O_CREATE, 0644)
	defer file.Close()
	temp.Execute(file, data)
}
