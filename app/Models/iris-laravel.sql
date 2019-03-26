# ************************************************************
# Sequel Pro SQL dump
# Version 4529
#
# http://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: 127.0.0.1 (MySQL 5.6.35)
# Database: iris-laravel
# Generation Time: 2019-03-26 14:15:04 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table log_admin_op
# ------------------------------------------------------------

DROP TABLE IF EXISTS `log_admin_op`;

CREATE TABLE `log_admin_op` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `uid` int(11) DEFAULT '0' COMMENT '用户编号',
  `url` varchar(255) DEFAULT '' COMMENT '请求ＵＲＬ',
  `request` text COMMENT '操作内容',
  `ip` varchar(20) DEFAULT NULL COMMENT 'ip',
  `ltime` int(11) DEFAULT '0' COMMENT '操作时间',
  PRIMARY KEY (`id`),
  KEY `ltime` (`ltime`),
  KEY `uid` (`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='管理员操作日志';



# Dump of table log_user_login
# ------------------------------------------------------------

DROP TABLE IF EXISTS `log_user_login`;

CREATE TABLE `log_user_login` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '编号',
  `uid` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '用户编号',
  `online` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '状态 0:上线;1:下线',
  `utype` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '登陆类型 0:fb;1:mobile;2web,3admin,4模拟登录',
  `ldate` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '时间',
  `ip` varchar(15) NOT NULL DEFAULT '' COMMENT 'ip',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 CHECKSUM=1 DELAY_KEY_WRITE=1 ROW_FORMAT=DYNAMIC COMMENT='网站用户登陆日志';



# Dump of table sys_admin_user
# ------------------------------------------------------------

DROP TABLE IF EXISTS `sys_admin_user`;

CREATE TABLE `sys_admin_user` (
  `user_id` int(11) NOT NULL COMMENT '用户',
  `role_id` int(11) NOT NULL COMMENT '角色',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态',
  `user_game` varchar(500) DEFAULT NULL COMMENT '游戏',
  PRIMARY KEY (`user_id`,`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='系统管理员';

LOCK TABLES `sys_admin_user` WRITE;
/*!40000 ALTER TABLE `sys_admin_user` DISABLE KEYS */;

INSERT INTO `sys_admin_user` (`user_id`, `role_id`, `status`, `user_game`)
VALUES
	(10001,1006,1,NULL);

/*!40000 ALTER TABLE `sys_admin_user` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table sys_admin_user_game
# ------------------------------------------------------------

DROP TABLE IF EXISTS `sys_admin_user_game`;

CREATE TABLE `sys_admin_user_game` (
  `user_id` int(11) NOT NULL,
  `game_id` int(11) NOT NULL,
  PRIMARY KEY (`user_id`,`game_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='管理员管理游戏';



# Dump of table sys_app
# ------------------------------------------------------------

DROP TABLE IF EXISTS `sys_app`;

CREATE TABLE `sys_app` (
  `app_id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '应用编号',
  `app_ename` varchar(50) DEFAULT NULL COMMENT '应用Code',
  `app_name` varchar(100) DEFAULT NULL COMMENT '应用名称',
  `app_img` varchar(200) DEFAULT NULL COMMENT '图片',
  `app_order` int(11) NOT NULL DEFAULT '0' COMMENT '排序',
  `app_tree_show` tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否在导航中显示',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`app_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='后台应用表';

LOCK TABLES `sys_app` WRITE;
/*!40000 ALTER TABLE `sys_app` DISABLE KEYS */;

INSERT INTO `sys_app` (`app_id`, `app_ename`, `app_name`, `app_img`, `app_order`, `app_tree_show`, `status`)
VALUES
	(1,'Setting','系统管理','setting16x16.gif',9000,1,1),
	(2,'Game','游戏设置','game16x16.gif',1000,1,1);

/*!40000 ALTER TABLE `sys_app` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table sys_app_function
# ------------------------------------------------------------

DROP TABLE IF EXISTS `sys_app_function`;

CREATE TABLE `sys_app_function` (
  `func_id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '功能编号',
  `app_id` int(11) NOT NULL COMMENT '应用编号',
  `func_name` varchar(50) DEFAULT NULL COMMENT '功能名称',
  `func_ename` varchar(100) DEFAULT NULL COMMENT '功能代码',
  `func_url` varchar(200) DEFAULT NULL COMMENT '地址',
  `func_img` varchar(200) DEFAULT NULL COMMENT '图标',
  `func_order` int(11) NOT NULL DEFAULT '0' COMMENT '排序',
  `status` int(3) NOT NULL DEFAULT '1' COMMENT '状态',
  PRIMARY KEY (`func_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='后台功能表';

LOCK TABLES `sys_app_function` WRITE;
/*!40000 ALTER TABLE `sys_app_function` DISABLE KEYS */;

INSERT INTO `sys_app_function` (`func_id`, `app_id`, `func_name`, `func_ename`, `func_url`, `func_img`, `func_order`, `status`)
VALUES
	(1,1,'管理员管理','AdminManage','admin/sys/admin_list','admin.gif',9010,1),
	(2,1,'用户角色管理','RoleManage','admin/sys/role_list','role.gif',9030,1),
	(3,1,'功能管理','FunctionManage','admin/sys/func_list','function.gif',9020,1),
	(4,2,'用户管理','UserManage','admin/sys/user_list','admin.gif',1,1);

/*!40000 ALTER TABLE `sys_app_function` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table sys_role
# ------------------------------------------------------------

DROP TABLE IF EXISTS `sys_role`;

CREATE TABLE `sys_role` (
  `role_id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '角色',
  `role_name` varchar(100) DEFAULT NULL COMMENT '角色名称',
  `role_ename` varchar(50) DEFAULT NULL COMMENT '角色代码',
  `role_funcnames` varchar(3000) DEFAULT NULL COMMENT '角色功能',
  `role_funcids` varchar(3000) DEFAULT NULL COMMENT '角色功能代码',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`role_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='后台角色表';

LOCK TABLES `sys_role` WRITE;
/*!40000 ALTER TABLE `sys_role` DISABLE KEYS */;

INSERT INTO `sys_role` (`role_id`, `role_name`, `role_ename`, `role_funcnames`, `role_funcids`, `status`)
VALUES
	(1000,'系统管理员','Administrator','用户管理;','view-add-edit-delete-4;',1),
	(1006,'超级管理员','SuperAdmin','用户管理;管理员管理;功能管理;用户角色管理;','view-add-edit-delete-4;view-add-edit-delete-1;view-add-edit-delete-3;view-add-edit-delete-2;',1);

/*!40000 ALTER TABLE `sys_role` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table sys_role_function
# ------------------------------------------------------------

DROP TABLE IF EXISTS `sys_role_function`;

CREATE TABLE `sys_role_function` (
  `role_id` int(11) NOT NULL COMMENT '角色编号',
  `func_id` int(11) NOT NULL COMMENT '功能编号',
  `func_op` varchar(100) DEFAULT NULL COMMENT '功能操作',
  PRIMARY KEY (`role_id`,`func_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='后台角色功能表';

LOCK TABLES `sys_role_function` WRITE;
/*!40000 ALTER TABLE `sys_role_function` DISABLE KEYS */;

INSERT INTO `sys_role_function` (`role_id`, `func_id`, `func_op`)
VALUES
	(1000,4,'view-add-edit-delete-4'),
	(1006,1,'view-add-edit-delete-1'),
	(1006,2,'view-add-edit-delete-2'),
	(1006,3,'view-add-edit-delete-3'),
	(1006,4,'view-add-edit-delete-4');

/*!40000 ALTER TABLE `sys_role_function` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table user_online
# ------------------------------------------------------------

DROP TABLE IF EXISTS `user_online`;

CREATE TABLE `user_online` (
  `user_id` int(11) NOT NULL COMMENT '在线用户编号',
  `user_url` varchar(200) DEFAULT NULL COMMENT '当前ＵＲＬ',
  `url_ip` varchar(45) DEFAULT NULL,
  `last_time` int(11) NOT NULL DEFAULT '0' COMMENT '最后一次操作时间',
  PRIMARY KEY (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='网站用户在线';



# Dump of table yly_member
# ------------------------------------------------------------

DROP TABLE IF EXISTS `yly_member`;

CREATE TABLE `yly_member` (
  `uid` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '用户编号',
  `pid` char(100) NOT NULL DEFAULT '' COMMENT '平台id',
  `username` varchar(64) NOT NULL DEFAULT '' COMMENT '用户名',
  `password` char(32) NOT NULL DEFAULT '' COMMENT '密码',
  `tel` varchar(20) NOT NULL DEFAULT '' COMMENT '手机',
  `email` char(32) NOT NULL DEFAULT '' COMMENT '邮箱',
  `birthday` varchar(50) NOT NULL DEFAULT '' COMMENT '生日',
  `reg_ip` char(15) NOT NULL DEFAULT '' COMMENT '注册ip',
  `reg_date` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '注册时间',
  `gender` enum('f','m') NOT NULL DEFAULT 'm' COMMENT '性别',
  `utype` char(10) NOT NULL DEFAULT '' COMMENT '用户类型',
  `nickname` char(20) NOT NULL DEFAULT '' COMMENT '昵称',
  `group_id` tinyint(3) unsigned NOT NULL DEFAULT '5' COMMENT '管理级别',
  `locale` varchar(20) NOT NULL COMMENT '语言',
  `avatar` varchar(255) NOT NULL COMMENT '头像',
  `upuid` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '上层uid，是哪个用户推荐过来的',
  `ad` char(100) NOT NULL DEFAULT '' COMMENT '用户来源，哪个广告',
  `login_ip` char(15) NOT NULL DEFAULT '' COMMENT '登录ip',
  `login_times` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '登录次数',
  `login_date` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '最后一次登陆时间',
  `remark` varchar(255) NOT NULL DEFAULT '' COMMENT '封号备注',
  `sign` varchar(255) NOT NULL DEFAULT '' COMMENT '签名',
  `regcity` varchar(45) NOT NULL DEFAULT '' COMMENT '城市',
  `regarea` varchar(100) NOT NULL DEFAULT '' COMMENT '区域',
  `new_user` tinyint(1) NOT NULL DEFAULT '1' COMMENT '新用户',
  `ps` varchar(100) DEFAULT NULL COMMENT 'pass',
  `user_salt` char(6) NOT NULL DEFAULT '' COMMENT 'salt',
  `appid` tinyint(4) NOT NULL DEFAULT '0' COMMENT 'appid',
  PRIMARY KEY (`uid`),
  UNIQUE KEY `username` (`username`) USING BTREE,
  KEY `logindate` (`login_date`),
  KEY `regdate` (`reg_date`),
  KEY `ad` (`ad`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 CHECKSUM=1 DELAY_KEY_WRITE=1 ROW_FORMAT=DYNAMIC COMMENT='用户表';

LOCK TABLES `yly_member` WRITE;
/*!40000 ALTER TABLE `yly_member` DISABLE KEYS */;

INSERT INTO `yly_member` (`uid`, `pid`, `username`, `password`, `tel`, `email`, `birthday`, `reg_ip`, `reg_date`, `gender`, `utype`, `nickname`, `group_id`, `locale`, `avatar`, `upuid`, `ad`, `login_ip`, `login_times`, `login_date`, `remark`, `sign`, `regcity`, `regarea`, `new_user`, `ps`, `user_salt`, `appid`)
VALUES
	(10001,'0','admin','a5010410a83cbc8498df1b9adfab5a1c','','','','180.166.124.2',1427109932,'m','web','admin',5,'zh_TW','',0,'','127.0.0.1',59,1553566069,'test_10000012','','','',0,'','7l0agm',0);

/*!40000 ALTER TABLE `yly_member` ENABLE KEYS */;
UNLOCK TABLES;



/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
