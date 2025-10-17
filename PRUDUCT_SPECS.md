# Viblog - 개인 블로그 프로젝트

## 프로젝트 개요
개인 블로그 플랫폼 (모노레포 구조)

## 기술 스택

### Backend
- **언어**: Go (Golang)
- **웹 프레임워크**: Gin
- **아키텍처**: Clean Architecture
- **데이터베이스**: PostgreSQL
- **ORM**: GORM
- **마이그레이션**: GORM AutoMigrate
- **의존성 주입**: Wire
- **인증**: JWT
- **API 문서**: OpenAPI 3.0
- **로깅**: Zap (Uber)
- **모니터링**: Prometheus + Grafana

### Frontend
- **프레임워크**: Next.js
- **스타일링**: Tailwind CSS
- **UI 컴포넌트**: shadcn/ui
- **마크다운 에디터**: react-markdown + react-simplemde-editor
- **상태 관리**: Zustand
- **렌더링 전략**: SSG/ISR (블로그 특성 최적화)

### Infrastructure
- **컨테이너**: Docker
- **개발환경**: Docker Compose (dev/prod 환경 분리)
- **배포**: 개인 서버
- **CI/CD**: GitHub Actions
- **테스트**: Unit Test, Integration Test, E2E Test (testify)
- **API 문서**: Swagger UI (OpenAPI 3.0 자동 생성)
- **보안**: 
  - Rate Limiting (API: 분당 100회, 댓글: 분당 5회, 인메모리 저장)
  - CORS 설정 (개발: localhost:30001, 프로덕션: 배포 도메인)
  - XSS/CSRF 방어
  - 비밀번호 해싱: bcrypt

## 개발 워크플로우
- **Git 전략**: Git Worktree 활용
- **에러 처리**: 표준화된 에러 응답 형식 및 에러 코드 체계

## 핵심 기능

### 블로그 기능
- 마크다운 에디터 (실시간 프리뷰)
- 태그 시스템 (글당 최대 10개)
- 카테고리 시스템 (계층 구조, 글당 1개)
- 댓글 시스템 (자체 구현)
  - 익명 댓글 허용 (닉네임 + 비밀번호, bcrypt 해싱)
  - 회원 댓글 지원
  - 대댓글(답글) 지원
  - 댓글 수정/삭제 기능
  - 댓글 좋아요
  - Rate Limiting (IP당 분당 5회)
  - 답글 알림 (인앱)
- 검색 기능 (PostgreSQL Full-text search)
  - 제목, 본문, 태그, 카테고리 검색
  - 검색어 하이라이팅
- 글 상태 관리 (임시저장, 발행, 예약발행)
- 글 조회수 (IP 기반 중복 방지, 24시간 캐시)
- 글 좋아요 (회원만 가능)
- 북마크/즐겨찾기 (회원만 가능)
- 페이지네이션 (Cursor 기반, 페이지당 20개)
- SEO 최적화
  - 메타 태그 관리
  - sitemap.xml 자동 생성
  - RSS 피드
- 다국어 지원 (한국어, 영어)

### 사용자 관리
- 회원가입/로그인 (일반 사용자)
  - 필수 정보: 이메일, 비밀번호, 닉네임
  - 프로필 이미지 (URL 형태)
  - 비밀번호 정책 (최소 8자, 영문+숫자+특수문자)
- JWT 기반 인증
  - Access Token (15분)
  - Refresh Token (7일)
  - HttpOnly Cookie 사용
- 권한 체계
  - 관리자: 글 작성/수정/삭제, 댓글 관리
  - 일반 사용자: 댓글 작성만 가능

### 관리 기능
- 별도 관리자 페이지
- 단일 관리자 체계
- 대시보드
  - 글/댓글/사용자 통계
  - 조회수/좋아요 추이
  - 최근 댓글/답글
  - 인기 글/태그

### API
- RESTful API
- 버저닝: `/api/v1/` 형식

## 제외 기능
- 이미지 업로드 (현재 지원 안 함)

## 프로젝트 구조

```
viblog/
├── .github/
│   └── workflows/
│       ├── backend-ci.yml
│       ├── frontend-ci.yml
│       └── deploy.yml
│
├── backend/
│   ├── cmd/
│   │   └── api/
│   │       └── main.go                    # 애플리케이션 엔트리포인트
│   │
│   ├── internal/
│   │   ├── domain/                        # 엔티티 & 비즈니스 로직
│   │   │   ├── entity/
│   │   │   │   ├── user.go
│   │   │   │   ├── post.go
│   │   │   │   ├── comment.go
│   │   │   │   ├── tag.go
│   │   │   │   ├── category.go
│   │   │   │   ├── like.go
│   │   │   │   ├── bookmark.go
│   │   │   │   └── notification.go
│   │   │   ├── repository/                # Repository 인터페이스
│   │   │   │   ├── user.go
│   │   │   │   ├── post.go
│   │   │   │   ├── comment.go
│   │   │   │   └── ...
│   │   │   └── service/                   # 도메인 서비스
│   │   │
│   │   ├── usecase/                       # 유즈케이스 (애플리케이션 로직)
│   │   │   ├── user/
│   │   │   │   ├── register.go
│   │   │   │   ├── login.go
│   │   │   │   └── update_profile.go
│   │   │   ├── post/
│   │   │   │   ├── create.go
│   │   │   │   ├── update.go
│   │   │   │   ├── delete.go
│   │   │   │   ├── list.go
│   │   │   │   └── search.go
│   │   │   ├── comment/
│   │   │   └── admin/
│   │   │
│   │   ├── interface/                     # 어댑터 레이어
│   │   │   ├── http/
│   │   │   │   ├── handler/               # HTTP 핸들러
│   │   │   │   │   ├── user.go
│   │   │   │   │   ├── post.go
│   │   │   │   │   ├── comment.go
│   │   │   │   │   └── admin.go
│   │   │   │   ├── middleware/            # 미들웨어
│   │   │   │   │   ├── auth.go
│   │   │   │   │   ├── cors.go
│   │   │   │   │   ├── ratelimit.go
│   │   │   │   │   ├── logger.go
│   │   │   │   │   └── error.go
│   │   │   │   ├── presenter/             # 응답 변환
│   │   │   │   │   ├── user.go
│   │   │   │   │   ├── post.go
│   │   │   │   │   └── error.go
│   │   │   │   ├── router/                # 라우터 설정
│   │   │   │   │   └── router.go
│   │   │   │   └── dto/                   # 요청/응답 DTO 및 검증
│   │   │   │       ├── user.go
│   │   │   │       ├── post.go
│   │   │   │       └── comment.go
│   │   │
│   │   └── infrastructure/                # 인프라 레이어
│   │       ├── database/
│   │       │   ├── postgres.go            # DB 연결
│   │       │   └── migration.go           # 마이그레이션
│   │       ├── repository/                # Repository 구현체
│   │       │   ├── user_repository.go
│   │       │   ├── post_repository.go
│   │       │   ├── comment_repository.go
│   │       │   └── ...
│   │       ├── auth/
│   │       │   └── jwt.go                 # JWT 처리
│   │       ├── cache/
│   │       │   └── memory.go              # 인메모리 캐시
│   │       ├── logger/
│   │       │   └── zap.go                 # Zap 로거 설정
│   │       └── monitoring/
│   │           └── prometheus.go          # Prometheus 메트릭
│   │
│   ├── pkg/                               # 공용 패키지
│   │   ├── errors/                        # 에러 정의
│   │   │   ├── errors.go
│   │   │   └── codes.go
│   │   ├── password/                      # 비밀번호 유틸 (bcrypt)
│   │   │   └── hash.go
│   │   ├── validator/                     # 공통 검증 유틸 (이메일, URL 등)
│   │   │   └── validator.go
│   │   └── utils/                         # 기타 유틸
│   │       └── pagination.go
│   │
│   ├── api/
│   │   └── openapi/
│   │       └── spec.yaml                  # OpenAPI 3.0 스펙
│   │
│   ├── test/                              # 테스트
│   │   ├── unit/
│   │   ├── integration/
│   │   └── e2e/
│   │
│   ├── wire/                              # Wire DI 설정
│   │   ├── wire.go
│   │   └── wire_gen.go
│   │
│   ├── go.mod
│   ├── go.sum
│   ├── Makefile
│   └── .air.toml                          # Hot reload 설정
│
├── frontend/
│   ├── app/                               # Next.js App Router
│   │   ├── page.tsx                       # 홈 (루트)
│   │   ├── posts/
│   │   │   ├── page.tsx                   # 글 목록
│   │   │   └── [id]/
│   │   │       └── page.tsx               # 글 상세
│   │   ├── categories/
│   │   │   └── [slug]/
│   │   │       └── page.tsx
│   │   ├── tags/
│   │   │   └── [slug]/
│   │   │       └── page.tsx
│   │   ├── search/
│   │   │   └── page.tsx
│   │   ├── login/
│   │   │   └── page.tsx
│   │   ├── register/
│   │   │   └── page.tsx
│   │   ├── profile/
│   │   │   └── page.tsx
│   │   ├── bookmarks/
│   │   │   └── page.tsx
│   │   ├── notifications/
│   │   │   └── page.tsx
│   │   ├── admin/                         # 관리자 페이지
│   │   │   ├── layout.tsx
│   │   │   ├── page.tsx                   # 대시보드
│   │   │   ├── posts/
│   │   │   │   ├── page.tsx               # 글 관리
│   │   │   │   ├── new/
│   │   │   │   │   └── page.tsx           # 글 작성
│   │   │   │   └── [id]/
│   │   │   │       └── edit/
│   │   │   │           └── page.tsx       # 글 수정
│   │   │   ├── comments/
│   │   │   │   └── page.tsx               # 댓글 관리
│   │   │   ├── categories/
│   │   │   │   └── page.tsx
│   │   │   └── tags/
│   │   │       └── page.tsx
│   │   ├── api/                           # API Routes (선택적)
│   │   │   └── sitemap/
│   │   │       └── route.ts
│   │   ├── layout.tsx                     # 루트 레이아웃
│   │   └── globals.css
│   │
│   ├── components/
│   │   ├── ui/                            # 기본 UI 컴포넌트
│   │   │   ├── Button.tsx
│   │   │   ├── Input.tsx
│   │   │   ├── Card.tsx
│   │   │   └── ...
│   │   ├── layout/                        # 레이아웃 컴포넌트
│   │   │   ├── Header.tsx
│   │   │   ├── Footer.tsx
│   │   │   └── Sidebar.tsx
│   │   ├── post/                          # 글 관련 컴포넌트
│   │   │   ├── PostCard.tsx
│   │   │   ├── PostList.tsx
│   │   │   ├── MarkdownEditor.tsx
│   │   │   └── MarkdownViewer.tsx
│   │   ├── comment/                       # 댓글 컴포넌트
│   │   │   ├── CommentForm.tsx
│   │   │   ├── CommentItem.tsx
│   │   │   └── CommentList.tsx
│   │   └── admin/                         # 관리자 컴포넌트
│   │       ├── Dashboard.tsx
│   │       └── Stats.tsx
│   │
│   ├── lib/
│   │   ├── api/                           # API 클라이언트
│   │   │   ├── client.ts
│   │   │   ├── auth.ts
│   │   │   ├── posts.ts
│   │   │   └── comments.ts
│   │   ├── hooks/                         # Custom Hooks
│   │   │   ├── useAuth.ts
│   │   │   ├── usePosts.ts
│   │   │   └── useComments.ts
│   │   ├── store/                         # 상태 관리 (Zustand)
│   │   │   ├── authStore.ts
│   │   │   └── uiStore.ts
│   │   ├── utils/                         # 유틸리티
│   │   │   ├── format.ts
│   │   │   └── validation.ts
│   │   └── constants/
│   │       └── index.ts
│   │
│   ├── types/                             # TypeScript 타입
│   │   ├── api.ts
│   │   ├── post.ts
│   │   ├── user.ts
│   │   └── comment.ts
│   │
│   ├── public/
│   │   ├── images/
│   │   └── fonts/
│   │
│   ├── .env.local
│   ├── .env.example
│   ├── next.config.js
│   ├── tailwind.config.js
│   ├── tsconfig.json
│   ├── package.json
│   └── postcss.config.js
│
├── docker/
│   ├── backend/
│   │   ├── Dockerfile.dev
│   │   └── Dockerfile.prod
│   ├── frontend/
│   │   ├── Dockerfile.dev
│   │   └── Dockerfile.prod
│   └── postgres/
│       └── init.sql
│
├── .env.dev
├── .env.prod
├── .env.example
├── docker-compose.yml
├── docker-compose.prod.yml
├── .gitignore
├── .editorconfig
├── Makefile                               # 공통 명령어
└── README.md
```

## 포트 설정
- Backend API: 30000
- Frontend Dev: 30001
- PostgreSQL: 30002
- Prometheus: 30003
- Grafana: 30004

## 환경 변수 관리
- `.env.dev`: 개발 환경 설정
- `.env.prod`: 프로덕션 환경 설정
- Git ignore 처리 (`.env.example` 템플릿 제공)

## 개발 환경 설정

### 필수 도구
- Go 1.23+
- Node.js 20.x LTS
- Docker & Docker Compose
- PostgreSQL 17+
