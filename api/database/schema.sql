DROP TABLE IF EXISTS expense_types;
CREATE TABLE expense_types (
  id int(10) unsigned NOT NULL AUTO_INCREMENT,
  type_name varchar(32) NOT NULL DEFAULT '',
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  UNIQUE KEY U_type_name (type_name)
);

DROP TABLE IF EXISTS expense_categories;
CREATE TABLE expense_categories (
  id int(10) unsigned NOT NULL AUTO_INCREMENT,
  category_name varchar(32) NOT NULL DEFAULT '',
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  UNIQUE KEY U_category_name (category_name)
);

DROP TABLE IF EXISTS user_accounts;
CREATE TABLE user_accounts (
  id int(10) unsigned NOT NULL AUTO_INCREMENT,
  user_account_email varchar(256) NOT NULL DEFAULT '',
  user_account_password varbinary(60) NOT NULL DEFAULT '0',
  PRIMARY KEY (id),
  UNIQUE KEY U_email (user_account_email)
);

DROP TABLE IF EXISTS user_expenses;
CREATE TABLE user_expenses (
  id int(10) unsigned NOT NULL AUTO_INCREMENT,
  user_account_id int(10) unsigned NOT NULL DEFAULT '0',
  expense_type_id int(10) unsigned NOT NULL DEFAULT '0',
  expense_category_id int(10) unsigned NOT NULL DEFAULT '0',
  expense_description varchar(256) NOT NULL DEFAULT '',
  expense_amount int(10) unsigned NOT NULL DEFAULT '0',
  expense_date datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);
