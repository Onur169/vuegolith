# Vuegolith Readme (In Progress)

## Introduction

Welcome to Vuegolith! This project combines a backend web server built with Go and a frontend web application built with Vue, TypeScript and Tailwind. Vuegolith aims to provide a user-friendly interface for various tasks, along with an easy-to-use file upload functionality.

<!-- TypeScript Logo -->

## Technologies Used

| TypeScript                                                                    | Vue.js                                                                     | Tailwind CSS                                                                          | Go                                                                      |
| ----------------------------------------------------------------------------- | -------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| <img src="assets/ts-logo.svg" alt="TypeScript Logo" width="100" height="100"> | <img src="assets/vue-logo.svg" alt="Vue.js Logo" width="100" height="100"> | <img src="assets/tailwind-logo.svg" alt="Tailwind CSS Logo" width="100" height="100"> | <img src="assets/go-logo.svg" alt="Go Logo" width="100" height="100" /> |

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

- `dev`: Starts the frontend development server using Vite with the `dev` mode. This is used during development to test and hot-reload changes.

- `prod`: Builds the frontend in production mode using Vite with the `prod` mode. This is used to generate optimized and minified assets for production deployment.

- `build`: Executes the `build.sh` shell script, responsible for building the frontend and moving the output to the `ui/dist` directory. It compiles TypeScript files, generates production-ready assets, and prepares the application for deployment.

- `publish`: Executes the `publish.sh` shell script, responsible for building and packaging the frontend and backend as a single executable. It creates an executable file named `vuegolith` and moves it to `/usr/local/bin/vuegolith`. The script also copies SSL certificate files (`server.key` and `server.crt`) to `/etc/vuegolith/ssl/`. It's important to run this script with elevated permissions using `sudo`.

- `publish-dev`: Sets the `VITE_ENV` environment variable to `dev` and then runs the `publish` script, which packages the application for development deployment.

- `publish-prod`: Sets the `VITE_ENV` environment variable to `prod` and then runs the `publish` script, which packages the application for production deployment.

- `build-dev`: Sets the `VITE_ENV` environment variable to `dev` and then runs the `build` script, which generates frontend assets in development mode.

- `build-prod`: Sets the `VITE_ENV` environment variable to `prod` and then runs the `build` script, which generates optimized frontend assets for production.

- `preview`: Previews the built frontend locally using Vite's preview feature. This allows you to test the production-ready assets locally before deploying them.

- `boilerplate`: Executes the `boilerplate.cjs` Node.js script to generate a new Vue.js component with the specified name in the "components" directory. To create a new component, you can use `npm run boilerplate ComponentName`.

- `format`: Formats Vue files using Prettier, ensuring consistent and standardized code formatting.

- `eslint`: Runs ESLint on TypeScript and Vue files, identifying potential code errors and enforcing coding standards.

### Using secure and unsecure webserver while development

To start local development, use the following commands based on your desired configuration:

1. **Unsecure Local Domain (HTTP)**:

   ```bash
   npm run dev
   ```

   This will start the frontend development server on `http://localhost:1234` with hot reloading. The development environment uses the `.env.dev` file, which is set to `VITE_USE_SECURE_LOCAL_DOMAIN=false`. This means you'll be running an unsecure HTTP web server during development by executing `go run main.go --secure=false`.

2. **Secure Local Domain (HTTPS)**:
   ```bash
   npm run prod
   ```
   This command starts the frontend development server on `https://localhost:1234` with hot reloading. The development environment uses the `.env.prod` file, which is set to `VITE_USE_SECURE_LOCAL_DOMAIN=true`. As a result, you'll be running a secure HTTPS web server during development by executing `go run main.go --secure=true`.

#### Publishing Vuegolith

When you are ready to publish your Vuegolith application, use the following scripts:

- **Publish for Development**:

  ```bash
  npm run publish-dev
  ```

  This script builds and publishes the frontend and backend as a single executable, `vuegolith`, to `/usr/local/bin/`. The Vuegolith executable will use the development environment settings, and you can start it with the `--secure=true` or `--secure=false` flag as needed.

- **Publish for Production**:
  ```bash
  npm run publish-prod
  ```
  This script builds and publishes the frontend and backend as a single executable, `vuegolith`, to `/usr/local/bin/`. The Vuegolith executable will use the production environment settings, and you can start it with the `--secure=true` or `--secure=false` flag based on your deployment requirements.

With the published Vuegolith executable, you can easily run your application using `vuegolith --secure=true` for a secure HTTPS server or `vuegolith --secure=false` for an unsecure HTTP server. The Vuegolith command is now available globally, allowing you to access your application from anywhere in the console.

### vuegolith.local Note

To access the secure local domain vuegolith.local, you need to add an entry to your system's hosts file. You can do this by running the following command in the terminal:

```bash
sudo nano /etc/hosts
```

Then, add the following line at the end of the file:

```bash
127.0.0.1    vuegolith.local
```

#### SSL Note

The "browser's trusted certificate store" refers to a collection of trusted root certificates that your web browser uses to validate the authenticity of SSL/TLS certificates presented by websites. These root certificates are issued by trusted Certificate Authorities (CAs) and are pre-installed in your browser or operating system.

When you access a website over HTTPS, the server presents its SSL/TLS certificate to your browser. Your browser then checks if the certificate is signed by a trusted root certificate. If it is, the browser considers the connection secure and displays a padlock icon or similar indicator to signify that the website is safe to use.

However, when you use a self-signed SSL certificate, like in local development scenarios, the certificate is not issued by a recognized Certificate Authority, so it is not automatically trusted by your browser. As a result, when you access a website using a self-signed certificate, your browser will display a security warning or error.

To avoid these warnings and make your browser trust the self-signed certificate used in local development, you can manually import the certificate into your browser's trusted certificate store. On macOS, you can use the Keychain Access application (Schlüsselbundverwaltung) to import the certificate.

## Conclusion

Vuegolith provides a powerful combination of Go and Vue.js, allowing you to build a feature-rich web application with a robust backend. By following the instructions in this readme, you should be able to set up Vuegolith and start developing your web application. If you have any questions or face any issues, feel free to reach out to the Vuegolith community for support.

Happy coding! 🚀
