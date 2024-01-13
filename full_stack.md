# Full Stack With Angular

## 1. 逻辑结构

![1704852387648](image/full_stack/1704852387648.png)

### 需求

> - dashboard users roles products orders
>
>
> - 用户管理， 商品目录， 订单内容
>
>   ``Role based access control``
>
> 权限：数据操作权限（CRUD）
>
> 不同角色：admin， sales，customers

### 1.1 前端

    - React

- json

  - {}
  - []

### 1.2 后端

- 业务逻辑
  - role

    - user story
  - infra

    - 表与关系
      - 一对一
      - 多对一
- 交互 Go lang
  - GORM (后端-infra交互)
  - Gin（前端-后端交互）

### 1.3 架构

- 数据
  - 持久化
  - 选择展示
  - 组织 检索
- CRUD
  - post， get， put， delete

## 2. 实现

### 2.1 后端

> database (连接数据库)
>
>     - connect 用于连接
>
> routes（连接前端，接受用户指令）
>
> models （定义不同user 不同角色 不同权限）
>
