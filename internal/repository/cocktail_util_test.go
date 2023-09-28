package repository

import (
	"fmt"
	"testing"

	"github.com/marcos-wz/capstone-go-bootcamp/internal/entity"
	"github.com/marcos-wz/capstone-go-bootcamp/internal/repository/mocks"

	"github.com/stretchr/testify/assert"
)

var _ HttpClient = &mocks.HttpClient{}

func TestDrink_GetIngredients(t *testing.T) {
	type fields struct {
		Alcoholic        string
		Category         string
		DateModified     string
		DrinkId          string
		DrinkName        string
		DrinkAlternate   string
		DrinkThumb       string
		Glass            string
		IBA              string
		ImageSource      string
		ImageAttribution string
		Instructions     string
		Ingredient1      string
		Ingredient2      string
		Ingredient3      string
		Ingredient4      string
		Ingredient5      string
		Ingredient6      string
		Ingredient7      string
		Ingredient8      string
		Ingredient9      string
		Ingredient10     string
		Ingredient11     string
		Ingredient12     string
		Ingredient13     string
		Ingredient14     string
		Ingredient15     string
		Measure1         string
		Measure2         string
		Measure3         string
		Measure4         string
		Measure5         string
		Measure6         string
		Measure7         string
		Measure8         string
		Measure9         string
		Measure10        string
		Measure11        string
		Measure12        string
		Measure13        string
		Measure14        string
		Measure15        string
		Tags             string
		Video            string
	}
	tests := []struct {
		name   string
		fields fields
		want   []entity.Ingredient
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := drink{
				Alcoholic:        tt.fields.Alcoholic,
				Category:         tt.fields.Category,
				DateModified:     tt.fields.DateModified,
				DrinkId:          tt.fields.DrinkId,
				DrinkName:        tt.fields.DrinkName,
				DrinkAlternate:   tt.fields.DrinkAlternate,
				DrinkThumb:       tt.fields.DrinkThumb,
				Glass:            tt.fields.Glass,
				IBA:              tt.fields.IBA,
				ImageSource:      tt.fields.ImageSource,
				ImageAttribution: tt.fields.ImageAttribution,
				Instructions:     tt.fields.Instructions,
				Ingredient1:      tt.fields.Ingredient1,
				Ingredient2:      tt.fields.Ingredient2,
				Ingredient3:      tt.fields.Ingredient3,
				Ingredient4:      tt.fields.Ingredient4,
				Ingredient5:      tt.fields.Ingredient5,
				Ingredient6:      tt.fields.Ingredient6,
				Ingredient7:      tt.fields.Ingredient7,
				Ingredient8:      tt.fields.Ingredient8,
				Ingredient9:      tt.fields.Ingredient9,
				Ingredient10:     tt.fields.Ingredient10,
				Ingredient11:     tt.fields.Ingredient11,
				Ingredient12:     tt.fields.Ingredient12,
				Ingredient13:     tt.fields.Ingredient13,
				Ingredient14:     tt.fields.Ingredient14,
				Ingredient15:     tt.fields.Ingredient15,
				Measure1:         tt.fields.Measure1,
				Measure2:         tt.fields.Measure2,
				Measure3:         tt.fields.Measure3,
				Measure4:         tt.fields.Measure4,
				Measure5:         tt.fields.Measure5,
				Measure6:         tt.fields.Measure6,
				Measure7:         tt.fields.Measure7,
				Measure8:         tt.fields.Measure8,
				Measure9:         tt.fields.Measure9,
				Measure10:        tt.fields.Measure10,
				Measure11:        tt.fields.Measure11,
				Measure12:        tt.fields.Measure12,
				Measure13:        tt.fields.Measure13,
				Measure14:        tt.fields.Measure14,
				Measure15:        tt.fields.Measure15,
				Tags:             tt.fields.Tags,
				Video:            tt.fields.Video,
			}
			assert.Equalf(t, tt.want, c.getIngredients(), "getIngredients()")
		})
	}
}

func TestDrink_ParseCocktail(t *testing.T) {
	type fields struct {
		Alcoholic        string
		Category         string
		DateModified     string
		DrinkId          string
		DrinkName        string
		DrinkAlternate   string
		DrinkThumb       string
		Glass            string
		IBA              string
		ImageSource      string
		ImageAttribution string
		Instructions     string
		Ingredient1      string
		Ingredient2      string
		Ingredient3      string
		Ingredient4      string
		Ingredient5      string
		Ingredient6      string
		Ingredient7      string
		Ingredient8      string
		Ingredient9      string
		Ingredient10     string
		Ingredient11     string
		Ingredient12     string
		Ingredient13     string
		Ingredient14     string
		Ingredient15     string
		Measure1         string
		Measure2         string
		Measure3         string
		Measure4         string
		Measure5         string
		Measure6         string
		Measure7         string
		Measure8         string
		Measure9         string
		Measure10        string
		Measure11        string
		Measure12        string
		Measure13        string
		Measure14        string
		Measure15        string
		Tags             string
		Video            string
	}
	tests := []struct {
		name    string
		fields  fields
		want    entity.Cocktail
		wantErr assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := drink{
				Alcoholic:        tt.fields.Alcoholic,
				Category:         tt.fields.Category,
				DateModified:     tt.fields.DateModified,
				DrinkId:          tt.fields.DrinkId,
				DrinkName:        tt.fields.DrinkName,
				DrinkAlternate:   tt.fields.DrinkAlternate,
				DrinkThumb:       tt.fields.DrinkThumb,
				Glass:            tt.fields.Glass,
				IBA:              tt.fields.IBA,
				ImageSource:      tt.fields.ImageSource,
				ImageAttribution: tt.fields.ImageAttribution,
				Instructions:     tt.fields.Instructions,
				Ingredient1:      tt.fields.Ingredient1,
				Ingredient2:      tt.fields.Ingredient2,
				Ingredient3:      tt.fields.Ingredient3,
				Ingredient4:      tt.fields.Ingredient4,
				Ingredient5:      tt.fields.Ingredient5,
				Ingredient6:      tt.fields.Ingredient6,
				Ingredient7:      tt.fields.Ingredient7,
				Ingredient8:      tt.fields.Ingredient8,
				Ingredient9:      tt.fields.Ingredient9,
				Ingredient10:     tt.fields.Ingredient10,
				Ingredient11:     tt.fields.Ingredient11,
				Ingredient12:     tt.fields.Ingredient12,
				Ingredient13:     tt.fields.Ingredient13,
				Ingredient14:     tt.fields.Ingredient14,
				Ingredient15:     tt.fields.Ingredient15,
				Measure1:         tt.fields.Measure1,
				Measure2:         tt.fields.Measure2,
				Measure3:         tt.fields.Measure3,
				Measure4:         tt.fields.Measure4,
				Measure5:         tt.fields.Measure5,
				Measure6:         tt.fields.Measure6,
				Measure7:         tt.fields.Measure7,
				Measure8:         tt.fields.Measure8,
				Measure9:         tt.fields.Measure9,
				Measure10:        tt.fields.Measure10,
				Measure11:        tt.fields.Measure11,
				Measure12:        tt.fields.Measure12,
				Measure13:        tt.fields.Measure13,
				Measure14:        tt.fields.Measure14,
				Measure15:        tt.fields.Measure15,
				Tags:             tt.fields.Tags,
				Video:            tt.fields.Video,
			}
			got, err := c.parseCocktail()
			if !tt.wantErr(t, err, "parseCocktail()") {
				return
			}
			assert.Equalf(t, tt.want, got, "parseCocktail()")
		})
	}
}

func TestCocktailCSVRec_ParseCocktail(t *testing.T) {
	tests := []struct {
		name    string
		record  cocktailCSVRec
		want    entity.Cocktail
		wantErr assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.record.parseCocktail()
			if !tt.wantErr(t, err, "parseCocktail()") {
				return
			}
			assert.Equalf(t, tt.want, got, "parseCocktail()")
		})
	}
}

func TestParseCocktailCsvRec(t *testing.T) {
	type args struct {
		c entity.Cocktail
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseCocktailCsvRec(tt.args.c)
			if !tt.wantErr(t, err, fmt.Sprintf("parseCocktailCsvRec(%v)", tt.args.c)) {
				return
			}
			assert.Equalf(t, tt.want, got, "parseCocktailCsvRec(%v)", tt.args.c)
		})
	}
}
