# Chatify

Chatify is a user-friendly and versatile chat application designed to facilitate seamless communication among multiple clients. It uses WebSockets to communicate between the frontend and backend. The frontend is built with React.js, TypeScript, and Tailwind CSS, while the backend is built with Go.

![Demo picture](/images/demo.png)

## Getting Started

### Prerequisites

-   Fill in the variables in the `.env` file in both `frontend` and `backend` directories. You can use the `.env.example` files as a template.

### Installation

1. Install node modules in `frontend` directory.

```bash
cd frontend
npm install
```

2. Install Go modules in `backend` directory.

```bash
cd backend
go mod download
```

### Usage

1. Start the frontend server.

```bash
cd frontend
npm start
```

2. Start the backend server.

```bash
cd backend
go run main.go
```

## License

Distributed under the MIT License. See `LICENSE` for more information.
