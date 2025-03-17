// Code generated by ent, DO NOT EDIT.

package ent

import (
	"company/biz/dal/mysql/ent/user"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	// primary key
	ID int64 `json:"id,omitempty"`
	// created time
	CreatedAt time.Time `json:"created_at,omitempty"`
	// last update time
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// 状态[0:禁用;1:正常]
	Status int64 `json:"status,omitempty"`
	// user's login name | 登录名
	Username string `json:"username,omitempty"`
	// password | 密码
	Password string `json:"password,omitempty"`
	// nickname | 昵称
	Nickname string `json:"nickname,omitempty"`
	// template mode | 布局方式
	SideMode string `json:"side_mode,omitempty"`
	// base color of template | 后台页面色调
	BaseColor string `json:"base_color,omitempty"`
	// active color of template | 当前激活的颜色设定
	ActiveColor string `json:"active_color,omitempty"`
	// role id | 角色ID
	RoleID int64 `json:"role_id,omitempty"`
	// mobile number | 手机号
	Mobile string `json:"mobile,omitempty"`
	// email | 邮箱号
	Email string `json:"email,omitempty"`
	// wecom | 微信号
	Wecom string `json:"wecom,omitempty"`
	// 职业
	Job string `json:"job,omitempty"`
	// 部门
	Organization string `json:"organization,omitempty"`
	// 登陆后默认场馆ID
	DefaultVenueID int64 `json:"default_venue_id,omitempty"`
	// avatar | 头像路径
	Avatar string `json:"avatar,omitempty"`
	// 性别 | [0:女性;1:男性;3:保密]
	Gender int64 `json:"gender,omitempty"`
	// 出生日期
	Birthday time.Time `json:"birthday,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges        UserEdges `json:"edges"`
	selectValues sql.SelectValues
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// UserEntry holds the value of the user_entry edge.
	UserEntry []*EntryLogs `json:"user_entry,omitempty"`
	// UserFace holds the value of the user_face edge.
	UserFace []*Face `json:"user_face,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// UserEntryOrErr returns the UserEntry value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) UserEntryOrErr() ([]*EntryLogs, error) {
	if e.loadedTypes[0] {
		return e.UserEntry, nil
	}
	return nil, &NotLoadedError{edge: "user_entry"}
}

// UserFaceOrErr returns the UserFace value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) UserFaceOrErr() ([]*Face, error) {
	if e.loadedTypes[1] {
		return e.UserFace, nil
	}
	return nil, &NotLoadedError{edge: "user_face"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldID, user.FieldStatus, user.FieldRoleID, user.FieldDefaultVenueID, user.FieldGender:
			values[i] = new(sql.NullInt64)
		case user.FieldUsername, user.FieldPassword, user.FieldNickname, user.FieldSideMode, user.FieldBaseColor, user.FieldActiveColor, user.FieldMobile, user.FieldEmail, user.FieldWecom, user.FieldJob, user.FieldOrganization, user.FieldAvatar:
			values[i] = new(sql.NullString)
		case user.FieldCreatedAt, user.FieldUpdatedAt, user.FieldBirthday:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = int64(value.Int64)
		case user.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				u.CreatedAt = value.Time
			}
		case user.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				u.UpdatedAt = value.Time
			}
		case user.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				u.Status = value.Int64
			}
		case user.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				u.Username = value.String
			}
		case user.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				u.Password = value.String
			}
		case user.FieldNickname:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field nickname", values[i])
			} else if value.Valid {
				u.Nickname = value.String
			}
		case user.FieldSideMode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field side_mode", values[i])
			} else if value.Valid {
				u.SideMode = value.String
			}
		case user.FieldBaseColor:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field base_color", values[i])
			} else if value.Valid {
				u.BaseColor = value.String
			}
		case user.FieldActiveColor:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field active_color", values[i])
			} else if value.Valid {
				u.ActiveColor = value.String
			}
		case user.FieldRoleID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field role_id", values[i])
			} else if value.Valid {
				u.RoleID = value.Int64
			}
		case user.FieldMobile:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mobile", values[i])
			} else if value.Valid {
				u.Mobile = value.String
			}
		case user.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				u.Email = value.String
			}
		case user.FieldWecom:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field wecom", values[i])
			} else if value.Valid {
				u.Wecom = value.String
			}
		case user.FieldJob:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field job", values[i])
			} else if value.Valid {
				u.Job = value.String
			}
		case user.FieldOrganization:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field organization", values[i])
			} else if value.Valid {
				u.Organization = value.String
			}
		case user.FieldDefaultVenueID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field default_venue_id", values[i])
			} else if value.Valid {
				u.DefaultVenueID = value.Int64
			}
		case user.FieldAvatar:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field avatar", values[i])
			} else if value.Valid {
				u.Avatar = value.String
			}
		case user.FieldGender:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field gender", values[i])
			} else if value.Valid {
				u.Gender = value.Int64
			}
		case user.FieldBirthday:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field birthday", values[i])
			} else if value.Valid {
				u.Birthday = value.Time
			}
		default:
			u.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the User.
// This includes values selected through modifiers, order, etc.
func (u *User) Value(name string) (ent.Value, error) {
	return u.selectValues.Get(name)
}

// QueryUserEntry queries the "user_entry" edge of the User entity.
func (u *User) QueryUserEntry() *EntryLogsQuery {
	return NewUserClient(u.config).QueryUserEntry(u)
}

// QueryUserFace queries the "user_face" edge of the User entity.
func (u *User) QueryUserFace() *FaceQuery {
	return NewUserClient(u.config).QueryUserFace(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return NewUserClient(u.config).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("created_at=")
	builder.WriteString(u.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(u.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", u.Status))
	builder.WriteString(", ")
	builder.WriteString("username=")
	builder.WriteString(u.Username)
	builder.WriteString(", ")
	builder.WriteString("password=")
	builder.WriteString(u.Password)
	builder.WriteString(", ")
	builder.WriteString("nickname=")
	builder.WriteString(u.Nickname)
	builder.WriteString(", ")
	builder.WriteString("side_mode=")
	builder.WriteString(u.SideMode)
	builder.WriteString(", ")
	builder.WriteString("base_color=")
	builder.WriteString(u.BaseColor)
	builder.WriteString(", ")
	builder.WriteString("active_color=")
	builder.WriteString(u.ActiveColor)
	builder.WriteString(", ")
	builder.WriteString("role_id=")
	builder.WriteString(fmt.Sprintf("%v", u.RoleID))
	builder.WriteString(", ")
	builder.WriteString("mobile=")
	builder.WriteString(u.Mobile)
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(u.Email)
	builder.WriteString(", ")
	builder.WriteString("wecom=")
	builder.WriteString(u.Wecom)
	builder.WriteString(", ")
	builder.WriteString("job=")
	builder.WriteString(u.Job)
	builder.WriteString(", ")
	builder.WriteString("organization=")
	builder.WriteString(u.Organization)
	builder.WriteString(", ")
	builder.WriteString("default_venue_id=")
	builder.WriteString(fmt.Sprintf("%v", u.DefaultVenueID))
	builder.WriteString(", ")
	builder.WriteString("avatar=")
	builder.WriteString(u.Avatar)
	builder.WriteString(", ")
	builder.WriteString("gender=")
	builder.WriteString(fmt.Sprintf("%v", u.Gender))
	builder.WriteString(", ")
	builder.WriteString("birthday=")
	builder.WriteString(u.Birthday.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User
