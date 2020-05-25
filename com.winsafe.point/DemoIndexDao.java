package com.winsafe.point.dao;


import com.winsafe.point.model.DemoIndex;
import org.apache.ibatis.session.SqlSession;
import org.mybatis.spring.SqlSessionTemplate;
import org.springframework.stereotype.Repository;

import javax.annotation.Resource;
import java.util.List;

@Repository
public class DemoIndexDao {

    @Resource(name = "sqlSessionTemplate")
    protected SqlSessionTemplate sqlSession;

    public SqlSession getSqlSession() {
        return sqlSession;
    }


    public int insert(DemoIndex demoIndex) {
        return getSqlSession().insert(DemoIndex.class.getSimpleName() + ".insert", demoIndex);
    }

    public int delete(DemoIndex demoIndex) {
        return getSqlSession().delete(DemoIndex.class.getSimpleName() + ".delete",  demoIndex);
    }

    public int update(DemoIndex demoIndex) {
        return getSqlSession().update(DemoIndex.class.getSimpleName() + ".update", demoIndex);
    }

    public List<DemoIndex> select(DemoIndex demoIndex) {
        return getSqlSession().selectList(DemoIndex.class.getSimpleName() + ".select", demoIndex);
    }

 
}
