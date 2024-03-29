create database `user` default charset utf8mb4;
create database goods_service default charset utf8mb4;
create databse inventory default charset utf8mb4;


use `user`;
CREATE TABLE `user` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `username` varchar(100) NOT NULL,
    `password` varchar(100) NOT NULL,
    `email` varchar(100) NOT NULL,
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

use goods_service;

CREATE TABLE `goods`
(
    `id`                int(11) NOT NULL AUTO_INCREMENT,
    `add_time`          datetime       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time`       datetime       NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `is_deleted`        tinyint(1) NOT NULL DEFAULT 0,
    `category_id`       int(11) NOT NULL,
    `brands_id`         int(11) NOT NULL,
    `on_sale`           tinyint(1) NOT NULL DEFAULT '0',
    `ship_free`         tinyint(1) NOT NULL DEFAULT '0',
    `is_new`            tinyint(1) NOT NULL DEFAULT '0',
    `is_hot`            tinyint(1) NOT NULL DEFAULT '0',
    `name`              varchar(50)    NOT NULL,
    `goods_sn`          varchar(50)    NOT NULL,
    `click_num`         int(11) NOT NULL DEFAULT '0',
    `sold_num`          int(11) NOT NULL DEFAULT '0',
    `fav_num`           int(11) NOT NULL DEFAULT '0',
    `market_price`      decimal(10, 2) NOT NULL default '0.00',
    `shop_price`        decimal(10, 2) NOT NULL default '0.00',
    `goods_brief`       varchar(100)   NOT NULL default '',
    `images`            varchar(1000)  NOT NULL default '',
    `desc_images`       varchar(1000)  NOT NULL default '',
    `goods_front_image` varchar(200)   NOT NULL default '',
    PRIMARY KEY (`id`),
    KEY                 `idx_goods_category` (`category_id`),
    KEY                 `idx_goods_brands` (`brands_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


CREATE TABLE `category`
(
    `id`                 int(11) NOT NULL AUTO_INCREMENT,
    `add_time`           datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time`        datetime    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `is_deleted`         tinyint(1) NOT NULL DEFAULT 0,
    `name`               varchar(20) NOT NULL,
    `parent_category_id` int(11) DEFAULT NULL,
    `level`              int(11) NOT NULL DEFAULT '1',
    `is_tab`             tinyint(1) NOT NULL DEFAULT '0',
    PRIMARY KEY (`id`),
    KEY                  `idx_category_sub_category` (`parent_category_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



CREATE TABLE `goodscategorybrand`
(
    `id`          int(11) NOT NULL AUTO_INCREMENT,
    `add_time`    datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `is_deleted`  tinyint(1) NOT NULL DEFAULT 0,
    `category_id` int(11) DEFAULT NULL,
    `brands_id`   int(11) DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_category_brand` (`category_id`,`brands_id`),
    KEY           `idx_goodscategorybrand_brands` (`brands_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



CREATE TABLE `banner`
(
    `id`          int(11) NOT NULL AUTO_INCREMENT,
    `add_time`    datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `is_deleted`  tinyint(1) NOT NULL DEFAULT 0,
    `image`       varchar(200) NOT NULL,
    `url`         varchar(200) NOT NULL,
    `index`       int(11) NOT NULL DEFAULT '1',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



CREATE TABLE `brands`
(
    `id`          int(11) NOT NULL AUTO_INCREMENT,
    `add_time`    datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `is_deleted`  tinyint(1) NOT NULL DEFAULT 0,
    `name`        varchar(20)  NOT NULL,
    `logo`        varchar(200) NOT NULL DEFAULT '',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4;


use inventory;



CREATE TABLE `inventory`(
    `id`          int(11) NOT NULL AUTO_INCREMENT,
    `add_time`    datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `is_deleted`  tinyint(1) NOT NULL DEFAULT 0,
    goods int(11) not null comment '商品id',
    stocks int(11) not null default '0' comment '库存数',
    version int(11) not null default '0' comment '版本号 乐观锁 只是用来学习用',
    freeze int(11) not null default '0' comment '冻结数,tcc使用',
    primary key(id),
    key idx_goods(goods)
)ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4;


CREATE TABLE `stockselldetail`(
          `id`          int(11) NOT NULL AUTO_INCREMENT,
          `add_time`    datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP,
          `update_time` datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
          `is_deleted`  tinyint(1) NOT NULL DEFAULT 0,
          `order_sn` varchar(100) not null comment '订单号',
          `detail`   varchar(200) not null default '' comment '商品信息',
          `status`  tinyint(1) not null default '0' comment '1 表示已扣减 2. 表示已归还',
          primary key(id),
          key idx_order_sn(order_sn)
)ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4;


