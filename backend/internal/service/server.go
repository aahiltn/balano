package service

import (
	"strconv"

	"palaam/models"

	"github.com/gofiber/fiber/v2"
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

// Services holds all service layer implementations
type Services struct {
	PatientService  PatientServiceInterface
	SessionService  SessionServiceInterface
	StaffService    StaffServiceInterface
	ActivityService ActivityServiceInterface
}

// Service interfaces based on your domain
type PatientServiceInterface interface {
	Create(patient *models.Patient) (*models.Patient, error)
	GetByID(id uint) (*models.Patient, error)
	Update(id uint, patient *models.Patient) (*models.Patient, error)
	Delete(id uint) error
	List(limit, offset int) ([]*models.Patient, int64, error)
	GetSessions(patientID uint, limit, offset int) ([]*models.Session, int64, error)
}

type SessionServiceInterface interface {
	Create(session *models.Session) (*models.Session, error)
	GetByID(id uint) (*models.Session, error)
	Update(id uint, session *models.Session) (*models.Session, error)
	Delete(id uint) error
	List(limit, offset int) ([]*models.Session, int64, error)
	GetDetails(id uint) (*models.SessionDetails, error)
	GetByPatientID(patientID uint, sessionID uint) (*models.Session, error)
}

type StaffServiceInterface interface {
	Create(staff *models.Staff) (*models.Staff, error)
	GetByID(id uint) (*models.Staff, error)
	Update(id uint, staff *models.Staff) (*models.Staff, error)
	Delete(id uint) error
	List(limit, offset int) ([]*models.Staff, int64, error)
	GetSessions(staffID uint, limit, offset int) ([]*models.Session, int64, error)
}

type ActivityServiceInterface interface {
	Create(activity *models.Activity) (*models.Activity, error)
	GetByID(id uint) (*models.Activity, error)
	Update(id uint, activity *models.Activity) (*models.Activity, error)
	Delete(id uint) error
	GetBySessionAndStaff(staffID, sessionID uint, limit, offset int) ([]*models.Activity, int64, error)
	GetSpecific(staffID, sessionID, activityID uint) (*models.Activity, error)
}

// Helper function to convert float32 to uint (since oapi-codegen uses float32 for IDs)
func float32ToUint(f float32) uint {
	return uint(f)
}

// Helper function to parse query parameters
func parseQueryParams(c *fiber.Ctx) (limit, offset int) {
	limit, _ = strconv.Atoi(c.Query("limit", "10"))
	offset, _ = strconv.Atoi(c.Query("offset", "0"))

	if limit <= 0 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}
	if offset < 0 {
		offset = 0
	}

	return limit, offset
}

// Patient handlers
func (s *Server) GetPatients(c *fiber.Ctx, params GetPatientsParams) error {
	limit, offset := parseQueryParams(c)

	patients, total, err := s.services.PatientService.List(limit, offset)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
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
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(createdPatient)
}

func (s *Server) GetPatientsId(c *fiber.Ctx, id float32) error {
	patient, err := s.services.PatientService.GetByID(float32ToUint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(patient)
}

func (s *Server) PutPatientsId(c *fiber.Ctx, id float32) error {
	var patient models.Patient

	if err := c.BodyParser(&patient); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	updatedPatient, err := s.services.PatientService.Update(float32ToUint(id), &patient)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(updatedPatient)
}

func (s *Server) DeletePatientsId(c *fiber.Ctx, id float32) error {
	err := s.services.PatientService.Delete(float32ToUint(id))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusNoContent).Send(nil)
}

func (s *Server) GetPatientsPatientIdSessions(c *fiber.Ctx, patientId float32, params GetPatientsPatientIdSessionsParams) error {
	limit, offset := parseQueryParams(c)

	sessions, total, err := s.services.PatientService.GetSessions(float32ToUint(patientId), limit, offset)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"data":   sessions,
		"total":  total,
		"limit":  limit,
		"offset": offset,
	})
}

func (s *Server) GetPatientsPatientIdSessionsSessionId(c *fiber.Ctx, patientId float32, sessionId float32) error {
	session, err := s.services.SessionService.GetByPatientID(float32ToUint(patientId), float32ToUint(sessionId))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(session)
}

// Session handlers
func (s *Server) GetSessions(c *fiber.Ctx, params GetSessionsParams) error {
	limit, offset := parseQueryParams(c)

	sessions, total, err := s.services.SessionService.List(limit, offset)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
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
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(createdSession)
}

func (s *Server) GetSessionsId(c *fiber.Ctx, id float32) error {
	session, err := s.services.SessionService.GetByID(float32ToUint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(session)
}

func (s *Server) PutSessionsId(c *fiber.Ctx, id float32) error {
	var session models.Session

	if err := c.BodyParser(&session); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	updatedSession, err := s.services.SessionService.Update(float32ToUint(id), &session)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(updatedSession)
}

func (s *Server) DeleteSessionsId(c *fiber.Ctx, id float32) error {
	err := s.services.SessionService.Delete(float32ToUint(id))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusNoContent).Send(nil)
}

func (s *Server) GetSessionsIdDetails(c *fiber.Ctx, id float32) error {
	details, err := s.services.SessionService.GetDetails(float32ToUint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(details)
}

// Staff handlers
func (s *Server) GetStaff(c *fiber.Ctx, params GetStaffParams) error {
	limit, offset := parseQueryParams(c)

	staff, total, err := s.services.StaffService.List(limit, offset)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
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
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(createdStaff)
}

func (s *Server) GetStaffId(c *fiber.Ctx, id float32) error {
	staff, err := s.services.StaffService.GetByID(float32ToUint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(staff)
}

func (s *Server) PutStaffId(c *fiber.Ctx, id float32) error {
	var staff models.Staff

	if err := c.BodyParser(&staff); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	updatedStaff, err := s.services.StaffService.Update(float32ToUint(id), &staff)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(updatedStaff)
}

func (s *Server) DeleteStaffId(c *fiber.Ctx, id float32) error {
	err := s.services.StaffService.Delete(float32ToUint(id))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusNoContent).Send(nil)
}

func (s *Server) GetStaffIdSessions(c *fiber.Ctx, id float32, params GetStaffIdSessionsParams) error {
	limit, offset := parseQueryParams(c)

	sessions, total, err := s.services.StaffService.GetSessions(float32ToUint(id), limit, offset)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"data":   sessions,
		"total":  total,
		"limit":  limit,
		"offset": offset,
	})
}

// Activity handlers
func (s *Server) GetStaffStaffIdSessionsSessionIdActivities(c *fiber.Ctx, staffId float32, sessionId float32, params GetStaffStaffIdSessionsSessionIdActivitiesParams) error {
	limit, offset := parseQueryParams(c)

	activities, total, err := s.services.ActivityService.GetBySessionAndStaff(
		float32ToUint(staffId),
		float32ToUint(sessionId),
		limit,
		offset,
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"data":   activities,
		"total":  total,
		"limit":  limit,
		"offset": offset,
	})
}

func (s *Server) PostStaffStaffIdSessionsSessionIdActivities(c *fiber.Ctx, staffId float32, sessionId float32) error {
	var activity models.Activity

	if err := c.BodyParser(&activity); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Set the staff and session IDs from the URL
	activity.StaffID = float32ToUint(staffId)
	activity.SessionID = float32ToUint(sessionId)

	createdActivity, err := s.services.ActivityService.Create(&activity)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(createdActivity)
}

func (s *Server) GetStaffStaffIdSessionsSessionIdActivitiesId(c *fiber.Ctx, staffId float32, sessionId float32, id float32) error {
	activity, err := s.services.ActivityService.GetSpecific(
		float32ToUint(staffId),
		float32ToUint(sessionId),
		float32ToUint(id),
	)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(activity)
}

func (s *Server) PutStaffStaffIdSessionsSessionIdActivitiesId(c *fiber.Ctx, staffId float32, sessionId float32, id float32) error {
	var activity models.Activity

	if err := c.BodyParser(&activity); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	updatedActivity, err := s.services.ActivityService.Update(float32ToUint(id), &activity)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(updatedActivity)
}

func (s *Server) DeleteStaffStaffIdSessionsSessionIdActivitiesId(c *fiber.Ctx, staffId float32, sessionId float32, id float32) error {
	err := s.services.ActivityService.Delete(float32ToUint(id))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusNoContent).Send(nil)
}
