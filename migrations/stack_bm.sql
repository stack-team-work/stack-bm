/*
Navicat MySQL Data Transfer

Source Server         : 本地
Source Server Version : 50744
Source Host           : localhost:3306
Source Database       : stack_bm

Target Server Type    : MYSQL
Target Server Version : 50744
File Encoding         : 65001

Date: 2026-05-25 22:49:32
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for sys_admin
-- ----------------------------
DROP TABLE IF EXISTS `sys_admin`;
CREATE TABLE `sys_admin` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户自增id',
  `username` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '登录账号',
  `password` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '登录密码',
  `salt` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '掩码',
  `user_type` int(64) DEFAULT '1' COMMENT '账号类型，1：渠道账号；2：内部账号',
  `group_id` int(11) NOT NULL DEFAULT '0' COMMENT '用户组id',
  `name` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '昵称',
  `phone` varchar(16) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '电话',
  `department_id` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '部门',
  `updated_at` int(11) DEFAULT NULL COMMENT '更新时间',
  `created_at` int(11) DEFAULT '0' COMMENT '创建用户时间戳',
  `expiration` int(11) DEFAULT '0' COMMENT '用户过期时间戳',
  `login_num` int(11) NOT NULL DEFAULT '0' COMMENT '累计登陆次数',
  `last_login_time` int(11) DEFAULT '0' COMMENT '最后登录时间',
  `last_login_ip` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `game_app_permit` mediumtext COLLATE utf8mb4_unicode_ci COMMENT '子游戏权限',
  `game_permit` mediumtext COLLATE utf8mb4_unicode_ci COMMENT '父游戏权限',
  `status` tinyint(1) DEFAULT '1' COMMENT '用户状态',
  `is_deleted` tinyint(4) DEFAULT '0',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `username` (`username`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=387 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=COMPACT COMMENT='用户表';

-- ----------------------------
-- Table structure for sys_admin_group
-- ----------------------------
DROP TABLE IF EXISTS `sys_admin_group`;
CREATE TABLE `sys_admin_group` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户组自增id',
  `mark` varchar(11) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '标识',
  `name` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '管理员组名称',
  `description` varchar(256) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `menu_permit` mediumtext COLLATE utf8mb4_unicode_ci COMMENT '菜单权限',
  `column_permit` mediumtext COLLATE utf8mb4_unicode_ci COMMENT '报表指标权限',
  `status` tinyint(4) DEFAULT '1',
  `is_deleted` tinyint(4) DEFAULT '0',
  `created_at` int(11) DEFAULT '0' COMMENT '创建用户组时间戳',
  `updated_at` int(11) DEFAULT NULL COMMENT '更新操作时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `mark` (`mark`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=COMPACT COMMENT='用户组';

-- ----------------------------
-- Table structure for sys_logs
-- ----------------------------
DROP TABLE IF EXISTS `sys_logs`;
CREATE TABLE `sys_logs` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '日志表自增id',
  `level` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `path` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '请求地址',
  `username` varchar(32) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `ip` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `desc` mediumtext COLLATE utf8mb4_unicode_ci,
  `created_at` int(11) DEFAULT '0' COMMENT '创建时间',
  `updated_at` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `username` (`username`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=163 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=COMPACT COMMENT='日志表';

-- ----------------------------
-- Table structure for sys_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_menu`;
CREATE TABLE `sys_menu` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '菜单自增id',
  `type` tinyint(4) NOT NULL DEFAULT '1' COMMENT '菜单类型，1、路由菜单，2，按钮菜单，3，表格菜单',
  `author` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL,
  `name` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '菜单名字',
  `path` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL,
  `parent` int(10) NOT NULL DEFAULT '0' COMMENT '上级菜单id',
  `icon` varchar(128) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `sort` int(11) NOT NULL DEFAULT '0' COMMENT '排序',
  `is_deleted` tinyint(4) DEFAULT '0',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态',
  `created_at` int(11) DEFAULT '0' COMMENT '创建菜单时间戳',
  `updated_at` int(11) DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `href` (`path`)
) ENGINE=InnoDB AUTO_INCREMENT=641 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=COMPACT COMMENT='菜单';
