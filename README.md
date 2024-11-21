# FoodieBaba

Welcome to **FoodieBaba**—a dynamic platform where food enthusiasts can discover, share, and review restaurants. Built with a powerful Go backend and a modern React frontend, FoodieBaba aims to create a vibrant community for food lovers.


---

## Table of Contents

- [Features](#features)
- [Demo](#demo)
- [Tech Stack](#tech-stack)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Backend Setup](#backend-setup)
  - [Frontend Setup](#frontend-setup)
- [Project Structure](#project-structure)
- [API Documentation](#api-documentation)
- [Contributing](#contributing)
- [License](#license)
- [Contact](#contact)

---

## Features

- **User Authentication**: Secure registration and login with JWT authentication.
- **Restaurant Listings**: View, add, and manage restaurant listings.
- **Reviews and Ratings**: Share reviews and rate restaurants.
- **Admin Panel**: Admins can approve or reject new restaurant submissions.
- **Responsive Design**: Optimized for both desktop and mobile devices.
- **Search and Filter**: Easily find restaurants by name, category, or location.
- **Photo Uploads**: Upload and view photos of restaurants.

---

## Demo

[Live Demo](#) *(Link to be available after deploying the application)*

---

## Tech Stack

**Frontend**:

- [React](https://reactjs.org/)
- [Vite](https://vitejs.dev/)
- [Axios](https://axios-http.com/)
- [React Router](https://reactrouter.com/)

**Backend**:

- [Go (Golang)](https://golang.org/)
- [Gin Web Framework](https://github.com/gin-gonic/gin)
- [Gorm ORM](https://gorm.io/)
- [PostgreSQL](https://www.postgresql.org/)
- [JWT (JSON Web Tokens)](https://jwt.io/)
- [Swagger](https://swagger.io/) for API documentation

**Others**:

- [Docker](https://www.docker.com/) (optional)
- [Prometheus](https://prometheus.io/) and [Grafana](https://grafana.com/) for monitoring
- [Nginx](https://www.nginx.com/) (optional for deployment)

---

## Getting Started

Follow these instructions to get a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- **Go**: Version 1.16 or higher
- **Node.js**: Version 14 or higher
- **npm** or **Yarn**
- **PostgreSQL**: Version 12 or higher
- **Git**

### Backend Setup

#### 1. Clone the Repository

```bash
git clone https://github.com/pp00x/foodiebaba.git
cd foodiebaba/backend
```

#### 2. Set Up Environment Variables

Create a `.env` file in the `backend` directory:

```env
DB_HOST=localhost
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_NAME=foodiebaba
DB_PORT=5432
SSL_MODE=disable
JWT_SECRET=your_jwt_secret
```

- Replace `your_db_user` and `your_db_password` with your PostgreSQL credentials.
- `JWT_SECRET` should be a strong, random string.

#### 3. Install Dependencies

```bash
go mod download
```

#### 4. Initialize the Database

- Ensure PostgreSQL is running.
- Create the database:

  ```sql
  CREATE DATABASE foodiebaba;
  ```

- Grant privileges to your database user if necessary.

#### 5. Run the Backend Server

```bash
go run cmd/server/main.go
```

The server should start on `http://localhost:8080`.

### Frontend Setup

#### 1. Navigate to Frontend Directory

```bash
cd ../frontend
```

#### 2. Install Dependencies

```bash
npm install
# or
yarn install
```

#### 3. Set Up Environment Variables

Create a `.env` file in the `frontend` directory:

```env
VITE_API_BASE_URL=http://localhost:8080
```

#### 4. Run the Frontend Development Server

```bash
npm run dev
# or
yarn dev
```

The application should be running on `http://localhost:5173`.

---

## Project Structure

```bash
foodiebaba/
├── backend/
│   ├── cmd/
│   │   └── server/
│   │       └── main.go
│   ├── internal/
│   │   ├── controllers/
│   │   ├── middlewares/
│   │   ├── models/
│   │   └── db/
│   ├── configs/
│   ├── pkg/
│   │   └── logger/
│   ├── go.mod
│   └── ...
├── frontend/
│   ├── src/
│   │   ├── components/
│   │   ├── contexts/
│   │   ├── pages/
│   │   ├── services/
│   │   ├── App.jsx
│   │   ├── main.jsx
│   │   └── index.css
│   ├── public/
│   ├── .env
│   ├── package.json
│   └── ...
├── README.md
└── ...
```

---

## API Documentation

The API is documented using Swagger.

- **Swagger UI**: Visit `http://localhost:8080/swagger/index.html` after starting the backend server.

---

## Contributing

We welcome contributions from the community!

### Steps to Contribute

1. **Fork the Repository**

   Click the "Fork" button on the top right of the repository page.

2. **Clone Your Fork**

   ```bash
   git clone https://github.com/pp00x/foodiebaba.git
   ```

3. **Create a Branch**

   ```bash
   git checkout -b feature/YourFeatureName
   ```

4. **Make Changes**

   Implement your feature or fix.

5. **Commit and Push**

   ```bash
   git add .
   git commit -m "Add Your Feature"
   git push origin feature/YourFeatureName
   ```

6. **Create a Pull Request**

   Go to the original repository and create a pull request from your fork.

### Coding Guidelines

- Follow the existing code style.
- Write meaningful commit messages.
- Update documentation and comments where necessary.

---

## License

This project is licensed under the [Apache-2.0 License](LICENSE).


---

## Acknowledgements

- Thanks to all contributors and the open-source community for their invaluable support.
- Special mentions to the authors of the libraries and frameworks used in this project.


---

## Additional Documentation and Resources

- **Go Documentation**: [https://golang.org/doc/](https://golang.org/doc/)
- **React Documentation**: [https://reactjs.org/docs/getting-started.html](https://reactjs.org/docs/getting-started.html)
- **Gin Web Framework**: [https://github.com/gin-gonic/gin](https://github.com/gin-gonic/gin)
- **Gorm ORM**: [https://gorm.io/docs/](https://gorm.io/docs/)

---

## Troubleshooting

### Common Issues

#### CORS Errors

If you encounter CORS errors when making API requests from the frontend, ensure that the CORS middleware is properly configured in your backend.

**Example CORS Configuration in `main.go`:**

```go
import (
    "github.com/gin-contrib/cors"
    // other imports
)

func main() {
    // other setup code
    r := gin.Default()

    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:5173"},
        AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        AllowCredentials: true,
        MaxAge:           12 * time.Hour,
    }))

    // rest of your code
}
```

#### Database Connection Issues

- Ensure PostgreSQL is running and accessible.
- Verify the database credentials in your `.env` file.
- Check that the database exists and the user has the necessary permissions.

---


## Future Enhancements

- **User Profiles**: Allow users to customize their profiles.
- **Social Sharing**: Enable sharing reviews and restaurants on social media.
- **Notifications**: Implement notification system for user interactions.
- **Advanced Search**: Add filters and sorting options for restaurant searches.
- **Mobile App**: Develop a mobile application version.