use tweat;

CREATE TABLE IF NOT EXISTS `users` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'ユーザーID',
    `name` varchar(255) NOT NULL COMMENT 'ユーザー名',
    `email` varchar(255) NOT NULL COMMENT 'メールアドレス',
    `password` varchar(255) NOT NULL COMMENT 'パスワード',
    `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '作成日',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日',
    PRIMARY KEY(`id`),
    UNIQUE(`email`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;

INSERT INTO `users` (`id`, `name`, `email`, `password`) VALUES
    (1, 'testname1', 'test1@test.email.com', '$2a$10$hEMqR7m/llJ2AaMl55t4Tu2XhXVxEVf41fFDd2SJkBrYhN.QFqPVK'),
    (2, 'testname2', 'test2@test.email.com', '$2a$10$hEMqR7m/llJ2AaMl55t4Tu2XhXVxEVf41fFDd2SJkBrYhN.QFqPVK'),
    (3, 'testname3', 'test3@test.email.com', '$2a$10$hEMqR7m/llJ2AaMl55t4Tu2XhXVxEVf41fFDd2SJkBrYhN.QFqPVK'),
    (4, 'testname4', 'test4@test.email.com', '$2a$10$hEMqR7m/llJ2AaMl55t4Tu2XhXVxEVf41fFDd2SJkBrYhN.QFqPVK'),
    (5, 'testname5', 'test5@test.email.com', '$2a$10$hEMqR7m/llJ2AaMl55t4Tu2XhXVxEVf41fFDd2SJkBrYhN.QFqPVK');