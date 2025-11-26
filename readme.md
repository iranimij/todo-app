# Todo App

A full-stack todo application with a Go backend (REST API + CLI) and React + TypeScript frontend.

## Features

✅ **Three Ways to Interact:**
- 🌐 Modern web interface (React UI)
- 💻 Interactive CLI (Terminal-based)
- 🔌 REST API (For integrations)

✅ **Core Functionality:**
- Create new todos
- List all todos (organized by pending/completed)
- Mark todos as done
- Delete todos
- Real-time sync across all interfaces

✅ **Technical Features:**
- One-command setup (auto-installs dependencies)
- Automatic port detection (tries 8080-8089)
- Auto-starts frontend dev server
- CORS enabled for API
- Interactive/non-interactive mode detection
- TypeScript type safety
- Responsive design with gradient UI

## Quick Start

### Prerequisites
- **Backend:** Go 1.24.6 or higher
- **Frontend:** Node.js 18+ and npm (will auto-install dependencies)
- SQLite3 (included with Go)

### One Command Setup! 🚀

Just run this in your terminal:

```bash
cd backend
go run .
```

**That's it!** The app will:
1. ✅ Automatically install frontend dependencies (first run only)
2. ✅ Start the frontend dev server
3. ✅ Start the backend API
4. ✅ Launch the interactive CLI

You'll see:
```
📦 Installing frontend dependencies...
   This may take a minute on first run...
✅ Frontend dependencies installed!

🚀 Starting frontend dev server...

🌐 Web Interface: http://localhost:3000

📋 To-Do List (SQLite)
----------------------------
1. Add task
2. List tasks
3. Mark task as done
4. Delete task
5. Exit
Choose an option:
```

### Access the App

- **Web UI:** Open `http://localhost:3000` in your browser
- **CLI:** Use the interactive menu in the terminal
- **API:** Available at `http://localhost:8080` (or next available port)

### Alternative: API Only Mode

To run only the API server without CLI or frontend:

```bash
cd backend
go run . < /dev/null
```

## Backend (Go)

### API Endpoints

All endpoints use JSON format with CORS enabled.

| Method | Endpoint | Description | Request Body | Response |
|--------|----------|-------------|--------------|----------|
| GET | `/todos` | Get all todos | - | `[{"id":1,"task":"...","done":false}]` |
| POST | `/todos` | Create new todo | `{"task":"Buy milk"}` | `{"message":"task created successfully"}` |
| PUT | `/todos/{id}/done` | Mark todo as done | - | `{"message":"task marked as done"}` |
| DELETE | `/todos/{id}` | Delete todo | - | `{"message":"task deleted successfully"}` |

### Example API Calls

```bash
# Get all todos
curl http://localhost:8080/todos

# Create a new todo
curl -X POST http://localhost:8080/todos \
  -H "Content-Type: application/json" \
  -d '{"task":"Buy groceries"}'

# Mark todo #1 as done
curl -X PUT http://localhost:8080/todos/1/done

# Delete todo #1
curl -X DELETE http://localhost:8080/todos/1
```

### Smart Port Detection

The backend automatically tries ports 8080-8089 and selects the first available one. If port 8080 is busy, it will use 8081, 8082, etc.

## Frontend (React + TypeScript)

### Development

```bash
cd frontend
npm install  # First time only
npm run dev  # Start dev server
```

The frontend runs on `http://localhost:3000` and automatically proxies API requests to `http://localhost:8080`.

### Build for Production

```bash
npm run build
```

Built files will be in the `dist/` directory.

### Features

- **Modern UI:** Gradient design with smooth animations
- **Responsive:** Works on desktop, tablet, and mobile
- **Organized View:** Separates pending and completed todos
- **Real-time Updates:** Changes reflect immediately
- **Error Handling:** User-friendly error messages

## Technologies Used

### Backend
- **Go** - Programming language
- **net/http** - Standard library HTTP server
- **SQLite3** - Embedded database
- **go-sqlite3** - SQLite driver for Go

### Frontend
- **React 18** - UI library
- **TypeScript** - Type-safe JavaScript
- **Vite** - Fast build tool and dev server
- **CSS3** - Styling with gradients and animations

## Architecture

### Data Flow

```
User Action (Web/CLI)
    ↓
API Request (HTTP)
    ↓
Backend Handler (Go)
    ↓
Database Operation (SQLite)
    ↓
Response (JSON)
    ↓
UI Update (React/CLI)
```

### Database Schema

```sql
CREATE TABLE todos (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    task TEXT NOT NULL,
    done BOOLEAN NOT NULL DEFAULT 0
);
```

## Troubleshooting

**"npm: command not found" error?**
- Make sure Node.js and npm are installed
- Install from: https://nodejs.org/
- Verify with: `npm --version`

**Port already in use?**
- The backend automatically tries ports 8080-8089
- Check which port it selected in the startup message

**Frontend dependencies not installing?**
- The backend will show an error if npm install fails
- You can manually run: `cd ../frontend && npm install`
- Check that you have write permissions in the frontend directory

**Frontend can't connect to API?**
- Ensure backend is running first
- Check the proxy configuration in `frontend/vite.config.ts`
- Verify CORS is enabled in `backend/api/handler.go`

**Database locked error?**
- Only one backend instance can access the database at a time
- Close other instances before starting a new one

## Development

### Adding New Features

1. **Backend:**
   - Add database operations in `backend/todo/service.go`
   - Create API handlers in `backend/api/handler.go`
   - Update routes in `backend/main.go`

2. **Frontend:**
   - Add API methods in `frontend/src/services/api.ts`
   - Create/update components in `frontend/src/components/`
   - Update types in `frontend/src/types/todo.ts`

### Testing

**Manual Testing:**
```bash
# Test API endpoints
curl http://localhost:8080/todos

# Test frontend
Open http://localhost:3000 in browser

# Test CLI
Run backend and use the menu options
```

## License

MIT

---

**Made with Go, React, and TypeScript** 🚀
