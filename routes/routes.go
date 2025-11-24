package routes

import (
    "BE_PROJECTUAS/middleware"
    "BE_PROJECTUAS/apps/repository"
    "BE_PROJECTUAS/apps/services"
    "BE_PROJECTUAS/utils"
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

    // AUTH REPO
    authRepo := repository.NewAuthRepository()
    authSvc = services.NewAuthService(authRepo)

    // USER REPO CRUD ADMIN
    userRepo := repository.NewUserRepository()
    userSvc = services.NewUserService(userRepo)

    // ACHIEVEMENT REPO
    achRepo := repository.NewAchievementRepository()
    achSvc = services.NewAchievementService(achRepo)

    // =============================
    // AUTH ROUTES (Public)
    // =============================
    api.Post("/auth/login",
        utils.WrapLogic[models.LoginRequest, models.LoginResponse](authSvc.Login))

    api.Post("/auth/refresh",
        utils.WrapRefresh(authSvc.Refresh))

    auth := api.Group("/auth", middleware.AuthRequired())
    auth.Post("/logout", utils.WrapLogout(authSvc.Logout))
    auth.Get("/profile", utils.WrapProfile(authSvc.Profile))

    // =============================
    // USER ROUTES (Admin Only)
    // =============================
    users := api.Group("/users",
        middleware.AuthRequired(),
        middleware.RequirePermission("user:manage"),
    )

    users.Get("/", utils.WrapNoBody(userSvc.List))
    users.Get("/:id", utils.WrapParamReturn(userSvc.Get))
    users.Post("/", utils.WrapLogic[models.CreateUserRequest, string](userSvc.Create))
    users.Put("/:id", utils.WrapUpdateResp[models.UpdateUserRequest, string](userSvc.Update))
    users.Delete("/:id", utils.WrapParamResp[string](userSvc.Delete))


    users.Put("/:id/role", utils.WrapLogicParam[models.UpdateRoleRequest, string](userSvc.UpdateRole))


    // =============================
    // ACHIEVEMENT ROUTES
    // =============================
    achievements := api.Group("/achievements", middleware.AuthRequired())

    achievements.Post("/",
        middleware.RequirePermission("achievement:create"),
        utils.WrapCreateAchievement(achSvc.Create),
    )

    achievements.Post("/:id/submit",
        middleware.RequirePermission("achievement:submit"),
        utils.WrapParam(achSvc.Submit))

    achievements.Post("/:id/verify",
        middleware.RequirePermission("achievement:verify"),
        utils.WrapParam(achSvc.Verify))

    achievements.Post("/:id/reject",
        middleware.RequirePermission("achievement:reject"),
        utils.WrapReject(achSvc.Reject))

    achievements.Get("/",
        middleware.RequirePermission("achievement:read_own"),
        utils.WrapListOwn(achSvc.ListForStudent))

    achievements.Get("/advisor",
        middleware.RequirePermission("achievement:read_advisee"),
        utils.WrapListOwn(achSvc.ListForAdvisor))

    achievements.Put("/:id",
        middleware.RequirePermission("achievement:update"),
        utils.WrapUpdate[models.CreateAchievementRequest](achSvc.Update))

    achievements.Delete("/:id",
        middleware.RequirePermission("achievement:delete"),
        utils.WrapParam(achSvc.Delete))

    achievements.Get("/:id/history",
        middleware.RequirePermission("achievement:read"),
        utils.WrapParam(achSvc.GetHistory))

    // achievements.Post("/:id/attachments",
    //     middleware.RequirePermission("achievement:upload"),
    //     utils.WrapParam(achSvc.UploadAttachment))
}
