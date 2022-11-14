use tweat;

CREATE TABLE IF NOT EXISTS `tweats` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ツイートID',
    `text` varchar(255) NOT NULL COMMENT 'ツイートテキスト',
    `user_id` bigint(20) NOT NULL COMMENT 'ユーザーID',
    `parent_id` bigint(20) DEFAULT NULL COMMENT '親ツイート',
    `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '作成日',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日',
    PRIMARY KEY(`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;

INSERT INTO `tweats` (`id`, `text`, `user_id`, `parent_id`) VALUES
    (1, 'text1', 1, NULL),
    (2, 'text2', 2, NULL),
    (3, 'text3', 3, NULL),
    (4, 'text4', 1, NULL),
    (5, 'text5', 1, NULL),
    (6, 'text6', 2, NULL),
    (7, 'comment1', 2, 2),
    (8, 'comment2', 1, 3),
    (9, 'comment3', 5, 2);

CREATE TABLE IF NOT EXISTS `likes` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'いいねID',
    `tweat_id` bigint(20) NOT NULL COMMENT 'ユーザーID',
    `user_id` bigint(20) NOT NULL COMMENT 'ユーザーID',
    PRIMARY KEY(`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;

INSERT INTO `likes` (`id`, `tweat_id`, `user_id`) VALUES
    (1, 1, 1),
    (2, 1, 2),
    (3, 1, 3),
    (4, 2, 1),
    (5, 2, 2),
    (6, 2, 3);

CREATE TABLE IF NOT EXISTS `comments` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'コメントID',
    `text` varchar(255) NOT NULL COMMENT 'コメント',
    `tweat_id` bigint(20) NOT NULL COMMENT 'ユーザーID',
    `user_id` bigint(20) NOT NULL COMMENT 'ユーザーID',
    PRIMARY KEY(`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;

INSERT INTO `comments` (`id`, `text`, `tweat_id`, `user_id`) VALUES
    (1, 'comment1', 1, 1),
    (2, 'comment2', 1, 2),
    (3, 'comment3', 1, 3),
    (4, 'comment4', 2, 2),
    (5, 'comment5', 2, 3);
