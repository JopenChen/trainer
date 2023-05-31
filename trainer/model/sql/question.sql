CREATE TABLE question (
    `id` int NOT NULL COMMENT '题目 ID',
    `question` text NOT NULL COMMENT '题目',
    `status` varchar(3) NOT NULL DEFAULT '未回答' COMMENT '回答状态，已回答、未回答',
    `group_id` int NOT NULL DEFAULT 0 COMMENT '知识点分组 ID',
    `create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `delete_at` timestamp NULL COMMENT '删除时间',
    PRIMARY KEY (id)
) CHARACTER SET utf8mb4 engine=Innodb;