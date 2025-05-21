# Capys-OJ

## Online Judge 在线判题系统

### 功能
用户模块
- 用户基本信息
  - 邮箱
  - 手机号
  - 密码
  - 安全码
  - 用户权限
    - 学生
    - 教师
    - 管理员
- 用户信息
  - 用户名
  - 学号/工号
  - 头像

## 用户凭证

### 注册功能

1. 写入数据到user
   1. email *
   2. phone *
   3. password *
   4. safe

2. 写入数据到user_info
   1. username *
   2. id_number *
   3. avatar



### 登录功能



### JWT

JWT token 发给用户

自定义 存储在服务器


用户id + 时间戳 转为 base64 = token

写在数据库中,

redis 2ms
mysql 200ms



## 分组

group 分组

user 用户


user 1 -> group n
group 1 -> n user

1:1 个人与个人的身份证
1:n 个人与手机号
n:n 一个人可以加入多个分组，一个分组有多个人

user 存储用户信息
group 存储分组信息

user_group 用户-分组中间表
user_id 指向user
group_id 指向group



## 数据表说明

user

存储用户基本信息：仅包含账号、密码、权限等基本信息

user_info

存储用户的信息：包含用户名等信息


### 表关系说明
