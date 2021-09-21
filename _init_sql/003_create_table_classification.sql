-- admin_users
CREATE TABLE IF NOT EXISTS `gql_db`.`classification`(
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `class_id` char(36) NOT NULL DEFAULT '' COMMENT 'class ID（UUIDv4）',
  `class_name` char(36) NOT NULL COMMENT 'class_name',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'created data',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'updated date',
  `deleted_at` datetime DEFAULT NULL COMMENT 'deleted date',
  PRIMARY KEY(`id`),
  INDEX idx_class (`class_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
