package handlers

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	genproto "github.com/saladin2098/month3/lesson11/voting_gateway/genproto"
)

// CreatePublic godoc
// @ID create_public
// @Summary Creates a new Public
// @Description Creates a new Public with the provided JSON data
// @Tags Public
// @Accept json
// @Produce json
// @Param Public body genproto.PublicCreate true "Public data"
// @Success 200 {object} string "Public created successfully"
// @Failure 400 {object} string "Failed to create public"
// @Router /public/create [POST]
func (h *Handler) CreatePublic(c *gin.Context) {
	var public genproto.PublicCreate
	if err := c.ShouldBindJSON(&public); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	_, err := h.Clients.PublicClient.CreatePublic(context.Background(), &public)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "public created successfully"})
}
// GetPublicById godoc
// @ID get_public_by_id
// @Summary Gets Public by ID
// @Description Gets Public by ID
// @Tags Public
// @Accept json
// @Produce json
// @Param id path string true "Public ID"
// @Success 200 {object} genproto.Public
// @Failure 400 {object} string "Failed to get public by ID"
// @Router /public/{id} [GET]
func (h *Handler) GetPublicById(c *gin.Context) {
	var byId genproto.ById
	id := c.Param("id")
	byId.Id = id
	res, err := h.Clients.PublicClient.GetByIdPublic(context.Background(), &byId)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, res)
}
// UpdatePublic godoc
// @ID update_public
// @Summary Updates an existing Public
// @Description Updates an existing Public with the provided JSON data
// @Tags Public
// @Accept json
// @Produce json
// @Param Public body genproto.PublicCreate true "Public data"
// @Success 200 {object} string "Public updated successfully"
// @Failure 400 {object} string "Failed to update public"
// @Router /public/update [PUT]
func (h *Handler) UpdatePublic(c *gin.Context) {
	var public genproto.PublicCreate
	if err := c.ShouldBindJSON(&public); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	_, err := h.Clients.PublicClient.UpdatePublic(context.Background(), &public)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "public updated succesfully"})
}
// DeletePublic godoc
// @ID delete_public
// @Summary Deletes a Public
// @Description Deletes a Public by ID
// @Tags Public
// @Accept json
// @Produce json
// @Param id path string true "Public ID"
// @Success 200 {object} string "Public deleted successfully"
// @Failure 400 {object} string "Failed to delete public"
// @Router /public/{id} [DELETE]
func (h *Handler) DeletePublic(c *gin.Context) {
	var byId genproto.ById
	id := c.Param("id")
	byId.Id = id
	_, err := h.Clients.PublicClient.DeletePublic(context.Background(), &byId)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "public deleted succesfully"})
}
// GetAllPublic godoc
// @ID get_all_public
// @Summary Gets all Publics
// @Description Gets all Publics without any filter
// @Tags Public
// @Accept json
// @Produce json
// @Success 200 {object} genproto.GetAllPublic
// @Failure 400 {object} string "Failed to get all publics"
// @Router /public/all [GET]
func (h *Handler) GetAllPublic(c *gin.Context) {
	var public genproto.Filter
	res, err := h.Clients.PublicClient.GetAllPublics(context.Background(), &public)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, res)
}
// GetPublicFilter godoc
// @ID get_public_filter
// @Summary Gets Publics by filter
// @Description Gets Publics by party, gender, nation, and age
// @Tags Public
// @Accept json
// @Produce json
// @Param party query string true "Party"
// @Param gender query string true "Gender"
// @Param nation query string true "Nation"
// @Param age query int true "Age"
// @Success 200 {object} genproto.GetAllPublic
// @Failure 400 {object} string "Failed to get publics by filter"
// @Router /public/filter [GET]
func (h *Handler) GetPublicFilter(c *gin.Context) {
	party := c.Param("party")
	gender := c.Param("gender")
	nation := c.Param("nation")
	ageStr := c.Param("age")
	age, _ := strconv.Atoi(ageStr)
	var filter genproto.Filter
	filter.Party = party
	filter.Gender = gender
	filter.Nation = nation
	filter.Age = int32(age)
	res, err := h.Clients.PublicClient.GetAllPublics(context.Background(), &filter)
	if err!= nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
	c.JSON(200, res)
}
