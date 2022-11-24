use wesoul;
CREATE TABLE if not exists `user`
(
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT 'pk',
    `uid`           varchar(32) NOT NULL COMMENT 'User ID',
    `password`      varchar(45)      NOT NULL COMMENT 'User Password',
    `username`      varchar(45)      NOT NULL COMMENT 'User Name',
    `nickname`  varchar(40) NOT NULL COMMENT 'User Nickname',
    `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Created Time',
    `update_at` datetime DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT 'Updated Time',
    `phone_number`  varchar(20)      NOT NULL COMMENT 'Phone Number',
    `wechat_number` varchar(45) COMMENT 'Wechat Number',
    `invite_code`   varchar(45) COMMENT 'Invite Code',
    `introduction`  varchar(200)     NOT NULL COMMENT 'Introduction',
    `avatar` blob DEFAULT NULL COMMENT '头像',
    `scores`        int(32) unsigned NOT NULL default (0) COMMENT 'Scores',
    PRIMARY KEY (`id`),
    UNIQUE KEY(`uid`)
    ) ENGINE = InnoDB
    DEFAULT CHARSET = utf8;

CREATE TABLE if not exists `user`(
    `uid`           varchar(32) NOT NULL COMMENT 'User ID',
    `link`  varchar(32) NOT NULL COMMENT 'Link',
    `link_type` int(32) NOT NULL COMMENT 'Link type'
    ) ENGINE = InnoDB
    DEFAULT CHARSET = utf8;

create table if not exists `poap`
(
    `poap_id`           int(64) unsigned NOT NULL AUTO_INCREMENT COMMENT 'Poap id',
    `miner`             int(64) unsigned NOT NULL COMMENT 'Miner',
    `poap_name`         varchar(45)      NOT NULL COMMENT 'Poap name',
    `poap_number`       int(64)          NOT NULL comment 'Poap number',
    `receive_condition` int(64)          NOT NULL comment 'Receive condition',
    `cover_pic`         varchar(400)     NOT NULL COMMENT 'Cover picture',
    `poap_intro`        varchar(200)     not null comment 'Poap introduction',
    `favour_number`     int(64) unsigned NOT NULL default (0) comment 'Favour_number',
    PRIMARY KEY (`poap_id`),
    foreign key (`miner`) references `user` (`uid`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

create table if not exists `hold`
(
    `uid`     int(64) unsigned NOT NULL COMMENT 'User ID',
    `poap_id` int(64) unsigned NOT NULL COMMENT 'Poap id',
    foreign key (`uid`) references `user` (`uid`),
    foreign key (`poap_id`) references `poap` (`poap_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

create table if not exists `favor`
(
    `uid`        int(64) unsigned NOT NULL COMMENT 'User ID',
    `poap_id`    int(64) unsigned NOT NULL COMMENT 'Poap id',
    `favor_time` datetime DEFAULT NULL COMMENT 'Favor Time',
    foreign key (`uid`) references `user` (`uid`),
    foreign key (`poap_id`) references `poap` (`poap_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

create table if not exists `follow`
(
    `followee`    int(64) unsigned NOT NULL COMMENT 'Followee ID',
    `follower`    int(64) unsigned NOT NULL COMMENT 'Follower id',
    `follow_time` datetime DEFAULT NULL COMMENT 'Follow Time',
    foreign key (`followee`) references `user` (`uid`),
    foreign key (`follower`) references `user` (`uid`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;

create table if not exists `operation`
(
    `uid`          int(64) unsigned NOT NULL COMMENT 'User ID',
    `operate_code` int(32) unsigned NOT NULL COMMENT 'Operate Code',
    `operate_time` datetime DEFAULT NULL COMMENT 'Operate Time',
    `score`        int(64) unsigned NOT NULL COMMENT 'Score',
    foreign key (`uid`) references `user` (`uid`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;
