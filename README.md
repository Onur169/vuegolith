# Vuegolith Readme

## Introduction

Welcome to Vuegolith! This project combines a backend web server built with Go and a frontend web application built with Vue, TypeScript and Tailwind. Vuegolith aims to provide a user-friendly interface for various tasks, along with an easy-to-use file upload functionality.

<!-- TypeScript Logo -->

## Technologies Used

<div style="display:flex; flex-direction: row; gap: 2rem">
    <img src="assets/ts-logo.svg" alt="TypeScript Logo" width="100" height="100">
    <img src="assets/vue-logo.svg" alt="Vue.js Logo" width="100" height="100">
    <img src="assets/tailwind-logo.svg" alt="Tailwind CSS Logo" width="100" height="100">
    <img src="assets/go-logo.svg" alt="Go Logo" width="100" height="100">
</div>

## Getting Started

To get started with Vuegolith, follow the steps below:

### Backend Setup

1. Clone the repository: `git clone git@github.com:Onur169/vuegolith.git`
2. Navigate to the project root directory: `cd vuegolith`
3. Make sure you have Go installed on your system.
4. Install Go dependencies: `go mod download`
5. Create a self-signed SSL certificate and key, and place them in the `/etc/vuegolith/ssl/` directory. Name the files `server.crt` and `server.key`, respectively.
6. Run the backend server: `sudo go run .`
   Note: Running the server with sudo is necessary to access the `/etc/vuegolith/ssl/` directory.

### Frontend Setup

1. Navigate to the frontend directory: `cd frontend`
2. Make sure you have Node.js and npm installed on your system.
3. Install frontend dependencies: `npm install`
4. Start the development server: `npm run dev`

## Important Notes

### Backend Root Directory

The backend of Vuegolith is built using the Go programming language. It serves as a web server providing APIs and handling file uploads. The backend's root directory is where you have cloned the repository (e.g., `/path/to/vuegolith`).

### Frontend Root Directory

The frontend of Vuegolith is built using Vue.js and Vite. It is located in the `frontend` directory. When working with the frontend, make sure to navigate to the `frontend` directory using `cd frontend` before running any npm commands.

### Scripts in package.json

Here is a quick overview of the scripts available in the `package.json` file:

- `dev`: Starts the frontend development server.
- `build`: Builds the frontend and moves the output to the `ui/dist` directory.
- `preview`: Previews the built frontend locally.
- `boilerplate`: (Optional) Runs a Node.js script for boilerplate generation (you may provide more details about this if needed).
- `format`: Formats Vue files using Prettier.
- `eslint`: Runs ESLint for linting TypeScript and Vue files.
- `publish`: Publishes the frontend and backend as a single executable to `/usr/local/bin/vuegolith` (make sure to run this script with elevated permissions using `sudo`).

### Recommended NPM Commands

When working on the frontend, it is essential to execute the npm commands from within the `frontend` directory. For example:

- To install new dependencies: `npm install <package-name>`
- To run the development server: `npm run dev`
- To build the frontend: `npm run build`
- To preview the built frontend: `npm run preview`

Remember not to execute `npm install` in the root directory. Doing so will install dependencies at the project root, which is not recommended.

## Conclusion

Vuegolith provides a powerful combination of Go and Vue.js, allowing you to build a feature-rich web application with a robust backend. By following the instructions in this readme, you should be able to set up Vuegolith and start developing your web application. If you have any questions or face any issues, feel free to reach out to the Vuegolith community for support.

Happy coding! 🚀
