<?xml version="1.0" encoding="UTF-8" ?>
<!DOCTYPE mapper PUBLIC "-//mybatis.org//DTD Mapper 3.0//EN" "http://mybatis.org/dtd/mybatis-3-mapper.dtd" >
<mapper namespace="{{.ClassName}}">
    <sql id="base-column-list"> 
    {{$first := true}}{{range .Fields}}{{if $first}}{{$first = false}}{{else}},{{end}}{{.Column}}{{end}}
    </sql>
    <sql id="table-name"> {{.TableName}} </sql>
    <sql id="from-table"> FROM {{.TableName}} </sql>

    <sql id="where-condition">
        WHERE 1=1
        {{range .Fields}}
        <if test="{{.Name}}!=null">AND {{.Column}}=#{ {{.Name}}, jdbcType={{.JdbcType}} }</if>{{ end }} 
    </sql>

    <select id="select" resultType="{{.ClassName}}">
		select * 
		<include refid="from-table"/>
        <include refid="where-condition"/>
    </select>

    <insert id="insert" >
        INSERT INTO
        <include refid="table-name"/>
        (
        {{$first := true}}{{range .Fields}}{{if $first}}{{$first = false}} {{else}}, {{end}}{{.Column}}{{end}}
        )
        VALUES (
        {{$first := true}} {{range .Fields}} {{if $first}}{{$first = false}}{{else}}, {{end}}#{ {{.Name}}, jdbcType={{.JdbcType}} } {{end}}
        )
    </insert>

    <update id="update" parameterType="{{.ClassName}}">
        UPDATE
        <include refid="table-name"/>
        <set>
            {{range .Fields}}
            <if test="{{.Name}}!=null">DATASTATUS = #{ {{.Name}},jdbcType= {{.Name}} },</if>{{ end }} 
        </set>
       WHERE {{.PrimaryKey}} =#{ {{.PrimaryName}},jdbcType={{.PrimaryJdbcType}} }
    </update>

    <delete id="delete" parameterType="{{.ClassName}}">
        DELETE
        <include refid="from-table"/>
        WHERE {{.PrimaryKey}} =#{ {{.PrimaryName}},jdbcType={{.PrimaryJdbcType}} }
    </delete>

</mapper>
