use wesoul;

CREATE TABLE if not exists `user`
(
    `id`            int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'pk',
    `uid`           int(64) unsigned NOT NULL COMMENT 'User ID',
    `password`      varchar(45)      NOT NULL COMMENT 'User Password',
    `username`      varchar(45)      NOT NULL COMMENT 'User Name',
    `nickname`      varchar(40)      NOT NULL COMMENT 'User Nickname',
    `create_at`     datetime         NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Created Time',
    `update_at`     datetime                  DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT 'Updated Time',
    `phone_number`  varchar(20)      NOT NULL COMMENT 'Phone Number',
    `wechat_number` varchar(45) COMMENT 'Wechat Number',
    `invite_code`   varchar(45) COMMENT 'Invite Code',
    `introduction`  varchar(200)     NOT NULL COMMENT 'Introduction',
    `avatar`        varchar(200)              DEFAULT NULL COMMENT '头像',
    PRIMARY KEY (`id`),
    UNIQUE KEY (`uid`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

CREATE TABLE if not exists `userlink`
(
    `id`        int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'pk',
    `uid`       int(64) unsigned NOT NULL COMMENT 'User ID',
    `link`      varchar(32)      NOT NULL COMMENT 'Link',
    `link_type` int(10)          NOT NULL COMMENT 'Link type',
    `create_at` datetime         NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Created Time',
    `update_at` datetime                  DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT 'Updated Time',
    PRIMARY KEY (`id`),
    UNIQUE KEY (`uid`, `link_type`),
    foreign key (`uid`) references `user` (`uid`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

create table if not exists `poap`
(
    `id`           int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'pk',
    `poap_id`      int(64) unsigned NOT NULL COMMENT 'Poap id',
    `miner`        int(64) unsigned NOT NULL COMMENT 'Miner',
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
    `uid`       int(64) unsigned NOT NULL COMMENT 'User ID',
    `poap_id`   int(64) unsigned NOT NULL COMMENT 'Poap id',
    `create_at` datetime         NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Created Time',
    `update_at` datetime                  DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT 'Updated Time',
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
    `uid`       int(64) unsigned NOT NULL COMMENT 'User ID',
    `poap_id`   int(64) unsigned NOT NULL COMMENT 'Poap id',
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
    `followee`  int(64) unsigned NOT NULL COMMENT 'Followee ID',
    `follower`  int(64) unsigned NOT NULL COMMENT 'Follower id',
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
    `id`           int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'pk',
    `uid`          int(64) unsigned NOT NULL COMMENT 'User ID',
    `opt_type`     int(32) unsigned NOT NULL COMMENT 'Operate Type',
    `score`        int(64) unsigned NOT NULL COMMENT 'Score',
    `create_at`    datetime         NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Created Time',
    `update_at`    datetime                  DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT 'Updated Time',
    foreign key (`uid`) references `user` (`uid`),
    index (`uid`),
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;
