CREATE DATABASE IF NOT EXISTS `wager_test` CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;

-- Grant admin
CREATE USER IF NOT EXISTS 'admin' IDENTIFIED BY 'WagerService123';
GRANT ALL ON `wager_test`.* TO 'admin'@'%';
