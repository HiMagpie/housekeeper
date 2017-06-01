# ************************************************************
# Sequel Pro SQL dump
# Version 4499
#
# http://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: 127.0.0.1 (MySQL 5.6.27)
# Database: himagpie
# Generation Time: 2016-09-14 07:26:16 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table hi_app
# ------------------------------------------------------------

DROP TABLE IF EXISTS `hi_app`;

CREATE TABLE `hi_app` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(40) NOT NULL COMMENT '应用名称',
  `app_id` varchar(32) NOT NULL DEFAULT '' COMMENT 'app唯一标示',
  `app_secret` varchar(32) NOT NULL DEFAULT '' COMMENT 'app密钥，用于手机客户端',
  `master_secret` varchar(32) NOT NULL DEFAULT '' COMMENT 'app密钥，用于服务端向HiMagpie发消息密钥',
  `enabled` tinyint(1) NOT NULL DEFAULT '1',
  `ctime` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `utime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `app_id` (`app_id`,`app_secret`,`master_secret`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;



# Dump of table hi_app_cid
# ------------------------------------------------------------

DROP TABLE IF EXISTS `hi_app_cid`;

CREATE TABLE `hi_app_cid` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `app_id` varchar(32) NOT NULL DEFAULT '' COMMENT '第三方应用唯一表示',
  `cid` varchar(32) NOT NULL DEFAULT '' COMMENT 'cid',
  `enabled` tinyint(1) NOT NULL DEFAULT '1',
  `ctime` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `utime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='第三方应用和 cid 的绑定关系';



# Dump of table hi_msg_cid_status
# ------------------------------------------------------------

DROP TABLE IF EXISTS `hi_msg_cid_status`;

CREATE TABLE `hi_msg_cid_status` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `cid` varchar(32) NOT NULL DEFAULT '' COMMENT '对应客户端id,',
  `msg_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '该消息是该cid对应的msg_id（作为分表依赖的字段）',
  `status` tinyint(4) unsigned NOT NULL DEFAULT '0' COMMENT '1：收到；2：放到目标服务器队列；3：服务器已推送，等待ACK；4：收到ACK，成功；5：客户端不在线，消息离线中；6：消息超时，失效；7：失败；',
  `enabled` tinyint(4) NOT NULL,
  `ctime` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `utime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `cid_msg_index` (`cid`,`msg_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# Dump of table hi_msg_info
# ------------------------------------------------------------

DROP TABLE IF EXISTS `hi_msg_info`;

CREATE TABLE `hi_msg_info` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `msg_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '给该消息分配的id（作为分表依赖的字段）',
  `cids` text NOT NULL COMMENT 'cid客户端id,多个用,分隔',
  `msg_ctime` int(11) NOT NULL COMMENT '收到该消息的时间',
  `ring` tinyint(4) unsigned NOT NULL DEFAULT '0' COMMENT '消息铃声',
  `vibrate` tinyint(4) unsigned NOT NULL DEFAULT '0',
  `cleanable` tinyint(4) unsigned NOT NULL DEFAULT '1',
  `trans` tinyint(4) unsigned NOT NULL DEFAULT '1' COMMENT '传输方式1：透传；2： 提醒',
  `begin` varchar(8) NOT NULL DEFAULT '00:00:00' COMMENT '接收的开始时间',
  `end` varchar(8) NOT NULL DEFAULT '00:00:00' COMMENT '接受的结束时间',
  `title` varchar(500) NOT NULL DEFAULT '' COMMENT '标题栏信息',
  `text` varchar(1000) NOT NULL DEFAULT '' COMMENT '正文',
  `logo` varchar(200) NOT NULL DEFAULT '' COMMENT '通知提醒的icon地址',
  `url` varchar(200) NOT NULL DEFAULT '' COMMENT '跳转网页的url',
  `status` tinyint(4) unsigned NOT NULL DEFAULT '1' COMMENT '1：收到；2：全部发送成功；3：部分失败；4：完全失败；',
  `enabled` tinyint(4) unsigned NOT NULL DEFAULT '1' COMMENT '0: 删除; 1: 正常',
  `ctime` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  `utime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;




/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
