CREATE TABLE group (
    `id` int NOT NULL COMMENT '知识点分组 ID',
    `name` varchar(10) NOT NULL DEFAULT '' COMMENT '名称',
    `create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `delete_at` timestamp NULL COMMENT '删除时间',
    PRIMARY KEY (id)
) CHARACTER SET utf8mb4 engine=Innodb;