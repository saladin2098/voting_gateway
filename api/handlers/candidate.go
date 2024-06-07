package handlers

import (
	"context"

	"github.com/gin-gonic/gin"
	genproto "github.com/saladin2098/month3/lesson11/voting_gateway/genproto"
)

// CreateCandidate godoc
// @ID create_candidate
// @Summary Creates Candidate
// @Description Creates candidate by reading from body
// @Tags Candidate
// @Accept json
// @Produce json
// @Param Candiate body genproto.CandidateCreate true "candiate body data"
// @Success 200 {object} string "Candiate created successfully"
// @Failure 500 {object} string "Failed to create candidate"
// @Router /candidate/create [POST]
func (h *Handler) CreateCandidate(c *gin.Context) {
	var can genproto.CandidateCreate
	if err := c.ShouldBindJSON(&can); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	_, err := h.Clients.CandidateClient.CreateCandidate(context.Background(),&can)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message":"Candiate created successfully"})
}
// DeleteCandidate godoc
// @ID delete_candidate
// @Summary deletes Candidate
// @Description Deletes candidate by reading id from body
// @Tags Candidate
// @Accept json
// @Produce json
// @Param id path string true "candiate id data"
// @Success 200 {object} string "Candiate deleted successfully"
// @Failure 500 {object} string "Failed to create candidate"
// @Router /candidate/{id} [DELETE]
func (h *Handler) DeleteCandidate(c *gin.Context) {
	var can genproto.ById
    id := c.Param("id")
    can.Id = id
    _, err := h.Clients.CandidateClient.DeleteCandidate(context.Background(),&can)
    if err!= nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, gin.H{"message":"Candiate deleted successfully"})
}
// UpdateCandidate godoc
// @ID updates_candidate
// @Summary updates Candidate
// @Description updates candidate by reading from body
// @Tags Candidate
// @Accept json
// @Produce json
// @Param candiate body  genproto.Candidate true "candiate data"
// @Success 200 {object} string "Candiate updated successfully"
// @Failure 500 {object} string "Failed to create candidate"
// @Router /candidate/update [PUT]
func (h *Handler) UpdateCandidate(c *gin.Context) {
	var can genproto.Candidate
    if err := c.ShouldBindJSON(&can); err!= nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    _, err := h.Clients.CandidateClient.UpdateCandidate(context.Background(),&can)
    if err!= nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, gin.H{"message":"Candiate updated successfully"})
}
// GetByIdCandidate godoc
// @ID get_candidate_by_id
// @Summary gets Candidate by id
// @Description Gets candidate by reading id from body
// @Tags Candidate
// @Accept json
// @Produce json
// @Param candiate path  genproto.ById true "candiate ID data"
// @Success 200 {object} genproto.Candidate
// @Failure 500 {object} string "Failed to get candidate by id"
// @Router /candidate/{id} [GET]
func (h *Handler) GetByIdCandidate(c *gin.Context) {
	var can genproto.ById
    id := c.Param("id")
    can.Id = id
    res, err := h.Clients.CandidateClient.GetByIdCandidate(context.Background(),&can)
    if err!= nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, res)
}
// GetAllCandidates godoc
// @ID get_all_candidate
// @Summary gets all Candidate 
// @Description Gets all candidates
// @Tags Candidate
// @Accept json
// @Produce json
// @Success 200 {object} genproto.GetAllCandidate
// @Failure 500 {object} string "Failed to get all candidates"
// @Router /candidate/all [GET]
func (h *Handler) GetAllCandidates(c *gin.Context) {
	var can genproto.Filter
    res, err := h.Clients.CandidateClient.GetAllCandidates(context.Background(),&can)
    if err!= nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, res)
}
// GetAllCandidatesByParty godoc
// @ID get_candidates_by_party  
// @Summary gets Candidate by the party id
// @Description Gets all candidates
// @Tags Candidate
// @Accept json
// @Produce json
// @Success 200 {object} genproto.GetAllCandidate
// @Failure 500 {object} string "Failed to get candidates by party"
// @Router /candidate/by-party [GET]
func (h *Handler) GetAllCandidatesByParty(c *gin.Context) {
	var can genproto.Filter
	party := c.Query("party")
	can.Party = party
    res, err := h.Clients.CandidateClient.GetAllCandidates(context.Background(),&can)
    if err!= nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, res)
}