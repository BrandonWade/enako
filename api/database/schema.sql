DROP TABLE IF EXISTS categories;
CREATE TABLE categories (
  id int(10) unsigned NOT NULL AUTO_INCREMENT,
  name varchar(32) NOT NULL DEFAULT '',
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  UNIQUE KEY U_name (name)
);

DROP TABLE IF EXISTS user_accounts;
CREATE TABLE user_accounts (
  id int(10) unsigned NOT NULL AUTO_INCREMENT,
  username varchar(32) NOT NULL DEFAULT '',
  email varchar(256) NOT NULL DEFAULT '',
  password varbinary(60) NOT NULL DEFAULT '0',
  PRIMARY KEY (id),
  UNIQUE KEY U_username (username),
  UNIQUE KEY U_email (email)
);

DROP TABLE IF EXISTS expenses;
CREATE TABLE expenses (
  id int(10) unsigned NOT NULL AUTO_INCREMENT,
  user_account_id int(10) unsigned NOT NULL DEFAULT '0',
  category_id int(10) unsigned NOT NULL DEFAULT '0',
  description varchar(256) NOT NULL DEFAULT '',
  amount int(10) unsigned NOT NULL DEFAULT '0',
  expense_date datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);
