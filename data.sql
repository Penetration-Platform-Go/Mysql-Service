create table if not exists `users` (
  id INT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
  username char(255) NOT NULL UNIQUE,
  nickname char(255),
  email char(255) NOT NULL,
  password char(255) NOT NULL,
  photo char(255)
) AUTO_INCREMENT = 1;