package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// TODO: estrutura modelo Price:
//  - ID Clave primaria
//  - ProductId Referencia a un producto
//  - Price Registro del precio modificado
//  - CreatedAt Fecha de creación
//  - UpdatedAt Fecha de última actualización
//  - DeletedAt FEcha de borrado
type Price struct {
	ID  uuid.UUID `gorm:"type:uuid;"`
	ProductId uuid.UUID  `gorm:"type:uuid"`
	Price float32  `json:"price"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
