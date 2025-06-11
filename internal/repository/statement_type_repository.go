package repository

import (
	"github.com/Bangdams/kelola-dokumen-BE/internal/entity"
	"gorm.io/gorm"
)

type StatementTypeRepository interface {
	Create(tx *gorm.DB, statementType *entity.StatementType) error
	Update(tx *gorm.DB, statementType *entity.StatementType) error
	Delete(tx *gorm.DB, statementType *entity.StatementType) error
	FindById(tx *gorm.DB, statementType *entity.StatementType) error
	FindByName(tx *gorm.DB, statementType *entity.StatementType) error
	FindAll(tx *gorm.DB, statementType *[]entity.StatementType) error
}

type StatementTypeRepositoryImpl struct {
	Repository[entity.StatementType]
}

func NewStatementTypeRepository() StatementTypeRepository {
	return &StatementTypeRepositoryImpl{}
}

// FindAll implements StatementTypeItemRepository.
func (repository *StatementTypeRepositoryImpl) FindAll(tx *gorm.DB, statementType *[]entity.StatementType) error {
	return tx.Find(statementType).Error
}

// FindById implements UserRepository.
func (repository *StatementTypeRepositoryImpl) FindById(tx *gorm.DB, statementType *entity.StatementType) error {
	return tx.First(statementType).Error
}

// FindByName implements UserRepository.
func (repository *StatementTypeRepositoryImpl) FindByName(tx *gorm.DB, statementType *entity.StatementType) error {
	return tx.First(statementType, "name=?", statementType.Name).Error
}
