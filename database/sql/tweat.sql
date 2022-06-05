use tweat;

CREATE TABLE IF NOT EXISTS `tweats` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ツイートID',
    `text` varchar(255) NOT NULL COMMENT 'ツイートテキスト',
    `user_id` bigint(20) NOT NULL COMMENT 'ユーザーID',
    `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '作成日',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日',
    PRIMARY KEY(`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;

INSERT INTO `tweats` (`id`, `text`, `user_id`) VALUES
    (1, 'text1', 1),
    (2, 'text2', 2),
    (3, 'text3', 3),
    (4, 'text4', 1),
    (5, 'text5', 1),
    (6, 'text6', 2);