package biz

import (
	"github.com/go-kratos/kratos/v2/log"
	"ucenter/internal/biz/entity"
)

//用仓储层操作集合
type MemberRepo interface {
	CreateMember(*entity.Member) error
	UpdateMember(*entity.Member) error
	GetMemberById(int64) (m *entity.Member, err error)
	GetMemberByName(string) (m *entity.Member, err error)
	GetMemberByPhone(string) (m *entity.Member, err error)
	GetMembers() (m *[]entity.Member, err error)
}

type MemberMgr struct {
	repo MemberRepo //用户相关repository
	log  *log.Helper
}

func NewMemberMgr(repo MemberRepo, logger log.Logger) *MemberMgr {
	return &MemberMgr{repo: repo, log: log.NewHelper("ucenter/Member", logger)}
}

func (uc *MemberMgr) Create(g *entity.Member) error {
	return uc.repo.CreateMember(g)
}

func (uc *MemberMgr) Update(g *entity.Member) error {
	return uc.repo.UpdateMember(g)
}

func (uc *MemberMgr) GetMemberById(id int64) (m *entity.Member, err error) {
	return uc.repo.GetMemberById(id)
}

func (uc *MemberMgr) GetMemberByName(name string) (m *entity.Member, err error) {
	return uc.repo.GetMemberByName(name)
}

func (uc *MemberMgr) GetMemberByPhone(phone string) (m *entity.Member, err error) {
	return uc.repo.GetMemberByPhone(phone)
}

func (uc *MemberMgr) List() (m *[]entity.Member, err error) {
	return uc.repo.GetMembers()
}
