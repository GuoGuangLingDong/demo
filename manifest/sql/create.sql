
use wesoul;
CREATE TABLE if not exists `user`
(
    `id`            int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'pk',
    `uid`      varchar(45) NOT NULL COMMENT 'User ID',
    `password`      varchar(45)      NOT NULL COMMENT 'User Password',
    `username`      varchar(45)      NOT NULL COMMENT 'User Name',
    `nickname`      varchar(40)      NOT NULL COMMENT 'User Nickname',
    `create_at`     datetime         NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Created Time',
    `update_at`     datetime                  DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT 'Updated Time',
    `phone_number`  varchar(20)      NOT NULL COMMENT 'Phone Number',
    `wechat_number` varchar(45) COMMENT 'Wechat Number',
    `invite_code`   varchar(45) COMMENT 'Invite Code',
    `introduction`  varchar(200)       COMMENT 'Introduction',
    `avatar`        varchar(200)              DEFAULT NULL COMMENT '头像',
    PRIMARY KEY (`id`),
    UNIQUE KEY (`uid`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

CREATE TABLE if not exists `userlink`
(
    `id`        int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'pk',
    `uid`      varchar(45) NOT NULL COMMENT 'User ID',
    `link`      varchar(32)      NOT NULL COMMENT 'Link',
    `link_type` int(10)          NOT NULL COMMENT 'Link type',
    `create_at` datetime         NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Created Time',
    `update_at` datetime                  DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT 'Updated Time',
    PRIMARY KEY (`id`),
    UNIQUE KEY (`uid`, `link_type`),
    foreign key (`uid`) references `user` (`uid`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;
select * from userlink;
create table if not exists `poap`
(
    `id`           int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'pk',
    `poap_id`      varchar(45) NOT NULL COMMENT 'Poap id',
    `miner`        varchar(45) NOT NULL COMMENT 'Miner',
    `poap_name`    varchar(45)      NOT NULL COMMENT 'Poap name',
    `poap_sum`     int(64)          NOT NULL comment 'Poap sum',
    `receive_cond` int(64)          NOT NULL comment 'Receive condition',
    `cover_img`    varchar(400)     NOT NULL COMMENT 'Cover picture',
    `poap_intro`   varchar(200)     not null comment 'Poap introduction',
    `create_at`    datetime         NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Created Time',
    `update_at`    datetime                  DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT 'Updated Time',
    PRIMARY KEY (`id`),
    UNIQUE KEY (`poap_id`),
    foreign key (`miner`) references `user` (`uid`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

create table if not exists `hold`
(
    `id`        int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'pk',
    `uid`      varchar(45) NOT NULL COMMENT 'User ID',
    `poap_id`   varchar(45) NOT NULL COMMENT 'Poap id',
    `create_at` datetime         NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Created Time',
    `update_at` datetime                  DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT 'Updated Time',
    `token_id`  varchar(20)               DEFAULT NULL COMMENT 'Poap tokenId',
    foreign key (`uid`) references `user` (`uid`),
    foreign key (`poap_id`) references `poap` (`poap_id`),
    PRIMARY KEY (`id`),
    index (`uid`),
    index (`poap_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

create table if not exists `like`
(
    `id`        int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'pk',
    `uid`      varchar(45) NOT NULL COMMENT 'User ID',
    `poap_id`   varchar(45) NOT NULL COMMENT 'Poap id',
    `create_at` datetime         NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Created Time',
    `update_at` datetime                  DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT 'Updated Time',
    foreign key (`uid`) references `user` (`uid`),
    foreign key (`poap_id`) references `poap` (`poap_id`),
    PRIMARY KEY (`id`),
    index (`poap_id`),
    UNIQUE KEY (`uid`, `poap_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

create table if not exists `follow`
(
    `id`        int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'pk',
    `followee`  varchar(45) NOT NULL COMMENT 'Followee ID',
    `follower`  varchar(45) NOT NULL COMMENT 'Follower id',
    `create_at` datetime         NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Created Time',
    `update_at` datetime                  DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT 'Updated Time',
    foreign key (`followee`) references `user` (`uid`),
    foreign key (`follower`) references `user` (`uid`),
    PRIMARY KEY (`id`),
    UNIQUE KEY (`followee`, `follower`),
    UNIQUE KEY (`follower`, `followee`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

create table if not exists `operation`
(
    `id`        int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'pk',
    `uid`      varchar(45) NOT NULL COMMENT 'User ID',
    `opt_type`  int(32) unsigned NOT NULL COMMENT 'Operate Type',
    `score`     int(64) unsigned NOT NULL COMMENT 'Score',
    `create_at` datetime         NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Created Time',
    `update_at` datetime                  DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT 'Updated Time',
    foreign key (`uid`) references `user` (`uid`),
    index (`uid`),
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

CREATE TABLE `publish`
(
    `id`            bigint           NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `poap_id`       varchar(45)    NOT NULL COMMENT 'Poap id',
    `token_id`      varchar(20)               DEFAULT NULL COMMENT 'Poap tokenId',
    `status`        varchar(10)      NOT NULL DEFAULT 'disable' COMMENT '状态 disable:未使用 used.已使用',
    `no`            bigint unsigned  NOT NULL DEFAULT '0' COMMENT '编号',
    `chain_status`  tinyint(1)       NOT NULL DEFAULT '0' COMMENT '上链状态  0.未上链 1.已上链  2.上链中 3.上链失败',
    `chain_hash`    varchar(255)     NOT NULL DEFAULT '' COMMENT '链上hash',
    `is_error`      int              NOT NULL DEFAULT '0' COMMENT '是否异常',
    `error_message` text COMMENT '错误信息',
    `lock_flag`     tinyint unsigned NOT NULL DEFAULT '1',
    `created_at`    datetime                  DEFAULT CURRENT_TIMESTAMP,
    `updated_at`    datetime                  DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`) USING BTREE,
    UNIQUE KEY `uk_token_id` (`token_id`),
    KEY `idx_poap_id` (`poap_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='poap发行';
