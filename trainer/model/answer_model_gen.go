// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	answerFieldNames          = builder.RawFieldNames(&Answer{})
	answerRows                = strings.Join(answerFieldNames, ",")
	answerRowsExpectAutoSet   = strings.Join(stringx.Remove(answerFieldNames, "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	answerRowsWithPlaceHolder = strings.Join(stringx.Remove(answerFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"
)

type (
	answerModel interface {
		Insert(ctx context.Context, data *Answer) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Answer, error)
		Update(ctx context.Context, data *Answer) error
		Delete(ctx context.Context, id int64) error
	}

	defaultAnswerModel struct {
		conn  sqlx.SqlConn
		table string
	}

	Answer struct {
		Id         int64        `db:"id"`          // 答案 ID
		Answer     string       `db:"answer"`      // 答案
		QuestionId int64        `db:"question_id"` // 题目 ID
		CreateAt   time.Time    `db:"create_at"`   // 创建时间
		UpdateAt   time.Time    `db:"update_at"`   // 更新时间
		DeleteAt   sql.NullTime `db:"delete_at"`   // 删除时间
	}
)

func newAnswerModel(conn sqlx.SqlConn) *defaultAnswerModel {
	return &defaultAnswerModel{
		conn:  conn,
		table: "`answer`",
	}
}

func (m *defaultAnswerModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultAnswerModel) FindOne(ctx context.Context, id int64) (*Answer, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", answerRows, m.table)
	var resp Answer
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultAnswerModel) Insert(ctx context.Context, data *Answer) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, answerRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Id, data.Answer, data.QuestionId, data.DeleteAt)
	return ret, err
}

func (m *defaultAnswerModel) Update(ctx context.Context, data *Answer) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, answerRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Answer, data.QuestionId, data.DeleteAt, data.Id)
	return err
}

func (m *defaultAnswerModel) tableName() string {
	return m.table
}
