# 1. 文献引用关系图

## 数据预处理

文献id 文献名 引用的文献id 引用的文献名

## 后端请求处理

```json
request：{
    docid
}
```

关系包括一级引用、二级引用以及三级引用

三次数据库查询

```json
response:
{
    rootID:'',根节点编号，即本文献编号，
    nodes:所有文献节点[
        {
            id:'',文献编号
            text:'',文献title
        }
    ]
    links:所有关系[
        {
            from:''上层节点的文献编号
            to:''下层节点的文献编号
            text:''第几级引用
        }
    ]
}
```

# 2. 作者关系图

关系包括合作伙伴关系图、同机构关系图

## 合作伙伴

### 数据预处理

作者1id 作者1姓名 作者2id 作者2姓名 合作的文献列表 并查集id

### 后端请求处理

两次数据库查询，第一次查询作者id所在的并查集id，第二次获取所有该并查集下的关系，即返回结果

```json
request：{
    authorid
}
response:
{
    rootID:'',根节点编号，即本作者id
    nodes:所有作者节点[
        {
            id:'',作者id
            text:'',作者名字
        }
    ]
    links:所有关系[
        {
            from:''作者1id
            to:''作者2id
            text:''合作的文献列表
        }
    ]
}
```

## 同机构

## 数据预处理

机构id 机构名字 作者列表[{id,name}]

## 后端请求处理

一次数据库查询，根据机构id找到所有作者

```json
request:{
    affiliation
}
response:
{
    rootID:'',根节点编号，即机构id
    nodes:所有节点，包括一个机构节点和所有的作者节点[
        {
            id:'',作者（机构）id
            text：'',作者（机构）名字
        }
    ]
    links：所有关系[
        {
            from:''机构id
            to:''作者id
            text:'',空
        }
    ]
}
```