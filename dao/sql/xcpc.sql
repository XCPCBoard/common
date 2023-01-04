DROP TABLE IF EXISTS users;
CREATE TABLE users(
                      id INT NOT NULL AUTO_INCREMENT  COMMENT '用户id' ,
                      account VARCHAR(64) NOT NULL   COMMENT '账号' ,
                      keyword VARCHAR(64) NOT NULL   COMMENT '密码' ,
                      email VARCHAR(64) NOT NULL   COMMENT '邮箱' ,
                      is_administrator VARCHAR(1)    COMMENT '管理员标签' ,
                      name VARCHAR(64) NOT NULL   COMMENT '姓名' ,
                      image_path VARCHAR(255)    COMMENT '头像路径' ,
                      signature VARCHAR(255)    COMMENT '个性签名' ,
                      phone_number VARCHAR(64)    COMMENT '手机号' ,
                      qq_number VARCHAR(64)    COMMENT 'qq号' ,
                      created_at DATETIME    COMMENT '创建时间' ,
                      updated_at DATETIME    COMMENT '更新时间' ,
                      PRIMARY KEY (id)
)  COMMENT = '用户表';


CREATE UNIQUE INDEX check_account ON users(account);
CREATE UNIQUE INDEX check_email ON users(email);

DROP TABLE IF EXISTS posts;
CREATE TABLE posts(
                      id INT NOT NULL AUTO_INCREMENT  COMMENT '主键id' ,
                      user_id INT NOT NULL   COMMENT '用户id' ,
                      title VARCHAR(64)    COMMENT '帖子标题' ,
                      content VARCHAR(1500)    COMMENT '帖子内容' ,
                      note VARCHAR(64)    COMMENT '备注（备用字段）' ,
                      created_at DATETIME    COMMENT '创建时间' ,
                      updated_at DATETIME    COMMENT '更新时间' ,
                      PRIMARY KEY (id)
)  COMMENT = '帖子表';

DROP TABLE IF EXISTS website_account;
CREATE TABLE website_account(
                                id INT NOT NULL   COMMENT 'id' ,
                                codeforces VARCHAR(64)    COMMENT 'codeforces' ,
                                nowcoder VARCHAR(64)    COMMENT 'nowcoder' ,
                                luogu VARCHAR(64)    COMMENT 'luogu' ,
                                atcoder VARCHAR(64)    COMMENT 'atcoder' ,
                                vjudge VARCHAR(64)    COMMENT 'vjudge' ,
                                created_at DATETIME    COMMENT '创建时间' ,
                                updated_at DATETIME    COMMENT '更新时间' ,
                                PRIMARY KEY (id)
)  COMMENT = '网站账户名';
