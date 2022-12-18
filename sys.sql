/*
SQLyog Ultimate v12.09 (64 bit)
MySQL - 5.7.25 : Database - nbainfosystem
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`nbainfosystem` /*!40100 DEFAULT CHARACTER SET utf8 */;

USE `nbainfosystem`;

/*Table structure for table `favor_players` */

DROP TABLE IF EXISTS `favor_players`;

CREATE TABLE `favor_players` (
  `uid` int(11) NOT NULL,
  `pid` int(11) NOT NULL,
  PRIMARY KEY (`uid`,`pid`),
  KEY `favor_players_ibfk_2` (`pid`),
  CONSTRAINT `favor_players_ibfk_1` FOREIGN KEY (`uid`) REFERENCES `user_infos` (`uid`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `favor_players_ibfk_2` FOREIGN KEY (`pid`) REFERENCES `player_infos` (`pid`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `favor_players` */

insert  into `favor_players`(`uid`,`pid`) values (8,9),(8,14),(8,20),(8,29),(8,35);

/*Table structure for table `favor_teams` */

DROP TABLE IF EXISTS `favor_teams`;

CREATE TABLE `favor_teams` (
  `uid` int(11) NOT NULL,
  `tid` smallint(6) NOT NULL,
  PRIMARY KEY (`uid`,`tid`),
  KEY `favor_teams_ibfk_2` (`tid`),
  CONSTRAINT `favor_teams_ibfk_1` FOREIGN KEY (`uid`) REFERENCES `user_infos` (`uid`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `favor_teams_ibfk_2` FOREIGN KEY (`tid`) REFERENCES `team_infos` (`tid`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/*Data for the table `favor_teams` */

insert  into `favor_teams`(`uid`,`tid`) values (3,2),(8,4),(8,5);

/*Table structure for table `player_infos` */

DROP TABLE IF EXISTS `player_infos`;

CREATE TABLE `player_infos` (
  `pid` int(11) NOT NULL AUTO_INCREMENT,
  `cn_name` char(16) NOT NULL,
  `en_name` char(32) NOT NULL,
  `height` float DEFAULT NULL,
  `weight` float DEFAULT NULL,
  `nation` char(32) NOT NULL,
  `img_path` char(128) DEFAULT NULL,
  `age` smallint(6) DEFAULT NULL,
  `tid` smallint(6) NOT NULL,
  PRIMARY KEY (`pid`),
  KEY `player_tid` (`tid`),
  CONSTRAINT `player_tid` FOREIGN KEY (`tid`) REFERENCES `team_infos` (`tid`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=45 DEFAULT CHARSET=utf8;

/*Data for the table `player_infos` */

insert  into `player_infos`(`pid`,`cn_name`,`en_name`,`height`,`weight`,`nation`,`img_path`,`age`,`tid`) values (9,'特雷-杨','Trae-Young',185,74.4,'United States','1670914196216158特雷杨.png',1998,1),(12,' 安东尼-戴维斯','Anthony-Davis',208,114.8,'United States','1670914381782950戴维斯.png',1993,5),(13,'奥斯汀-里维斯','Austin-Reaves',196,89.4,'United States','1670914429553367里维斯.png',1998,5),(14,'勒布朗-詹姆斯','LeBron-James',206,113.4,'United States','1670914481038961詹姆斯.png',1984,5),(15,' 拉塞尔-威斯布鲁克','Russell-Westbrook',190,90.7,'United States','1670914548102223威斯布鲁克.png',1988,5),(16,' 凯文-杜兰特','Kevin-Durant',208,108.9,'United States','1670914612603796杜兰特.png',1988,4),(17,'凯里-欧文','Kyrie-Irving',188,88.5,'Australia','1670914655335225欧文.png',1992,4),(18,'渡边雄太','Yuta-Watanabe',203,97.5,'Japan','1670914694771975渡边雄太.png',1994,4),(19,'凯斯勒-爱德华兹','Kessler-Edwards',201,92.1,'United States','1670914745766505爱德华兹.png',2000,4),(20,'本-西蒙斯','Ben-Simmons',208,108.9,'Australia','1670914859955434西蒙斯.png',1996,4),(21,'塞斯-库里','Seth-Curry',185,83.9,'United States','1670914919283190塞斯库里.png',1990,4),(22,'帕蒂-米尔斯','Patty-Mills',183,81.6,'Australia','1670914966324803米尔斯.png',1988,4),(23,'尼古拉斯-克拉克斯顿','Nic-Claxton',211,97.5,'United States','1670915056793929克拉克斯顿.png',1999,4),(24,'杰森-塔图姆','Jayson-Tatum',203,95.3,'United States','1670915373432947塔图姆.png',1998,6),(25,'杰伦-布朗','Jaylen-Brown',198,101.2,'United States','1670915406933874布朗.png',1996,6),(26,'格兰特-威廉姆斯','Grant-Williams',198,107,'United States','1670915436048976格兰特.png',1998,6),(27,'艾尔-霍福德','Al-Horford',206,108.9,'Dominican Republic','1670915481336584霍福德.png',1986,6),(28,' 马库斯-斯马特','Marcus-Smart',193,99.8,'United States','1670915515035192塔图姆.png',1994,6),(29,'斯蒂芬-库里','Stephen-Curry',188,83.9,'United States','1670915582035648斯蒂夫库里.png',1988,2),(30,'乔丹-普尔','Jordan-Poole',193,88,'United States','1670915620784410普尔.png',1999,2),(31,'詹姆斯-怀斯曼','James-Wiseman',213,108.9,'United States','1670915654046547怀斯曼.png',2001,2),(32,'乔纳森-库明加','Jonathan-Kuminga',201,102.1,'DRC','1670915690438797库明佳.png',2002,2),(33,'安德鲁-维金斯','Andrew-Wiggins',201,89.4,'Canada','1670915730192705威金斯.png',1995,2),(34,'凯文-鲁尼','Kevon-Looney',206,100.7,'United States','1670915767990910鲁尼.png',1996,2),(35,'尼古拉-约基奇','Nikola-Jokic',211,128.8,'Serbia','1670915845515192约基奇.png',1995,10),(36,'阿隆-戈登','Aaron-Gordon',203,106.6,'United States','1670915881057319戈登.png',1995,10),(37,'德马尔-德罗赞','DeMar-DeRozan',198,99.8,'United States','1670915923130315德罗赞.png',1989,7),(38,'扎克-拉文','Zach-LaVine',196,90.7,'United States','1670915960755415拉文.png',1995,7),(39,'朗佐-鲍尔','Lonzo-Ball',198,86.2,'United States','1670916001112200鲍尔.png',1997,7),(40,'尼古拉-武切维奇','Nikola-Vucevic',208,117.9,'Montenegro','1670916037496062武切维奇.png',1990,7),(41,'贾-莫兰特','Ja-Morant',188,78.9,'United States','1670916345175950莫兰特.png',1999,8),(42,'狄龙-布鲁克斯','Dillon-Brooks',198,102.1,'Canada','1670916378019853狄龙.png',1996,8),(43,'卢卡-东契奇','Luka-Doncic',201,104.3,'Slovenia','1670916423515325东契奇.png',1999,9),(44,'扬尼斯-安特托昆博','Giannis-Antetokounmpo',213,110.2,'Greece','1670916490911950阿德托昆博.png',1994,3);

/*Table structure for table `team_infos` */

DROP TABLE IF EXISTS `team_infos`;

CREATE TABLE `team_infos` (
  `tid` smallint(6) NOT NULL AUTO_INCREMENT,
  `tname` char(32) NOT NULL,
  `tcity` char(32) NOT NULL,
  `tzone` char(32) DEFAULT NULL,
  `arena` char(32) DEFAULT NULL,
  `logo` char(128) DEFAULT NULL,
  PRIMARY KEY (`tid`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8;

/*Data for the table `team_infos` */

insert  into `team_infos`(`tid`,`tname`,`tcity`,`tzone`,`arena`,`logo`) values (1,'老鹰','亚特兰大','东南分区','菲利普体育馆','ATL.png'),(2,'勇士','金州','太平洋分区','奥克兰体育馆','GSW.png'),(3,'雄鹿','密尔沃','中央分区','布拉德利中心','MIL.png'),(4,'篮网','布鲁克林','大西洋分区','巴克莱中心','BKN.png'),(5,'湖人','洛杉矶','太平洋分区','斯台普斯中心','LAL.png'),(6,'凯尔特人','波士顿','大西洋分区','TD北岸花园球馆','BOS.png'),(7,'公牛','芝加哥','中央分区','联合中心球馆','CHI.png'),(8,'灰熊','孟菲斯','西南分区','联邦快递球馆','MEM.png'),(9,'独行侠','达拉斯','西南分区','美国航线中心','DAL.png'),(10,'掘金','丹佛','西北分区','丹佛百事中心','DEN.png');

/*Table structure for table `user_infos` */

DROP TABLE IF EXISTS `user_infos`;

CREATE TABLE `user_infos` (
  `uid` int(11) NOT NULL AUTO_INCREMENT,
  `username` char(32) NOT NULL,
  `password` char(32) NOT NULL,
  `utype` smallint(6) NOT NULL,
  `headshot` char(128) DEFAULT NULL,
  PRIMARY KEY (`uid`),
  UNIQUE KEY `UNIQUENAME` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8;

/*Data for the table `user_infos` */

insert  into `user_infos`(`uid`,`username`,`password`,`utype`,`headshot`) values (3,'chen','111',1,'chen.jpg'),(8,'广工斯蒂芬','gdut',0,'1670918372072136stars.jpg');

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
