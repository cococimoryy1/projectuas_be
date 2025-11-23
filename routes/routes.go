package routes

import (
    "BE_PROJECTUAS/middleware"
    "BE_PROJECTUAS/apps/repository"
    "BE_PROJECTUAS/apps/services"
    "BE_PROJECTUAS/utils"

    "BE_PROJECTUAS/apps/models" // Tambah: Untuk type di wrap

    "github.com/gofiber/fiber/v2"
)

var authSvc *services.AuthService
var achSvc *services.AchievementService

func SetupRoutes(app *fiber.App) {
    api := app.Group("/api/v1")

    // =============================
    // INIT REPO & SERVICE
    // =============================
    userRepo := repository.NewUserRepository() // Fix: Hapus /apps/
    authSvc = services.NewAuthService(userRepo)

    achRepo := repository.NewAchievementRepository()
    achSvc = services.NewAchievementService(achRepo)

    // =============================
    // AUTH (Public)
    // =============================
    api.Post("/auth/login", utils.WrapLogic[models.LoginRequest, models.LoginResponse](authSvc.Login)) // Fix: Type match

	// POST /api/v1/auth/refresh
    api.Post("/auth/refresh", utils.WrapRefresh(authSvc.Refresh))

    // POST /api/v1/auth/logout
    auth := api.Group("/auth", middleware.AuthRequired())
    auth.Post("/logout", utils.WrapLogout(authSvc.Logout))

    // GET /api/v1/auth/profile
    auth.Get("/profile", utils.WrapProfile(authSvc.Profile))
    // =============================
    // ACHIEVEMENT ROUTES
    // =============================
    achievements := api.Group("/achievements", middleware.AuthRequired())

    // Create (FR-003)
    achievements.Post("/", middleware.RequirePermission("achievement:create"), utils.WrapWithUser[models.CreateAchievementRequest, models.AchievementResponse](achSvc.Create))

    // Submit (FR-004)
    achievements.Post("/:id/submit", middleware.RequirePermission("achievement:submit"), utils.WrapParam(achSvc.Submit))

    // Verify (FR-007)
    achievements.Post("/:id/verify", middleware.RequirePermission("achievement:verify"), utils.WrapParam(achSvc.Verify))

    // Reject (FR-008)
    achievements.Post("/:id/reject", middleware.RequirePermission("achievement:reject"), utils.WrapReject(achSvc.Reject))

    // List Own (FR-006 adapt)
    achievements.Get("/", middleware.RequirePermission("achievement:read_own"), utils.WrapListOwn(achSvc.ListForStudent))

    // List Advisor (FR-006)
    achievements.Get("/advisor", middleware.RequirePermission("achievement:read_advisee"), utils.WrapListOwn(achSvc.ListForAdvisor)) // Reuse, pass advisorID as userID

    // Update (FR-003)
    achievements.Put("/:id", middleware.RequirePermission("achievement:update"), utils.WrapUpdate[models.CreateAchievementRequest](achSvc.Update))

    // Delete (FR-005)
    achievements.Delete("/:id", middleware.RequirePermission("achievement:delete"), utils.WrapParam(achSvc.Delete))

    // History
    achievements.Get("/:id/history", middleware.RequirePermission("achievement:read"), utils.WrapParam(achSvc.GetHistory))

    // Attachments
    achievements.Post("/:id/attachments", middleware.RequirePermission("achievement:upload"), utils.WrapParam(achSvc.UploadAttachment))
}