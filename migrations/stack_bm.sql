/*
Navicat MySQL Data Transfer

Source Server         : 本地
Source Server Version : 50744
Source Host           : localhost:3306
Source Database       : stack_bm

Target Server Type    : MYSQL
Target Server Version : 50744
File Encoding         : 65001

Date: 2026-07-27 21:37:56
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
) ENGINE=InnoDB AUTO_INCREMENT=390 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=COMPACT COMMENT='用户表';

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
-- Table structure for sys_column
-- ----------------------------
DROP TABLE IF EXISTS `sys_column`;
CREATE TABLE `sys_column` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增id',
  `report_type` tinyint(4) NOT NULL DEFAULT '1' COMMENT '报表类型:1投放报表',
  `indicator_type` tinyint(4) NOT NULL DEFAULT '1' COMMENT '1：属性指标，2：媒体指标，3，bm指标，4，n日指标',
  `name` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '表格列名',
  `mark` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `field` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '字段名称',
  `default` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否默认选中',
  `status` tinyint(4) NOT NULL DEFAULT '1',
  `admin_id` int(11) NOT NULL,
  `created_at` int(11) DEFAULT '0' COMMENT '创建时间',
  `updated_at` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `report_type` (`report_type`,`indicator_type`,`field`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=COMPACT COMMENT='权限表';

-- ----------------------------
-- Table structure for sys_feishu_apps
-- ----------------------------
DROP TABLE IF EXISTS `sys_feishu_apps`;
CREATE TABLE `sys_feishu_apps` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `app_id` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '应用id',
  `app_secret` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '应用密钥',
  `app_name` varchar(100) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '应用名称',
  `admin_id` int(11) NOT NULL COMMENT '创建人',
  `mark` varchar(100) CHARACTER SET utf8mb4 NOT NULL COMMENT '应用标识',
  `status` int(11) NOT NULL DEFAULT '1' COMMENT '应用状态',
  `created_at` int(11) DEFAULT NULL COMMENT '创建时间',
  `updated_at` int(11) DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='飞书应用表';

-- ----------------------------
-- Table structure for sys_feishu_chats
-- ----------------------------
DROP TABLE IF EXISTS `sys_feishu_chats`;
CREATE TABLE `sys_feishu_chats` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `type` int(11) NOT NULL COMMENT '机器人类型：1：普通机器人，2：应用机器人',
  `chat_id` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '群聊天id',
  `default_at_list` text COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '默认艾特飞书用户，格式{sys_user_id:飞书用户id}',
  `at_list` text COLLATE utf8mb4_unicode_ci COMMENT '选择艾特飞书用户，格式{sys_user_id:飞书用户id}',
  `at_type` int(11) NOT NULL COMMENT '艾特方式：1：艾特所配置的全部，2：艾特所配置对应的数据负责人',
  `created_at` int(11) DEFAULT NULL COMMENT '创建时间',
  `updated_at` int(11) DEFAULT NULL COMMENT '更新时间',
  `status` int(11) NOT NULL DEFAULT '1' COMMENT '状态：1启用，0关闭',
  `feishu_app_id` int(11) DEFAULT NULL COMMENT '对应飞书应用app_id,关联表sys_feishu_apps的id',
  `call_action` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '读取配置key',
  `admin_id` int(11) NOT NULL COMMENT '创建人',
  `action_title` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '对话标题',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='飞书机器人';

-- ----------------------------
-- Table structure for sys_feishu_users
-- ----------------------------
DROP TABLE IF EXISTS `sys_feishu_users`;
CREATE TABLE `sys_feishu_users` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `admin_id` int(11) NOT NULL DEFAULT '0' COMMENT 'BM用户id',
  `feishu_user_id` varchar(100) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '飞书用户id',
  `created_at` int(11) DEFAULT NULL COMMENT '创建时间',
  `updated_at` int(11) DEFAULT NULL COMMENT '更新时间',
  `status` int(11) NOT NULL DEFAULT '1' COMMENT '状态：1正常，0：异常',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Table structure for sys_logs
-- ----------------------------
DROP TABLE IF EXISTS `sys_logs`;
CREATE TABLE `sys_logs` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '日志表自增id',
  `level` tinyint(4) DEFAULT '1' COMMENT '日志等级，1；info，2：error，3：debug',
  `path` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '请求地址',
  `username` varchar(32) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `ip` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `desc` mediumtext COLLATE utf8mb4_unicode_ci,
  `created_at` int(11) DEFAULT '0' COMMENT '创建时间',
  `updated_at` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `username` (`username`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=183 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=COMPACT COMMENT='日志表';

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
