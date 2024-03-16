package repositories

import (
	"context"
	"fmt"
	"log"
	repositories "obserbooks/domain/domain-repositories"
	"obserbooks/domain/model"
	"obserbooks/infrastructure/dto"
	database "obserbooks/infrastructure/mysql"
	"os"

	"github.com/99designs/gqlgen/graphql"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type BookRepository struct {
	db       *database.DB
	resolver graphql.Resolver
}

var _ repositories.BookRepository = &BookRepository{}

func NewBookRepository(db *database.DB) *BookRepository {
	return &BookRepository{
		db: db,
	}
}

func (br *BookRepository) CreateBook(ctx context.Context, input model.NewBook) (*dto.BookDTO, error) {
	sub := ParseJwt(ctx)
	bookmodel := &dto.BookDTO{
		UserId:         sub,
		Title:          input.Title,
		Author:         input.Author,
		JanCode:        input.JanCode,
		PublisherName:  input.PublisherName,
		ItemUrl:        input.ItemUrl,
		LargeImageUrl:  input.LargeImageUrl,
		MediumImageUrl: input.MediumImageUrl,
		MyBook:         input.MyBook,
		AttentionBook:  input.AttentionBook,
	}
	result, err := br.db.CreateBook(bookmodel)
	if err != nil {
		return bookmodel, err
	}
	bookmodel = result

	return bookmodel, nil
}

func (br *BookRepository) CreateAttentionBook(ctx context.Context, input model.NewBook) (*dto.BookDTO, error) {
	sub := ParseJwt(ctx)
	bookmodel := &dto.BookDTO{
		UserId:         sub,
		Title:          input.Title,
		Author:         input.Author,
		JanCode:        input.JanCode,
		PublisherName:  input.PublisherName,
		ItemUrl:        input.ItemUrl,
		LargeImageUrl:  input.LargeImageUrl,
		MediumImageUrl: input.MediumImageUrl,
		AttentionBook:  input.AttentionBook,
	}
	result, err := br.db.CreateAttentionBook(bookmodel)
	if err != nil {
		return bookmodel, err
	}
	bookmodel = result

	return bookmodel, nil
}

func (br *BookRepository) FindMyBooks(ctx context.Context) ([]*dto.BookDTO, error) {
	sub := ParseJwt(ctx)
	book := &dto.BookDTO{
		UserId: sub,
	}
	result := br.db.FindBooks(book)
	if result == nil {
		fmt.Println("該当レコードが存在しませんでした")
	}

	return result, nil
}

func (br *BookRepository) FindBook(ctx context.Context, input model.FindBook) (*dto.BookDTO, error) {
	book := dto.BookDTO{
		UserId:  input.UserId,
		JanCode: input.JanCode,
	}

	result, err := br.db.FindBook(book)
	book = result
	if err != nil {
		return &book, err
	}

	return &book, nil
}

func (br *BookRepository) UpdateBook(ctx context.Context, input model.UpdateBook) (*dto.BookDTO, error) {
	book := dto.BookDTO{
		UserId:         input.UserId,
		Title:          input.Title,
		Author:         input.Author,
		JanCode:        input.JanCode,
		ItemUrl:        input.ItemUrl,
		LargeImageUrl:  input.LargeImageUrl,
		MediumImageUrl: input.MediumImageUrl,
		MyBook:         input.MyBook,
		AttentionBook:  input.AttentionBook,
	}

	result, err := br.db.UpdateBook(book)
	fmt.Println(result)
	if err != nil {
		return &book, err
	}

	return &book, nil
}

func (br *BookRepository) DeleteBook(ctx context.Context, input model.DeleteBook) (*dto.BookDTO, error) {
	book := dto.BookDTO{
		UserId:  input.UserId,
		JanCode: input.JanCode,
	}

	result, err := br.db.DeleteBook(book)
	book = result
	if err != nil {
		return &book, err
	}

	return &book, nil
}

func ParseJwt(ctx context.Context) string {
	gc, err := GinContextFromContext(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	tokenString := gc.GetHeader("Authorization")
	if err != nil {
		log.Fatalln(err)
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("Secret")), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// JWTが有効であり、クレームが存在する場合
		sub, ok := claims["sub"].(string)
		if ok == true {
			return sub
		}
	} else {
		fmt.Println("Token is invalid.")
	}
	fmt.Println(err)

	return ""
}

func GinContextFromContext(ctx context.Context) (*gin.Context, error) {
	ginContext := ctx.Value("GinContextKey")
	if ginContext == nil {
		err := fmt.Errorf("could not retrieve gin.Context")
		return nil, err
	}

	gc, ok := ginContext.(*gin.Context)
	if !ok {
		err := fmt.Errorf("gin.Context has wrong type")
		return nil, err
	}
	return gc, nil
}
