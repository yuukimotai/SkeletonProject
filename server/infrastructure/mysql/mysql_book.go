package database

import (
	"errors"
	"fmt"
	"obserbooks/domain/model"
	"obserbooks/infrastructure/dto"
)

func (d *DB) CreateBook(DTO *dto.BookDTO) (*dto.BookDTO, error) {
	book := model.Book{}
	result := d.connection.Where("user_id = ?", DTO.UserId).Where("jan_code", DTO.JanCode).First(&book).Statement.Model
	if book.JanCode != "" {
		fmt.Println(result)
		return DTO, errors.New("既にその本の情報があります")
	} else {
		d.connection.Create(DTO)
		return DTO, nil
	}
}

func (d *DB) CreateAttentionBook(DTO *dto.BookDTO) (*dto.BookDTO, error) {
	book := model.Book{}
	result := d.connection.Where("user_id = ?", DTO.UserId).Where("jan_code", DTO.JanCode).First(&book).Statement.Model
	if book.JanCode != "" {
		fmt.Println(result)
		return DTO, errors.New("既にその本の情報があります")
	} else {
		d.connection.Create(DTO)
		return DTO, nil
	}
}

func (d *DB) FindBooks(book *dto.BookDTO) []*dto.BookDTO {
	var books []*dto.BookDTO
	d.connection.
		Where("user_id = ?", book.UserId).
		Where("my_book = ?", true).Find(&books)

	return books
}

func (d *DB) FindBook(book dto.BookDTO) (dto.BookDTO, error) {
	err := d.connection.Where(&model.FindBook{}).First(&book).Error
	//fmt.Println(book.UserId)
	if err != nil {
		fmt.Println(err)
	}
	return book, nil
}

func (d *DB) UpdateBook(book dto.BookDTO) (model.Book, error) {
	var target = model.Book{}
	result := d.connection.Where("user_id = ?", book.UserId).Where("jan_code", book.JanCode).First(&target).Statement.Model
	println(result)
	//resultの有無のチェック忘れずに
	targetColumn := []string{}
	targetColumn = updateColumnCheck(book, targetColumn)
	newBox := new(model.Book)
	makeUpdatePropertyList(book, newBox)
	updateResult := d.connection.
		Model(&result).
		Where("user_id = ?", book.UserId).Where("jan_code", book.JanCode).
		First(&target).Select(targetColumn).
		Updates(newBox)
	println(updateResult)
	return target, nil
}

func (d *DB) DeleteBook(book dto.BookDTO) (dto.BookDTO, error) {
	var target = model.Book{}
	deleteResult := d.connection.
		Model(&target).Where("user_id = ?", book.UserId).Where("jan_code", book.JanCode).
		Delete(&target)
	print(deleteResult)
	//fmt.Println(book.UserId)
	if deleteResult == nil {
		fmt.Println(deleteResult) //エラーハンドリングに書き換え必要
	}
	return book, nil
}

// 以下はここ以外では呼び出すつもりがないメソッド
func updateColumnCheck(book dto.BookDTO, targetColumnList []string) []string {
	if book.Title != "" {
		targetColumnList = append(targetColumnList, "Title")
	}
	if book.Author != "" {
		targetColumnList = append(targetColumnList, "Author")
	}
	if book.JanCode != "" {
		targetColumnList = append(targetColumnList, "JanCode")
	}

	return targetColumnList
}

func makeUpdatePropertyList(updateRequest dto.BookDTO, book *model.Book) *model.Book {
	if updateRequest.Title != "" {
		book.Title = updateRequest.Title
	}
	if updateRequest.Author != "" {
		book.Author = updateRequest.Author
	}
	if updateRequest.JanCode != "" {
		book.JanCode = updateRequest.JanCode
	}

	return book
}
