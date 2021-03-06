package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"ucenter/internal/biz"
	"ucenter/internal/biz/entity"
)

type memberRepo struct {
	data *Data
	log  *log.Helper
}

func (r *memberRepo) CreateMember(member *entity.Member) (err error) {
	_, err = r.data.db.Insert(member)
	return
}

func (r *memberRepo) UpdateMember(member *entity.Member) error {
	panic("implement me")
}

func (r *memberRepo) GetMemberById(i int64) (m *entity.Member, err error) {
	panic("implement me")
}

func (r *memberRepo) GetMemberByName(s string) (m *entity.Member, err error) {
	panic("implement me")
}

func (r *memberRepo) GetMemberByPhone(s string) (m *entity.Member, err error) {
	panic("implement me")
}

func (r *memberRepo) GetMembers() (m *[]entity.Member, err error) {
	panic("implement me")
}

// NewGreeterRepo .
func NewMemberRepo(data *Data, logger log.Logger) biz.MemberRepo {
	return &memberRepo{
		data: data,
		log:  log.NewHelper("data/member", logger),
	}
}
