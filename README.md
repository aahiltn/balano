# Paalam 🌉

Paalam is a comprehensive system designed for the **Dr. Beema Clinic for Child Development**. It aims to bridge the gap between patients and providers by enabling therapists, supervisors, and staff to track therapy sessions, monitor progress, and manage activities effectively. The system adheres to standards like ABA (Applied Behavior Analysis) and IBT (Individualized Behavioral Therapy).

Paalam, meaning bridge in Malayalam, will bridge the gap between patients, guardians, and hospital staff in organization.

## 🌟 Purpose

The primary goal of Paalam is to streamline therapy management and reporting for child development clinics. Key features include:

- **Session Tracking**: Record and monitor therapy sessions for patients.
- **Activity Management**: Log activities performed during sessions, categorized by therapy type.
- **Progress Monitoring**: Enable supervisors to generate reports and track patient progress.
- **Staff and Patient Management**: Manage staff roles, schedules, and patient assignments.
- **Medication Tracking**: Track prescribed medicines for patients.

## 🛠️ Technology Stack

Paalam is built using the following technologies:

- **Backend**: Go (Golang) with GORM for database interactions and Fiber for the web framework.
- **Frontend**: Astro with integrations for React and Vue.
- **Database**: MySQL.
- **API Documentation**: OpenAPI 3.0.
- **Deployment**: Designed to be cloud-ready with Azure best practices.

## 🚀 Getting Started

### For Clients

If you are a client (e.g., a clinic administrator or therapist), you can:

1. Access the system through the provided web interface.
2. Use the API for integrations with other systems (refer to the API documentation at `/docs`).

### For Contributors

If you are a developer looking to contribute, follow these steps:

### Prerequisites

- Install [Go](https://golang.org/) (v1.23 or higher).
- Install [Node.js](https://nodejs.org/) (v16 or higher).
- Install [MySQL](https://www.mysql.com/).

### Backend Setup

1. Clone the repository:
   ```sh
   git clone https://github.com/your-org/Paalam.git
   cd Paalam/backend
   ```
2. Set up environment variables:
   Create a `.env` file in the `backend` directory with the following:
   ```env
   DB_USER=<your-db-user>
   DB_PASSWORD=<your-db-password>
   DB_HOST=localhost
   DB_PORT=3306
   DB_NAME=Paalam
   APPLICATION_PORT=3000
   ```
3. Run the backend:
   ```sh
   go run ./cmd/server/main.go
   ```

### Frontend Setup

1. Navigate to the frontend directory:
   ```sh
   cd ../frontend
   ```
2. Install dependencies:
   ```sh
   npm install
   ```
3. Start the development server:
   ```sh
   npm run dev
   ```

### API Documentation

The API documentation is available at `/docs` when the backend server is running.

## ❤️ Acknowledgments

This project is built for the **Dr. Beema Clinic for Child Development** with the goal of improving therapy management and patient outcomes. Special thanks to all contributors and supporters of this initiative.

Built with ❤️ by Aahil Nishad.
