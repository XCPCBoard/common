DROP TABLE IF EXISTS users;
CREATE TABLE users(
                      `id` INT NOT NULL AUTO_INCREMENT  COMMENT '用户id' ,
                      `account` VARCHAR(128) NOT NULL   COMMENT '账号' ,
                      `keyword` VARCHAR(128) NOT NULL   COMMENT '密码' ,
                      `email` VARCHAR(128) NOT NULL   COMMENT '邮箱' ,
                      `is_administrator` INT(2) NOT NULL   COMMENT '管理员标签' ,
                      `name` VARCHAR(128)    COMMENT '姓名' ,
                      `image_path` VARCHAR(255)    COMMENT '头像路径' ,
                      `signature` VARCHAR(255)    COMMENT '个性签名' ,
                      `phone_number` VARCHAR(128)    COMMENT '手机号' ,
                      `qq_number` VARCHAR(128)    COMMENT 'qq号' ,
                      `created_at` TIMESTAMP    COMMENT '创建时间' ,
                      `updated_at` TIMESTAMP    COMMENT '更新时间' ,
                      PRIMARY KEY (id)
)  COMMENT = '用户表';


CREATE UNIQUE INDEX check_account ON users(account);
CREATE UNIQUE INDEX check_email ON users(email);

DROP TABLE IF EXISTS posts;
CREATE TABLE posts(
                      `id` INT NOT NULL AUTO_INCREMENT  COMMENT '主键id' ,
                      `user_id` INT NOT NULL   COMMENT '用户id' ,
                      `title` VARCHAR(128)    COMMENT '帖子标题' ,
                      `content` VARCHAR(1500)    COMMENT '帖子内容' ,
                      `note` VARCHAR(128)    COMMENT '备注（备用字段）' ,
                      `created_at` TIMESTAMP    COMMENT '创建时间' ,
                      `updated_at` TIMESTAMP    COMMENT '更新时间' ,
                      PRIMARY KEY (id)
)  COMMENT = '帖子表';

DROP TABLE IF EXISTS website_account;
CREATE TABLE website_account(
                                `id` INT NOT NULL   COMMENT 'id' ,
                                `codeforces` VARCHAR(128)    COMMENT 'codeforces' ,
                                `nowcoder` VARCHAR(128)    COMMENT 'nowcoder' ,
                                `luogu` VARCHAR(128)    COMMENT 'luogu' ,
                                `atcoder` VARCHAR(128)    COMMENT 'atcoder' ,
                                `vjudge` VARCHAR(128)    COMMENT 'vjudge' ,
                                `rank` INT    COMMENT '排名' ,
                                `total` INT    COMMENT '总过题数' ,
                                `created_at` TIMESTAMP    COMMENT '创建时间' ,
                                `updated_at` TIMESTAMP    COMMENT '更新时间' ,
                                PRIMARY KEY (id)
)  COMMENT = '网站账户名';

