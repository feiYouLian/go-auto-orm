package {{.PackageName}}.dao;


import {{.PackageName}}.model.{{.ClassName}};
import org.apache.ibatis.session.SqlSession;
import org.mybatis.spring.SqlSessionTemplate;
import org.springframework.stereotype.Repository;

import javax.annotation.Resource;
import java.util.List;

@Repository
public class {{.ClassName}}Dao {

    @Resource(name = "sqlSessionTemplate")
    protected SqlSessionTemplate sqlSession;

    public SqlSession getSqlSession() {
        return sqlSession;
    }


    public int insert({{.ClassName}} {{.VariableName}}) {
        return getSqlSession().insert({{.ClassName}}.class.getSimpleName() + ".insert", {{.VariableName}});
    }

    public int delete({{.ClassName}} {{.VariableName}}) {
        return getSqlSession().delete({{.ClassName}}.class.getSimpleName() + ".delete",  {{.VariableName}});
    }

    public int update({{.ClassName}} {{.VariableName}}) {
        return getSqlSession().update({{.ClassName}}.class.getSimpleName() + ".update", {{.VariableName}});
    }

    public List<{{.ClassName}}> select({{.ClassName}} {{.VariableName}}) {
        return getSqlSession().selectList({{.ClassName}}.class.getSimpleName() + ".select", {{.VariableName}});
    }

 
}
