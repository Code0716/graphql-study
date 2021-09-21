-- person
CREATE TABLE IF NOT EXISTS `gql_db`.`person`(
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
  `person_id` char(36) NOT NULL DEFAULT '' COMMENT 'ID（UUIDv4）',
  `name` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '名前',
  `postal_code` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '郵便番号',
  `address` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '住所',
  `mail_address` VARCHAR(255) NOT NULL DEFAULT '' COMMENT 'Email address',
  `phone_number` VARCHAR(255) DEFAULT NULL COMMENT '電話番号',
  `class_name` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '関係性区分',
  `birthday` datetime DEFAULT NULL COMMENT 'birthday',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'created data',
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'updated date',
  `deleted_at` datetime DEFAULT NULL COMMENT 'deleted date',
  PRIMARY KEY(`id`),
  INDEX person_id_index (`person_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;
