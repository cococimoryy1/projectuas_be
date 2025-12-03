package mocks

import (
    "context"
    "BE_PROJECTUAS/apps/models"
)

type MockReportRepo struct {
    StatsAllFunc       func(ctx context.Context) (*models.ReportStatisticsResponse, error)
    StatsByAdvisorFunc func(ctx context.Context, advisorID string) (*models.ReportStatisticsResponse, error)
    StatsByStudentFunc func(ctx context.Context, studentID string) (*models.ReportStatisticsResponse, error)
    StudentSummaryFunc func(ctx context.Context, studentID string) (*models.ReportStudentDetail, error)
}

func (m *MockReportRepo) StatsAll(ctx context.Context) (*models.ReportStatisticsResponse, error) {
    return m.StatsAllFunc(ctx)
}
func (m *MockReportRepo) StatsByAdvisor(ctx context.Context, id string) (*models.ReportStatisticsResponse, error) {
    return m.StatsByAdvisorFunc(ctx, id)
}
func (m *MockReportRepo) StatsByStudent(ctx context.Context, id string) (*models.ReportStatisticsResponse, error) {
    return m.StatsByStudentFunc(ctx, id)
}
func (m *MockReportRepo) StudentSummary(ctx context.Context, id string) (*models.ReportStudentDetail, error) {
    return m.StudentSummaryFunc(ctx, id)
}
