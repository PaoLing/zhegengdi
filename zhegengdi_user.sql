# 用户数据表
CREATE TABLE zgd_users_table (
	user_id INT NOT NULL AUTO_INCREMENT,
	user_name VARCHAR(50) NOT NULL,
	user_mobile VARCHAR(20) NOT NULL,
	password VARCHAR(50) NOT NULL,
	email VARCHAR(100),
	nickname VARCHAR(50) NOT NULL DEFAULT '折更低用户',
	level TINYINT NOT NULL DEFAULT 3,
	locked TINYINT(1) NOT NULL DEFAULT false,
	create_time TIMESTAMP NOT NULL DEFAULT current_timestamp(),
	comment VARCHAR(255),
	PRIMARY KEY(user_id)
);


# 商品数据表