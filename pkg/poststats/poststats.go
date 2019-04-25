package poststats

import (
	pb "github.com/andreymgn/RSOI-poststats/pkg/poststats/proto"
	"github.com/google/uuid"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	statusNotFound     = status.Error(codes.NotFound, "post not found")
	statusInvalidUUID  = status.Error(codes.InvalidArgument, "invalid UUID")
	statusInvalidToken = status.Errorf(codes.Unauthenticated, "invalid token")
)

func internalError(err error) error {
	return status.Error(codes.Internal, err.Error())
}

// SinglePostStats converts PostStats to SinglePostStats
func (ps *PostStats) SinglePostStats() (*pb.SinglePostStats, error) {
	res := new(pb.SinglePostStats)
	res.PostUid = ps.UID.String()
	res.NumLikes = ps.NumLikes
	res.NumDislikes = ps.NumDislikes
	res.NumViews = ps.NumViews

	return res, nil
}

// GetPostStats returns post stats
func (s *Server) GetPostStats(ctx context.Context, req *pb.GetPostStatsRequest) (*pb.SinglePostStats, error) {
	uid, err := uuid.Parse(req.PostUid)
	if err != nil {
		return nil, statusInvalidUUID
	}

	postStats, err := s.db.get(uid)
	switch err {
	case nil:
		return postStats.SinglePostStats()
	case errNotFound:
		return nil, statusNotFound
	default:
		return nil, internalError(err)
	}
}

// CreatePostStats creates a new post statistics record
func (s *Server) CreatePostStats(ctx context.Context, req *pb.CreatePostStatsRequest) (*pb.SinglePostStats, error) {
	uid, err := uuid.Parse(req.PostUid)
	if err != nil {
		return nil, statusInvalidUUID
	}

	postStats, err := s.db.create(uid)
	if err != nil {
		return nil, internalError(err)
	}

	return postStats.SinglePostStats()
}

// LikePost increases number of post likes
func (s *Server) LikePost(ctx context.Context, req *pb.LikePostRequest) (*pb.LikePostResponse, error) {
	postUID, err := uuid.Parse(req.PostUid)
	if err != nil {
		return nil, statusInvalidUUID
	}

	userUID, err := uuid.Parse(req.UserUid)
	if err != nil {
		return nil, statusInvalidUUID
	}

	success, firstTime, err := s.db.like(postUID, userUID)
	switch err {
	case nil:
		result := new(pb.LikePostResponse)
		result.Success = success
		result.FirstTime = firstTime
		return result, nil
	case errNotFound:
		return nil, statusNotFound
	default:
		return nil, internalError(err)
	}
}

// DislikePost increases number of post dislikes
func (s *Server) DislikePost(ctx context.Context, req *pb.DislikePostRequest) (*pb.DislikePostResponse, error) {
	postUID, err := uuid.Parse(req.PostUid)
	if err != nil {
		return nil, statusInvalidUUID
	}

	userUID, err := uuid.Parse(req.UserUid)
	if err != nil {
		return nil, statusInvalidUUID
	}

	success, firstTime, err := s.db.dislike(postUID, userUID)
	switch err {
	case nil:
		result := new(pb.DislikePostResponse)
		result.Success = success
		result.FirstTime = firstTime
		return result, nil
	case errNotFound:
		return nil, statusNotFound
	default:
		return nil, internalError(err)
	}
}

// IncreaseViews increases number of post views
func (s *Server) IncreaseViews(ctx context.Context, req *pb.IncreaseViewsRequest) (*pb.IncreaseViewsResponse, error) {
	uid, err := uuid.Parse(req.PostUid)
	if err != nil {
		return nil, statusInvalidUUID
	}

	err = s.db.view(uid)
	switch err {
	case nil:
		return new(pb.IncreaseViewsResponse), nil
	case errNotFound:
		return nil, statusNotFound
	default:
		return nil, internalError(err)
	}
}

// DeletePostStats deletes stats of a post
func (s *Server) DeletePostStats(ctx context.Context, req *pb.DeletePostStatsRequest) (*pb.DeletePostStatsResponse, error) {
	uid, err := uuid.Parse(req.PostUid)
	if err != nil {
		return nil, statusInvalidUUID
	}

	err = s.db.delete(uid)
	switch err {
	case nil:
		return new(pb.DeletePostStatsResponse), nil
	case errNotFound:
		return nil, statusNotFound
	default:
		return nil, internalError(err)
	}
}
