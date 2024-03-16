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

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ReadingMemoRepository struct {
	db *database.DB
}

var _ repositories.ReadingMemoRepository = &ReadingMemoRepository{}

func NewReadingMemoRepository(db *database.DB) *ReadingMemoRepository {
	return &ReadingMemoRepository{
		db: db,
	}
}

func (mr *ReadingMemoRepository) CreateReadingMemo(ctx context.Context, input model.NewReadingMemo) (*model.ReadingMemo, error) {
	sub := ParseJwt(ctx)
	memodto := &dto.ReadingMemoDTO{
		MemoId:  uuid.New().String(),
		UserId:  sub,
		JanCode: input.JanCode,
		Title:   input.Title,
		Author:  input.Author,
		Content: input.Content,
	}
	dto, err := mr.db.CreateMemo(memodto)
	memomodel := &model.ReadingMemo{
		UserId:  dto.UserId,
		Title:   dto.Title,
		Author:  dto.Author,
		Content: dto.Content,
	}

	return memomodel, err
}

func (mr *ReadingMemoRepository) GetReadingMemo(ctx context.Context, input model.GetMemo) (*model.ReadingMemo, error) {
	sub := ParseJwt(ctx)
	memodto := &dto.ReadingMemoDTO{
		MemoId:  input.MemoId,
		UserId:  sub,
		JanCode: input.JanCode,
	}
	dto, err := mr.db.GetReadingMemo(memodto)

	memo := &model.ReadingMemo{
		MemoId:  dto.MemoId,
		UserId:  dto.UserId,
		Title:   dto.Title,
		Author:  dto.Author,
		Content: dto.Content,
	}

	return memo, err
}

func (mr *ReadingMemoRepository) GetAllReadingMemo(ctx context.Context, input model.GetAllMemo) ([]*model.ReadingMemo, error) {
	sub := ParseJwt(ctx)
	memodto := &dto.ReadingMemoDTO{
		UserId:  sub,
		JanCode: input.JanCode,
	}
	dto, err := mr.db.GetAllReadingMemo(memodto)
	readingmemos := []*model.ReadingMemo{}
	for i := 0; i < len(dto); i++ {
		memo := &model.ReadingMemo{
			JanCode: dto[i].JanCode,
			MemoId:  dto[i].MemoId,
			UserId:  dto[i].UserId,
			Title:   dto[i].Title,
			Author:  dto[i].Author,
			Content: dto[i].Content,
		}
		readingmemos = append(readingmemos, memo)
	}

	return readingmemos, err
}

func (mr *ReadingMemoRepository) UpdateReadingMemo(ctx context.Context, input model.UpdateReadingMemo) (*model.ReadingMemo, error) {
	sub := ParseJwt(ctx)
	memodto := &dto.ReadingMemoDTO{
		UserId:  sub,
		JanCode: input.JanCode,
		MemoId:  input.MemoId,
		Title:   input.Title,
		Author:  input.Author,
		Content: input.Content,
	}
	dto, err := mr.db.UpdateMemo(memodto)
	memomodel := &model.ReadingMemo{
		UserId:  dto.UserId,
		Title:   dto.Title,
		Author:  dto.Author,
		Content: dto.Content,
	}

	return memomodel, err
}

func (mr *ReadingMemoRepository) DeleteReadingMemo(ctx context.Context, input model.DeleteReadingMemo) (int, error) {
	sub := ParseJwt(ctx)
	memodto := &dto.ReadingMemoDTO{
		MemoId:  input.MemoId,
		UserId:  sub,
		JanCode: input.JanCode,
	}
	err := mr.db.DeleteMemo(memodto)
	if err != nil {
		fmt.Println("メモの削除に失敗しました")
	}

	return 204, nil
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
