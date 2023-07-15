DROP TABLE IF EXISTS articles;

CREATE TABLE articles (
  id int NOT NULL AUTO_INCREMENT,
  title varchar(200),
  content text,
  category varchar(200),
  status ENUM('publish', 'draft', 'trash') NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (id)
);