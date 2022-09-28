CREATE TABLE IF NOT EXISTS `carps`
(
 `id`      char(32) NOT NULL ,
 `word`    varchar(256) NOT NULL ,
 `typeId`  char(32) NOT NULL ,
 `comment` varchar(512) NULL ,
 `url`     varchar(256) NULL ,
PRIMARY KEY (`id`),
KEY `FK_2` (`typeId`),
CONSTRAINT `FK_1` FOREIGN KEY `FK_2` (`typeId`) REFERENCES `types` (`id`)
);
