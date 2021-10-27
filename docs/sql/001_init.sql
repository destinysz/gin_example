CREATE DATABASE IF NOT EXISTS shan_user DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_general_ci;

USE shan_user;


CREATE TABLE `tb_depart` (
 `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'id',
 `create_time` datetime(6) NOT NULL COMMENT '创建时间',
 `update_time` datetime(6) NOT NULL COMMENT '更新时间',
 `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '逻辑删除 1未删除 0已删除',
 `name` varchar(30) NOT NULL DEFAULT '' COMMENT '部门名称',
 `leader` varchar(30) NOT NULL DEFAULT '' COMMENT '部门领导',
 `email` varchar(128) NOT NULL DEFAULT '' COMMENT '部门邮箱',
 PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='部门表';


CREATE TABLE `tb_user` (
 `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'id',
 `create_time` datetime(6) NOT NULL COMMENT '创建时间',
 `update_time` datetime(6) NOT NULL COMMENT '更新时间',
 `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '逻辑删除 1未删除 0已删除',
 `name_cn` varchar(30) NOT NULL DEFAULT '' COMMENT '中文名',
 `name_en` varchar(30) NOT NULL DEFAULT '' COMMENT '英文名',
 `mobile` varchar(20) NOT NULL DEFAULT '' COMMENT '手机',
 `email` varchar(128) NOT NULL DEFAULT '' COMMENT '邮箱',
 `password` varchar(50) NOT NULL DEFAULT '' COMMENT '密码', 
 `depart_id` int(11) NOT NULL COMMENT '部门id',
 FOREIGN KEY (`depart_id`) REFERENCES `tb_depart` (`id`),
 PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户表';


CREATE TABLE `tb_permission_group` (
 `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'id',
 `create_time` datetime(6) NOT NULL COMMENT '创建时间',
 `update_time` datetime(6) NOT NULL COMMENT '更新时间',
 `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '逻辑删除 1未删除 0已删除',
 `name` varchar(30) COMMENT '权限组名称',
 PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='权限组表';


CREATE TABLE `tb_permission` (
 `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'id',
 `create_time` datetime(6) NOT NULL COMMENT '创建时间',
 `update_time` datetime(6) NOT NULL COMMENT '更新时间',
 `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '逻辑删除 1未删除 0已删除',
 `name` varchar(30) NOT NULL COMMENT '权限名称',
 `interface` varchar(100) NOT NULL COMMENT '权限名称',
 `permission_group_id` int(11) NOT NULL COMMENT '权限组id',
 PRIMARY KEY (`id`),
 FOREIGN KEY (`permission_group_id`) REFERENCES `tb_permission_group` (`id`),
 KEY (`interface`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='权限表';


CREATE TABLE `tb_role` (
 `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'id',
 `create_time` datetime(6) NOT NULL COMMENT '创建时间',
 `update_time` datetime(6) NOT NULL COMMENT '更新时间',
 `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '逻辑删除 1未删除 0已删除',
 `name` varchar(30) NOT NULL COMMENT '角色名称',
 `default` tinyint(1) NOT NULL DEFAULT 1 COMMENT '是否默认角色 1是 0否',
 PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COMMENT='角色表';

INSERT INTO `tb_role` VALUES 
 (1,'2021-09-01 09:00:00','2021-09-01 09:00:00',1,'超级管理员',0),
 (2,'2021-09-01 09:00:00','2021-09-01 09:00:00',1,'安全管理员',0),
 (3,'2021-09-01 09:00:00','2021-09-01 09:00:00',1,'安全测试员',0),
 (4,'2021-09-01 09:00:00','2021-09-01 09:00:00',1,'普通用户',1);


CREATE TABLE `tb_role_permission_ref` (
 `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'id',
 `create_time` datetime(6) NOT NULL COMMENT '创建时间',
 `update_time` datetime(6) NOT NULL COMMENT '更新时间',
 `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '逻辑删除 1未删除 0已删除',
 `role_id` int(11) NOT NULL COMMENT '角色id',
 `permission_id` int(11) NOT NULL COMMENT '权限id',
 PRIMARY KEY (`id`),
 FOREIGN KEY (`role_id`) REFERENCES `tb_role` (`id`),
 FOREIGN KEY (`permission_id`) REFERENCES `tb_permission` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='角色权限关联表';



CREATE TABLE `tb_user_role_ref` (
 `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'id',
 `create_time` datetime(6) NOT NULL COMMENT '创建时间',
 `update_time` datetime(6) NOT NULL COMMENT '更新时间',
 `status` tinyint(1) NOT NULL DEFAULT 1 COMMENT '逻辑删除 1未删除 0已删除',
 `user_id` int(11) NOT NULL COMMENT '用户id',
 `role_id` int(11) NOT NULL COMMENT '角色id',
 PRIMARY KEY (`id`),
 FOREIGN KEY (`user_id`) REFERENCES `tb_user` (`id`),
 FOREIGN KEY (`role_id`) REFERENCES `tb_role` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户角色关联表';
