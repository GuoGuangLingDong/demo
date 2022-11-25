CREATE TABLE `user`
(
    `id`        int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'pk',
    `uid`       varchar(32) NOT NULL COMMENT 'User ID',
    `username`  varchar(20) NOT NULL COMMENT 'Username',
    `password`  varchar(20) NOT NULL COMMENT 'User Password',
    `nickname`  varchar(40) NOT NULL COMMENT 'User Nickname',
    `phone_number` varchar(20) DEFAULT NULL COMMENT 'Phone Number',
    `wechat_number` varchar(20) DEFAULT NULL COMMENT 'Wechat Number',
    `invite_code` varchar(20) DEFAULT NULL COMMENT '邀请码',
    `introduction` varchar(200) DEFAULT NULL COMMENT '自我介绍',
    `avatar` blob DEFAULT NULL COMMENT '头像',
    `score` int(10) DEFAULT NULL COMMENT '积分',
    `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Created Time',
    `update_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT 'Updated Time',
    PRIMARY KEY (`id`),
    UNIQUE KEY(`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
