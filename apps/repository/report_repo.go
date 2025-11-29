package repository

import (
    "context"
    "BE_PROJECTUAS/apps/models"
    "BE_PROJECTUAS/database"

    "go.mongodb.org/mongo-driver/bson"
)


type reportRepo struct{}

func NewReportRepository() ReportRepository {
    return &reportRepo{}
}

func (r *reportRepo) StatsAll(ctx context.Context) (*models.ReportStatisticsResponse, error) {

    result := models.ReportStatisticsResponse{}

    col := database.MongoDB.Collection("achievements")

    pipeline := []bson.M{
        {"$group": bson.M{
            "_id":   "$achievementType",
            "total": bson.M{"$sum": 1},
        }},
    }

    cursor, err := col.Aggregate(ctx, pipeline)
    if err == nil {
        defer cursor.Close(ctx)
        for cursor.Next(ctx) {
            var doc struct {
                ID    string `bson:"_id"`
                Total int    `bson:"total"`
            }
            cursor.Decode(&doc)
            result.TypeStats = append(result.TypeStats, models.AchievementTypeStat{
                Type:  doc.ID,
                Total: doc.Total,
            })
        }
    }

    rows2, err := database.PostgresDB.QueryContext(ctx, `
        SELECT TO_CHAR(created_at, 'YYYY-MM') AS month, COUNT(*)
        FROM achievement_references
        GROUP BY month
        ORDER BY month;
    `)
    if err == nil {
        defer rows2.Close()
        for rows2.Next() {
            var p models.AchievementPeriodStat
            rows2.Scan(&p.Month, &p.Total)
            result.PeriodStats = append(result.PeriodStats, p)
        }
    }

    rows3, err := database.PostgresDB.QueryContext(ctx, `
        SELECT s.id, u.full_name, COUNT(ar.id) AS total
        FROM students s
        JOIN users u ON u.id = s.user_id
        JOIN achievement_references ar ON ar.student_id = s.id
        GROUP BY s.id, u.full_name
        ORDER BY total DESC
        LIMIT 10;
    `)
    if err == nil {
        defer rows3.Close()
        for rows3.Next() {
            var t models.TopStudentStat
            rows3.Scan(&t.StudentID, &t.FullName, &t.TotalAwards)
            result.TopStudents = append(result.TopStudents, t)
        }
    }

    return &result, nil
}

func (r *reportRepo) StatsByAdvisor(ctx context.Context, lecturerID string) (*models.ReportStatisticsResponse, error) {

    result := models.ReportStatisticsResponse{}

    // Ambil statistik hanya yang student.advisor_id = lecturerID

    // Per bulan
    rows, err := database.PostgresDB.QueryContext(ctx, `
        SELECT TO_CHAR(ar.created_at, 'YYYY-MM'), COUNT(*)
        FROM achievement_references ar
        JOIN students s ON s.id = ar.student_id
        WHERE s.advisor_id = $1
        GROUP BY 1 ORDER BY 1;
    `, lecturerID)
    if err == nil {
        defer rows.Close()
        for rows.Next() {
            var p models.AchievementPeriodStat
            rows.Scan(&p.Month, &p.Total)
            result.PeriodStats = append(result.PeriodStats, p)
        }
    }

    return &result, nil
}

func (r *reportRepo) StatsByStudent(ctx context.Context, studentID string) (*models.ReportStatisticsResponse, error) {

    result := models.ReportStatisticsResponse{}

    // Per bulan
    rows, err := database.PostgresDB.QueryContext(ctx, `
        SELECT TO_CHAR(created_at, 'YYYY-MM'), COUNT(*)
        FROM achievement_references
        WHERE student_id = $1
        GROUP BY 1 ORDER BY 1;
    `, studentID)

    if err == nil {
        defer rows.Close()
        for rows.Next() {
            var p models.AchievementPeriodStat
            rows.Scan(&p.Month, &p.Total)
            result.PeriodStats = append(result.PeriodStats, p)
        }
    }

    return &result, nil
}

func (r *reportRepo) StudentSummary(ctx context.Context, studentID string) (*models.ReportStudentDetail, error) {

    var resp models.ReportStudentDetail

    err := database.PostgresDB.QueryRowContext(ctx, `
        SELECT s.id, u.full_name, COUNT(ar.id)
        FROM students s
        JOIN users u ON u.id = s.user_id
        LEFT JOIN achievement_references ar ON ar.student_id = s.id
        WHERE s.id = $1
        GROUP BY s.id, u.full_name;
    `, studentID).Scan(
        &resp.StudentID,
        &resp.FullName,
        &resp.TotalAchievements,
    )

    if err != nil {
        return nil, err
    }

    return &resp, nil
}
