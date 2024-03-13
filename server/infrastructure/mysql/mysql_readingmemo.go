package database

import (
	"fmt"
	"obserbooks/domain/model"
	"obserbooks/infrastructure/dto"
)

func (d *DB) CreateMemo(DTO *dto.ReadingMemoDTO) (*dto.ReadingMemoDTO, error) {
	// var count int64
	// d.connection.Model(&DTO).
	// 	Where("user_id = ?", DTO.UserId).
	// 	Where("jan_code = ?", DTO.JanCode).
	// 	Count(&count)
	// memoNumber := d.connection.
	// 	Where("user_id = ?", DTO.UserId).
	// 	Where("jan_code = ?", DTO.JanCode).
	// 	Order("memo_number").Last(&DTO)
	// fmt.Println(memoNumber)
	// DTO.MemoNumber = 0
	var memo model.ReadingMemo
	d.connection.
		Where("user_id = ?", DTO.UserId).
		Where("jan_code = ?", DTO.JanCode).
		Order("memo_id DESC").Last(&memo)
	memo.MemoId = DTO.MemoId
	err := d.connection.Create(DTO).Error
	if err != nil {
		fmt.Println("GORMでDBへの保存に失敗しました")
	}

	return DTO, err
}

func (d *DB) GetReadingMemo(DTO *dto.ReadingMemoDTO) (*dto.ReadingMemoDTO, error) {
	var memo *dto.ReadingMemoDTO
	err := d.connection.
		Where("user_id = ?", DTO.UserId).
		Where("jan_code = ?", DTO.JanCode).
		Where("memo_id = ?", DTO.MemoId).
		Find(&memo).Error

	return memo, err
}

func (d *DB) GetAllReadingMemo(DTO *dto.ReadingMemoDTO) ([]*dto.ReadingMemoDTO, error) {
	var memos []*dto.ReadingMemoDTO
	err := d.connection.
		Where("user_id = ?", DTO.UserId).
		Where("jan_code = ?", DTO.JanCode).
		Find(&memos).Error

	return memos, err
}

func (d *DB) UpdateMemo(DTO *dto.ReadingMemoDTO) (*dto.ReadingMemoDTO, error) {
	memomodel := model.ReadingMemo{}
	err := d.connection.Model(&memomodel).
		Where("user_id = ?", DTO.UserId).
		Where("jan_code = ?", DTO.JanCode).
		Where("memo_id = ?", DTO.MemoId).
		Updates(model.ReadingMemo{
			UserId: DTO.UserId, Title: DTO.Title, Author: DTO.Author,
			JanCode: DTO.JanCode, Content: DTO.Content}).Error
	if err != nil {
		fmt.Println("GORMでDB更新に失敗しました")
	}

	return DTO, err
}

func (d *DB) DeleteMemo(DTO *dto.ReadingMemoDTO) error {
	memomodel := model.ReadingMemo{}
	err := d.connection.
		Where("user_id = ?", DTO.UserId).
		Where("jan_code = ?", DTO.JanCode).
		Where("memo_id = ?", DTO.MemoId).
		Delete(&memomodel).Error

	return err
}
