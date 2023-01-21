-- create "users" table
CREATE TABLE `users` (`id` char(36) NOT NULL, `user_email` varchar(255) NOT NULL, `user_name` varchar(255) NOT NULL, `pass_hash` varchar(255) NOT NULL, `created_at` timestamp NOT NULL, `updated_at` timestamp NOT NULL, PRIMARY KEY (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
