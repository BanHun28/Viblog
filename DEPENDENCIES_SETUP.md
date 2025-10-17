# Required Dependencies Setup

## Frontend Dependencies

The following packages need to be installed for the authentication system to work:

### Core Dependencies

```bash
cd frontend

# HTTP client
npm install axios

# Class name utilities (for UserAvatar component)
npm install clsx tailwind-merge

# State management (if not already installed)
npm install zustand
```

### Type Definitions

```bash
# Type definitions (if using TypeScript)
npm install -D @types/node
```

## Verification

After installation, verify the dependencies:

```bash
npm list axios zustand clsx tailwind-merge
```

## Environment Variables

Create or update `.env.local` file in the frontend directory:

```env
# API Configuration
NEXT_PUBLIC_API_URL=http://localhost:30000/api/v1
```

## Running the Application

```bash
# Development mode
npm run dev

# Build for production
npm run build

# Start production server
npm start
```

## Testing the Implementation

1. **Start Backend Server** (Port 30000)
   ```bash
   cd backend
   make run-dev
   ```

2. **Start Frontend Server** (Port 30001)
   ```bash
   cd frontend
   npm run dev
   ```

3. **Access the Application**
   - Frontend: http://localhost:30001
   - Backend API: http://localhost:30000/api/v1
   - Swagger Docs: http://localhost:30000/swagger

## Quick Test Checklist

- [ ] Register a new user at `/register`
- [ ] Login with credentials at `/login`
- [ ] View profile at `/profile`
- [ ] Update profile information
- [ ] Check header shows user info
- [ ] Test logout functionality
- [ ] Verify protected routes redirect to login
- [ ] Test token refresh (wait 15+ minutes)

## Troubleshooting

### CORS Issues
If you encounter CORS errors:
1. Check backend CORS configuration
2. Verify API_URL in `.env.local`
3. Ensure credentials are properly sent

### Token Issues
If tokens aren't persisting:
1. Check localStorage in browser DevTools
2. Verify `auth-storage` key exists
3. Check browser console for errors

### API Connection
If API requests fail:
1. Verify backend is running on port 30000
2. Check network tab in DevTools
3. Verify API_URL environment variable
