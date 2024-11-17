// controllers/products_test.go
package controllers_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"jessie_miniproject/controllers"
	"jessie_miniproject/models"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

// MockDB adalah mock dari DBInterface
type MockDB struct {
	mock.Mock
}

func (m *MockDB) Find(out interface{}, where ...interface{}) *gorm.DB {
	args := m.Called(out, where)
	return args.Get(0).(*gorm.DB)
}

func (m *MockDB) SetFindResult(result interface{}, err error) {
	m.On("Find", mock.Anything, mock.Anything).Run(func(args mock.Arguments) {
		arg := args.Get(0).(*[]models.ProductLog)
		*arg = result.([]models.ProductLog)
	}).Return(&gorm.DB{Error: err})
}

func TestAddProduct_Success(t *testing.T) {
	e := echo.New()
	productPayload := `{
		"product_name": "Sample Product",
		"material": "Plastic",
		"is_plastic": true
	}`
	req := httptest.NewRequest(http.MethodPost, "/products", strings.NewReader(productPayload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Menambahkan mock user ID ke context
	c.Set("user_id", 1)

	// Memanggil controller dengan db mock
	err := controllers.AddProduct(c)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Memastikan respons sesuai yang diharapkan
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "Product added successfully")
}

func TestGetAllProducts_Success(t *testing.T) {
	e := echo.New()

	// Sample data untuk produk
	// products := []models.ProductLog{
	// 	{ID: 1, ProductName: "Product 1", Material: "Material 1", IsPlastic: true},
	// 	{ID: 2, ProductName: "Product 2", Material: "Material 2", IsPlastic: false},
	// }

	// Membuat instance mock DB
	// mockDB := new(MockDB)
	// mockDB.SetFindResult(products, nil) // Set hasil yang dikembalikan untuk Find

	// Membuat konteks test
	req := httptest.NewRequest(http.MethodGet, "/products", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Memanggil controller dengan db mock
	err := controllers.GetAllProducts(c)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Memastikan respons sesuai yang diharapkan
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.NotEmpty(t, rec.Body.String())

	// Memastikan mockDB telah dipanggil
	// mockDB.AssertExpectations(t)
}

func TestGetByID_Success(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/products/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("2")

	// Memanggil controller dengan db mock
	err := controllers.GetByID(c)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Memastikan respons sesuai yang diharapkan
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "Product fetched successfully")
}

func TestAddProductWithAI_Success(t *testing.T) {
	e := echo.New()
	productPayload := `{
		"product_name": "AI Generated Product",
		"material": "Composite",
		"is_plastic": false
	}`
	req := httptest.NewRequest(http.MethodPost, "/products/ai", strings.NewReader(productPayload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Memanggil controller dengan db mock
	err := controllers.AddProductWithAI(c)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// Memastikan respons sesuai yang diharapkan
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "Product added successfully with AI description")
}

// func TestGetAllProducts_DBError(t *testing.T) {
// 	e := echo.New()

// 	// Membuat instance mock DB
// 	mockDB := new(MockDB)
// 	mockDB.SetFindResult(nil, fmt.Errorf("DB error"))

// 	// Membuat konteks test
// 	req := httptest.NewRequest(http.MethodGet, "/products", nil)
// 	rec := httptest.NewRecorder()
// 	c := e.NewContext(req, rec)

// 	// Memanggil controller dengan db mock
// 	err := controllers.GetAllProducts(c)
// 	if err != nil {
// 		t.Fatalf("expected no error, got %v", err)
// 	}

// 	// Memastikan respons sesuai dengan error
// 	assert.Equal(t, http.StatusInternalServerError, rec.Code)
// 	assert.Contains(t, rec.Body.String(), "Could not fetch products")

// 	// Memastikan mockDB telah dipanggil
// 	mockDB.AssertExpectations(t)
// }
