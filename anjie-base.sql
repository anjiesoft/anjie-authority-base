/*
 Navicat Premium Data Transfer

 Source Server         : 本地 - docker8.0
 Source Server Type    : MySQL
 Source Server Version : 80004
 Source Host           : localhost:3306
 Source Schema         : admin

 Target Server Type    : MySQL
 Target Server Version : 80004
 File Encoding         : 65001

 Date: 28/08/2025 14:34:02
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for anjie_action_log
-- ----------------------------
DROP TABLE IF EXISTS `anjie_action_log`;
CREATE TABLE `anjie_action_log`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '内容',
  `admin_id` int(11) NOT NULL COMMENT '操作人',
  `admin_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '真实姓名',
  `type` smallint(3) NOT NULL COMMENT '类型，以常量配置文件为准，目前，1为添加，2为编辑，3为删除，4为强退，5为下载',
  `module_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '操作模块名，如管理员',
  `ip` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'IP',
  `browser_info` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '浏览器详细信息',
  `browser_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '浏览器名称',
  `browser_version` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '版本',
  `project_id` int(11) NULL DEFAULT 0 COMMENT '项目ID',
  `reason` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '原因',
  `create_time` datetime NULL DEFAULT NULL COMMENT '添加时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 81 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '操作日志表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of anjie_action_log
-- ----------------------------

-- ----------------------------
-- Table structure for anjie_admin
-- ----------------------------
DROP TABLE IF EXISTS `anjie_admin`;
CREATE TABLE `anjie_admin`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `username` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '账号',
  `avatar` varchar(255) CHARACTER SET sjis COLLATE sjis_japanese_ci NULL DEFAULT '' COMMENT '头像',
  `name` varchar(30) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '' COMMENT '真实姓名',
  `phone` char(20) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '电话',
  `email` varchar(255) CHARACTER SET sjis COLLATE sjis_japanese_ci NULL DEFAULT '' COMMENT '邮箱',
  `password` char(62) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '密码',
  `department_id` int(11) NULL DEFAULT 0 COMMENT '部门ID',
  `post_id` int(11) NOT NULL DEFAULT 0 COMMENT '岗位ID',
  `salt` char(5) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '与密码混合搭配加密',
  `roles` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT '' COMMENT '角色',
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '状态，1为正常，2为禁用',
  `last_time` datetime NULL DEFAULT NULL COMMENT '最后一次登录时间',
  `fail_number` tinyint(2) NULL DEFAULT 0 COMMENT '连续登录失败次数',
  `reason` varchar(255) CHARACTER SET sjis COLLATE sjis_japanese_ci NULL DEFAULT '' COMMENT '理由',
  `admin_id` int(11) NOT NULL DEFAULT 0 COMMENT '操作人',
  `admin_name` varchar(100) CHARACTER SET sjis COLLATE sjis_japanese_ci NULL DEFAULT '' COMMENT '操作人姓名',
  `create_time` datetime NULL DEFAULT NULL COMMENT '添加时间',
  `update_time` datetime NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `username`(`username`) USING BTREE COMMENT '账号与公司绑定唯一',
  INDEX `status`(`status`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = sjis COLLATE = sjis_japanese_ci COMMENT = '公司内部人员表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of anjie_admin
-- ----------------------------
INSERT INTO `anjie_admin` VALUES (1, 'admin', '1111', 'admin', '13888888888', 'bigapegao@163.com', '$2a$10$SY81FUVIUIutuUW9OoSGG.nhbq.YRIWlN2aQN8szqRCzvgXNjtqTm', 0, 1, 'DR753', '2,1,7', 1, NULL, 0, '', 1, 'admin', '2025-08-09 15:28:23', '2025-08-27 10:20:35');
INSERT INTO `anjie_admin` VALUES (2, 'test', '', 'test', '13999998888', 'test@test.test', '$2a$10$hYWKAyRiLXHsFW4v/kadt.5nzy6amsyuQtF4EJUY2FxrpeLwOH4A2', 0, 0, 'zRxsi', '5', 1, NULL, 0, '', 1, 'admin', '2025-08-09 18:17:26', '2025-08-12 09:09:04');

-- ----------------------------
-- Table structure for anjie_authority
-- ----------------------------
DROP TABLE IF EXISTS `anjie_authority`;
CREATE TABLE `anjie_authority`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '菜单',
  `parent_id` int(11) NOT NULL DEFAULT 0 COMMENT '父级菜单',
  `parent_ids` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '多级父级菜单',
  `level` tinyint(2) NOT NULL DEFAULT 0 COMMENT '级别',
  `type` tinyint(2) NOT NULL COMMENT '类型，1目录，2页面，3按钮',
  `project_id` int(11) NOT NULL COMMENT '项目ID',
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '页面请求地址',
  `api` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '接口地址',
  `view_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '页面文件路径',
  `identification` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '权限标识',
  `icon` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '图标',
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '状态，1为正常，2为禁用',
  `reason` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '理由',
  `sort` smallint(3) NULL DEFAULT 0 COMMENT '排序',
  `is_show` tinyint(1) NOT NULL DEFAULT 1 COMMENT '菜单中是否显示，1为显示，2为不显示',
  `remarks` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '备注',
  `admin_id` int(11) NULL DEFAULT NULL COMMENT '操作人ID',
  `admin_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '操作人员姓名',
  `create_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime NULL DEFAULT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `project_id`(`project_id`, `identification`) USING BTREE COMMENT '项目与权限唯一',
  INDEX `parent_id`(`project_id`) USING BTREE,
  INDEX `status`(`status`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 76 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '权限表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of anjie_authority
-- ----------------------------
INSERT INTO `anjie_authority` VALUES (1, '权限管理', 0, '0', 1, 1, 1, '/auth', '#', '#', 'auth:index', 'perm', 1, '', 0, 1, '', 1, NULL, '2025-07-31 06:21:02', NULL);
INSERT INTO `anjie_authority` VALUES (3, '权限列表', 1, '0,1', 2, 2, 1, '/authority/list', 'admin/authority/items', 'authority/authority/index', 'auth:list', 'category', 1, '', 0, 1, '', 1, NULL, '2025-07-31 06:21:02', NULL);
INSERT INTO `anjie_authority` VALUES (4, '权限下拉', 3, '0,1,3', 3, 3, 1, '/', 'admin/authority/name_items', '#', 'auth:titles', '', 1, '', 0, 2, '下拉的数据', 1, NULL, '2025-07-31 06:21:02', NULL);
INSERT INTO `anjie_authority` VALUES (14, '权限编辑', 3, '0,1,3', 3, 3, 1, '/', 'admin/authority/edit', '#', 'auth:edit', '', 1, '', 0, 2, '', 1, NULL, '2025-07-31 06:21:02', NULL);
INSERT INTO `anjie_authority` VALUES (19, '权限详情', 3, '0,1,3', 3, 3, 1, '/', 'admin/authority/info', '#', 'auth:detail', '', 1, '', 0, 2, '', 1, NULL, '2025-07-31 06:21:02', NULL);
INSERT INTO `anjie_authority` VALUES (30, '成员管理', 0, '0', 1, 1, 1, '/admin', '#', '#', 'admin:index', 'Customer management', 1, '', 0, 1, '', 1, NULL, '2025-07-31 06:21:02', NULL);
INSERT INTO `anjie_authority` VALUES (31, '权限状态', 3, '0,1,3', 3, 3, 1, '#', 'admin/authority/status', '#', 'auth:status', '', 1, '', 0, 2, '', 1, 'admin', '2025-07-31 06:21:02', '2025-08-19 14:33:26');
INSERT INTO `anjie_authority` VALUES (32, '列表数据', 3, '0,1,3', 3, 3, 1, '#', 'admin/authority/listdata', '#', 'auth:list:list', '', 1, '', 0, 1, '', 1, NULL, '2025-07-31 06:21:02', NULL);
INSERT INTO `anjie_authority` VALUES (33, '成员列表', 30, '0,30', 2, 2, 1, '/admin/list', 'admin/admin/items', 'admin/index', 'admin:list', 'category', 1, '', 0, 1, '', 1, NULL, '2025-07-31 06:21:02', NULL);
INSERT INTO `anjie_authority` VALUES (34, '成员编辑', 33, '0,30,33', 3, 3, 1, '#', 'admin/admin/edit', '#', 'admin:edit', '', 1, '', 0, 1, '', 1, 'admin', '2025-07-31 06:21:02', '2025-08-19 14:33:43');
INSERT INTO `anjie_authority` VALUES (35, '角色管理', 0, '0', 1, 1, 1, '/role', '#', '#', 'role:index', 'role', 1, '', 0, 1, '', 1, NULL, '2025-07-31 06:21:02', NULL);
INSERT INTO `anjie_authority` VALUES (37, '项目管理', 0, '0', 1, 1, 1, '/project', '#', '#', 'project:index', 'client', 1, '', 0, 1, '', 1, NULL, '2025-07-31 06:21:02', NULL);
INSERT INTO `anjie_authority` VALUES (38, '项目列表', 37, '0,37', 2, 2, 1, '/project/list', 'admin/project/items', 'company/project/index', 'project:list', 'category', 1, '', 0, 1, '', 1, NULL, '2025-07-31 06:21:02', NULL);
INSERT INTO `anjie_authority` VALUES (39, '公共功能', 0, '0', 1, 1, 1, '/public', '#', '#', 'public:index', '', 1, '', 0, 2, '', 1, NULL, '2025-07-31 06:21:02', NULL);
INSERT INTO `anjie_authority` VALUES (40, '功能列表', 39, '0,39', 2, 2, 1, '/', 'public/list', 'public/list', 'public:list', '', 1, '', 0, 2, '', 1, NULL, '2025-07-31 06:21:02', NULL);
INSERT INTO `anjie_authority` VALUES (41, '项目下拉', 40, '0,39,40', 3, 3, 1, '#', 'admin/project/name_items', '#', 'project:sortList', '', 1, '', 0, 2, '', 1, NULL, '2025-07-31 06:21:02', NULL);
INSERT INTO `anjie_authority` VALUES (42, '角色列表', 35, '0,35', 2, 2, 1, '/role/list', 'admin/role/items', 'authority/role/index', 'role:list', 'category', 1, '', 0, 1, '', 1, NULL, '2025-07-31 06:21:02', NULL);
INSERT INTO `anjie_authority` VALUES (43, '角色状态', 42, '0,35,42', 3, 3, 1, '#', 'admin/role/status', '#', 'role:status', '', 1, '', 0, 1, '', 1, NULL, '2025-07-31 06:21:02', NULL);
INSERT INTO `anjie_authority` VALUES (44, '角色编辑', 42, '0,35,42', 3, 3, 1, '#', 'admin/role/edit', '#', 'role:edit', '', 1, '', 0, 1, '', 1, NULL, '2025-07-31 06:21:02', NULL);
INSERT INTO `anjie_authority` VALUES (45, '角色详情', 42, '0,35,42', 3, 3, 1, '#', 'admin/role/info', '#', 'role:detail', '', 1, '', 0, 1, '', 1, NULL, '2025-07-31 06:21:02', NULL);
INSERT INTO `anjie_authority` VALUES (46, '权限下拉', 42, '0,35,42', 3, 3, 1, '#', 'admin/role/name_items', '#', 'role:auths', '', 1, '', 0, 1, '', 1, 'admin', '2025-07-31 06:21:02', '2025-08-19 13:56:59');
INSERT INTO `anjie_authority` VALUES (47, '第二个项目的', 0, '0', 1, 1, 2, '/index', '#', '#', 'auth:111', '', 1, '', 0, 1, '', 1, NULL, '2025-07-31 06:21:02', NULL);
INSERT INTO `anjie_authority` VALUES (51, '项目编辑', 38, '0,37,38', 3, 3, 1, '#', 'admin/project/edit', '#', 'project:edit', '', 1, '', 0, 1, '', 1, NULL, '2025-07-31 06:21:02', NULL);
INSERT INTO `anjie_authority` VALUES (52, '项目状态', 38, '0,37,38', 3, 3, 1, '#', 'admin/project/status', '#', 'project:status', '', 1, '', 0, 1, '', 1, NULL, '2025-07-31 06:21:02', NULL);
INSERT INTO `anjie_authority` VALUES (54, '项目详情', 38, '0,37,38', 3, 3, 1, '#', 'admin/project/info', '#', 'project:detail', '', 1, '', 0, 1, '', 1, NULL, '2025-07-31 06:21:02', NULL);
INSERT INTO `anjie_authority` VALUES (60, '管理员下拉', 33, '0,30,33', 3, 3, 1, '#', 'admin/admin/sortList', '#', 'admin:sortList', '', 1, '', 0, 1, '', 1, NULL, '2025-07-31 06:21:02', NULL);
INSERT INTO `anjie_authority` VALUES (61, '项目成员', 38, '0,37,38', 3, 3, 1, '#', 'admin/project/member', '#', 'project:member', '', 1, '', 0, 1, '', 1, NULL, '2025-07-31 06:21:02', NULL);
INSERT INTO `anjie_authority` VALUES (62, '项目成员编辑', 38, '0,37,38', 3, 3, 1, '#', 'admin/project/memberedit', '#', 'project:memberedit', '', 1, '', 0, 1, '', 1, NULL, '2025-07-31 06:21:02', NULL);
INSERT INTO `anjie_authority` VALUES (63, '成员状态', 33, '0,30,33', 3, 3, 1, '#', 'admin/admin/status', '#', 'admin:status', '', 1, '', 0, 1, '', 1, NULL, '2025-07-31 06:21:02', NULL);
INSERT INTO `anjie_authority` VALUES (64, '成员密码', 33, '0,30,33', 3, 3, 1, '#', 'admin/admin/password', '#', 'admin:password', '', 1, '', 0, 1, '', 1, 'admin', '2025-07-31 06:21:02', '2025-08-11 08:45:46');
INSERT INTO `anjie_authority` VALUES (65, '成员详情', 33, '0,30,33', 3, 3, 1, '#', 'admin/admin/info', '#', 'admin:detail', '', 1, '', 0, 1, '', 1, NULL, '2025-07-31 06:21:02', NULL);
INSERT INTO `anjie_authority` VALUES (66, '成员角色', 33, '0,30,33', 3, 3, 1, '#', 'admin/admin/role', '#', 'admin:role', '', 1, '', 0, 1, '', 1, NULL, '2025-07-31 06:21:02', NULL);
INSERT INTO `anjie_authority` VALUES (67, '成员公司', 33, '0,30,33', 3, 3, 1, '#', 'admin/admin/company', '#', 'admin:company', '', 1, '', 0, 1, '', 1, NULL, '2025-07-31 06:21:02', NULL);
INSERT INTO `anjie_authority` VALUES (68, '官方网站', 0, '0', 1, 4, 1, '/https://www.anjiesoft.com', '', '#', 'baidu:index', 'link', 1, '', 2, 1, '', 1, 'admin', '2025-08-11 18:28:15', '2025-08-11 18:28:15');
INSERT INTO `anjie_authority` VALUES (73, '日志管理', 0, '0', 1, 1, 1, '/log', '#', '#', 'log:index', 'order', 1, '', 0, 1, '', 1, 'admin', '2025-08-15 10:28:21', '2025-08-15 10:28:21');
INSERT INTO `anjie_authority` VALUES (74, '登录日志', 73, '0,73', 2, 2, 1, '/loginlog', '/admin/log/login', 'log/login', 'log:login', '', 1, '', 0, 1, '', 1, 'admin', '2025-08-15 10:35:21', '2025-08-15 10:59:32');
INSERT INTO `anjie_authority` VALUES (75, '操作日志', 73, '0,73', 2, 2, 1, '/action', '/admin/log/action', 'log/action', 'log:action', '', 1, '', 0, 1, '', 1, 'admin', '2025-08-15 10:36:04', '2025-08-15 10:59:35');

-- ----------------------------
-- Table structure for anjie_login_log
-- ----------------------------
DROP TABLE IF EXISTS `anjie_login_log`;
CREATE TABLE `anjie_login_log`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `admin_id` int(11) NOT NULL DEFAULT 0 COMMENT '登录人ID',
  `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '用户名',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '真实姓名',
  `browser_info` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '浏览器详细信息',
  `browser_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '浏览器名称',
  `browser_version` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '浏览器版本',
  `ip` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'IP',
  `status` tinyint(1) NULL DEFAULT 1 COMMENT '状态，1为成功，2为失败',
  `create_time` datetime NULL DEFAULT NULL COMMENT '添加时间',
  `reason` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '失败原因',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 586 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '登录日志表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of anjie_login_log
-- ----------------------------

-- ----------------------------
-- Table structure for anjie_project
-- ----------------------------
DROP TABLE IF EXISTS `anjie_project`;
CREATE TABLE `anjie_project`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '项目描述',
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '状态，1为正常，2为禁用',
  `remarks` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '描述',
  `logo` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT 'logo',
  `reason` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '理由',
  `admin_id` int(11) NULL DEFAULT 0 COMMENT '操作人员ID',
  `admin_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '操作人员姓名',
  `create_time` datetime NULL DEFAULT NULL,
  `update_time` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '项目表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of anjie_project
-- ----------------------------
INSERT INTO `anjie_project` VALUES (1, '权限系统', 1, '默认', '', '', 1, 'admin', '2025-07-31 00:00:00', '2025-08-19 14:31:27');
INSERT INTO `anjie_project` VALUES (2, '测试项目', 1, '默认', '', '', 1, 'admin', '2025-07-31 00:00:00', '2025-08-19 08:41:22');

-- ----------------------------
-- Table structure for anjie_role
-- ----------------------------
DROP TABLE IF EXISTS `anjie_role`;
CREATE TABLE `anjie_role`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '名称',
  `project_id` int(11) NOT NULL COMMENT '项目ID',
  `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '状态，1为正常，2为禁用',
  `rules` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL COMMENT '权限内容',
  `reason` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '' COMMENT '理由',
  `admin_id` int(11) NULL DEFAULT NULL,
  `admin_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT NULL COMMENT '操作人员姓名',
  `create_time` datetime NULL DEFAULT NULL,
  `update_time` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `project_id`(`project_id`, `name`) USING BTREE,
  INDEX `company_id`(`project_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci COMMENT = '角色表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of anjie_role
-- ----------------------------
INSERT INTO `anjie_role` VALUES (1, '权限管理', 1, 1, '4,1,3,14,19,31,32,36', '', 1, 'admin', NULL, '2025-08-26 21:59:15');
INSERT INTO `anjie_role` VALUES (2, '超级管理员', 1, 1, '4,1,3,14,19,31,32,34,30,33,41,39,40,43,35,42,44,45,46,51,37,38,52,54,60,61,62,63,64,65,66,67,68', '', 1, 'admin', '2025-08-09 15:28:04', '2025-08-11 18:28:28');
INSERT INTO `anjie_role` VALUES (3, '项目管理', 1, 1, '51,37,38,52,53,54,61,62', '', 1, 'admin', '2025-08-09 16:42:52', '2025-08-09 18:10:37');
INSERT INTO `anjie_role` VALUES (4, '角色管理', 1, 1, '43,35,42,44,45,46,50', '', 1, 'admin', '2025-08-09 16:58:03', '2025-08-09 17:00:23');
INSERT INTO `anjie_role` VALUES (5, '成员管理', 1, 1, '34,30,33,60,63,64,65,66,67', '', 1, 'admin', '2025-08-09 18:11:32', '2025-08-09 18:11:32');
INSERT INTO `anjie_role` VALUES (7, '日志管理', 1, 1, '74,73,75', '', 1, 'admin', '2025-08-19 14:18:55', '2025-08-19 14:18:55');

SET FOREIGN_KEY_CHECKS = 1;
