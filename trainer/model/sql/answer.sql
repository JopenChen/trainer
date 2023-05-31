CREATE TABLE answer (
    `id` int NOT NULL COMMENT '答案 ID',
    `answer` json NOT NULL default '' COMMENT '答案',
    `question_id` int NOT NULL DEFAULT 0 COMMENT '题目 ID',
    `create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    `delete_at` timestamp NULL COMMENT '删除时间',
    PRIMARY KEY (id)
) CHARACTER SET utf8mb4 engine=Innodb;