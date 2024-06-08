package servers

import (
	"log/slog"

	pb "github.com/saladin2098/month3/lesson11/voting_gateway/genproto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)
type Clients struct {
	CandidateClient pb.CandidateServiceClient
	ElectionClient pb.ElectionServiceClient
	PublicVoteClient pb.PublicVoteServiceClient
	PublicClient pb.PublicServiceClient
	PartyClient pb.PartyServiceClient
}

func NewClients() *Clients {
	conn,err := grpc.NewClient("localhost:8080",grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err!=nil{
        slog.Error("error:",err)
    }
	conn2,err := grpc.NewClient("localhost:8083",grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err!=nil{
        slog.Error("error:",err)
    }
	clientS := pb.NewCandidateServiceClient(conn)
	clientE := pb.NewElectionServiceClient(conn)
	publicVoteS := pb.NewPublicVoteServiceClient(conn)
	publicS := pb.NewPublicServiceClient(conn2)
	partyS := pb.NewPartyServiceClient(conn2)
	return &Clients{
		CandidateClient: clientS,
        ElectionClient: clientE,
        PublicVoteClient: publicVoteS,
        PublicClient: publicS,
        PartyClient: partyS,
	}
}