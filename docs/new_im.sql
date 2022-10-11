/*
 Navicat Premium Data Transfer

 Source Server         : 81.71.162.89
 Source Server Type    : MySQL
 Source Server Version : 50738
 Source Host           : 81.71.162.89:3306
 Source Schema         : new_im

 Target Server Type    : MySQL
 Target Server Version : 50738
 File Encoding         : 65001

 Date: 12/10/2022 06:05:19
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for im_friend_records
-- ----------------------------
DROP TABLE IF EXISTS `im_friend_records`;
CREATE TABLE `im_friend_records`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `form_id` int(11) NOT NULL,
  `to_id` int(11) NOT NULL,
  `status` tinyint(1) NULL DEFAULT NULL COMMENT '0 等待通过 1 已通过 2 已拒绝',
  `created_at` timestamp NULL DEFAULT NULL,
  `information` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '请求信息',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 88 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for im_friends
-- ----------------------------
DROP TABLE IF EXISTS `im_friends`;
CREATE TABLE `im_friends`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `form_id` int(11) NULL DEFAULT NULL,
  `to_id` int(11) NULL DEFAULT NULL,
  `created_at` datetime NULL DEFAULT NULL,
  `note` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `top_time` datetime NULL DEFAULT NULL,
  `status` tinyint(1) NULL DEFAULT 0 COMMENT '0.未置顶 1.已置顶',
  `uid` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 65 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for im_group_messages
-- ----------------------------
DROP TABLE IF EXISTS `im_group_messages`;
CREATE TABLE `im_group_messages`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `message` json NOT NULL COMMENT '消息实体',
  `send_time` bigint(20) NULL DEFAULT NULL COMMENT '消息添加时间',
  `message_id` bigint(20) NULL DEFAULT NULL COMMENT '服务端消息id',
  `client_message_id` bigint(20) NULL DEFAULT NULL COMMENT '客户端消息id',
  `form_id` int(11) NULL DEFAULT NULL COMMENT '消息发送者id',
  `group_id` int(11) NULL DEFAULT NULL COMMENT '群聊id',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for im_group_offline_messages
-- ----------------------------
DROP TABLE IF EXISTS `im_group_offline_messages`;
CREATE TABLE `im_group_offline_messages`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `message` json NULL COMMENT '消息体',
  `send_time` int(11) NULL DEFAULT NULL COMMENT '消息接收时间',
  `status` tinyint(1) NULL DEFAULT NULL COMMENT '消息状态 0.未推送 1.已推送',
  `receive_id` int(11) NULL DEFAULT NULL COMMENT '接受id',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 128 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for im_group_user_messages
-- ----------------------------
DROP TABLE IF EXISTS `im_group_user_messages`;
CREATE TABLE `im_group_user_messages`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NULL DEFAULT NULL,
  `group_id` int(11) NULL DEFAULT NULL,
  `status` tinyint(1) NULL DEFAULT 0 COMMENT '0 未读 1 已读',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for im_group_users
-- ----------------------------
DROP TABLE IF EXISTS `im_group_users`;
CREATE TABLE `im_group_users`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `group_id` int(11) NULL DEFAULT NULL,
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 219 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for im_groups
-- ----------------------------
DROP TABLE IF EXISTS `im_groups`;
CREATE TABLE `im_groups`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '群聊id',
  `user_id` int(11) NULL DEFAULT NULL COMMENT '创建者',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '群聊名称',
  `created_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '添加时间',
  `info` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '群聊描述',
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '群聊头像',
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '密码',
  `is_pwd` tinyint(1) NULL DEFAULT 0 COMMENT '是否加密 0 否 1 是 ',
  `hot` int(10) NULL DEFAULT NULL COMMENT '热度',
  `theme` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '群聊主题',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `user_id`(`user_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 22 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for im_messages
-- ----------------------------
DROP TABLE IF EXISTS `im_messages`;
CREATE TABLE `im_messages`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `msg` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `form_id` int(11) NULL DEFAULT NULL,
  `to_id` int(11) NULL DEFAULT NULL,
  `is_read` tinyint(1) NULL DEFAULT NULL COMMENT '0 未读 1已读',
  `msg_type` tinyint(1) NULL DEFAULT 1,
  `status` tinyint(1) NULL DEFAULT NULL,
  `data` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 668 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for im_offline_messages
-- ----------------------------
DROP TABLE IF EXISTS `im_offline_messages`;
CREATE TABLE `im_offline_messages`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `message` json NULL COMMENT '消息体',
  `send_time` int(11) NULL DEFAULT NULL COMMENT '消息接收时间',
  `status` tinyint(1) NULL DEFAULT NULL COMMENT '消息状态 0.未推送 1.已推送',
  `receive_id` int(11) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 133 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for im_sessions
-- ----------------------------
DROP TABLE IF EXISTS `im_sessions`;
CREATE TABLE `im_sessions`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '会话表',
  `form_id` int(11) NOT NULL,
  `to_id` int(11) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `top_status` tinyint(1) NULL DEFAULT 0 COMMENT '0.否 1.是',
  `top_time` timestamp NULL DEFAULT NULL,
  `note` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '备注',
  `channel_type` tinyint(1) NULL DEFAULT 0 COMMENT '0.单聊 1.群聊',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '会话名称',
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '会话头像',
  `status` tinyint(1) NULL DEFAULT 0 COMMENT '会话状态 0.正常 1.禁用',
  `group_id` int(11) NULL DEFAULT NULL COMMENT '群ID',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 149 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for im_users
-- ----------------------------
DROP TABLE IF EXISTS `im_users`;
CREATE TABLE `im_users`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `email` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL,
  `created_at` datetime NULL DEFAULT NULL,
  `updated_at` datetime NULL DEFAULT NULL,
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '头像',
  `oauth_id` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '第三方id',
  `bound_oauth` tinyint(1) NULL DEFAULT 0 COMMENT '1\\github 2\\gitee',
  `oauth_type` tinyint(1) NULL DEFAULT NULL COMMENT '1.微博 2.github',
  `status` tinyint(1) NULL DEFAULT 0 COMMENT '0 离线 1 在线',
  `bio` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT '用户简介',
  `sex` tinyint(1) NULL DEFAULT 0 COMMENT '0 未知 1.男 2.女',
  `client_type` tinyint(1) NULL DEFAULT NULL COMMENT '1.web 2.pc 3.app',
  `age` int(3) NULL DEFAULT NULL,
  `last_login_time` timestamp NULL DEFAULT NULL COMMENT '最后登录时间',
  `uid` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT NULL COMMENT 'uid 关联',
  `user_json` json NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 48 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
