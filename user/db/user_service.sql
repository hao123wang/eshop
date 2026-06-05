
CREATE TABLE IF NOT EXISTS `user` (
    `user_id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '用户id',
    `nick_name` varchar(20) NOT NULL DEFAULT "" COMMENT '用户昵称',
    `password` varchar(255) NOT NULL COMMENT "用户密码",
    `mobile` char(11) NOT NULL DEFAULT "" COMMENT '用户手机号',
    `email` varchar(20) NOT NULL DEFAULT "" COMMENT '用户邮箱',
    `birthday` date COMMENT '出生日期',
    `gender` tinyint unsigned NOT NULL DEFAULT 0 COMMENT '性别：0-女，1-男',
    `role` tinyint unsigned NOT NULL DEFAULT 1 COMMENT '用户角色：1-普通用户，2-管理员',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
    PRIMARY KEY (`user_id`) USING BTREE,
    KEY `idx_user_mobile` (`mobile`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC COMMENT='user 用户表';