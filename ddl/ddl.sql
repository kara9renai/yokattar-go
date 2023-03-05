CREATE TABLE `account` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `username` varchar(255) NOT NULL UNIQUE,
    `password_hash` varchar(255) NOT NULL,
    `display_name` varchar(255),
    `avatar` text,
    `header` text,
    `note` text,
    `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `followers_count` bigint(20) DEFAULT 0,
    `following_count` bigint(20) DEFAULT 0,
    PRIMARY KEY (`id`)
);

CREATE TABLE `status` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `account_id` bigint(20) NOT NULL,
    `content` text NOT NULL,
    `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    INDEX `idx_account_id` (`account_id`),
    CONSTRAINT `fk_status_account_id` FOREIGN KEY (`account_id`) REFERENCES `account` (`id`)
);

CREATE TABLE `relation` (
    `follower_id` bigint(20) NOT NULL,
    `followee_id` bigint(20) NOT NULL,
    CONSTRAINT `fk_relation_follower_id` FOREIGN KEY (`follower_id`) REFERENCES `account` (`id`),
    CONSTRAINT `fk_relation_followee_id` FOREIGN KEY (`followee_id`) REFERENCES `account` (`id`),
    PRIMARY KEY (`follower_id`, `followee_id`)
);

CREATE TABLE `attachment` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `type` text NOT NULL,
    `url` text NOT NULL,
    `description` varchar(420),
    PRIMARY KEY (`id`)
);

CREATE TABLE `favorite` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `account_id` bigint(20) NOT NULL,
    `status_id` bigint(20) NOT NULL,
    `create_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT `fk_like_account_id` FOREIGN KEY (`account_id`) REFERENCES `account` (`id`),
    CONSTRAINT `fk_like_status_id` FOREIGN KEY (`status_id`) REFERENCES `status` (`id`)
);
