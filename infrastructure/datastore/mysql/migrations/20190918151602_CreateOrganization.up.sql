CREATE TABLE IF NOT EXISTS organizations (
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR (128) NOT NULL,
    PRIMARY KEY (id),
    UNIQUE KEY(name)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
