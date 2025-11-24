package models
import ("time"

        "go.mongodb.org/mongo-driver/bson/primitive")

type Achievement struct {
    ID                 string `json:"id"`
    StudentID          string `json:"student_id"`
    MongoAchievementID string `json:"mongo_achievement_id"`
    Status             string `json:"status"`
    SubmittedAt        *time.Time `json:"submitted_at"`
    VerifiedAt         *time.Time `json:"verified_at"`
    VerifiedBy         *string `json:"verified_by"`
    RejectionNote      *string `json:"rejection_note"`
    CreatedAt          time.Time `json:"created_at"`
    UpdatedAt          time.Time `json:"updated_at"`
    DeletedAt          *time.Time `json:"deleted_at"`
    DeletedBy          *string `json:"deleted_by"`
}

type AchievementMongo struct {
    ID          primitive.ObjectID       `bson:"_id,omitempty"`
    StudentID   string                   `bson:"studentId"`
    Title       string                   `bson:"title"`
    Description string                   `bson:"description"`
    AchievementType string               `bson:"achievementType"`
    Details     map[string]interface{}   `bson:"details"`
    Attachments []AttachmentMongo        `bson:"attachments"`
    Tags        []string                 `bson:"tags"`
    Points      int                      `bson:"points"`
    CreatedAt   time.Time                `bson:"createdAt"`
    UpdatedAt   time.Time                `bson:"updatedAt"`
}

type AttachmentMongo struct {
    FileName   string    `bson:"fileName"`
    FileUrl    string    `bson:"fileUrl"`
    FileType   string    `bson:"fileType"`
    UploadedAt time.Time `bson:"uploadedAt"`
}


type CreateAchievementRequest struct {
    AchievementType string                 `json:"achievementType"` // 'academic', 'competition', etc. SRS hal.6
    Title           string                 `json:"title"`
    Description     string                 `json:"description"`
    Details         map[string]interface{} `json:"details"` // Dynamic fields SRS hal.6-7
    Tags            []string               `json:"tags"`
}

type AchievementResponse struct {
    ID               string `json:"id"`
    MongoID          string `json:"mongo_id"`
    StudentID        string `json:"student_id"`
    Title            string `json:"title"`
    Description      string `json:"description"`
    Category         string `json:"category"`
    Status           string `json:"status"`
    CreatedAt        string `json:"created_at"`
}

type RejectRequest struct {
    Note string `json:"rejection_note"` // SRS hal.5
}
type CreateAchievementParsed struct {
    Title           string
    Description     string
    AchievementType string
    Details         map[string]interface{}
    Tags            []string
    FilePath        string // file disimpan sebagai path
    FileType        string // mime type
}
