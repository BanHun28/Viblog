# Viblog

개인 블로그 플랫폼 (모노레포 구조)

## 기술 스택

### Backend
- Go (Golang) + Gin
- Clean Architecture
- PostgreSQL + GORM
- JWT 인증
- Prometheus + Grafana

### Frontend
- Next.js + TypeScript
- Tailwind CSS + shadcn/ui
- Zustand (상태 관리)
- SSG/ISR

## 시작하기

### 필수 요구사항
- Go 1.23+
- Node.js 20.x LTS
- Docker & Docker Compose
- PostgreSQL 17+

### 개발 환경 실행

```bash
# 전체 서비스 시작
make dev

# 개별 서비스 시작
cd backend && make run
cd frontend && npm run dev
```

### 포트 설정
- Backend API: 30000
- Frontend Dev: 30001
- PostgreSQL: 30002
- Prometheus: 30003
- Grafana: 30004

## 프로젝트 구조

```
viblog/
├── backend/         # Go 백엔드
├── frontend/        # Next.js 프론트엔드
├── docker/          # Docker 설정
└── .github/         # CI/CD 워크플로우
```

## 라이선스

MIT
