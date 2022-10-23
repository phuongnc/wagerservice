SET FOREIGN_KEY_CHECKS = 0;
-- ----------------------------
-- Table structure for wagers
-- ----------------------------
CREATE TABLE IF NOT EXISTS `wagers` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `total_wager_value` bigint unsigned DEFAULT NULL,
  `odds` bigint unsigned DEFAULT NULL,
  `selling_percentage` tinyint unsigned DEFAULT NULL,
  `selling_price` float DEFAULT NULL,
  `current_selling_price` float DEFAULT NULL,
  `percentage_sold` float DEFAULT NULL,
  `amount_sold` bigint unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_wagers_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Table structure for wager_transactions
-- ----------------------------
CREATE TABLE IF NOT EXISTS `wager_transactions` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `wager_id` bigint unsigned DEFAULT NULL,
  `buying_price` float DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_wager_transactions_deleted_at` (`deleted_at`),
  KEY `fk_wagers_wager_transactions` (`wager_id`),
  CONSTRAINT `fk_wagers_wager_transactions` FOREIGN KEY (`wager_id`) REFERENCES `wagers` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;

-- ----------------------------
-- Records of wagers
-- ----------------------------
DELETE FROM `wagers`;
INSERT INTO `wagers` 
VALUES 
(1, '2022-10-22 11:30:41.685', '2022-10-22 11:30:41.685', NULL, 1500, 10, 80, 1200.55, 1200.55, NULL, NULL),
(2, '2022-10-22 11:31:42.435', '2022-10-22 11:31:42.435', NULL, 2500, 20, 90, 2300.7, 2300.7, NULL, NULL),
(3, '2022-10-22 11:33:07.287', '2022-10-22 11:33:07.287', NULL, 1250, 15, 60, 800.05, 800.05, NULL, NULL),
(4, '2022-10-22 11:33:55.788', '2022-10-22 11:33:55.788', NULL, 1150, 4, 70, 810.3, 810.3, NULL, NULL);

SET FOREIGN_KEY_CHECKS=1;

