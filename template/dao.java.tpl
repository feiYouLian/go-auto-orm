package {{.PackageName}}.dao;

import {{.PackageName}}.model.{{.ClassName}};

import java.util.List;

public interface {{.ClassName}}Dao {

    public int insert({{.ClassName}} {{.VariableName}});

    public int delete({{.ClassName}} {{.VariableName}});

    public int update({{.ClassName}} {{.VariableName}});

    public List<{{.ClassName}}> select({{.ClassName}} {{.VariableName}});

 
}
