CREATE TABLE field (
    `id` int NOT NULL AUTO_INCREMENT,
    `view_id` int NOT NULL,
    `title` varchar(64) NOT NULL, 
    `name` varchar(64) NOT NULL, 
    `desc` varchar(128),
    `indicator_type` enum('DIMENSION','MEASUREMENT','CUSTOMIZED'),
    `group` varchar(64),
    `order` smallint DEFAULT 0,
    `selected` tinyint(1) DEFAULT 0,
    `create_time` datetime DEFAULT NULL,
    `update_time` datetime DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY view_id (`view_id`)
)