package v1

import (
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/pkg/errors"

	"github.com/bytebase/bytebase/backend/common"
	"github.com/bytebase/bytebase/backend/store"
	v1pb "github.com/bytebase/bytebase/proto/generated-go/v1"
)

func (s *DatabaseService) ListRevisions(ctx context.Context, request *v1pb.ListRevisionsRequest) (*v1pb.ListRevisionsResponse, error) {
	instanceID, databaseName, err := common.GetInstanceDatabaseID(request.Parent)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to get instance and database from %v, err: %v", request.Parent, err)
	}

	instance, err := s.store.GetInstanceV2(ctx, &store.FindInstanceMessage{ResourceID: &instanceID})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to find instance %v, err: %v", instanceID, err)
	}
	if instance == nil {
		return nil, status.Errorf(codes.NotFound, "instance %v not found", instanceID)
	}

	database, err := s.store.GetDatabaseV2(ctx, &store.FindDatabaseMessage{
		InstanceID:   &instanceID,
		DatabaseName: &databaseName,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to find database %v, err: %v", request.Parent, err)
	}
	if database == nil {
		return nil, status.Errorf(codes.NotFound, "database %v not found", request.Parent)
	}

	limit, offset, err := parseLimitAndOffset(request.PageToken, int(request.PageSize))
	if err != nil {
		return nil, err
	}
	limitPlusOne := limit + 1

	find := &store.FindRevisionMessage{
		DatabaseUID: database.UID,
		Limit:       &limitPlusOne,
		Offset:      &offset,
	}

	revisions, err := s.store.ListRevisions(ctx, find)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to find revisions, err: %v", err)
	}

	var nextPageToken string
	if len(revisions) == limitPlusOne {
		pageToken, err := getPageToken(limit, offset+limit)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to get next page token, error: %v", err)
		}
		nextPageToken = pageToken
		revisions = revisions[:limit]
	}

	converted, err := convertToRevisions(ctx, s.store, request.Parent, revisions)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to convert to revisions, err: %v", err)
	}

	return &v1pb.ListRevisionsResponse{
		Revisions:     converted,
		NextPageToken: nextPageToken,
	}, nil
}

func convertToRevisions(ctx context.Context, s *store.Store, parent string, revisions []*store.RevisionMessage) ([]*v1pb.Revision, error) {
	var rs []*v1pb.Revision
	for _, revision := range revisions {
		r, err := convertToRevision(ctx, s, parent, revision)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to convert to revisions")
		}
		rs = append(rs, r)
	}
	return rs, nil
}

func convertToRevision(ctx context.Context, s *store.Store, parent string, revision *store.RevisionMessage) (*v1pb.Revision, error) {
	creator, err := s.GetUserByID(ctx, revision.CreatorUID)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get creator")
	}
	_, sheetUID, err := common.GetProjectResourceIDSheetUID(revision.Payload.Sheet)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get sheetUID from %q", revision.Payload.Sheet)
	}
	sheet, err := s.GetSheet(ctx, &store.FindSheetMessage{UID: &sheetUID})
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get sheet %q", revision.Payload.Sheet)
	}
	if sheet == nil {
		return nil, errors.Errorf("sheet %q not found", revision.Payload.Sheet)
	}

	taskRunName, issueName := revision.Payload.TaskRun, ""
	if taskRunName != "" {
		_, rolloutUID, _, _, _, err := common.GetProjectIDRolloutIDStageIDTaskIDTaskRunID(taskRunName)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to get rollout UID from %q", taskRunName)
		}
		issue, err := s.GetIssueV2(ctx, &store.FindIssueMessage{PipelineID: &rolloutUID})
		if err != nil {
			return nil, errors.Wrapf(err, "failed to get issue by rollout %q", rolloutUID)
		}
		if issue == nil {
			return nil, errors.Errorf("issue not found by rollout %q", rolloutUID)
		}
		issueName = fmt.Sprintf("%s%s/%s%d", common.ProjectNamePrefix, issue.Project.ResourceID, common.IssueNamePrefix, issue.UID)
	}

	return &v1pb.Revision{
		Name:          fmt.Sprintf("%s/%s%d", parent, common.RevisionNamePrefix, revision.UID),
		Release:       revision.Payload.Release,
		CreateTime:    timestamppb.New(revision.CreatedTime),
		Creator:       common.FormatUserEmail(creator.Email),
		Sheet:         revision.Payload.Sheet,
		SheetSha256:   revision.Payload.SheetSha256,
		Statement:     sheet.Statement,
		StatementSize: sheet.Size,
		Type:          v1pb.ReleaseFileType(revision.Payload.Type),
		Version:       revision.Payload.Version,
		File:          revision.Payload.File,
		Issue:         issueName,
		TaskRun:       taskRunName,
	}, nil
}