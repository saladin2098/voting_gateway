package handlers

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"

	genproto "github.com/saladin2098/month3/lesson11/voting_gateway/genproto"
)

// CreateParty godoc
// @ID create_party
// @Summary Creates Party
// @Description Create party by reading from body
// @Tags Party
// @Accept json
// @Produce json
// @Param Party body genproto.Party true "party body data"
// @Success 200 {object} string "party created successfully"
// @Failure 500 {object} string "Failed to create party"
// @Router /party/create [POST]
func (h *Handler) CreateParty(c *gin.Context) {
	var party genproto.Party
	if err := c.ShouldBindJSON(&party); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	_, err := h.Clients.PartyClient.CreateParty(context.Background(), &party)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "party created successfully"})
}

// GetPartyByID godoc
// @ID get_party_by_id
// @Summary Gets Party by ID
// @Description Gets Party by ID
// @Tags Party
// @Accept json
// @Produce json
// @Param id path string true "Party ID data"
// @Success 200 {object} genproto.Party
// @Failure 500 {object} string "Failed to get party by ID"
// @Router /party/{id} [GET]
func (h *Handler) GetPartyByID(c *gin.Context) {
	var party genproto.ById
    id := c.Param("id")
    party.Id = id
	log.Println(id)
    res, err := h.Clients.PartyClient.GetByIdParty(context.Background(), &party)
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, res)
}
// GetAllPartiesBydate godoc
// @ID get_party_by_date_and_slogan
// @Summary Gets Party by date and slogan
// @Description Gets Party by date and slogan
// @Tags Party
// @Accept json
// @Produce json
// @Param date query string true "Party open date"
// @Param slogan query string true "Party slogan"
// @Success 200 {object} genproto.GetAllParty
// @Failure 500 {object} string "Failed to get party by ID"
// @Router /party/date-slogan [GET]
func (h *Handler) GetAllPartiesBydate(c *gin.Context) {
	var party genproto.Filter
	date := c.Query("date")
	slogan := c.Query("slogan")
	party.Date = date
	party.Slogan = slogan
    res,err := h.Clients.PartyClient.GetAllPartys(context.Background(),&party)
    if err!= nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, gin.H{"object": res})
}

// GetAllParties godoc
// @ID get_all_parties
// @Summary Gets all Parties
// @Description Gets all Parties
// @Tags Party
// @Accept json
// @Produce json
// @Success 200 {object} genproto.GetAllParty
// @Failure 500 {object} string "Failed to get party by ID"
// @Router /party/all [GET]
func (h *Handler) GetAllParties(c *gin.Context) {
	var party genproto.Filter
    res,err := h.Clients.PartyClient.GetAllPartys(context.Background(),&party)
    if err!= nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, gin.H{"object": res})
}
// UpdateParty godoc
// @ID update_party
// @Summary Updates a Party
// @Description Updates a Party with the provided JSON data
// @Tags Party
// @Accept json
// @Produce json
// @Param Party body genproto.Party true "Party data"
// @Success 200 {object} string "Party updated successfully"
// @Failure 400 {object} string "Failed to update party"
// @Router /party/update [PUT]
func (h *Handler) UpdateParty(c *gin.Context) {
	var party genproto.Party
    if err := c.ShouldBindJSON(&party); err!= nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    _, err := h.Clients.PartyClient.UpdateParty(context.Background(), &party)
    if err!= nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, gin.H{"message":"Party updated successfully"})
}

// DeleteParty godoc
// @ID delete_party
// @Summary Deletes a Party
// @Description Deletes a Party by ID
// @Tags Party
// @Accept json
// @Produce json
// @Param id path string true "Party ID"
// @Success 200 {object} string "Party deleted successfully"
// @Failure 400 {object} string "Failed to delete party"
// @Router /party/{id} [DELETE]
func (h *Handler) Deleteparty(c *gin.Context) {
	var byId genproto.ById
	id := c.Param("id")
	byId.Id = id
	_,err := h.Clients.PartyClient.DeleteParty(context.Background(),&byId)
	if err!= nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, gin.H{"message":"Party deleted successfully"})
}