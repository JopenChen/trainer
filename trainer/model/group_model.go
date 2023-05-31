package model

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"strings"
)

var _ GroupModel = (*customGroupModel)(nil)

type (
	// GroupModel is an interface to be customized, add more methods here,
	// and implement the added methods in customGroupModel.
	GroupModel interface {
		groupModel
		GetByCondition(ctx context.Context, conditionMap map[string]interface{}) (resp *Group, err error)
		FindByCondition(ctx context.Context, conditionMap map[string]interface{}) (resp []Group, err error)
		DeleteByCondition(ctx context.Context, conditionMap map[string]interface{}) (err error)
	}

	customGroupModel struct {
		*defaultGroupModel
	}
)

func (c customGroupModel) GetByCondition(ctx context.Context, conditionMap map[string]interface{}) (resp *Group, err error) {
	resp = new(Group)
	var conditionBuilder strings.Builder

	for key, val := range conditionMap {
		splited := strings.Split(key, "_")

		var prefix string
		var field string
		if splited[0] == "not" {
			prefix = strings.Join(splited[:1], "_")
			field = strings.Join(splited[2:], "_")
		} else {
			prefix = splited[0]
			field = strings.Join(splited[1:], "_")
		}

		switch val.(type) {
		case int, int64, int32, float32, float64:
			switch prefix {
			case "eq":
				conditionBuilder.WriteString(fmt.Sprintf("and %s = %v", field, val))
			case "not_eq":
				conditionBuilder.WriteString(fmt.Sprintf("and %s != %v", field, val))
			case "lt":
				conditionBuilder.WriteString(fmt.Sprintf("and %s < %v", field, val))
			case "gt":
				conditionBuilder.WriteString(fmt.Sprintf("and %s > %v", field, val))
			}

		case string:
			switch prefix {
			case "eq":
				conditionBuilder.WriteString(fmt.Sprintf("and %s = '%v'", field, val))
			case "not_eq":
				conditionBuilder.WriteString(fmt.Sprintf("and %s != '%v'", field, val))
			}
		default:
			return nil, errors.New("unknown field type")
		}
	}

	conditionRaw := conditionBuilder.String()

	query := fmt.Sprintf("select %s from %s where 1 = 1 %s and delete_at is NULL limit 1", groupRows, c.table, conditionRaw)
	err = c.conn.QueryRowCtx(ctx, resp, query)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (c customGroupModel) FindByCondition(ctx context.Context, conditionMap map[string]interface{}) (resp []Group, err error) {
	resp = make([]Group, 5)
	var conditionBuilder strings.Builder

	for key, val := range conditionMap {
		splited := strings.Split(key, "_")

		var prefix string
		var field string
		if splited[0] == "not" {
			prefix = strings.Join(splited[:1], "_")
			field = strings.Join(splited[2:], "_")
		} else {
			prefix = splited[0]
			field = strings.Join(splited[1:], "_")
		}

		switch val.(type) {
		case int, int64, int32, float32, float64:
			switch prefix {
			case "eq":
				conditionBuilder.WriteString(fmt.Sprintf("and %s = %v", field, val))
			case "not_eq":
				conditionBuilder.WriteString(fmt.Sprintf("and %s != %v", field, val))
			case "lt":
				conditionBuilder.WriteString(fmt.Sprintf("and %s < %v", field, val))
			case "gt":
				conditionBuilder.WriteString(fmt.Sprintf("and %s > %v", field, val))
			}

		case string:
			switch prefix {
			case "eq":
				conditionBuilder.WriteString(fmt.Sprintf("and %s = '%v'", field, val))
			case "not_eq":
				conditionBuilder.WriteString(fmt.Sprintf("and %s != '%v'", field, val))
			}
		default:
			return nil, errors.New("unknown field type")
		}
	}

	conditionRaw := conditionBuilder.String()

	query := fmt.Sprintf("select %s from %s where 1 = 1 %s and delete_at is NULL", groupRows, c.table, conditionRaw)
	err = c.conn.QueryRowCtx(ctx, &resp, query)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (c customGroupModel) DeleteByCondition(ctx context.Context, conditionMap map[string]interface{}) (err error) {
	var conditionBuilder strings.Builder

	for key, val := range conditionMap {
		splited := strings.Split(key, "_")

		var prefix string
		var field string
		if splited[0] == "not" {
			prefix = strings.Join(splited[:1], "_")
			field = strings.Join(splited[2:], "_")
		} else {
			prefix = splited[0]
			field = strings.Join(splited[1:], "_")
		}

		switch val.(type) {
		case int, int64, int32, float32, float64:
			switch prefix {
			case "eq":
				conditionBuilder.WriteString(fmt.Sprintf("and %s = %v", field, val))
			case "not_eq":
				conditionBuilder.WriteString(fmt.Sprintf("and %s != %v", field, val))
			case "lt":
				conditionBuilder.WriteString(fmt.Sprintf("and %s < %v", field, val))
			case "gt":
				conditionBuilder.WriteString(fmt.Sprintf("and %s > %v", field, val))
			}

		case string:
			switch prefix {
			case "eq":
				conditionBuilder.WriteString(fmt.Sprintf("and %s = '%v'", field, val))
			case "not_eq":
				conditionBuilder.WriteString(fmt.Sprintf("and %s != '%v'", field, val))
			}
		default:
			return errors.New("unknown field type")
		}
	}

	conditionRaw := conditionBuilder.String()

	query := fmt.Sprintf("delete from %s where 1 = 1 %s and delete_at is NULL", c.table, conditionRaw)
	_, err = c.conn.ExecCtx(ctx, query)
	return err
}

// NewGroupModel returns a model for the database table.
func NewGroupModel(conn sqlx.SqlConn) GroupModel {
	return &customGroupModel{
		defaultGroupModel: newGroupModel(conn),
	}
}
