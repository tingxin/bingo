CREATE TABLE IF NOT EXISTS resource (
	id INT NOT NULL AUTO_INCREMENT,
	name VARCHAR(64) NOT NULL,
	`desc` VARCHAR(128),
	creator VARCHAR(64) NOT NULL,
	editor VARCHAR(64) NOT NULL,
	create_time DATETIME DEFAULT NULL,
  	update_time DATETIME DEFAULT NULL,
  	visible TINYINT NOT NULL DEFAULT 1,
  	visible_time DATETIME DEFAULT NULL,
  	`order` INT NOT NULL DEFAULT 0,
  	PRIMARY KEY(id)
)