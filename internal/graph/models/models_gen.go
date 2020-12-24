// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type Cost struct {
	ID          string    `json:"id"`
	Owner       *User     `json:"owner"`
	Amount      int       `json:"amount"`
	OccurDate   time.Time `json:"occurDate"`
	Category    Category  `json:"category"`
	Description *string   `json:"description"`
	Vote        []*User   `json:"vote"`
}

type CostInput struct {
	Amount      *int       `json:"amount"`
	Date        *time.Time `json:"date"`
	Category    *Category  `json:"category"`
	Description *string    `json:"description"`
}

type Income struct {
	ID          string    `json:"id"`
	Owner       *User     `json:"owner"`
	Amount      int       `json:"amount"`
	OccurDate   time.Time `json:"occurDate"`
	Category    Category  `json:"category"`
	Description *string   `json:"description"`
	Vote        []*User   `json:"vote"`
}

type IncomeInput struct {
	Amount      *int       `json:"amount"`
	Date        *time.Time `json:"date"`
	Category    *Category  `json:"category"`
	Description *string    `json:"description"`
}

// List current or historical portfolio
type Portfolio struct {
	Total  int       `json:"total"`
	Income []*Income `json:"income"`
	Cost   []*Cost   `json:"cost"`
}

type User struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	UserID    string    `json:"userId"`
	NickName  *string   `json:"nickName"`
	CreatedAt time.Time `json:"createdAt"`
	// granted permission to friends to view portfolio
	Friends []*User `json:"friends"`
	// permission to view followers portfolio
	Followers []*User `json:"followers"`
}

type UserInput struct {
	Email    *string `json:"email"`
	UserID   *string `json:"userId"`
	NickName *string `json:"nickName"`
}

type Category string

const (
	CategoryInvestment Category = "INVESTMENT"
	CategorySalory     Category = "SALORY"
	CategoryOthers     Category = "OTHERS"
	CategoryDaily      Category = "DAILY"
	CategoryLearning   Category = "LEARNING"
	CategoryCharity    Category = "CHARITY"
)

var AllCategory = []Category{
	CategoryInvestment,
	CategorySalory,
	CategoryOthers,
	CategoryDaily,
	CategoryLearning,
	CategoryCharity,
}

func (e Category) IsValid() bool {
	switch e {
	case CategoryInvestment, CategorySalory, CategoryOthers, CategoryDaily, CategoryLearning, CategoryCharity:
		return true
	}
	return false
}

func (e Category) String() string {
	return string(e)
}

func (e *Category) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Category(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Category", str)
	}
	return nil
}

func (e Category) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
