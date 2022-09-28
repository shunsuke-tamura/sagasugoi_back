CREATE TABLE IF NOT EXISTS `types`
(
 `id`   char(32) NOT NULL ,
 `size` double NOT NULL ,
 `r`    int NOT NULL ,
 `g`    int NOT NULL ,
 `b`    int NOT NULL ,

PRIMARY KEY (`id`)
);
