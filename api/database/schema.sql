CREATE TABLE `expense_types` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `expense_type_name` varchar(32) NOT NULL DEFAULT '',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `U_expense_type_name` (`expense_type_name`)
);

CREATE TABLE `expense_categories` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `expense_category_name` varchar(32) NOT NULL DEFAULT '',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `U_expense_category_name` (`expense_category_name`)
);

CREATE TABLE `user_accounts` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_account_email` varchar(256) NOT NULL DEFAULT '',
  `user_account_password` varbinary(60) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  UNIQUE KEY `U_email` (`user_account_email`)
);

CREATE TABLE `user_expenses` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_account_id` int(10) unsigned NOT NULL DEFAULT '0',
  `expense_type` varchar(32) NOT NULL DEFAULT '',
  `expense_category` varchar(32) NOT NULL DEFAULT '',
  `expense_description` varchar(256) NOT NULL DEFAULT '',
  `expense_amount` int(10) unsigned NOT NULL DEFAULT '0',
  `expense_date` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
);
