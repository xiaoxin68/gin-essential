/*
Navicat MySQL Data Transfer

Source Server         : zhang
Source Server Version : 50717
Source Host           : localhost:3306
Source Database       : ginessential

Target Server Type    : MYSQL
Target Server Version : 50717
File Encoding         : 65001

Date: 2020-06-19 15:38:11
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for categories
-- ----------------------------
DROP TABLE IF EXISTS `categories`;
CREATE TABLE `categories` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of categories
-- ----------------------------
INSERT INTO `categories` VALUES ('3', '李四', '2020-05-15 10:28:28', '2020-05-20 10:28:31');
INSERT INTO `categories` VALUES ('4', '啦啦', '2020-05-23 02:22:30', '2020-05-23 02:22:30');
INSERT INTO `categories` VALUES ('5', '哈哈', '2020-05-23 10:26:08', '2020-05-23 10:26:08');
INSERT INTO `categories` VALUES ('6', '返回实例的空间', '2020-05-23 02:28:54', '2020-05-23 02:28:54');
INSERT INTO `categories` VALUES ('7', 'hjadfweh', null, null);
INSERT INTO `categories` VALUES ('8', '234253', '2020-05-23 10:51:39', '2020-05-23 10:51:39');
INSERT INTO `categories` VALUES ('9', 'wqrwer', '2020-05-23 10:56:25', '2020-05-23 10:56:25');

-- ----------------------------
-- Table structure for posts
-- ----------------------------
DROP TABLE IF EXISTS `posts`;
CREATE TABLE `posts` (
  `id` char(36) NOT NULL,
  `user_id` int(10) unsigned NOT NULL,
  `category_id` int(10) unsigned NOT NULL,
  `title` varchar(50) NOT NULL,
  `head_img` varchar(255) DEFAULT NULL,
  `content` text NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of posts
-- ----------------------------
INSERT INTO `posts` VALUES ('f6e83e62-c1fb-4539-8f0e-283ea2ff0477', '9', '3', '啦啦啦', '哈哈哈哈', '小安安', '2020-05-23 15:08:55', '2020-05-23 15:08:55');

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `password` varchar(255) NOT NULL DEFAULT '' COMMENT '密码',
  `telephone` varchar(11) NOT NULL DEFAULT '' COMMENT '密码',
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `name` varchar(255) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_users_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES ('9', '$2a$10$Ji6YBbvMttUQDV2vl27U5elwICyV0CyfXyyGyvhibtUAE4ARMPxee', '12345678901', '2020-05-21 12:44:22', '2020-05-21 12:44:22', null, 'q1mi');
INSERT INTO `users` VALUES ('12', '$2a$10$4iXOEIQISxwBAQ6f.2h8.e6HuewnTSGdEZHd49cUSzzuOitsQQ/Wq', '12345678921', '2020-05-22 09:39:53', '2020-05-22 09:39:53', null, 'q1mi');
INSERT INTO `users` VALUES ('13', '$2a$10$8ArZuvg6tSSAhpSB0f6cW.MO5rVAyojgJYSW4fX18ip2583Wm8pVu', '12345678903', '2020-05-23 02:09:36', '2020-05-23 02:09:36', null, 'q1mi');
