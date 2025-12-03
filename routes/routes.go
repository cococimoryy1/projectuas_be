package routes

import (
    "BE_PROJECTUAS/middleware"
    "BE_PROJECTUAS/apps/repository"
    "BE_PROJECTUAS/apps/services"
    "BE_PROJECTUAS/helper"
    "BE_PROJECTUAS/apps/models"

    "github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

    api := app.Group("/api/v1")

    // === INIT REPOSITORIES (PAKAI SATU INSTANSI SAJA) ===
    authRepo := repository.NewAuthRepository()
    userRepo := repository.NewUserRepository()
    studentRepo := repository.NewStudentRepository()
    lecturerRepo := repository.NewLecturerRepository()
    achievementRepo := repository.NewAchievementRepository()
    reportRepo := repository.NewReportRepository()

    // === INIT SERVICES (PAKAI REPO YANG SAMA) ===
    authSvc := services.NewAuthService(authRepo)
    userSvc := services.NewUserService(userRepo)
    achSvc := services.NewAchievementService(achievementRepo)
    studentSvc := services.NewStudentService(studentRepo, authRepo, achievementRepo)
    lecturerSvc := services.NewLecturerService(lecturerRepo, studentRepo)
    reportSvc := services.NewReportService(reportRepo, studentRepo)

    // === AUTH ===
    api.Post("/auth/login", helper.ParseBody[models.LoginRequest](), helper.WrapLogic(authSvc.Login))
    api.Post("/auth/refresh", helper.ParseBody[models.RefreshRequest](), helper.WrapRefresh(authSvc.Refresh))

    auth := api.Group("/auth", middleware.AuthRequired())
    auth.Post("/logout", helper.WrapLogout(authSvc.Logout))
    auth.Get("/profile", helper.WrapProfile(authSvc.Profile))

    // === USERS ===
    users := api.Group("/users", middleware.AuthRequired(), middleware.RequirePermission("user:manage"))
    users.Get("/", helper.WrapNoBody(userSvc.List))
    users.Get("/:id", helper.WrapParamReturn(userSvc.Get))
    users.Post("/", helper.ParseBody[models.CreateUserRequest](), helper.WrapLogic(userSvc.Create))
    users.Put("/:id", helper.ParseBody[models.UpdateUserRequest](), helper.WrapUpdateResp(userSvc.Update))
    users.Delete("/:id", helper.WrapParamResp(userSvc.Delete))
    users.Put("/:id/role", helper.ParseBody[models.UpdateRoleRequest](), helper.WrapLogicParam(userSvc.UpdateRole))

    // === ACHIEVEMENTS ===
    achievements := api.Group("/achievements", middleware.AuthRequired())
    achievements.Post("/", middleware.RequirePermission("achievement:create"), helper.WrapCreateAchievement(achSvc.Create))
    achievements.Post("/:id/submit", middleware.RequirePermission("achievement:submit"), helper.WrapParam(achSvc.Submit))
    achievements.Post("/:id/verify", middleware.RequirePermission("achievement:verify"), helper.WrapParam(achSvc.Verify))
    achievements.Post("/:id/reject", middleware.RequirePermission("achievement:reject"), helper.ParseBody[models.RejectRequest](), helper.WrapReject(achSvc.Reject))
    achievements.Get("/", middleware.RequirePermission("achievement:read_own"), helper.WrapListStudent(achSvc.ListForStudent))
    achievements.Get("/advisor", middleware.RequirePermission("achievement:read_advisee"), helper.WrapListAdvisor(achSvc.ListForAdvisor))
    achievements.Put("/:id", middleware.RequirePermission("achievement:update_own"), helper.ParseBody[models.UpdateAchievementRequest](), helper.WrapLogicParam(achSvc.Update))
    achievements.Delete("/:id", middleware.RequirePermission("achievement:delete_own"), helper.WrapDeleteDraft(achSvc.Delete))
    achievements.Post("/:id/attachments", middleware.RequirePermission("achievement:update_own"), helper.WrapUploadAttachment(achSvc.UploadAttachment))
    achievements.Get("/:id/history", middleware.RequireAnyPermission("achievement:read_own","achievement:view_advisee","achievement:read_all"), helper.WrapParamReturn(achSvc.GetHistory))

    // === STUDENTS ===
    students := api.Group("/students", middleware.AuthRequired(), middleware.RequireAnyPermission("student:read", "student:read_detail", "lecturer:read_advisees"),)
    students.Get("/", middleware.RequirePermission("student:read"),helper.WrapNoBody(studentSvc.List),)
    // DETAIL STUDENT
    students.Get("/:id", middleware.RequirePermission("student:read_detail"), helper.WrapParamReturn(studentSvc.GetByID),)
    students.Get("/:id/achievements", middleware.RequirePermission("student:read_achievements"), helper.WrapParamReturnList(studentSvc.GetStudentAchievements))
    students.Put("/:id/advisor", middleware.RequirePermission("student:update_advisor"), helper.ParseBody[models.UpdateAdvisorRequest](), helper.WrapParamBody(studentSvc.UpdateAdvisor))

    // === LECTURERS ===
    lecturers := api.Group("/lecturers", middleware.AuthRequired(), middleware.RequirePermission("lecturer:read"))
    lecturers.Get("/",middleware.RequirePermission("lecturer:read"), helper.WrapNoBody(lecturerSvc.List))
    lecturers.Get("/:id/advisees", middleware.RequirePermission("lecturer:read_advisees"), helper.WrapParamReturn(lecturerSvc.GetAdvisees))

    // === REPORTS ===
    reports := api.Group("/reports", middleware.AuthRequired())
    reports.Get("/statistics", middleware.RequirePermission("report:view_statistics"), helper.WrapNoBody(reportSvc.Statistics))
    reports.Get("/student/:id", middleware.RequirePermission("report:view_student"), helper.WrapParamReturn(reportSvc.StudentReport))
}

