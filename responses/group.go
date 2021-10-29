package responses

import (
	"bytes"

	"github.com/Bananenpro/hbank-api/models"
	"github.com/Bananenpro/hbank-api/services"
	"github.com/google/uuid"
)

type CreateGroupSuccess struct {
	Base
	Id string `json:"id"`
}

type Balance struct {
	Base
	Balance int `json:"balance"`
}

type DeleteFailedBecauseOfSoleGroupAdmin struct {
	Base
	GroupIds []string `json:"group_ids"`
}

type group struct {
	Id             string `json:"id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	GroupPictureId string `json:"group_picture_id"`
}

type groupDetailed struct {
	Id             string `json:"id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	GroupPictureId string `json:"group_picture_id"`
	Member         bool   `json:"member"`
	Admin          bool   `json:"admin"`
}

type transaction struct {
	Id          string `json:"id"`
	Time        int64  `json:"time"`
	Title       string `json:"title"`
	Description string `json:"description"`

	GroupId string `json:"group_id"`

	BalanceDifference int `json:"balance_difference"`
	NewBalance        int `json:"new_balance"`

	SenderId   string `json:"sender_id,omitempty"`
	ReceiverId string `json:"receiver_id,omitempty"`
}

type bankTransaction struct {
	Id          string `json:"id"`
	Time        int64  `json:"time"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Amount      int    `json:"amount"`

	GroupId string `json:"group_id"`

	SenderId   string `json:"sender_id,omitempty"`
	ReceiverId string `json:"receiver_id,omitempty"`
}

type invitation struct {
	Id                string `json:"id"`
	Created           int64  `json:"created"`
	InvitationMessage string `json:"invitation_message"`
	GroupId           string `json:"group_id,omitempty"`
	UserId            string `json:"user_id,omitempty"`
}

func NewInvitations(invitations []models.GroupInvitation) interface{} {
	dtos := make([]invitation, len(invitations))
	for i, in := range invitations {
		dtos[i].Id = in.Id.String()
		dtos[i].Created = in.Created
		dtos[i].InvitationMessage = in.Message
		dtos[i].UserId = in.UserId.String()
		dtos[i].GroupId = in.GroupId.String()
	}

	type invitationsResp struct {
		Base
		Invitations []invitation `json:"invitations"`
	}

	return invitationsResp{
		Base: Base{
			Success: true,
		},
		Invitations: dtos,
	}
}

func NewInvitation(invitationModel *models.GroupInvitation) interface{} {
	type invitationResp struct {
		Base
		invitation
	}

	return invitationResp{
		Base: Base{
			Success: true,
		},
		invitation: invitation{
			Id:                invitationModel.Id.String(),
			Created:           invitationModel.Created,
			InvitationMessage: invitationModel.Message,
			GroupId:           invitationModel.GroupId.String(),
			UserId:            invitationModel.UserId.String(),
		},
	}
}

func NewGroups(groups []models.Group) interface{} {
	groupDTOs := make([]group, len(groups))
	for i, g := range groups {
		groupDTOs[i].Id = g.Id.String()
		groupDTOs[i].Name = g.Name
		groupDTOs[i].Description = g.Description
		groupDTOs[i].GroupPictureId = g.GroupPictureId.String()
	}

	type groupsResp struct {
		Base
		Groups []group `json:"groups"`
	}

	return groupsResp{
		Base: Base{
			Success: true,
		},
		Groups: groupDTOs,
	}
}

func NewGroup(group *models.Group, isMember, isAdmin bool) interface{} {
	type groupResp struct {
		Base
		groupDetailed
	}

	return groupResp{
		Base: Base{
			Success: true,
		},
		groupDetailed: groupDetailed{
			Id:             group.Id.String(),
			Name:           group.Name,
			Description:    group.Description,
			GroupPictureId: group.GroupPictureId.String(),
			Member:         isMember,
			Admin:          isAdmin,
		},
	}
}

func NewTransaction(transactionModel *models.TransactionLogEntry, user *models.User) interface{} {
	type transactionResp struct {
		Base
		transaction
	}

	isSender := bytes.Equal(user.Id[:], transactionModel.SenderId[:])

	balanceDifference := transactionModel.Amount
	if isSender {
		balanceDifference = -transactionModel.Amount
	}

	newBalance := transactionModel.NewBalanceReceiver
	if isSender {
		newBalance = transactionModel.NewBalanceSender
	}

	transactionDTO := transaction{
		Id:                transactionModel.Id.String(),
		Time:              transactionModel.Created,
		Title:             transactionModel.Title,
		Description:       transactionModel.Description,
		BalanceDifference: balanceDifference,
		NewBalance:        newBalance,
		GroupId:           transactionModel.GroupId.String(),
	}

	if isSender {
		if transactionModel.ReceiverIsBank {
			transactionDTO.ReceiverId = "bank"
		} else {
			transactionDTO.ReceiverId = transactionModel.ReceiverId.String()
		}
	} else {
		if transactionModel.SenderIsBank {
			transactionDTO.SenderId = "bank"
		} else {
			transactionDTO.SenderId = transactionModel.SenderId.String()
		}
	}

	return transactionResp{
		Base: Base{
			Success: true,
		},
		transaction: transactionDTO,
	}
}

func NewBankTransaction(transactionModel *models.TransactionLogEntry) interface{} {
	type transactionResp struct {
		Base
		bankTransaction
	}

	transactionDTO := bankTransaction{
		Id:          transactionModel.Id.String(),
		Time:        transactionModel.Created,
		Title:       transactionModel.Title,
		Description: transactionModel.Description,
		Amount:      transactionModel.Amount,
		GroupId:     transactionModel.GroupId.String(),
	}

	if transactionModel.ReceiverIsBank {
		transactionDTO.ReceiverId = "bank"
	} else {
		transactionDTO.ReceiverId = transactionModel.ReceiverId.String()
	}

	if transactionModel.SenderIsBank {
		transactionDTO.SenderId = "bank"
	} else {
		transactionDTO.SenderId = transactionModel.SenderId.String()
	}

	return transactionResp{
		Base: Base{
			Success: true,
		},
		bankTransaction: transactionDTO,
	}
}

func NewTransactionLog(log []models.TransactionLogEntry, user *models.User) interface{} {
	type transactionsResp struct {
		Base
		Transactions []transaction `json:"transactions"`
	}

	transactionDTOs := make([]transaction, len(log))

	for i, entry := range log {
		isSender := bytes.Equal(user.Id[:], entry.SenderId[:])

		balanceDifference := entry.Amount
		if isSender {
			balanceDifference = -entry.Amount
		}

		newBalance := entry.NewBalanceReceiver
		if isSender {
			newBalance = entry.NewBalanceSender
		}

		transactionDTO := transaction{
			Id:                entry.Id.String(),
			Time:              entry.Created,
			Title:             entry.Title,
			Description:       entry.Description,
			BalanceDifference: balanceDifference,
			NewBalance:        newBalance,
			GroupId:           entry.GroupId.String(),
		}

		if isSender {
			if entry.ReceiverIsBank {
				transactionDTO.ReceiverId = "bank"
			} else {
				transactionDTO.ReceiverId = entry.ReceiverId.String()
			}
		} else {
			if entry.SenderIsBank {
				transactionDTO.SenderId = "bank"
			} else {
				transactionDTO.SenderId = entry.SenderId.String()
			}
		}

		transactionDTOs[i] = transactionDTO
	}

	return transactionsResp{
		Base: Base{
			Success: true,
		},
		Transactions: transactionDTOs,
	}
}

func NewBankTransactionLog(log []models.TransactionLogEntry) interface{} {
	type transactionsResp struct {
		Base
		Transactions []bankTransaction `json:"transactions"`
	}

	transactionDTOs := make([]bankTransaction, len(log))

	for i, entry := range log {
		transactionDTO := bankTransaction{
			Id:          entry.Id.String(),
			Time:        entry.Created,
			Title:       entry.Title,
			Description: entry.Description,
			Amount:      entry.Amount,
			GroupId:     entry.GroupId.String(),
		}

		if entry.ReceiverIsBank {
			transactionDTO.ReceiverId = "bank"
		} else {
			transactionDTO.ReceiverId = entry.ReceiverId.String()
		}

		if entry.SenderIsBank {
			transactionDTO.SenderId = "bank"
		} else {
			transactionDTO.SenderId = entry.SenderId.String()
		}

		transactionDTOs[i] = transactionDTO
	}

	return transactionsResp{
		Base: Base{
			Success: true,
		},
		Transactions: transactionDTOs,
	}
}

func NewDeleteFailedBecauseOfSoleGroupAdmin(groupIds []uuid.UUID, lang string) interface{} {
	ids := make([]string, len(groupIds))
	for i := range groupIds {
		ids[i] = groupIds[i].String()
	}

	return &DeleteFailedBecauseOfSoleGroupAdmin{
		Base: Base{
			Success: false,
			Message: services.Tr("Failed to delete user because he is the only admin of one or more groups", lang),
		},
		GroupIds: ids,
	}
}
