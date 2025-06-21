package service

import (
	"palaam/internal/models"
	"palaam/internal/repository"
	"palaam/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Server implements the generated ServerInterface
type Server struct {
	services *Services
}

// NewServer creates a new server with the service dependencies
func NewServer(services *Services) ServerInterface {
	return &Server{
		services: services,
	}
}

func InitApp(router fiber.Router, db *gorm.DB) {
	// Initialize repository with DB connection
	repo := repository.NewRepository(db)

	server := NewServer(&Services{
		PatientService:  NewPatientService(repo),
		SessionService:  NewSessionService(repo),
		StaffService:    NewStaffService(repo),
		ActivityService: NewActivityService(repo),
	})
	RegisterHandlers(router, server)
}

// Services holds all service layer implementations
type Services struct {
	PatientService  PatientServiceInterface
	SessionService  SessionServiceInterface
	StaffService    StaffServiceInterface
	ActivityService ActivityServiceInterface
}

/** SESSION HANDLERS **/
func (s *Server) GetSessions(c *fiber.Ctx, params GetSessionsParams) error {
	limit, offset := utils.ParseQueryParams(c)

	sessions, total, err := s.services.SessionService.List(limit, offset)
	if err != nil {
		return s.handleError(c, err, "Failed to fetch sessions")
	}

	return c.JSON(fiber.Map{
		"data":   sessions,
		"total":  total,
		"limit":  limit,
		"offset": offset,
	})
}

func (s *Server) PostSessions(c *fiber.Ctx) error {
	var session models.Session

	if err := c.BodyParser(&session); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	createdSession, err := s.services.SessionService.Create(&session)
	if err != nil {
		return s.handleError(c, err, "Failed to create session")
	}

	return c.Status(fiber.StatusCreated).JSON(createdSession)
}

func (s *Server) GetSessionsId(c *fiber.Ctx, id float32) error {
	sessionID := utils.Float32ToUint(id)
	if sessionID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid session ID",
		})
	}

	session, err := s.services.SessionService.GetByID(sessionID)
	if err != nil {
		return s.handleError(c, err, "Session not found")
	}

	return c.JSON(session)
}

func (s *Server) PutSessionsId(c *fiber.Ctx, id float32) error {
	sessionID := utils.Float32ToUint(id)
	if sessionID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid session ID",
		})
	}

	// Parse the request body into a map for partial updates
	var updates map[string]interface{}
	if err := c.BodyParser(&updates); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	updatedSession, err := s.services.SessionService.Update(sessionID, updates)
	if err != nil {
		return s.handleError(c, err, "Failed to update session")
	}

	return c.JSON(updatedSession)
}

func (s *Server) DeleteSessionsId(c *fiber.Ctx, id float32) error {
	sessionID := utils.Float32ToUint(id)
	if sessionID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid session ID",
		})
	}

	if err := s.services.SessionService.Delete(sessionID); err != nil {
		return s.handleError(c, err, "Failed to delete session")
	}

	return c.Status(fiber.StatusNoContent).Send(nil)
}

func (s *Server) GetSessionsIdDetails(c *fiber.Ctx, id float32) error {
	sessionID := utils.Float32ToUint(id)
	if sessionID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid session ID",
		})
	}

	details, err := s.services.SessionService.GetDetails(sessionID)
	if err != nil {
		return s.handleError(c, err, "Failed to fetch session details")
	}

	return c.JSON(details)
}

/** PATIENT HANDLERS **/
func (s *Server) GetPatients(c *fiber.Ctx, params GetPatientsParams) error {
	limit, offset := utils.ParseQueryParams(c)

	patients, total, err := s.services.PatientService.List(limit, offset)
	if err != nil {
		return s.handleError(c, err, "Failed to fetch patients")
	}

	return c.JSON(fiber.Map{
		"data":   patients,
		"total":  total,
		"limit":  limit,
		"offset": offset,
	})
}

func (s *Server) PostPatients(c *fiber.Ctx) error {
	var patient models.Patient

	if err := c.BodyParser(&patient); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	createdPatient, err := s.services.PatientService.Create(&patient)
	if err != nil {
		return s.handleError(c, err, "Failed to create patient")
	}

	return c.Status(fiber.StatusCreated).JSON(createdPatient)
}

func (s *Server) GetPatientsId(c *fiber.Ctx, id float32) error {
	patientID := utils.Float32ToUint(id)
	if patientID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid patient ID",
		})
	}

	patient, err := s.services.PatientService.GetByID(patientID)
	if err != nil {
		return s.handleError(c, err, "Patient not found")
	}

	return c.JSON(patient)
}

func (s *Server) PutPatientsId(c *fiber.Ctx, id float32) error {
	patientID := utils.Float32ToUint(id)
	if patientID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid patient ID",
		})
	}

	var patient models.Patient
	if err := c.BodyParser(&patient); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	updatedPatient, err := s.services.PatientService.Update(patientID, &patient)
	if err != nil {
		return s.handleError(c, err, "Failed to update patient")
	}

	return c.JSON(updatedPatient)
}

func (s *Server) DeletePatientsId(c *fiber.Ctx, id float32) error {
	patientID := utils.Float32ToUint(id)
	if patientID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid patient ID",
		})
	}

	if err := s.services.PatientService.DeletePatientsId(patientID); err != nil {
		return s.handleError(c, err, "Failed to delete patient")
	}

	return c.Status(fiber.StatusNoContent).Send(nil)
}

func (s *Server) GetPatientsPatientIdSessions(c *fiber.Ctx, patientId float32, params GetPatientsPatientIdSessionsParams) error {
	limit, offset := utils.ParseQueryParams(c)
	patientID := utils.Float32ToUint(patientId)

	if patientID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid patient ID",
		})
	}

	sessions, total, err := s.services.PatientService.GetSessions(patientID, limit, offset)
	if err != nil {
		return s.handleError(c, err, "Failed to fetch patient sessions")
	}

	return c.JSON(fiber.Map{
		"data":   sessions,
		"total":  total,
		"limit":  limit,
		"offset": offset,
	})
}

func (s *Server) GetPatientsPatientIdSessionsSessionId(c *fiber.Ctx, patientId float32, sessionId float32) error {
	patientID := utils.Float32ToUint(patientId)
	sessionID := utils.Float32ToUint(sessionId)

	if patientID == 0 || sessionID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid patient or session ID",
		})
	}

	session, err := s.services.SessionService.GetByPatientID(patientID, sessionID)
	if err != nil {
		return s.handleError(c, err, "Session not found")
	}

	return c.JSON(session)
}

/** STAFF HANDLERS **/
func (s *Server) GetStaff(c *fiber.Ctx, params GetStaffParams) error {
	limit, offset := utils.ParseQueryParams(c)

	staff, total, err := s.services.StaffService.List(limit, offset)
	if err != nil {
		return s.handleError(c, err, "Failed to fetch staff")
	}

	return c.JSON(fiber.Map{
		"data":   staff,
		"total":  total,
		"limit":  limit,
		"offset": offset,
	})
}

func (s *Server) PostStaff(c *fiber.Ctx) error {
	var staff models.Staff

	if err := c.BodyParser(&staff); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	createdStaff, err := s.services.StaffService.Create(&staff)
	if err != nil {
		return s.handleError(c, err, "Failed to create staff")
	}

	return c.Status(fiber.StatusCreated).JSON(createdStaff)
}

func (s *Server) GetStaffId(c *fiber.Ctx, id float32) error {
	staffID := utils.Float32ToUint(id)
	if staffID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid staff ID",
		})
	}

	staff, err := s.services.StaffService.GetByID(staffID)
	if err != nil {
		return s.handleError(c, err, "Staff not found")
	}

	return c.JSON(staff)
}

func (s *Server) PutStaffId(c *fiber.Ctx, id float32) error {
	staffID := utils.Float32ToUint(id)
	if staffID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid staff ID",
		})
	}

	var staff models.Staff
	if err := c.BodyParser(&staff); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	updatedStaff, err := s.services.StaffService.Update(staffID, &staff)
	if err != nil {
		return s.handleError(c, err, "Failed to update staff")
	}

	return c.JSON(updatedStaff)
}

func (s *Server) DeleteStaffId(c *fiber.Ctx, id float32) error {
	staffID := utils.Float32ToUint(id)
	if staffID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid staff ID",
		})
	}

	if err := s.services.StaffService.Delete(staffID); err != nil {
		return s.handleError(c, err, "Failed to delete staff")
	}

	return c.Status(fiber.StatusNoContent).Send(nil)
}

func (s *Server) GetStaffIdSessions(c *fiber.Ctx, id float32, params GetStaffIdSessionsParams) error {
	limit, offset := utils.ParseQueryParams(c)
	staffID := utils.Float32ToUint(id)

	if staffID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid staff ID",
		})
	}

	sessions, total, err := s.services.StaffService.GetSessions(staffID, limit, offset)
	if err != nil {
		return s.handleError(c, err, "Failed to fetch staff sessions")
	}

	return c.JSON(fiber.Map{
		"data":   sessions,
		"total":  total,
		"limit":  limit,
		"offset": offset,
	})
}

/** ACTIVITY HANDLERS **/
func (s *Server) GetStaffStaffIdSessionsSessionIdActivities(c *fiber.Ctx, staffId float32, sessionId float32, params GetStaffStaffIdSessionsSessionIdActivitiesParams) error {
	limit, offset := utils.ParseQueryParams(c)
	staffID := utils.Float32ToUint(staffId)
	sessionID := utils.Float32ToUint(sessionId)

	if staffID == 0 || sessionID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid staff or session ID",
		})
	}

	activities, total, err := s.services.ActivityService.GetBySessionAndStaff(staffID, sessionID, limit, offset)
	if err != nil {
		return s.handleError(c, err, "Failed to fetch activities")
	}

	return c.JSON(fiber.Map{
		"data":   activities,
		"total":  total,
		"limit":  limit,
		"offset": offset,
	})
}

func (s *Server) PostStaffStaffIdSessionsSessionIdActivities(c *fiber.Ctx, staffId float32, sessionId float32) error {
	staffID := utils.Float32ToUint(staffId)
	sessionID := utils.Float32ToUint(sessionId)

	if staffID == 0 || sessionID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid staff or session ID",
		})
	}

	var activity models.Activity
	if err := c.BodyParser(&activity); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Set the staff and session IDs from the URL
	activity.StaffID = staffID
	activity.SessionID = sessionID

	createdActivity, err := s.services.ActivityService.Create(&activity)
	if err != nil {
		return s.handleError(c, err, "Failed to create activity")
	}

	return c.Status(fiber.StatusCreated).JSON(createdActivity)
}

func (s *Server) GetStaffStaffIdSessionsSessionIdActivitiesId(c *fiber.Ctx, staffId float32, sessionId float32, id float32) error {
	staffID := utils.Float32ToUint(staffId)
	sessionID := utils.Float32ToUint(sessionId)
	activityID := utils.Float32ToUint(id)

	if staffID == 0 || sessionID == 0 || activityID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid staff, session, or activity ID",
		})
	}

	activity, err := s.services.ActivityService.GetSpecific(staffID, sessionID, activityID)
	if err != nil {
		return s.handleError(c, err, "Activity not found")
	}

	return c.JSON(activity)
}

func (s *Server) PutStaffStaffIdSessionsSessionIdActivitiesId(c *fiber.Ctx, staffId float32, sessionId float32, id float32) error {
	activityID := utils.Float32ToUint(id)
	if activityID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid activity ID",
		})
	}

	var activity models.Activity
	if err := c.BodyParser(&activity); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	updatedActivity, err := s.services.ActivityService.Update(activityID, &activity)
	if err != nil {
		return s.handleError(c, err, "Failed to update activity")
	}

	return c.JSON(updatedActivity)
}

func (s *Server) DeleteStaffStaffIdSessionsSessionIdActivitiesId(c *fiber.Ctx, staffId float32, sessionId float32, id float32) error {
	activityID := utils.Float32ToUint(id)
	if activityID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid activity ID",
		})
	}

	if err := s.services.ActivityService.Delete(activityID); err != nil {
		return s.handleError(c, err, "Failed to delete activity")
	}

	return c.Status(fiber.StatusNoContent).Send(nil)
}

// Helper method for consistent error handling
func (s *Server) handleError(c *fiber.Ctx, err error, message string) error {
	// Map common business logic errors to appropriate HTTP status codes
	switch err.Error() {
	case "patient not found", "staff member not found", "session not found", "activity not found":
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	case "patient ID is required", "staff ID is required", "invalid session ID", "invalid patient ID", "invalid staff ID":
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	case "staff member has overlapping session at this time", "cannot delete session with existing activities", "cannot delete sessions older than 24 hours":
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"error": err.Error(),
		})
	case "session does not belong to the specified patient":
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": err.Error(),
		})
	default:
		// For unknown errors, return the service error message but log the details
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
}
