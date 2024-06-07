package handlers

import (
	"context"

	"github.com/gin-gonic/gin"
	genprotos "github.com/saladin2098/month3/lesson11/voting_gateway/genproto"
)
// CreatePublicVote godoc
// @ID create_public_vote
// @Summary Creates Public Vote
// @Description Create Public Vote by reading from body
// @Tags Public-Votes
// @Accept json
// @Produce json
// @Param PublicVote body genprotos.PublicVoteCreate true "Public Vote body data"
// @Success 200 {object} string "Created successfully"
// @Failure 500 {object} string "Failed to create public vote"
// @Router /public-vote/create [POST]
func (h *Handler) CreatePublicVote(c *gin.Context) {
	var vote genprotos.PublicVoteCreate
	if err := c.ShouldBindJSON(&vote); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	_, err := h.Clients.PublicVoteClient.CreatePublicVote(context.Background(), &vote)
	if err!= nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
	c.JSON(200, gin.H{"success": "Created successfully"})
}
// GetAllPublicVotesByElection godoc
// @ID get_public_vote_byelection
// @Summary gets all Public Vote by election 
// @Description gets all Public Vote by election 
// @Tags Public-Votes
// @Accept json
// @Produce json
// @Param electionID query string true "election id  data"
// @Success 200 {object} genprotos.GetAllPublicVote
// @Failure 500 {object} string "Failed to get public votes by election"
// @Router /public-vote/{id} [GET]
func (h *Handler) GetAllPublicVotesByElection(c *gin.Context) {
	var vote genprotos.Filter
	election := c.Query("election")
	vote.Election = election
    res, err := h.Clients.PublicVoteClient.GetAllPublucVotes(context.Background(), &vote)
    if err!= nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, res)
}
// GetAllPublicVotes godoc
// @ID get_all_public_votes
// @Summary Gets all Public Votes
// @Description Gets all Public Votes without any filter
// @Tags Public-Votes
// @Accept json
// @Produce json
// @Success 200 {object} genprotos.GetAllPublicVote
// @Failure 500 {object} string "Failed to get public votes"
// @Router /public-vote/all [GET]
func (h *Handler) GetAllPublicVotes(c *gin.Context) {
	var vote genprotos.Filter
    res, err := h.Clients.PublicVoteClient.GetAllPublucVotes(context.Background(), &vote)
    if err!= nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, res)
}
// UpdatePublicVotes godoc
// @ID update_public_votes
// @Summary Updates a Public Vote
// @Description Updates a Public Vote by ID
// @Tags Public-Votes
// @Accept json
// @Produce json
// @Param id query string true "Public Vote ID"
// @Param election query string true "Election ID"
// @Param public query string true "Public ID"
// @Success 200 {object} string "Updated successfully"
// @Failure 500 {object} string "Failed to update public vote"
// @Router /public-vote/update [PUT]
func (h *Handler) UpdatePublicVotes(c *gin.Context) {
	id := c.Query("id")
	election := c.Query("election")
	public := c.Query("public")
	var vote genprotos.PublicVote
	vote.Id = id
	vote.ElectionId = election
	vote.PublicId = public
	_, err := h.Clients.PublicVoteClient.UpdatePublicVote(context.Background(), &vote)
	if err!= nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
	c.JSON(200, gin.H{"success": "Updated successfully"})
}
// DeletePublicVotes godoc
// @ID delete_public_votes
// @Summary Deletes a Public Vote
// @Description Deletes a Public Vote by ID
// @Tags Public-Votes
// @Accept json
// @Produce json
// @Param id query string true "Public Vote ID"
// @Success 200 {object} string "Deleted successfully"
// @Failure 500 {object} string "Failed to delete public vote"
// @Router /public-vote/delete [DELETE]
func (h *Handler) DeletePublicVotes(c *gin.Context) {
	id := c.Query("id")
    var vote genprotos.ById
    vote.Id = id
    _, err := h.Clients.PublicVoteClient.DeletePublicVote(context.Background(), &vote)
    if err!= nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, gin.H{"success": "Deleted successfully"})
}

// GetByIdPublicVotes godoc
// @ID get_by_id_public_votes
// @Summary Gets a Public Vote by ID
// @Description Gets a Public Vote by its ID
// @Tags Public-Votes
// @Accept json
// @Produce json
// @Param id path string true "Public Vote ID"
// @Success 200 {object} genprotos.PublicVote
// @Failure 500 {object} string "Failed to get public vote by ID"
// @Router /public-vote/{id} [GET]
func (h *Handler) GetByIdPublicVotes(c *gin.Context) {
	var vote genprotos.ById
    id := c.Param("id")
    vote.Id = id
    res, err := h.Clients.PublicVoteClient.GetByIdPublicVote(context.Background(), &vote)
    if err!= nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, res)
}
// FindWinner godoc
// @ID fin_the_winner
// @Summary find the winner
// @Description finds the cnadidate withmost botes
// @Tags Public-Votes
// @Accept json
// @Produce json
// @Param election query string true "election ID"
// @Success 200 {object} genprotos.Winner
// @Failure 500 {object} string "Failed to get public vote by ID"
// @Router /public-vote/find-winner [GET]
func (h *Handler) FindWinner(c *gin.Context) {
	var ellection genprotos.WhichElection
	election := c.Query("election")
	ellection.ElectionId = election
	res, err := h.Clients.PublicVoteClient.FindWinner(context.Background(), &ellection)
	if err!= nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
	c.JSON(200, res)
}
