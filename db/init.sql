-- CREATE USER 'user1'@'127.0.0.1' IDENTIFIED BY 'password';
-- GRANT SELECT, INSERT, UPDATE, DELETE ON family.* TO 'user1'@'127.0.0.1';
-- FLUSH PRIVILEGES;

-- IDENTIFIED WITH caching_sha2_password BY 'password';
-- grant all privileges on *.* to user1@localhost identified by 'password' with grant option;
CREATE TABLE User (
    id          INT AUTO_INCREMENT PRIMARY KEY,
    username    VARCHAR(30) NOT NULL,
    email       VARCHAR(55) NOT NULL,
    password    VARCHAR(255) NOT NULL
);

CREATE TABLE Member (
    id          INT AUTO_INCREMENT PRIMARY KEY,
    user_id     INT,
    firstname   VARCHAR(30) NOT NULL,
    middlename  VARCHAR(30),
    lastname    VARCHAR(30) NOT NULL,
    birthdate   DATE,
    gender      ENUM('Male', 'Female', 'Other'),
    city        VARCHAR(30),
    country     VARCHAR(30),
    description VARCHAR(255),
    photo       LONGBLOB,
    FOREIGN KEY (user_id) REFERENCES User(id)
);

CREATE TABLE Relationship (
    id                INT AUTO_INCREMENT PRIMARY KEY,
    relationship_type ENUM('Parent', 'Child', 'Sibling', 'Spouse', 'Cousin', 'Uncle', 'Aunt') NOT NULL,
    member_id         INT, -- The member who has the relationship
    related_member_id INT, -- The member related to the member in question
    FOREIGN KEY (member_id) REFERENCES Member(id),
    FOREIGN KEY (related_member_id) REFERENCES Member(id)
);