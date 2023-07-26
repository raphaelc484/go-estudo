package schemas

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Opening struct {
	ID uuid.UUID `gorm:"type:char(36);primaryKey"`
	gorm.Model
	Role     string
	Company  string
	Location string
	Remote   bool
	Link     string
	Salary   int64
}

/***
A estrutura abaixo utilizando os `json:...` pode ser usada na resposta diretamente
no Opening mas afim organizacional ele criou uma outra estrutura para o retorno
de resposta.
O OpeningResponse tem como responsabilidade apenas o retorno da resposta
***/

type OpeningResponse struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt,omitempty"`
	Role      string    `json:"role"`
	Company   string    `json:"company"`
	Location  string    `json:"location"`
	Remote    bool      `json:"remote"`
	Link      string    `json:"link"`
	Salary    int64     `json:"salary"`
}
