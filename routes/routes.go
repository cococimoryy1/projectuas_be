package routes

import (
    "BE_PROJECTUAS/middleware"
    "BE_PROJECTUAS/apps/repository"
    "BE_PROJECTUAS/apps/services"
    "BE_PROJECTUAS/helper"
    "BE_PROJECTUAS/apps/models"

    "github.com/gofiber/fiber/v2"
)

var authSvc *services.AuthService
var userSvc *services.UserService
var achSvc *services.AchievementService

func SetupRoutes(app *fiber.App) {

    api := app.Group("/api/v1")

    // =============================
    // INIT REPOSITORY & SERVICE
    // =============================

    authRepo := repository.NewAuthRepository()
    authSvc = services.NewAuthService(authRepo)

    userRepo := repository.NewUserRepository()
    userSvc = services.NewUserService(userRepo)

    achRepo := repository.NewAchievementRepository()
    achSvc = services.NewAchievementService(achRepo)

    // =============================
    // AUTH ROUTES (Public)
    // =============================
    api.Post("/auth/login",
        wrappers.WrapLogic(authSvc.Login))

    api.Post("/auth/refresh",
        wrappers.WrapRefresh(authSvc.Refresh))

    auth := api.Group("/auth", middleware.AuthRequired())
    auth.Post("/logout", wrappers.WrapLogout(authSvc.Logout))
    auth.Get("/profile", wrappers.WrapProfile(authSvc.Profile))

    // =============================
    // USER ROUTES (Admin Only)
    // =============================
    users := api.Group("/users",
        middleware.AuthRequired(),
        middleware.RequirePermission("user:manage"),
    )

    users.Get("/", wrappers.WrapNoBody(userSvc.List))
    users.Get("/:id", wrappers.WrapParamReturn(userSvc.Get))

    users.Post("/",
        wrappers.WrapLogic[models.CreateUserRequest, string](userSvc.Create))

    users.Put("/:id",
        wrappers.WrapUpdateResp[models.UpdateUserRequest, string](userSvc.Update))

    users.Delete("/:id",
        wrappers.WrapParamResp[string](userSvc.Delete))

    users.Put("/:id/role",
        wrappers.WrapLogicParam[models.UpdateRoleRequest, string](userSvc.UpdateRole))

    // =============================
    // ACHIEVEMENT ROUTES
    // =============================
    achievements := api.Group("/achievements", middleware.AuthRequired())

    // CREATE ACHIEVEMENT (Mahasiswa)
    achievements.Post("/",
        middleware.RequirePermission("achievement:create"),
        wrappers.WrapCreateAchievement(achSvc.Create))

    // SUBMIT
    achievements.Post("/:id/submit",
        middleware.RequirePermission("achievement:submit"),
        wrappers.WrapParam(achSvc.Submit))

    // VERIFY (Dosen Wali)
    achievements.Post("/:id/verify",
        middleware.RequirePermission("achievement:verify"),
        wrappers.WrapParam(achSvc.Verify))

    // REJECT (Dosen Wali)
    achievements.Post("/:id/reject",
        middleware.RequirePermission("achievement:reject"),
        wrappers.WrapReject(achSvc.Reject))

    // MAHASISWA LIST OWN
    achievements.Get("/",
        middleware.RequirePermission("achievement:read_own"),
        wrappers.WrapListStudent(achSvc.ListForStudent))

    // DOSEN WALI LIST ADVISEES
    achievements.Get("/advisor",
        middleware.RequirePermission("achievement:read_advisee"),
        wrappers.WrapListAdvisor(achSvc.ListForAdvisor))

    // UPDATE ACHIEVEMENT (Mahasiswa)
    achievements.Put("/:id",
        middleware.RequirePermission("achievement:update_own"),
        wrappers.WrapLogicParam[models.UpdateAchievementRequest, models.AchievementResponse](achSvc.Update),
    )


    // DELETE ACHIEVEMENT (Mahasiswa soft delete)
    achievements.Delete("/:id",
            middleware.RequirePermission("achievement:delete_own"),
            wrappers.WrapDeleteDraft(achSvc.Delete),
        )
    achievements.Post("/:id/attachments",
    middleware.RequirePermission("achievement:upload"),
    wrappers.WrapUploadAttachment(achSvc.UploadAttachment),
)

    // HISTORY
achievements.Get("/:id/history",
    middleware.RequireAnyPermission(
        "achievement:read_own",
        "achievement:view_advisee",
        "achievement:read_all",
    ),
    wrappers.WrapParamReturn(achSvc.GetHistory),
)

}
