CREATE TABLE IF NOT EXISTS resource_tag_map(
	resource_id INT NOT NULL,
	tag_id INT NOT NULL,
	UNIQUE KEY id (resource_id, tag_id),
	KEY resource_id (resource_id),
	KEY tag_id (tag_id)
)