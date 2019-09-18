CREATE TABLE IF NOT EXISTS users (
    id INT NOT NULL AUTO_INCREMENT,
    organziation_id INT NOT NULL,
    display_name VARCHAR(128) NOT NULL,

    PRIMARY KEY (id),
    INDEX (organziation_id),
    FOREIGN KEY (organziation_id) REFERENCES organizations(id) ON DELETE CASCADE

) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
