/*
 Navicat Premium Data Transfer

 Source Server         : 本地库
 Source Server Type    : MySQL
 Source Server Version : 50650
 Source Host           : 127.0.0.1:3306
 Source Schema         : gohub

 Target Server Type    : MySQL
 Target Server Version : 50650
 File Encoding         : 65001

 Date: 06/07/2022 17:02:44
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for gohub_migration
-- ----------------------------
DROP TABLE IF EXISTS `gohub_migration`;
CREATE TABLE `gohub_migration` (
  `id` bigint(20) NOT NULL DEFAULT '0',
  `migration` varchar(255) DEFAULT NULL,
  `batch` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='数据库结构维护记录';

-- ----------------------------
-- Records of gohub_migration
-- ----------------------------
BEGIN;
INSERT INTO `gohub_migration` (`id`, `migration`, `batch`) VALUES (0, '2022_06_19_140933_add_users_table', 1);
COMMIT;

-- ----------------------------
-- Table structure for gohub_projects
-- ----------------------------
DROP TABLE IF EXISTS `gohub_projects`;
CREATE TABLE `gohub_projects` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(30) DEFAULT NULL COMMENT '项目名称',
  `status` tinyint(2) DEFAULT '1' COMMENT '状态 1启用 0禁用',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='项目记录';

-- ----------------------------
-- Records of gohub_projects
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for gohub_user
-- ----------------------------
DROP TABLE IF EXISTS `gohub_user`;
CREATE TABLE `gohub_user` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `email` varchar(255) DEFAULT NULL,
  `phone` varchar(20) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `status` tinyint(2) DEFAULT '1',
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_users_updated_at` (`updated_at`),
  KEY `idx_users_created_at` (`created_at`),
  KEY `idx_gohub_user_created_at` (`created_at`),
  KEY `idx_gohub_user_updated_at` (`updated_at`),
  KEY `idx_gohub_user_deleted_at` (`deleted_at`),
  KEY `idx_gohub_user_email` (`email`(191)),
  KEY `idx_gohub_user_phone` (`phone`),
  KEY `idx_gohub_user_name` (`name`(191))
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of gohub_user
-- ----------------------------
BEGIN;
INSERT INTO `gohub_user` (`id`, `name`, `email`, `phone`, `password`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES (5, 'ripperT', 'wangyifan@testing.com', '', '$2a$14$52vGo9NMnmN2wYl5QzF9.eX04IX63NaCzE7Tgdwnxc7Yzxln8VY9i', 1, '2022-06-12 13:15:52.000', '2022-06-19 14:57:37.451', NULL);
INSERT INTO `gohub_user` (`id`, `name`, `email`, `phone`, `password`, `status`, `created_at`, `updated_at`, `deleted_at`) VALUES (6, 'ripper', '', '00088888888', '$2a$14$kxFTekwysSYBZblMGjOULuZR8NbIJNyhLbOfkHsK9a2EZZZSb5BXK', 1, '2022-06-12 13:26:37.000', '2022-06-19 12:13:36.000', NULL);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
