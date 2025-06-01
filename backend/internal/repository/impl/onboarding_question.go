package impl

// backend/internal/repository/impl/onboarding_question.go

import (
	"palaam/internal/models"

	"gorm.io/gorm"
)

type OnboardingQuestionRepository struct {
	db *gorm.DB
}

func NewOnboardingQuestionRepository(db *gorm.DB) *OnboardingQuestionRepository {
	return &OnboardingQuestionRepository{db: db}
}

// Create a new onboarding question
func (r *OnboardingQuestionRepository) Create(question *models.OnboardingQuestion) error {
	return r.db.Create(question).Error
}

// Find an onboarding question by text
func (r *OnboardingQuestionRepository) FindByText(text string) (*models.OnboardingQuestion, error) {
	var question models.OnboardingQuestion
	if err := r.db.First(&question, "text = ?", text).Error; err != nil {
		return nil, err
	}
	return &question, nil
}

// Find onboarding questions by assessment ID
func (r *OnboardingQuestionRepository) FindByAssessmentID(assessmentID int) ([]*models.OnboardingQuestion, error) {
	var questions []*models.OnboardingQuestion
	if err := r.db.Where("assessment_id = ?", assessmentID).Find(&questions).Error; err != nil {
		return nil, err
	}
	return questions, nil
}

// Get all onboarding questions
func (r *OnboardingQuestionRepository) GetAll() ([]*models.OnboardingQuestion, error) {
	var questions []*models.OnboardingQuestion
	if err := r.db.Find(&questions).Error; err != nil {
		return nil, err
	}
	return questions, nil
}

// Update an onboarding question
func (r *OnboardingQuestionRepository) Update(text string, updates map[string]interface{}) error {
	return r.db.Model(&models.OnboardingQuestion{}).Where("text = ?", text).Updates(updates).Error
}

// Delete an onboarding question
func (r *OnboardingQuestionRepository) Delete(text string) error {
	return r.db.Delete(&models.OnboardingQuestion{}, "text = ?", text).Error
}

// Create default questions for a specific assessment
func (r *OnboardingQuestionRepository) CreateDefaultQuestionsForAssessment(assessmentID int, assessmentName string) error {
	var defaultQuestions []models.OnboardingQuestion

	switch assessmentName {
	case "VB-MAPP":
		defaultQuestions = []models.OnboardingQuestion{
			// Group 1: Animal sounds & songs fill-ins
			{Text: "A kitty says...", AssessmentID: assessmentID, Group: 1},
			{Text: "Twinkle, twinkle, little...", AssessmentID: assessmentID, Group: 1},
			{Text: "Ready, set...", AssessmentID: assessmentID, Group: 1},
			{Text: "The wheels on the bus go...", AssessmentID: assessmentID, Group: 1},
			{Text: "A dog says...", AssessmentID: assessmentID, Group: 1},

			// Group 2: Name, fill-ins, associations
			{Text: "What is your name?", AssessmentID: assessmentID, Group: 2},
			{Text: "You brush your...", AssessmentID: assessmentID, Group: 2},
			{Text: "Shoes and...", AssessmentID: assessmentID, Group: 2},
			{Text: "You ride a...", AssessmentID: assessmentID, Group: 2},
			{Text: "You eat...", AssessmentID: assessmentID, Group: 2},

			// Group 3: Simple What questions
			{Text: "What can you drink?", AssessmentID: assessmentID, Group: 3},
			{Text: "What can fly?", AssessmentID: assessmentID, Group: 3},
			{Text: "What are some numbers?", AssessmentID: assessmentID, Group: 3},
			{Text: "What are some colors?", AssessmentID: assessmentID, Group: 3},
			{Text: "What are some animals?", AssessmentID: assessmentID, Group: 3},

			// Group 4: Simple Who, Where, & How old
			{Text: "Who is your teacher?", AssessmentID: assessmentID, Group: 4},
			{Text: "Where do you wash your hands?", AssessmentID: assessmentID, Group: 4},
			{Text: "Who lives on a farm?", AssessmentID: assessmentID, Group: 4},
			{Text: "How old are you?", AssessmentID: assessmentID, Group: 4},
			{Text: "Why do you use a bandaid?", AssessmentID: assessmentID, Group: 4},

			// Group 5: Categories, function, features
			{Text: "What shape are wheels?", AssessmentID: assessmentID, Group: 5},
			{Text: "What grows outside?", AssessmentID: assessmentID, Group: 5},
			{Text: "What can sting you?", AssessmentID: assessmentID, Group: 5},
			{Text: "What do you smell with?", AssessmentID: assessmentID, Group: 5},
			{Text: "What color are wheels?", AssessmentID: assessmentID, Group: 5},

			// Group 6: Adjectives, prepositions, adverbs
			{Text: "What do you wear on your head?", AssessmentID: assessmentID, Group: 6},
			{Text: "What do you eat with?", AssessmentID: assessmentID, Group: 6},
			{Text: "What's above a house?", AssessmentID: assessmentID, Group: 6},
			{Text: "What are some hot things?", AssessmentID: assessmentID, Group: 6},
			{Text: "What's under a house?", AssessmentID: assessmentID, Group: 6},

			// Group 7: Multiple part questions
			{Text: "What makes you sad?", AssessmentID: assessmentID, Group: 7},
			{Text: "What animal has a long neck?", AssessmentID: assessmentID, Group: 7},
			{Text: "Tell me something that is not a food.", AssessmentID: assessmentID, Group: 7},
			{Text: "What do you do with money?", AssessmentID: assessmentID, Group: 7},
			{Text: "What's something that is sticky?", AssessmentID: assessmentID, Group: 7},

			// Group 8: Multiple part questions
			{Text: "Where do you put your dirty clothes?", AssessmentID: assessmentID, Group: 8},
			{Text: "What do you take to a birthday party?", AssessmentID: assessmentID, Group: 8},
			{Text: "What day is today?", AssessmentID: assessmentID, Group: 8},
			{Text: "Why do people wear glasses?", AssessmentID: assessmentID, Group: 8},
			{Text: "How do you know if someone is sick?", AssessmentID: assessmentID, Group: 8},
			{Text: "What do you see in a city?", AssessmentID: assessmentID, Group: 8},
		}
	case "ESFLS":
		defaultQuestions = []models.OnboardingQuestion{
			{Text: "Can the patient make requests for essential items?", AssessmentID: assessmentID},
			{Text: "Is the patient able to tolerate specific situations?", AssessmentID: assessmentID},
			{Text: "Can the patient engage in daily living activities?", AssessmentID: assessmentID},
		}
	case "ABLLS-R":
		defaultQuestions = []models.OnboardingQuestion{
			{Text: "How would you rate the patient's visual performance skills?", AssessmentID: assessmentID},
			{Text: "Can the patient follow instructions?", AssessmentID: assessmentID},
			{Text: "Does the patient demonstrate language comprehension?", AssessmentID: assessmentID},
		}
	}

	for _, question := range defaultQuestions {
		var count int64
		r.db.Model(&models.OnboardingQuestion{}).Where("text = ?", question.Text).Count(&count)

		if count == 0 {
			if err := r.db.Create(&question).Error; err != nil {
				return err
			}
		}
	}

	return nil
}
