# 📝 Go To-Do CLI App

A simple **command-line To-Do List application** written in **Go (Golang)**, designed as a beginner-friendly project that demonstrates:

- Clean project structure
- Package organization
- SQLite database integration
- Basic CRUD operations (Create, Read, Update, Delete)

---

## 🚀 Features

✅ Add new tasks  
✅ List all tasks  
✅ Mark tasks as done  
✅ Delete tasks  
✅ Persistent data storage using SQLite  
✅ Modular structure (logic, UI, and data separated)

---

## 🗂️ Project Structure

```
todo-app/
│
├── go.mod              # Go module file (dependencies)
├── go.sum              # Dependency checksums
├── main.go             # App entry point
│
├── todo/               # Business logic (data layer)
│   ├── todo.go         # ToDo model definition
│   └── service.go      # Database functions (CRUD)
│
└── ui/                 # User Interface (CLI)
    └── cli.go          # Command-line interface loop
```

---

## 🧰 Requirements

- [Go 1.20+](https://go.dev/dl/)
- [SQLite](https://www.sqlite.org/download.html) (optional, used automatically via Go driver)

---

## ⚙️ Installation

### 1️⃣ Clone the repository
```bash
git clone https://github.com/<your-username>/todo-app.git
cd todo-app
```

### 2️⃣ Initialize Go modules (if needed)
```bash
go mod tidy
```

### 3️⃣ Run the app
```bash
go run .
```

---

## 🧠 Usage

When you run the app, you’ll see:

```
📋 To-Do List (SQLite)
----------------------------
1. Add task
2. List tasks
3. Mark task as done
4. Delete task
5. Exit
Choose an option:
```

Example session:
```
1. Add task
Enter new task: Learn Go
✅ Task added successfully!

2. List tasks
1. Learn Go [❌]

3. Mark task as done
Enter task ID to mark as done: 1
✅ Task marked as done!
```

Tasks are saved in a local SQLite database file:
```
./todo.db
```

---

## 🧱 Database

The app uses **SQLite** through the Go driver [`github.com/mattn/go-sqlite3`](https://github.com/mattn/go-sqlite3).  
If the `todo.db` file doesn’t exist, it will be created automatically.

You can inspect it using any SQLite client:
```bash
sqlite3 todo.db
sqlite> SELECT * FROM todos;
```

---

## 🧩 Next Steps (Ideas to Expand)

- [ ] Add a REST API with `net/http`
- [ ] Add unit tests
- [ ] Support PostgreSQL or MySQL
- [ ] Add categories or due dates
- [ ] Export tasks to JSON or CSV
- [ ] Build a web interface (React or Vue)

---

## 🧑‍💻 Author

**Iman Aboheydary**
- 🏙️ Frankfurt, Germany
- 💼 Developer at CustomGento & Founder of فارسی کیت
- 🌐 [I.MAN.News on Instagram](https://instagram.com/I.MAN.News)

---

## 📄 License

This project is open-source under the [MIT License](LICENSE).
