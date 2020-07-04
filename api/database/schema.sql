DROP TABLE IF EXISTS accounts;
CREATE TABLE accounts (
  id int unsigned NOT NULL AUTO_INCREMENT,
  username varchar(32) NOT NULL DEFAULT '',
  email varchar(256) NOT NULL DEFAULT '',
  password varbinary(60) NOT NULL DEFAULT '0',
  is_activated tinyint(1) NOT NULL DEFAULT 0,
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  UNIQUE KEY U_username (username),
  UNIQUE KEY U_email (email)
);

DROP TABLE IF EXISTS account_activation_tokens;
CREATE TABLE account_activation_tokens (
  id int unsigned NOT NULL AUTO_INCREMENT,
  account_id int unsigned NOT NULL DEFAULT 0,
  activation_token char(64) NOT NULL DEFAULT '',
  is_used tinyint(1) NOT NULL DEFAULT 0,
  last_sent_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  UNIQUE KEY U_activation_token (activation_token)
);

DROP TABLE IF EXISTS password_reset_tokens;
CREATE TABLE password_reset_tokens (
  id int unsigned NOT NULL AUTO_INCREMENT,
  account_id int unsigned NOT NULL DEFAULT 0,
  reset_token char(64) NOT NULL DEFAULT '',
  is_used tinyint(1) NOT NULL DEFAULT 0,
  expires_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  UNIQUE KEY U_reset_token (reset_token)
);

DROP TABLE IF EXISTS categories;
CREATE TABLE categories (
  id int unsigned NOT NULL AUTO_INCREMENT,
  name varchar(32) NOT NULL DEFAULT '',
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  UNIQUE KEY U_name (name)
);

DROP TABLE IF EXISTS expenses;
CREATE TABLE expenses (
  id int unsigned NOT NULL AUTO_INCREMENT,
  account_id int unsigned NOT NULL DEFAULT '0',
  category_id int unsigned NOT NULL DEFAULT '0',
  description varchar(256) NOT NULL DEFAULT '',
  amount int unsigned NOT NULL DEFAULT '0',
  expense_date datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);
