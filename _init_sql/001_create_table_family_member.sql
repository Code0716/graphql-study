-- admin_users
CREATE TABLE IF NOT EXISTS `member`(
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `member_id` char(36) NOT NULL DEFAULT '' COMMENT 'ID（UUIDv4）',
  `name` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '名前',
  `postal_code` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '郵便番号',
  `address` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '住所',
  `mail_address` VARCHAR(255) NOT NULL DEFAULT '' COMMENT 'Email address',
  `phone_number` VARCHAR(255) DEFAULT NULL COMMENT '電話番号',
  `status` enum('FRIEND', 'FAMILY','BUSINESS', 'OTHER') NOT NULL DEFAULT 'OTHER' COMMENT 'status',
  `birthday` datetime DEFAULT NULL COMMENT 'birthday',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'created data',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'updated date',
  `deleted_at` datetime DEFAULT NULL COMMENT 'deleted date',
  PRIMARY KEY(`id`),
  INDEX member_id_index (`member_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
