# IT KMITL Workshop Registration Portal

ระบบสมัครเข้าร่วมกิจกรรมและ Workshop ของคณะเทคโนโลยีสารสนเทศ สถาบันเทคโนโลยีพระจอมเกล้าเจ้าคุณทหารลาดกระบัง

## Tech Stack

| Layer | Technology |
|-------|-----------|
| Frontend | SvelteKit + TailwindCSS v4 |
| Backend | Go Fiber v2 |
| Database | PostgreSQL 16 |
| Containerization | Docker + Docker Compose |
| Target Cloud | AWS (EC2, ELB, ASG, RDS, S3, VPC) |

## Quick Start (Development)

### Prerequisites
- Docker & Docker Compose
- Node.js 22+
- Go 1.25+

### Option 1: Docker Compose (Recommended)

```bash
docker compose up --build
```

- Frontend: http://localhost:3000
- Backend API: http://localhost:8080
- PostgreSQL: localhost:5432

### Option 2: Run Locally

**Database:**
```bash
docker compose up db
```

**Backend:**
```bash
cd backend
cp ../.env.example .env
go run ./cmd/main.go
```

**Frontend:**
```bash
cd frontend
npm install
npm run dev
```

Frontend dev server: http://localhost:5173

## Project Structure

```
├── frontend/          # SvelteKit + TailwindCSS v4
│   ├── src/
│   │   ├── lib/       # Components, stores, utils
│   │   └── routes/    # Pages (/, /register, /login, /home, /activity/[id])
│   └── Dockerfile
├── backend/           # Go Fiber REST API
│   ├── cmd/           # Entry point
│   ├── internal/      # Config, DB, handlers, middleware, models, routes
│   └── Dockerfile
├── docker-compose.yml
└── .env.example
```

## API Endpoints

| Method | Endpoint | Auth | Description |
|--------|----------|------|-------------|
| GET | /api/health | No | Health check |
| POST | /api/auth/register | No | Register |
| POST | /api/auth/login | No | Login |
| POST | /api/auth/logout | No | Logout |
| GET | /api/auth/me | Yes | Current user |
| GET | /api/activities | No | List activities |
| GET | /api/activities/:id | No | Activity detail |
| POST | /api/registrations | Yes | Register for activity |
| GET | /api/registrations/me | Yes | My registrations |
| DELETE | /api/registrations/:id | Yes | Cancel registration |
| GET | /api/registrations/check/:activityId | Yes | Check if registered |

## AWS Architecture (Target)

- **VPC**: Multi-AZ with public/private subnets
- **EC2**: Docker containers running in ASG
- **ALB**: Application Load Balancer for traffic distribution
- **ASG**: Auto Scaling based on CPU metrics
- **RDS**: PostgreSQL Multi-AZ for HA
- **S3**: Static assets / activity images
