package handlers

import (
	"context"

	"github.com/gin-gonic/gin"
	genproto "github.com/saladin2098/month3/lesson11/voting_gateway/genproto"
)

// CreateElection godoc
// @ID create_election
// @Summary Creates Election
// @Description Ccreate election by reading from body
// @Tags Election
// @Accept json
// @Produce json
// @Param Election body genproto.Election true "election body data"
// @Success 200 {object} genproto.Void
// @Failure 500 {object} string "Failed to get product by id"
// @Router /election/create [POST]
func (h *Handler) CreateElection(c *gin.Context) {
	var election genproto.Election
    if err := c.ShouldBindJSON(&election); err!= nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    res, err := h.Clients.ElectionClient.CreateElection(context.Background(),&election)
    if err!= nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, res)
}
// GetElection godoc
// @ID get_election_by_id
// @Summary Gets Election by id
 // @Description gets election by reading id  from path
// @Tags Election
// @Accept json
// @Produce json
// @Param id path string true "election ID data"
// @Success 200 {object} genproto.Election
// @Failure 500 {object} string "Failed to get product by id"
// @Router /election/{id} [GET]
func (h *Handler) GetElection(c *gin.Context) {
	id := c.Param("id")
	var election genproto.ById
	election.Id = id
    res, err := h.Clients.ElectionClient.GetByIdElection(context.Background(),&election)
    if err!= nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, res)
}
// UpdateElection godoc
// @ID update_election
// @Summary Update Election 
 // @Description Updates election by reading election from body
// @Tags Election
// @Accept json
// @Produce json
// @Param election body genproto.Election true "election data"
// @Success 200 {object} genproto.Void
// @Failure 500 {object} string "Failed to update election data"
// @Router /election/update [PUT]
func (h *Handler) UpdateElection(c *gin.Context) {
	var election genproto.Election
    if err := c.ShouldBindJSON(&election); err!= nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    res, err := h.Clients.ElectionClient.UpdateElection(context.Background(),&election)
    if err!= nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, res)
}
// DeleteElection godoc
// @ID delete_election
// @Summary Deletes Election 
 // @Description Deletes election by reading id
// @Tags Election
// @Accept json
// @Produce json
// @Param id path string true "election ID data"
// @Success 200 {object} string "Deleted successfully"
// @Failure 500 {object} string "Failed to delete election data"
// @Router /election/{id} [DELETE]
func (h  *Handler) DeleteElection(c *gin.Context) {
	id := c.Param("id")
	var election genproto.ById
	election.Id = id
	_,err := h.Clients.ElectionClient.DeleteElection(context.Background(),&election)
	if err!= nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
	c.JSON(200, gin.H{"success": "Deleted successfully"})
}
// GetElectionsByDate godoc
// @ID get_elections_by_date
// @Summary Getss Elections by date 
 // @Description Getss Elections by date 
// @Tags Election
// @Accept json
// @Produce json
// @Param date query string true "election date"
// @Success 200 {object} genproto.GetAllElection
// @Failure 500 {object} string "Failed to Get elections data"
// @Router /election/by-date [GET]
func (h *Handler) GetElectionsByDate(c *gin.Context) {
	var election genproto.Filter
	date := c.Query("date")
	election.Date = date
    res,err := h.Clients.ElectionClient.GetAllElections(context.Background(),&election)
    if err!= nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, res)
}
// GetAllElections godoc
// @ID get_all_elections
// @Summary Gets all Elections
 // @Description Gets all Elections
// @Tags Election
// @Accept json
// @Produce json
// @Success 200 {object} genproto.GetAllElection
// @Failure 500 {object} string "Failed to Get all elections"
// @Router /election/all [GET]
func (h *Handler) GetAllElections(c *gin.Context) {
	var election genproto.Filter
    res,err := h.Clients.ElectionClient.GetAllElections(context.Background(),&election)
    if err!= nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, gin.H{"object": res})
}