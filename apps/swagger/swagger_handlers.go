package swagger

import "github.com/gofiber/fiber/v2"



//
// =======================
//  AUTH
// =======================

// Login User
// @Summary Login User
// @Description Mengautentikasi user dan menghasilkan JWT access token + refresh token
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body models.LoginRequest true "Login Credentials"
// @Success 200 {object} models.LoginResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /auth/login [post]
func SwaggerAuthLogin(c *fiber.Ctx) error { return nil }

// Refresh Token
// @Summary Refresh Token
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body models.RefreshRequest true "Refresh Token Request"
// @Success 200 {object} models.LoginResponse
// @Router /auth/refresh [post]
func SwaggerAuthRefresh(c *fiber.Ctx) error { return nil }

// Profile
// @Summary Profile Pengguna
// @Tags Auth
// @Security BearerAuth
// @Success 200 {object} models.UserResponse
// @Router /auth/profile [get]
func SwaggerAuthProfile(c *fiber.Ctx) error { return nil }

// Logout
// @Summary Logout User
// @Tags Auth
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Router /auth/logout [post]
func SwaggerAuthLogout(c *fiber.Ctx) error { return nil }

//
// =======================
//  USERS
// =======================

// @Summary List User
// @Tags Users
// @Security BearerAuth
// @Success 200 {array} models.User
// @Router /users [get]
func SwaggerUsersList(c *fiber.Ctx) error { return nil }

// @Summary Detail User
// @Tags Users
// @Security BearerAuth
// @Param id path string true "User ID"
// @Success 200 {object} models.User
// @Router /users/{id} [get]
func SwaggerUsersGet(c *fiber.Ctx) error { return nil }

// @Summary Create User
// @Tags Users
// @Security BearerAuth
// @Accept json
// @Param request body models.CreateUserRequest true "New User Data"
// @Success 201 {object} map[string]interface{}
// @Router /users [post]
func SwaggerUsersCreate(c *fiber.Ctx) error { return nil }

// @Summary Update User
// @Tags Users
// @Security BearerAuth
// @Param id path string true "User ID"
// @Param request body models.UpdateUserRequest true "Update Data"
// @Success 200 {object} map[string]interface{}
// @Router /users/{id} [put]
func SwaggerUsersUpdate(c *fiber.Ctx) error { return nil }

// @Summary Delete User
// @Tags Users
// @Security BearerAuth
// @Param id path string true "User ID"
// @Success 200 {object} map[string]interface{}
// @Router /users/{id} [delete]
func SwaggerUsersDelete(c *fiber.Ctx) error { return nil }

// @Summary Update User Role
// @Tags Users
// @Security BearerAuth
// @Param id path string true "User ID"
// @Param request body models.UpdateRoleRequest true "Role Payload"
// @Success 200 {object} map[string]interface{}
// @Router /users/{id}/role [put]
func SwaggerUsersRole(c *fiber.Ctx) error { return nil }

//
// =======================
//  ACHIEVEMENTS
// =======================

// @Summary Create Achievement (Draft)
// @Tags Achievements
// @Security BearerAuth
// @Accept json
// @Param request body models.CreateAchievementParsed true "Achievement draft"
// @Success 201 {object} models.AchievementResponse
// @Router /achievements [post]
func SwaggerAchCreate(c *fiber.Ctx) error { return nil }

// @Summary Submit Achievement
// @Tags Achievements
// @Security BearerAuth
// @Param id path string true "Achievement ID"
// @Router /achievements/{id}/submit [post]
func SwaggerAchSubmit(c *fiber.Ctx) error { return nil }

// @Summary Verify Achievement
// @Tags Achievements
// @Security BearerAuth
// @Param id path string true "Achievement ID"
// @Router /achievements/{id}/verify [post]
func SwaggerAchVerify(c *fiber.Ctx) error { return nil }

// @Summary Reject Achievement
// @Tags Achievements
// @Security BearerAuth
// @Param id path string true "Achievement ID"
// @Param request body models.RejectRequest true "Rejection note"
// @Router /achievements/{id}/reject [post]
func SwaggerAchReject(c *fiber.Ctx) error { return nil }

// @Summary List Student Achievements
// @Tags Achievements
// @Security BearerAuth
// @Success 200 {array} models.AchievementResponse
// @Router /achievements [get]
func SwaggerAchListStudent(c *fiber.Ctx) error { return nil }

// @Summary List Advisee Achievements
// @Tags Achievements
// @Security BearerAuth
// @Success 200 {array} models.AchievementResponse
// @Router /achievements/advisor [get]
func SwaggerAchListAdvisor(c *fiber.Ctx) error { return nil }

// @Summary Update Achievement Draft
// @Tags Achievements
// @Security BearerAuth
// @Param id path string true "Achievement ID"
// @Param request body models.UpdateAchievementRequest true "Update Draft"
// @Router /achievements/{id} [put]
func SwaggerAchUpdate(c *fiber.Ctx) error { return nil }

// @Summary Delete Achievement Draft
// @Tags Achievements
// @Security BearerAuth
// @Param id path string true "Achievement ID"
// @Router /achievements/{id} [delete]
func SwaggerAchDelete(c *fiber.Ctx) error { return nil }

// @Summary Upload Attachment
// @Tags Achievements
// @Security BearerAuth
// @Param id path string true "Achievement ID"
// @Router /achievements/{id}/attachments [post]
func SwaggerAchAttachment(c *fiber.Ctx) error { return nil }

// @Summary Achievement History
// @Tags Achievements
// @Security BearerAuth
// @Param id path string true "Achievement ID"
// @Router /achievements/{id}/history [get]
func SwaggerAchHistory(c *fiber.Ctx) error { return nil }

//
// =======================
//  STUDENTS
// =======================

// @Summary List Students
// @Tags Students
// @Security BearerAuth
// @Success 200 {array} models.StudentListResponse
// @Router /students [get]
func SwaggerStudentList(c *fiber.Ctx) error { return nil }

// @Summary Detail Mahasiswa
// @Tags Students
// @Security BearerAuth
// @Param id path string true "Student ID"
// @Success 200 {object} models.StudentDetailResponse
// @Router /students/{id} [get]
func SwaggerStudentDetail(c *fiber.Ctx) error { return nil }

// @Summary Student Achievements
// @Tags Students
// @Security BearerAuth
// @Param id path string true "Student ID"
// @Success 200 {array} models.AchievementResponse
// @Router /students/{id}/achievements [get]
func SwaggerStudentAchievements(c *fiber.Ctx) error { return nil }

// @Summary Update Advisor
// @Tags Students
// @Security BearerAuth
// @Param id path string true "Student ID"
// @Param request body models.UpdateAdvisorRequest true "Payload"
// @Router /students/{id}/advisor [put]
func SwaggerStudentUpdateAdvisor(c *fiber.Ctx) error { return nil }

//
// =======================
//  LECTURERS
// =======================

// @Summary List Lecturers
// @Tags Lecturers
// @Security BearerAuth
// @Success 200 {array} models.LecturerListResponse
// @Router /lecturers [get]
func SwaggerLectList(c *fiber.Ctx) error { return nil }

// @Summary Lecturer Advisees
// @Tags Lecturers
// @Security BearerAuth
// @Param id path string true "Lecturer ID"
// @Success 200 {array} models.StudentListResponse
// @Router /lecturers/{id}/advisees [get]
func SwaggerLectAdvisees(c *fiber.Ctx) error { return nil }

//
// =======================
//  REPORTS
// =======================

// @Summary Statistik Global
// @Tags Reports
// @Security BearerAuth
// @Success 200 {object} models.ReportStatisticsResponse
// @Router /reports/statistics [get]
func SwaggerReportStatistics(c *fiber.Ctx) error { return nil }

// @Summary Report Per Mahasiswa
// @Tags Reports
// @Security BearerAuth
// @Param id path string true "Student ID"
// @Success 200 {object} models.ReportStudentDetail
// @Router /reports/student/{id} [get]
func SwaggerReportStudent(c *fiber.Ctx) error { return nil }
