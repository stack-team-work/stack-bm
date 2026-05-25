/*
Navicat MySQL Data Transfer

Source Server         : 本地
Source Server Version : 50744
Source Host           : localhost:3306
Source Database       : stack_api

Target Server Type    : MYSQL
Target Server Version : 50744
File Encoding         : 65001

Date: 2026-05-25 22:49:24
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for game
-- ----------------------------
DROP TABLE IF EXISTS `game`;
CREATE TABLE `game` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '游戏名称，对内',
  `web_name` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '游戏名称，对外',
  `icon` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '游戏icon',
  `is_web_show` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否官网显示',
  `type_id` int(11) DEFAULT NULL COMMENT '游戏类型',
  `style_id` int(11) DEFAULT NULL COMMENT '游戏风格',
  `cp_id` int(11) NOT NULL COMMENT '研发id',
  `server_url` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '服务列表地址',
  `role_url` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '角色查询地址',
  `auth_name` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '授权方',
  `author` mediumtext COLLATE utf8mb4_unicode_ci COMMENT '著作人',
  `mark` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '备注',
  `status` int(11) NOT NULL DEFAULT '1' COMMENT '状态',
  `created_at` int(11) DEFAULT NULL COMMENT '创建时间',
  `updated_at` int(11) DEFAULT NULL COMMENT '更新时间',
  `is_deleted` tinyint(4) DEFAULT '0',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `m` (`mark`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=COMPACT COMMENT='父级游戏表';

-- ----------------------------
-- Table structure for game_app
-- ----------------------------
DROP TABLE IF EXISTS `game_app`;
CREATE TABLE `game_app` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `pid` int(11) NOT NULL COMMENT '父游戏id',
  `name` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '包名称对内',
  `package_name` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '应用包名',
  `app_name` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '应用名称',
  `os` int(11) NOT NULL DEFAULT '1' COMMENT '系统1：安卓，2ios:3:双端',
  `is_verify` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否开启实名认证，1是：0：否',
  `age` int(11) NOT NULL DEFAULT '18' COMMENT '可玩年龄',
  `is_open_charge` tinyint(4) NOT NULL DEFAULT '1' COMMENT '是否开启支付，1是：0：否',
  `is_open_register` tinyint(4) NOT NULL DEFAULT '1' COMMENT '是否开启登录，1是：0：否',
  `is_alert_email` tinyint(4) NOT NULL DEFAULT '1' COMMENT '是否弹出绑定邮箱：1:是；0：否',
  `is_alert_phone` tinyint(4) NOT NULL DEFAULT '1' COMMENT '是否弹出绑定手机：1:是；0：否',
  `is_alert_auth` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否弹出自定义授权：1:是；0：否',
  `is_open_float` tinyint(4) NOT NULL DEFAULT '1' COMMENT '是否开启悬浮窗：1:是；0：否',
  `is_deleted` tinyint(4) DEFAULT '0',
  `sdk_ver` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'sdk_ver版本',
  `app_ver` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'app_ver版本',
  `app_key` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'app_key',
  `app_secret` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'app_secret',
  `callback_url` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '回调地址',
  `api_domain` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'api域名',
  `pay_domain` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '支付域名',
  `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态',
  `cs_params` mediumtext COLLATE utf8mb4_unicode_ci COMMENT '客服参数',
  `pay_params` mediumtext COLLATE utf8mb4_unicode_ci COMMENT '支付参数配置',
  `h5_params` mediumtext COLLATE utf8mb4_unicode_ci COMMENT 'H5参数',
  `created_at` int(11) DEFAULT NULL,
  `updated_at` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=COMPACT COMMENT='子游戏表';

-- ----------------------------
-- Table structure for game_cp
-- ----------------------------
DROP TABLE IF EXISTS `game_cp`;
CREATE TABLE `game_cp` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '研发名称',
  `mark` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '备注',
  `phone` varchar(16) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '联系人',
  `addr` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '联系地址',
  `is_deleted` tinyint(4) DEFAULT '0',
  `status` tinyint(4) DEFAULT '1',
  `created_at` int(11) DEFAULT NULL,
  `updated_at` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=COMPACT;

-- ----------------------------
-- Table structure for game_tags
-- ----------------------------
DROP TABLE IF EXISTS `game_tags`;
CREATE TABLE `game_tags` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `type` tinyint(4) NOT NULL DEFAULT '1' COMMENT '1:风格，2：类型',
  `mark` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
  `status` int(11) NOT NULL DEFAULT '1',
  `is_deleted` tinyint(4) DEFAULT '0',
  `created_at` int(11) DEFAULT NULL,
  `updated_at` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=COMPACT;

-- ----------------------------
-- Table structure for game_variable
-- ----------------------------
DROP TABLE IF EXISTS `game_variable`;
CREATE TABLE `game_variable` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '名称',
  `key` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL,
  `value` mediumtext COLLATE utf8mb4_unicode_ci NOT NULL,
  `mark` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '备注',
  `status` tinyint(4) NOT NULL DEFAULT '1',
  `is_deleted` tinyint(4) NOT NULL DEFAULT '0',
  `updated_at` int(11) DEFAULT NULL,
  `created_at` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `key` (`key`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=COMPACT;

-- ----------------------------
-- Table structure for sys_logs
-- ----------------------------
DROP TABLE IF EXISTS `sys_logs`;
CREATE TABLE `sys_logs` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '日志表自增id',
  `pid` int(10) DEFAULT NULL COMMENT '父游戏id',
  `app_id` int(10) DEFAULT NULL COMMENT '子游戏id',
  `type` int(11) NOT NULL DEFAULT '1' COMMENT '日志类型：注册日志，登录日志等',
  `level` int(50) NOT NULL DEFAULT '1' COMMENT '日志等级',
  `ip` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '发生ip',
  `desc` mediumtext COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '日志描述',
  `create_time` int(11) DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  KEY `create_time` (`create_time`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=138 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=COMPACT COMMENT='日志表';

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `client_id` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '客户端和服务端自定义激活id',
  `user_id` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '玩家id',
  `pid` int(11) NOT NULL COMMENT '父游戏id',
  `app_id` int(11) NOT NULL COMMENT '子游戏id',
  `app_ver` varchar(16) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '应用包版本',
  `sdk_ver` varchar(16) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'sdk版本',
  `package_name` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '包名',
  `auth_type` int(11) NOT NULL COMMENT '授权方式',
  `account` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '注册账号',
  `nickname` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '昵称',
  `reg_from` tinyint(4) NOT NULL DEFAULT '1' COMMENT '注册来源1：SDK，2：官网',
  `verify_pi` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '国家实名系统pi',
  `is_phone` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '0' COMMENT '是否绑定手机',
  `is_email` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '0' COMMENT '是否绑定邮箱',
  `is_inner` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否内部玩家',
  `is_verify` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否已实名认证',
  `is_login` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否允许登录',
  `is_charge` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否允许充值',
  `sdk_level` int(11) DEFAULT '0' COMMENT 'sdk系统等级',
  `vip_level` int(11) DEFAULT '0' COMMENT 'sdk系统vip等级',
  `real_name` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '实名真实名字',
  `real_id` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '实名真实身份号码',
  `coin` int(11) DEFAULT '0' COMMENT '游戏币',
  `sex` tinyint(4) DEFAULT NULL COMMENT '性别',
  `age` int(11) DEFAULT NULL COMMENT '年龄',
  `ad_id` int(11) NOT NULL DEFAULT '0' COMMENT '分包标识/广告位id',
  `account_id` int(11) NOT NULL DEFAULT '0' COMMENT '渠道账号id',
  `media_id` int(11) NOT NULL DEFAULT '0' COMMENT '媒体渠道id',
  `media_sub_id` int(11) NOT NULL DEFAULT '0' COMMENT '媒体渠道子id',
  `cpid` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '0' COMMENT '渠道计划第一层级',
  `aid` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '0' COMMENT '渠道计划第二层级',
  `cid` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '0' COMMENT '渠道计划第三层级',
  `tid` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '0' COMMENT '渠道计划投放位置',
  `ip` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'ip',
  `mac` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'mac',
  `sys_brand` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '设备品牌',
  `sys_model` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '设备机型',
  `sys_name` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '设备系统名称',
  `sys_ver` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '设备版本',
  `screen_size` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '屏幕尺寸',
  `network` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '网络',
  `op` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'ip运营商',
  `country` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '国家',
  `province` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '省份',
  `city` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '城市',
  `idfa` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'idfa',
  `idfv` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'idfv',
  `odid` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'odid',
  `oaid` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'oaid',
  `udid` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'udid',
  `vaid` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'vaid',
  `aaid` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'aaid',
  `gaid` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'gaid',
  `caid` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'caid',
  `imei` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'imei',
  `ua` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'ua`',
  `android_id` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'android_id',
  `created_at` int(11) NOT NULL COMMENT '注册时间',
  `updated_at` int(11) DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `created_at` (`created_at`) USING BTREE,
  KEY `app_id` (`app_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=COMPACT COMMENT='用户表';

-- ----------------------------
-- Table structure for user_actives
-- ----------------------------
DROP TABLE IF EXISTS `user_actives`;
CREATE TABLE `user_actives` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `client_id` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '客户端和服务端自定义激活id',
  `pid` int(11) NOT NULL COMMENT '父游戏id',
  `app_id` int(11) NOT NULL COMMENT '子游戏id',
  `app_ver` varchar(16) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '应用包版本',
  `sdk_ver` varchar(16) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'sdk版本',
  `ad_id` int(11) NOT NULL DEFAULT '0' COMMENT '分包标识/广告位id',
  `account_id` int(11) NOT NULL DEFAULT '0' COMMENT '渠道账号id',
  `media_id` int(11) NOT NULL DEFAULT '0' COMMENT '媒体渠道id',
  `media_sub_id` int(11) NOT NULL DEFAULT '0' COMMENT '媒体渠道子id',
  `cpid` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '0' COMMENT '渠道计划第一层级',
  `aid` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '0' COMMENT '渠道计划第二层级',
  `cid` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '0' COMMENT '渠道计划第三层级',
  `tid` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '0' COMMENT '渠道计划投放位置',
  `ip` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'ip',
  `mac` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'mac',
  `sys_brand` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '设备品牌',
  `sys_model` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '设备机型',
  `sys_name` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '设备系统名称',
  `sys_ver` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '设备版本',
  `screen_size` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '屏幕尺寸',
  `network` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '网络',
  `op` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'ip运营商',
  `country` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '国家',
  `province` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '省份',
  `city` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '城市',
  `idfa` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'idfa',
  `idfv` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'idfv',
  `odid` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'odid',
  `oaid` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'oaid',
  `udid` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'udid',
  `vaid` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'vaid',
  `aaid` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'aaid',
  `gaid` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'gaid',
  `caid` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'caid',
  `imei` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'imei',
  `ua` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'ua`',
  `android_id` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'android_id',
  `created_at` int(11) NOT NULL COMMENT '流水时间',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `created_at` (`created_at`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=COMPACT COMMENT='用户激活表';

-- ----------------------------
-- Table structure for user_auth
-- ----------------------------
DROP TABLE IF EXISTS `user_auth`;
CREATE TABLE `user_auth` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '玩家唯一id',
  `auth_type` tinyint(4) NOT NULL COMMENT '注册类型1，邮箱，2手机，3游客',
  `account` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '注册账号',
  `password` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '密码',
  `salt` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '密码盐值',
  `created_at` int(11) NOT NULL COMMENT '注册时间',
  `updated_at` int(11) NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `account` (`account`) USING BTREE,
  KEY `user_id` (`user_id`) USING BTREE,
  KEY `type` (`auth_type`,`account`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=COMPACT COMMENT='用户授权表';

-- ----------------------------
-- Table structure for user_code
-- ----------------------------
DROP TABLE IF EXISTS `user_code`;
CREATE TABLE `user_code` (
  `id` bigint(16) NOT NULL AUTO_INCREMENT,
  `account` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '账号',
  `code` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '验证码',
  `expired_at` int(11) NOT NULL COMMENT '过期时间',
  `created_at` int(11) DEFAULT NULL,
  `updated_at` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `account` (`account`) USING BTREE,
  KEY `code` (`code`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=COMPACT COMMENT='手机短信码表';

-- ----------------------------
-- Table structure for user_logins
-- ----------------------------
DROP TABLE IF EXISTS `user_logins`;
CREATE TABLE `user_logins` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `client_id` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '客户端和服务端自定义激活id',
  `user_id` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '玩家id',
  `pid` int(11) NOT NULL COMMENT '父游戏id',
  `app_id` int(11) NOT NULL COMMENT '子游戏id',
  `app_ver` varchar(16) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '应用包版本',
  `sdk_ver` varchar(16) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'sdk版本',
  `package_name` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '包名',
  `is_reg` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否注册1：注册，0登录',
  `auth_type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '注册授权类型',
  `ad_id` int(11) NOT NULL DEFAULT '0' COMMENT '分包标识/广告位id',
  `account_id` int(11) NOT NULL DEFAULT '0' COMMENT '渠道账号id',
  `media_id` int(11) NOT NULL DEFAULT '0' COMMENT '媒体渠道id',
  `media_sub_id` int(11) NOT NULL DEFAULT '0' COMMENT '媒体渠道子id',
  `cpid` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '0' COMMENT '渠道计划第一层级',
  `aid` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '0' COMMENT '渠道计划第二层级',
  `cid` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '0' COMMENT '渠道计划第三层级',
  `tid` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '0' COMMENT '渠道计划投放位置',
  `ip` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'ip',
  `mac` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'mac',
  `sys_brand` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '设备品牌',
  `sys_model` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '设备机型',
  `sys_name` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '设备系统名称',
  `sys_ver` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '设备版本',
  `screen_size` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '屏幕尺寸',
  `network` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '网络',
  `op` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'ip运营商',
  `country` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '国家',
  `province` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '省份',
  `city` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '城市',
  `idfa` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'idfa',
  `idfv` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'idfv',
  `odid` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'odid',
  `oaid` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'oaid',
  `udid` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'udid',
  `vaid` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'vaid',
  `aaid` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'aaid',
  `gaid` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'gaid',
  `caid` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'caid',
  `imei` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'imei',
  `ua` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'ua`',
  `android_id` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'android_id',
  `created_at` int(11) NOT NULL COMMENT '流水时间',
  PRIMARY KEY (`id`),
  KEY `ct` (`created_at`) USING BTREE,
  KEY `user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=COMPACT COMMENT='用户登录注册日志表';

-- ----------------------------
-- Table structure for user_orders
-- ----------------------------
DROP TABLE IF EXISTS `user_orders`;
CREATE TABLE `user_orders` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `order_num` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'sdk订单号',
  `third_order_num` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '第三方订单号',
  `extend_order_num` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '扩展订单号',
  `cp_order_num` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '研发订单号',
  `client_id` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '客户端和服务端自定义激活id',
  `user_id` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '玩家id',
  `pid` int(11) NOT NULL COMMENT '父游戏id',
  `app_id` int(11) NOT NULL COMMENT '子游戏id',
  `app_ver` varchar(16) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '应用包版本',
  `sdk_ver` varchar(16) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'sdk版本',
  `package_name` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '包名',
  `ad_id` int(11) NOT NULL DEFAULT '0' COMMENT '分包标识/广告位id',
  `account_id` int(11) NOT NULL DEFAULT '0' COMMENT '渠道账号id',
  `media_id` int(11) NOT NULL DEFAULT '0' COMMENT '媒体渠道id',
  `media_sub_id` int(11) NOT NULL DEFAULT '0' COMMENT '媒体渠道子id',
  `cpid` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '0' COMMENT '渠道计划第一层级',
  `aid` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '0' COMMENT '渠道计划第二层级',
  `cid` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '0' COMMENT '渠道计划第三层级',
  `tid` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '0' COMMENT '渠道计划投放位置',
  `server_id` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '0' COMMENT '服务器标识',
  `server_name` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT '0' COMMENT '服务器名字',
  `role_id` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '0' COMMENT '角色id',
  `role_vip` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT '0' COMMENT '角色vip',
  `role_level` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT '0' COMMENT '角色等级',
  `role_name` varchar(50) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '0' COMMENT '角色名称',
  `coupon_id` int(11) NOT NULL DEFAULT '0' COMMENT '优惠券id',
  `gold` int(11) NOT NULL DEFAULT '0' COMMENT '充值钻石',
  `currency` varchar(16) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '币种',
  `discount` decimal(16,4) DEFAULT NULL COMMENT '折扣',
  `origin_total_fee` decimal(16,4) NOT NULL COMMENT '原始金额',
  `total_fee` decimal(16,4) NOT NULL COMMENT '实际支付金额',
  `product` varchar(32) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '购买产品',
  `product_id` varchar(32) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '购买产品id',
  `is_first` int(11) DEFAULT '0' COMMENT '是否首冲',
  `pay_status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '支付状态:1：未支付，2：已支付',
  `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '订单状态:1：未支付，2：未回掉，3：已发货，4：发货失败',
  `pay_domain` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '支付域名',
  `pay_number` int(11) NOT NULL COMMENT '支付标识(后台读取配置)',
  `pay_way` tinyint(4) NOT NULL DEFAULT '1' COMMENT '支付通道1sdk，2官网',
  `pay_type` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '支付方式:alipay,wechat',
  `cp_ext` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'cp扩展参数',
  `callback_at` int(11) DEFAULT NULL COMMENT '接受到第三方支付回调时间',
  `reg_at` int(11) DEFAULT '0' COMMENT '玩家注册时间',
  `pay_at` int(11) DEFAULT NULL COMMENT '订单支付时间',
  `sync_at` int(11) DEFAULT NULL COMMENT '回调给研发时间',
  `signature` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'sdk订单验证',
  `is_test` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否测试订单1是，0否',
  `sync_msg` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `is_send` tinyint(4) DEFAULT '0' COMMENT '是否已经生成订单发送给用户（微信小游戏,公众号支付）',
  `ip` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'ip',
  `mac` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'mac',
  `sys_brand` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '设备品牌',
  `sys_model` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '设备机型',
  `sys_name` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '设备系统名称',
  `sys_ver` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '设备版本',
  `screen_size` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '屏幕尺寸',
  `network` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '网络',
  `op` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'ip运营商',
  `country` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '国家',
  `province` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '省份',
  `city` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '城市',
  `idfa` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'idfa',
  `idfv` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'idfv',
  `odid` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'odid',
  `oaid` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'oaid',
  `udid` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'udid',
  `vaid` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'vaid',
  `aaid` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'aaid',
  `gaid` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'gaid',
  `caid` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'caid',
  `imei` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'imei',
  `ua` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'ua`',
  `android_id` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT 'android_id',
  `created_at` int(11) NOT NULL COMMENT '流水时间',
  `updated_at` int(11) DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `order_num` (`order_num`) USING BTREE,
  KEY `created_at` (`created_at`) USING BTREE,
  KEY `user_id` (`user_id`) USING BTREE,
  KEY `pay_status` (`pay_status`) USING BTREE,
  KEY `status` (`status`) USING BTREE,
  KEY `reg_at` (`reg_at`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=COMPACT COMMENT='用户订单表';
