-- admin_users
CREATE TABLE IF NOT EXISTS `gql_db`.`testimony`(
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `testimony_id` char(36) NOT NULL DEFAULT '' COMMENT 'testimony id（UUIDv4）',
  `person_id` char(36) NOT NULL DEFAULT '' COMMENT 'person ID（UUIDv4）',
  `testimony` text NOT NULL COMMENT '証言内容',
  `status` enum('DRAFT', 'PUBLIC', 'OTHER') NOT NULL DEFAULT 'DRAFT' COMMENT 'status',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'created data',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'updated date',
  `deleted_at` datetime DEFAULT NULL COMMENT 'deleted date',
  PRIMARY KEY(`id`),
  INDEX idx_told_me (`person_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
