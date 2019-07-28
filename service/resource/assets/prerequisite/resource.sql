CREATE TABLE `resource` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(64) NOT NULL,
  `desc` varchar(128) DEFAULT NULL,
  `creator` varchar(64) NOT NULL,
  `editor` varchar(64) NOT NULL,
  `create_time` datetime DEFAULT NULL,
  `update_time` datetime DEFAULT NULL,
  `visible` tinyint(4) NOT NULL DEFAULT '1',
  `visible_time` datetime DEFAULT NULL,
  `order` int(11) NOT NULL DEFAULT '0',
  `kind` enum('VIEW','DASHBOARD','CHART') DEFAULT NULL,
  PRIMARY KEY (`id`)

)