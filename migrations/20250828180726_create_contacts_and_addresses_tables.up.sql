CREATE TABLE `contacts` (
    `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
    `user_id` int(11) unsigned NOT NULL,
    `first_name` varchar(100) NOT NULL,
    `last_name` varchar(100) DEFAULT NULL,
    `email` varchar(200) DEFAULT NULL,
    `phone` varchar(20) DEFAULT NULL,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `fk_contacts_user` (`user_id`),
    CONSTRAINT `fk_contacts_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;

CREATE TABLE `addresses` (
    `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
    `contact_id` int(11) unsigned NOT NULL,
    `street` varchar(255) DEFAULT NULL,
    `city` varchar(100) DEFAULT NULL,
    `province` varchar(100) DEFAULT NULL,
    `country` varchar(100) NOT NULL,
    `postal_code` varchar(10) DEFAULT NULL,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `fk_addresses_contact` (`contact_id`),
    CONSTRAINT `fk_addresses_contact` FOREIGN KEY (`contact_id`) REFERENCES `contacts` (`id`) ON DELETE CASCADE
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;