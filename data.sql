create table if not exists `users` (
  id INT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
  email char(255) NOT NULL,
  password char(255) NOT NULL,
  username char(255) NOT NULL,
  nickname char(255),
  photo char(255)
) AUTO_INCREMENT = 1;