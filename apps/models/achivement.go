package models

type Achievement struct {
    ID        string `json:"id"`
    StudentID string `json:"student_id"`
    MongoID   string `json:"mongo_id"`
    Status    string `json:"status"` // ENUM: 'draft', 'submitted', 'verified', 'rejected' SRS hal.5
}

type CreateAchievementRequest struct {
    AchievementType string                 `json:"achievementType"` // 'academic', 'competition', etc. SRS hal.6
    Title           string                 `json:"title"`
    Description     string                 `json:"description"`
    Details         map[string]interface{} `json:"details"` // Dynamic fields SRS hal.6-7
    Tags            []string               `json:"tags"`
}

type AchievementResponse struct {
    ID     string `json:"id"`
    Status string `json:"status"`
    // Tambah Mongo details jika fetch
}

type RejectRequest struct {
    Note string `json:"rejection_note"` // SRS hal.5
}
