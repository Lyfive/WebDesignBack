/*
 Navicat Premium Data Transfer

 Source Server         : MySQL
 Source Server Type    : MySQL
 Source Server Version : 80030 (8.0.30)
 Source Host           : localhost:3301
 Source Schema         : webdesign

 Target Server Type    : MySQL
 Target Server Version : 80030 (8.0.30)
 File Encoding         : 65001

 Date: 20/02/2023 15:56:40
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for classes
-- ----------------------------
DROP TABLE IF EXISTS `classes`;
CREATE TABLE `classes` (
  `s_id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `d_id` bigint unsigned NOT NULL,
  `session` char(2) NOT NULL,
  `scode` char(2) NOT NULL,
  `name` longtext NOT NULL,
  PRIMARY KEY (`s_id`),
  KEY `classes_departments_d_id_fk` (`d_id`),
  CONSTRAINT `classes_departments_d_id_fk` FOREIGN KEY (`d_id`) REFERENCES `departments` (`d_id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of classes
-- ----------------------------
BEGIN;
INSERT INTO `classes` (`s_id`, `d_id`, `session`, `scode`, `name`) VALUES (1, 1, '20', '07', '计算机科学与技术七班');
INSERT INTO `classes` (`s_id`, `d_id`, `session`, `scode`, `name`) VALUES (2, 2, '20', '01', '数据科学与大数据技术一班');
COMMIT;

-- ----------------------------
-- Table structure for courses
-- ----------------------------
DROP TABLE IF EXISTS `courses`;
CREATE TABLE `courses` (
  `c_id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `title` longtext NOT NULL,
  PRIMARY KEY (`c_id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of courses
-- ----------------------------
BEGIN;
INSERT INTO `courses` (`c_id`, `title`) VALUES (1, 'JavaWeb');
INSERT INTO `courses` (`c_id`, `title`) VALUES (2, '人工智能');
INSERT INTO `courses` (`c_id`, `title`) VALUES (3, '分布式系统');
COMMIT;

-- ----------------------------
-- Table structure for dcs
-- ----------------------------
DROP TABLE IF EXISTS `dcs`;
CREATE TABLE `dcs` (
  `d_id` bigint unsigned DEFAULT NULL,
  `c_id` bigint unsigned DEFAULT NULL,
  KEY `dcs_departments_d_id_fk` (`d_id`),
  KEY `dcs_courses_c_id_fk` (`c_id`),
  CONSTRAINT `dcs_courses_c_id_fk` FOREIGN KEY (`c_id`) REFERENCES `courses` (`c_id`) ON DELETE CASCADE,
  CONSTRAINT `dcs_departments_d_id_fk` FOREIGN KEY (`d_id`) REFERENCES `departments` (`d_id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of dcs
-- ----------------------------
BEGIN;
INSERT INTO `dcs` (`d_id`, `c_id`) VALUES (1, 1);
INSERT INTO `dcs` (`d_id`, `c_id`) VALUES (2, 2);
INSERT INTO `dcs` (`d_id`, `c_id`) VALUES (2, 3);
COMMIT;

-- ----------------------------
-- Table structure for departments
-- ----------------------------
DROP TABLE IF EXISTS `departments`;
CREATE TABLE `departments` (
  `d_id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `f_id` bigint unsigned NOT NULL,
  `dcode` char(2) NOT NULL,
  `name` longtext NOT NULL,
  PRIMARY KEY (`d_id`),
  UNIQUE KEY `dcode` (`dcode`),
  KEY `departments_faculties_f_id_fk` (`f_id`),
  CONSTRAINT `departments_faculties_f_id_fk` FOREIGN KEY (`f_id`) REFERENCES `faculties` (`f_id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of departments
-- ----------------------------
BEGIN;
INSERT INTO `departments` (`d_id`, `f_id`, `dcode`, `name`) VALUES (1, 1, '01', '计算机科学与技术');
INSERT INTO `departments` (`d_id`, `f_id`, `dcode`, `name`) VALUES (2, 1, '06', '数据科学与大数据技术');
COMMIT;

-- ----------------------------
-- Table structure for faculties
-- ----------------------------
DROP TABLE IF EXISTS `faculties`;
CREATE TABLE `faculties` (
  `f_id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `fcode` char(2) NOT NULL,
  `name` longtext NOT NULL,
  PRIMARY KEY (`f_id`),
  UNIQUE KEY `fcode` (`fcode`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of faculties
-- ----------------------------
BEGIN;
INSERT INTO `faculties` (`f_id`, `fcode`, `name`) VALUES (1, '05', '计算机科学与工程');
COMMIT;

-- ----------------------------
-- Table structure for grades
-- ----------------------------
DROP TABLE IF EXISTS `grades`;
CREATE TABLE `grades` (
  `g_id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `number` char(10) NOT NULL,
  `c_id` bigint unsigned NOT NULL,
  `mark` tinyint NOT NULL,
  PRIMARY KEY (`g_id`),
  KEY `grades_students_number_fk` (`number`),
  KEY `grades_courses_c_id_fk` (`c_id`),
  CONSTRAINT `grades_courses_c_id_fk` FOREIGN KEY (`c_id`) REFERENCES `courses` (`c_id`) ON DELETE CASCADE,
  CONSTRAINT `grades_students_number_fk` FOREIGN KEY (`number`) REFERENCES `students` (`number`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of grades
-- ----------------------------
BEGIN;
INSERT INTO `grades` (`g_id`, `number`, `c_id`, `mark`) VALUES (5, '2005060101', 3, 100);
INSERT INTO `grades` (`g_id`, `number`, `c_id`, `mark`) VALUES (6, '2005060101', 2, 100);
COMMIT;

-- ----------------------------
-- Table structure for messages
-- ----------------------------
DROP TABLE IF EXISTS `messages`;
CREATE TABLE `messages` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(12) NOT NULL,
  `password` varchar(128) NOT NULL,
  `level` tinyint NOT NULL,
  `name` varchar(20) NOT NULL,
  `head` longtext NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `username` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of messages
-- ----------------------------
BEGIN;
INSERT INTO `messages` (`id`, `username`, `password`, `level`, `name`, `head`) VALUES (1, 'admin', '21232f297a57a5a743894a0e4a801fc3', 1, 'lyfive', '47f9f39376e7a419bef260e6749cb121');
INSERT INTO `messages` (`id`, `username`, `password`, `level`, `name`, `head`) VALUES (2, 'test', '098f6bcd4621d373cade4e832627b4f6', 3, '李映飞', '30616305ef290e389d9019a0683a8046');
COMMIT;

-- ----------------------------
-- Table structure for sessions
-- ----------------------------
DROP TABLE IF EXISTS `sessions`;
CREATE TABLE `sessions` (
  `d_id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `session` varchar(191) NOT NULL,
  PRIMARY KEY (`d_id`,`session`),
  CONSTRAINT `sessions_departments_d_id_fk` FOREIGN KEY (`d_id`) REFERENCES `departments` (`d_id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of sessions
-- ----------------------------
BEGIN;
INSERT INTO `sessions` (`d_id`, `session`) VALUES (1, '20');
INSERT INTO `sessions` (`d_id`, `session`) VALUES (2, '20');
COMMIT;

-- ----------------------------
-- Table structure for students
-- ----------------------------
DROP TABLE IF EXISTS `students`;
CREATE TABLE `students` (
  `number` char(10) NOT NULL,
  `name` varchar(6) NOT NULL,
  `sex` longtext NOT NULL,
  `birthday` date NOT NULL,
  `s_id` bigint unsigned NOT NULL,
  PRIMARY KEY (`number`),
  KEY `students_classes_s_id_fk` (`s_id`),
  CONSTRAINT `students_classes_s_id_fk` FOREIGN KEY (`s_id`) REFERENCES `classes` (`s_id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of students
-- ----------------------------
BEGIN;
INSERT INTO `students` (`number`, `name`, `sex`, `birthday`, `s_id`) VALUES ('2005060101', 'a1', '男', '2023-02-02', 1);
COMMIT;

-- ----------------------------
-- Table structure for systems
-- ----------------------------
DROP TABLE IF EXISTS `systems`;
CREATE TABLE `systems` (
  `create_time` datetime(3) DEFAULT NULL,
  `version` longtext,
  `visits_number` bigint unsigned DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- ----------------------------
-- Records of systems
-- ----------------------------
BEGIN;
INSERT INTO `systems` (`create_time`, `version`, `visits_number`) VALUES ('2023-02-20 14:58:00.594', '1.0.0', 2);
INSERT INTO `systems` (`create_time`, `version`, `visits_number`) VALUES ('2023-02-20 15:29:10.312', '1.0.0', 0);
INSERT INTO `systems` (`create_time`, `version`, `visits_number`) VALUES ('2023-02-20 15:53:38.203', '1.0.0', 0);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
