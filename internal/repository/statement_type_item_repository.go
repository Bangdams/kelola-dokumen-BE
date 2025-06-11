package repository

import (
	"github.com/Bangdams/kelola-dokumen-BE/internal/entity"
	"gorm.io/gorm"
)

type StatementTypeItemRepository interface {
	Create(tx *gorm.DB, statementTypeItem *entity.StatementTypeItem) error
	Update(tx *gorm.DB, statementTypeItem *entity.StatementTypeItem) error
	Delete(tx *gorm.DB, statementTypeItem *entity.StatementTypeItem) error
	FindById(tx *gorm.DB, statementTypeItem *entity.StatementTypeItem) error
	FindByNameStatementTypeId(tx *gorm.DB, statementTypeItem *entity.StatementTypeItem) error
	FindAll(tx *gorm.DB, statementTypeItems *[]entity.StatementTypeItem, statementTypeId uint) error
}

type StatementTypeItemRepositoryImpl struct {
	Repository[entity.StatementTypeItem]
}

func NewStatementTypeItemRepository() StatementTypeItemRepository {
	return &StatementTypeItemRepositoryImpl{}
}

// FindAll implements StatementTypeItemRepository.
func (repository *StatementTypeItemRepositoryImpl) FindAll(tx *gorm.DB, statementTypeItems *[]entity.StatementTypeItem, statementTypeId uint) error {
	return tx.Where("statement_type_id = ? ", statementTypeId).Find(statementTypeItems).Error
}

// FindById implements UserRepository.
func (repository *StatementTypeItemRepositoryImpl) FindById(tx *gorm.DB, statementTypeItem *entity.StatementTypeItem) error {
	return tx.First(statementTypeItem).Error
}

// FindByName implements UserRepository.
func (repository *StatementTypeItemRepositoryImpl) FindByNameStatementTypeId(tx *gorm.DB, statementTypeItem *entity.StatementTypeItem) error {
	return tx.First(statementTypeItem, "name=? and statement_type_id=?", statementTypeItem.Name, statementTypeItem.StatementTypeId).Error
}
