# api开发文档
====

## 接口定义:

## 编写工具
swagger

## 命名规范

url命名要求：
  - 尽量用标准英文单词（查字典）
  - 小写,组合单词用‘-’连接  

url命名举例:  
  - 增加单一类型数据 PUT /admin/question-libraries
  - 修改单一类型数据 PATCH /admin/question-libraries
  - 删除           DELETE /admin/question
  - 查询           GET  /admin/question
  - 非增加或修改单类型数据的动作,复杂业务逻辑 
                  POST /admin/some-action
  - 数据更新或者查询返回多条记录的 单词用复数

返回值字段命名:
  - 下划线  
  

url划分
 -date 后台管理接口  /admin/*
 - 用户接口  /user/*
