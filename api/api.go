package api

import (
	"github.com/gin-gonic/gin"
	"github.com/saladin2098/month3/lesson11/voting_gateway/api/handlers"
	_ "github.com/saladin2098/month3/lesson11/voting_gateway/docs"
	"github.com/swaggo/files"        
    "github.com/swaggo/gin-swagger"
)

func NewGin(h *handlers.Handler) *gin.Engine {
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	election := r.Group("/election")
	{
		election.POST("/create", h.CreateElection)
		election.GET("/:id", h.GetElection)
		election.PUT("/update",h.UpdateElection)
		election.DELETE("/:id",h.DeleteElection)
		election.GET("/by-date",h.GetElectionsByDate)
		election.GET("/all",h.GetAllElections)
   }
   candidate := r.Group("/candidate")
   {
	    candidate.POST("/create", h.CreateCandidate)
		candidate.GET("/:id", h.GetByIdCandidate)
		candidate.PUT("/update",h.UpdateCandidate)
		candidate.DELETE("/:id",h.DeleteCandidate)
		candidate.GET("/all",h.GetAllCandidates)
		candidate.GET("/by-party",h.GetAllCandidatesByParty)
   }
   public_vote := r.Group("/public-vote")
   {
	    public_vote.POST("/create", h.CreatePublicVote)
		public_vote.GET("/:id", h.GetByIdPublicVotes)
		public_vote.PUT("/update",h.UpdatePublicVotes)
		public_vote.DELETE("/:id",h.DeletePublicVotes)
		public_vote.GET("/all",h.GetAllPublicVotes)
		public_vote.GET("/find-winner",h.FindWinner)
		public_vote.GET("/by-election",h.GetAllPublicVotesByElection)
   }

    return r
}