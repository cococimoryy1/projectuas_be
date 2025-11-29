package models

type AchievementTypeStat struct {
    Type  string `json:"type"`
    Total int    `json:"total"`
}

type AchievementPeriodStat struct {
    Month string `json:"month"`
    Total int    `json:"total"`
}

type TopStudentStat struct {
    StudentID   string `json:"student_id"`
    FullName    string `json:"full_name"`
    TotalAwards int    `json:"total_awards"`
}

type CompetitionLevelDist struct {
    Level string `json:"level"`
    Total int    `json:"total"`
}

type ReportStatisticsResponse struct {
    TypeStats       []AchievementTypeStat     `json:"type_stats"`
    PeriodStats     []AchievementPeriodStat   `json:"period_stats"`
    TopStudents     []TopStudentStat          `json:"top_students"`
    LevelDistribution []CompetitionLevelDist  `json:"level_distribution"`
}

type ReportStudentDetail struct {
    StudentID         string `json:"student_id"`
    FullName          string `json:"full_name"`
    TotalAchievements int    `json:"total_achievements"`
}