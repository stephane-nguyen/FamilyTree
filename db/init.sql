CREATE USER 'user1'@'127.0.0.1' IDENTIFIED BY 'password';
GRANT SELECT, INSERT, UPDATE, DELETE ON family.* TO 'user1'@'127.0.0.1';
FLUSH PRIVILEGES;

-- IDENTIFIED WITH caching_sha2_password BY 'password';
-- grant all privileges on *.* to user1@localhost identified by 'password' with grant option;
CREATE TABLE Person (
    id          INT AUTO_INCREMENT PRIMARY KEY,
    firstname   VARCHAR(30) NOT NULL,
    middlename  VARCHAR(30),
    lastname    VARCHAR(30) NOT NULL,
    birthdate   DATE,
    gender      ENUM('Male', 'Female'),
    city        VARCHAR(30),
    country     VARCHAR(30),
    photo       LONGBLOB
);


CREATE TABLE Relationship (
    id          INT AUTO_INCREMENT PRIMARY KEY,
    relationship_type ENUM('Parent', 'Child', 'Sibling', 'Spouse', 'Cousin', 'Uncle', 'Aunt'),
    person_id   INT,
    related_person_id INT,
    FOREIGN KEY (person_id) REFERENCES Person(id),
    FOREIGN KEY (related_person_id) REFERENCES Person(id)
);