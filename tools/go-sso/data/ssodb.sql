-- mysql version: 8.0.19
-- 数据库： `ssodb`

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";

CREATE TABLE `device`
(
    `id`     bigint(20) UNSIGNED NOT NULL COMMENT '主键',
    `uid`    bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '用户主键',
    `client` varchar(50)         NOT NULL DEFAULT '' COMMENT '客户端',
    `model`  varchar(50)         NOT NULL DEFAULT '' COMMENT '设备型号',
    `ip`     int(10) UNSIGNED    NOT NULL DEFAULT '0' COMMENT 'ip地址',
    `ext`    varchar(1000)       NOT NULL DEFAULT '' COMMENT '扩展信息',
    `ctime`  int(10) UNSIGNED    NOT NULL DEFAULT '0' COMMENT '注册时间'
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci;

ALTER TABLE `device`
    ADD PRIMARY KEY (`id`),
    ADD KEY `uid` (`uid`);
ALTER TABLE `device`
    MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键';

CREATE TABLE `trace`
(
    `id`    bigint(20) UNSIGNED NOT NULL COMMENT '主键',
    `uid`   bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '用户主键',
    `type`  tinyint(4)          NOT NULL DEFAULT '0' COMMENT '类型(0:注册1:登录2:退出3:修改4:删除)',
    `ip`    int(10) UNSIGNED    NOT NULL COMMENT 'ip',
    `ext`   varchar(1000)       NOT NULL COMMENT '扩展字段',
    `ctime` int(11) UNSIGNED    NOT NULL DEFAULT '0' COMMENT '注册时间'
) ENGINE = MyISAM
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci;

ALTER TABLE `trace`
    ADD PRIMARY KEY (`id`),
    ADD KEY `UT` (`uid`, `type`) USING BTREE;
ALTER TABLE `trace`
    MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
    AUTO_INCREMENT = 2;

CREATE TABLE `users`
(
    `id`     bigint(20) UNSIGNED NOT NULL COMMENT '主键',
    `name`   varchar(50)         NOT NULL DEFAULT '' COMMENT '用户名',
    `email`  varchar(100)        NOT NULL DEFAULT '' COMMENT '邮箱',
    `mobile` varchar(20)         NOT NULL DEFAULT '' COMMENT '手机号',
    `passwd` varchar(40)         NOT NULL COMMENT '密码',
    `salt`   char(4)             NOT NULL COMMENT '盐值',
    `ext`    text                NOT NULL COMMENT '扩展字段',
    `status` tinyint(4)          NOT NULL DEFAULT '0' COMMENT '状态（0：未审核,1:通过 10删除）',
    `ctime`  int(10) UNSIGNED    NOT NULL DEFAULT '0' COMMENT '创建时间',
    `mtime`  timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间'
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_general_ci;

ALTER TABLE `users`
    ADD PRIMARY KEY (`id`),
    ADD KEY `ctime` (`ctime`);
ALTER TABLE `users`
    MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键';

COMMIT;
