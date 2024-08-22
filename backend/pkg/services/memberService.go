package services

import (
	"database/sql"
	"fmt"
	"socialNetwork/pkg/db/sqlite"
	"socialNetwork/pkg/models"
	"socialNetwork/utils"
)

type MemberService struct {
	db *sql.DB
}

func NewMemberService() *MemberService {
	dbs := sqlite.GlobalDB

	return &MemberService{
		db: dbs.GetDB(),
	}
}

func (m *MemberService) GetDB() *sql.DB {
	return m.db
}

func (m *MemberService) CreateMember(member models.Member) error {

	query := `INSERT INTO Membership (userId,groupId,role) VALUES (?,?,?)`

	_, err := m.GetDB().Exec(query, member.UserId, member.GroupId, member.Role)

	if err != nil {
		return fmt.Errorf("failed to create groupe: %w", err)
	}

	return nil
}

func (m *MemberService) GetMembership(groupId string) ([]models.Member, error) {
	query := `SELECT * FROM Membership WHERE groupId = ? ORDER BY joinDate DESC`

	rows, err := m.GetDB().Query(query,groupId)
	if err != nil {
		return nil, fmt.Errorf("failed to get membership: %w", err)
	}
	defer rows.Close()

	var membership []models.Member
	for rows.Next() {
		var member models.Member
		
		if err = rows.Scan(&member.UserId, &member.GroupId, &member.Role, &member.JoinDate); err != nil {
			return nil, fmt.Errorf("failed to scan member: %w", err)
		}
		memInfo, err := utils.GetAuthor(m.db, member.UserId)
		if err != nil {
			return nil, fmt.Errorf("failed to get member: %w", err)
		}
		member.MemberInfo = memInfo
		membership = append(membership, member)

	}

	return membership, nil
}
