/*
Navicat MySQL Data Transfer

Source Server         : 本地
Source Server Version : 50744
Source Host           : localhost:3306
Source Database       : stack_mkt

Target Server Type    : MYSQL
Target Server Version : 50744
File Encoding         : 65001

Date: 2026-08-21 14:09:51
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for media
-- ----------------------------
DROP TABLE IF EXISTS `media`;
CREATE TABLE `media` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `mark` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `status` tinyint(4) NOT NULL DEFAULT '1',
  `is_deleted` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` int(11) DEFAULT NULL,
  `updated_at` int(11) DEFAULT NULL,
  `admin_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `mark` (`mark`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=COMPACT;

-- ----------------------------
-- Records of media
-- ----------------------------
INSERT INTO `media` VALUES ('1', '头条', 'tt', '1', '0', '1779791396', '1779791396', null);
INSERT INTO `media` VALUES ('2', '腾讯', 'tc', '1', '0', null, null, null);
INSERT INTO `media` VALUES ('3', 'B站', 'bili', '1', '0', null, null, null);
INSERT INTO `media` VALUES ('4', '快手', 'ks', '1', '0', null, null, null);

-- ----------------------------
-- Table structure for media_agent
-- ----------------------------
DROP TABLE IF EXISTS `media_agent`;
CREATE TABLE `media_agent` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `mark` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `subject_id` int(11) NOT NULL COMMENT '主体id，对应media_subject.id',
  `name` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `status` tinyint(4) NOT NULL DEFAULT '1',
  `is_deleted` tinyint(4) DEFAULT NULL,
  `updated_at` int(11) DEFAULT NULL,
  `created_at` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=COMPACT COMMENT='代理';

-- ----------------------------
-- Records of media_agent
-- ----------------------------

-- ----------------------------
-- Table structure for media_application
-- ----------------------------
DROP TABLE IF EXISTS `media_application`;
CREATE TABLE `media_application` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `media_id` int(11) NOT NULL COMMENT '媒体id，对应media.id',
  `name` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '管家名称',
  `app_id` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '应用id',
  `app_secret` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '应用密钥',
  `status` tinyint(4) NOT NULL COMMENT '状态',
  `remark` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '备注',
  `extra` text COLLATE utf8mb4_unicode_ci COMMENT '扩展参数',
  `admin_id` int(11) NOT NULL COMMENT '对应admin.id',
  `is_deleted` tinyint(4) DEFAULT '0',
  `updated_at` int(11) DEFAULT NULL,
  `created_at` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=COMPACT COMMENT='mkt应用表';

-- ----------------------------
-- Records of media_application
-- ----------------------------

-- ----------------------------
-- Table structure for media_manager
-- ----------------------------
DROP TABLE IF EXISTS `media_manager`;
CREATE TABLE `media_manager` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `media_id` int(11) NOT NULL COMMENT '媒体id，对应media.id',
  `name` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '应用名称',
  `application_id` int(11) NOT NULL COMMENT '应用id，对应media_application.id',
  `account` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '管家账号',
  `account_id` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '管家id',
  `account_num` int(11) NOT NULL DEFAULT '0' COMMENT '绑定账户数',
  `auth_status` tinyint(4) NOT NULL DEFAULT '0',
  `status` tinyint(4) NOT NULL DEFAULT '1',
  `is_deleted` tinyint(4) NOT NULL DEFAULT '0',
  `remark` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `extra` text COLLATE utf8mb4_unicode_ci,
  `admin_id` int(11) NOT NULL,
  `updated_at` int(11) DEFAULT NULL,
  `created_at` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=COMPACT COMMENT='mkt管家表';

-- ----------------------------
-- Records of media_manager
-- ----------------------------

-- ----------------------------
-- Table structure for media_sub
-- ----------------------------
DROP TABLE IF EXISTS `media_sub`;
CREATE TABLE `media_sub` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `media_id` int(11) NOT NULL,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `mark` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `status` tinyint(4) NOT NULL DEFAULT '1',
  `is_deleted` tinyint(4) DEFAULT '0',
  `admin_id` int(11) DEFAULT NULL,
  `created_at` int(11) DEFAULT NULL,
  `updated_at` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=COMPACT;

-- ----------------------------
-- Records of media_sub
-- ----------------------------

-- ----------------------------
-- Table structure for media_subject
-- ----------------------------
DROP TABLE IF EXISTS `media_subject`;
CREATE TABLE `media_subject` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `mark` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `name` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL,
  `status` tinyint(4) NOT NULL DEFAULT '1',
  `is_deleted` tinyint(4) NOT NULL DEFAULT '0',
  `updated_at` int(11) DEFAULT NULL,
  `created_at` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=COMPACT COMMENT='主体';

-- ----------------------------
-- Records of media_subject
-- ----------------------------
