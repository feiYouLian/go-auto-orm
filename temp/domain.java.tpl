package {{.PackageName}}.model;

public class {{.ClassName}}{
    {{range .Fields}}
    private {{.JavaType}} {{.Name}};
    {{ end }}

    {{range .Fields}}
    public void set{{.UpperName}}({{.JavaType}} {{.Name}}){
        this.{{.Name}} = {{.Name}};
    }
    public {{.JavaType}} get{{.UpperName}}(){
       return this.{{.Name}};
    }
    {{ end }}

}
