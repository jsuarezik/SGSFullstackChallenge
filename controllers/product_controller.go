package controllers

import (
	"math"
	"net/http"
	"sgs_fullstack_challenge/repositories"
	"sgs_fullstack_challenge/responses"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProductController struct{}

func (ctrl *ProductController) getRepo(c *gin.Context) repositories.ProductRepository {
	db, _ := c.MustGet("db").(*mongo.Database)
	return repositories.NewProductRepository(db)
}

func (ctrl *ProductController) GetAllProducts(c *gin.Context) {
	repo := ctrl.getRepo(c)

	// Pagination
	page := 1
	size := 10
	maxSize := 100

	if val, ok := c.GetQuery("page"); ok {
		pageValue, _ := strconv.Atoi(val)
		if pageValue > 0 {
			page = pageValue
		}
	}

	if val, ok := c.GetQuery("size"); ok {
		pageSizeValue, _ := strconv.Atoi(val)
		if pageSizeValue > 0 {
			if pageSizeValue > maxSize { // Check the maximum size of the page
				pageSizeValue = maxSize
			}
			size = pageSizeValue
		}
	}

	// Sorting
	sortBy := c.DefaultQuery("sortBy", "")           // default is empty (no sorting)
	sortOrderStr := c.DefaultQuery("sortOrder", "1") // default is ascending

	sortOrder, err := strconv.Atoi(sortOrderStr)
	if err != nil || (sortOrder != 1 && sortOrder != -1) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sortOrder value"})
		return
	}

	// Filtering
	query := c.DefaultQuery("q", "")

	products, totalCount, err := repo.GetAll(page, size, sortBy, sortOrder, query)

	totalPages := int(math.Ceil(float64(totalCount) / float64(size)))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	paginationMetaData := responses.PaginationMetadata{
		CurrentPage: page,
		PageSize:    size,
		TotalPages:  totalPages,
		TotalCount:  int(totalCount),
	}
	response := responses.ProductResponse{
		Status:     200,
		Message:    "Success",
		Data:       products,
		Pagination: paginationMetaData,
	}
	c.JSON(http.StatusOK, response)
}
