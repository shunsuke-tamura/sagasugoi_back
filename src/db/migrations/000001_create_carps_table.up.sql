CREATE TABLE `carps`
(
 `id`         int NOT NULL AUTO_INCREMENT ,
 `word`       varchar(256) NOT NULL ,
 `comment`    varchar(512) NOT NULL ,
 `url`        varchar(256) NULL ,
 `updated_at` datetime NOT NULL ,
 `created_at` datetime NOT NULL ,
 `deleted_at` datetime NULL ,
 `size`       int NOT NULL ,
 `r`          int NOT NULL ,
 `g`          int NOT NULL ,
 `b`          int NOT NULL ,

PRIMARY KEY (`id`)
);
