# 🚀 FeatherJet

[![Go Version](https://img.shields.io/badge/Go-1.21+-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![Cross Platform](https://img.shields.io/badge/Platform-Linux%20%7C%20Windows%20%7C%20macOS-lightgrey.svg)]()

**FeatherJet** is a modern, lightweight web server written in Go, perfect for serving static websites and REST APIs. Think of it as a simple, fast alternative to heavyweight servers, with minimal configuration needed.

## ✨ Key Features

- 🌐 **Serve Static Files**: HTML, CSS, JavaScript, images - all with proper caching
- 🔄 **REST API Support**: Built-in handlers for JSON APIs
- 📱 **Modern UI**: Clean interface with light/dark mode support
- 🛠️ **Easy Configuration**: Simple YAML files - no complex setups
- 🔒 **Security Ready**: CORS, security headers, and rate limiting included
- 🚀 **Fast & Light**: ~10MB memory usage, <1s startup time
- 📊 **Built-in Logging**: Request tracking and error monitoring
- 💻 **Cross-Platform**: One binary for Linux, Windows, and macOS

## 🏃‍♂️ Quick Start

1. **Download & Run:**
   ```bash
   # Clone the repo
   git clone https://github.com/featherjet/featherjet.git
   cd featherjet

   # Build and run
   go build ./cmd/featherjet
   ./featherjet
   ```

2. **Visit your site:**
   - Open http://localhost:8081
   - Try the demo page with dark mode support
   - Check API endpoints at /api/hello

## � Project Structure

```
FeatherJet/
├── cmd/featherjet/          # Main application
├── internal/                # Core logic
│   ├── server/             # HTTP server & handlers
│   ├── config/             # Configuration
│   └── middleware/         # HTTP middlewares
├── public/                 # Demo website
└── config.yaml            # Main config file
```

## ⚙️ Setup Instructions

### Prerequisites

- **Go 1.21 or later**: [Download Go](https://golang.org/dl/)
- **Git**: For cloning the repository

### Installation

1. **Clone the repository**:
   ```bash
   git clone https://github.com/featherjet/featherjet.git
   cd featherjet
   ```

2. **Download dependencies**:
   ```bash
   go mod download
   ```

3. **Build the application**:
   ```bash
   # For your current platform
   go build -o featherjet ./cmd/featherjet
   
   # Or cross-compile for different platforms
   # Linux
   GOOS=linux GOARCH=amd64 go build -o featherjet-linux ./cmd/featherjet
   
   # Windows  
   GOOS=windows GOARCH=amd64 go build -o featherjet.exe ./cmd/featherjet
   
   # macOS
   GOOS=darwin GOARCH=amd64 go build -o featherjet-macos ./cmd/featherjet
   ```

### Running FeatherJet

1. **Start with default configuration**:
   ```bash
   ./featherjet
   ```
   Server will start on `http://localhost:8081`

2. **Use custom configuration**:
   ```bash
   ./featherjet -config custom-config.yaml
   ```

3. **View the demo application**:
   Open your browser to `http://localhost:8081` to see the included demo application.

### Platform-Specific Instructions

#### Linux/macOS
```bash
# Make executable (if needed)
chmod +x featherjet

# Run in background
nohup ./featherjet > featherjet.log 2>&1 &

# Stop the server
pkill featherjet
```

#### Windows
```cmd
# Run directly
featherjet.exe

# Run as background service (requires additional setup)
# Consider using NSSM or similar service wrapper
```

## 🔧 Configuration

FeatherJet uses a YAML configuration file (`config.yaml` by default):

```yaml
# Server settings
server:
  host: "localhost"          # Bind address (0.0.0.0 for all interfaces)
  port: 8081                # Port to listen on
  read_timeout: "30s"       # Request read timeout
  write_timeout: "30s"      # Response write timeout  
  idle_timeout: "120s"      # Keep-alive timeout

# Static file serving
static:
  directory: "./public"     # Directory containing static files
  cache_max_age: "3600"    # Cache-Control header value (seconds)

# Logging configuration
logging:
  level: "info"            # Log level: debug, info, warn, error
  enable_request_logging: true  # Log all HTTP requests

# Middleware settings
middleware:
  enable_cors: true        # Enable CORS headers
  enable_compression: false # Enable gzip compression (future)
```

### Configuration Options

| Section | Option | Default | Description |
|---------|--------|---------|-------------|
| `server.host` | string | `localhost` | Server bind address |
| `server.port` | int | `8081` | Server port |
| `server.read_timeout` | duration | `30s` | Request read timeout |
| `server.write_timeout` | duration | `30s` | Response write timeout |
| `server.idle_timeout` | duration | `120s` | Connection idle timeout |
| `static.directory` | string | `./public` | Static files directory |
| `static.cache_max_age` | string | `3600` | Cache-Control max-age |
| `logging.level` | string | `info` | Log level |
| `logging.enable_request_logging` | bool | `true` | Enable request logging |
| `middleware.enable_cors` | bool | `true` | Enable CORS middleware |
| `middleware.enable_compression` | bool | `false` | Enable compression |

## 🚀 Deploying Applications

### Static Frontend Application

1. **Prepare your static files**:
   ```
   my-app/
   ├── index.html
   ├── css/
   │   └── styles.css
   ├── js/
   │   └── app.js
   └── images/
       └── logo.png
   ```

2. **Update configuration**:
   ```yaml
   static:
     directory: "./my-app"
   ```
3. **OR replace the Featherjet public to Velocity Task static content**:
   ```bash
      cp -r $HOME/VelocityTasks/web/* $HOME/FeatherJet/public/
      ```

4. **Start FeatherJet**:
   ```bash
   ./featherjet
   ```

### Backend API Application

1. **Create custom handlers** in `internal/server/server.go`:
   ```go
   // Add to setupRoutes() method
   s.mux.HandleFunc("/api/users", s.handleUsers)
   s.mux.HandleFunc("/api/products", s.handleProducts)
   s.mux.HandleFunc("/api/tasks/", s.handleTasksProxy)
	s.mux.HandleFunc("/api/tasks", s.handleTasksProxy)
   ```

2. **Implement handler methods**:
  
   ```go
   func (s *Server) handleTasksProxy(w http.ResponseWriter, r *http.Request) {
    target, _ := url.Parse("http://localhost:8080") // Replace PORT with VelocityTasks port
    proxy := httputil.NewSingleHostReverseProxy(target)
    proxy.ServeHTTP(w, r)
   }
   ```

3. **Rebuild and deploy**:
   ```bash
   go build -o featherjet ./cmd/featherjet
   ./featherjet
   ```

### Full-Stack Application

Combine both approaches:
- Place frontend files in the the featherjet public folder
- Add API endpoints for backend functionality
- Frontend JavaScript can call `/api/*` endpoints

## ⚠️Note on Port Configuration
By default, the FeatherJet server runs on port 8080. In this setup, it has been reconfigured to run on port 8081 to allow the Velocity Tasks application to use port 8080.

If your application only relies on FeatherJet capabilities and does not require Velocity Taska, you may revert FeatherJet back to its default port from `config.yaml`

## 🧪 Running Tests

### Unit Tests
```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests with detailed output
go test -v ./...

# Run specific test file
go test ./tests/server_test.go
```

### Integration Tests
```bash
# Start server in test mode
./featherjet -config test-config.yaml &

# Run your integration tests
curl http://localhost:8081/api/hello
curl http://localhost:8081/api/status
curl http://localhost:8081/api/info

# Stop test server
pkill featherjet
```

### Load Testing
```bash
# Using Apache Bench (ab)
ab -n 1000 -c 10 http://localhost:8081/

# Using wrk
wrk -t12 -c400 -d30s http://localhost:8081/

# Using curl for API endpoints
for i in {1..100}; do curl http://localhost:8081/api/hello; done
```

## 🏗️ Development

### Adding New Features

1. **API Endpoints**: Add handlers in `internal/server/server.go`
2. **Middleware**: Add middleware functions in `internal/middleware/middleware.go`
3. **Configuration**: Extend the config struct in `internal/config/config.go`
4. **Static Assets**: Place files in the `public/` directory

### Project Structure Guidelines

- `cmd/`: Application entrypoints
- `internal/`: Private application code
- `pkg/`: Public library code (if needed)
- `tests/`: Test files
- `examples/`: Example applications
- `docs/`: Additional documentation

### Code Style

- Follow standard Go conventions
- Use `gofmt` for formatting
- Add comments for public functions
- Include tests for new features
- Keep functions small and focused

## 🤝 Contributing

We welcome contributions! Here's how to get started:

### Fork and Setup
```bash
# Fork the repository on GitHub
# Clone your fork
git clone https://github.com/YOUR_USERNAME/featherjet.git
cd featherjet

# Add upstream remote
git remote add upstream https://github.com/featherjet/featherjet.git
```

### Development Workflow
```bash
# Create a feature branch
git checkout -b feature/your-feature-name

# Make your changes
# Add tests for new functionality
# Ensure all tests pass
go test ./...

# Commit your changes
git commit -m "Add your feature description"

# Push to your fork
git push origin feature/your-feature-name

# Create a Pull Request on GitHub
```

### Contribution Guidelines

- **Issues**: Use GitHub issues for bugs and feature requests
- **Pull Requests**: Include tests and documentation
- **Code Review**: All changes require review
- **Commit Messages**: Use clear, descriptive commit messages
- **Breaking Changes**: Must be documented and discussed

### Areas for Contribution

- 🔧 **Features**: Compression, authentication, rate limiting
- 📚 **Documentation**: Tutorials, examples, API docs  
- 🧪 **Testing**: More test coverage, benchmarks
- 🎨 **Examples**: Sample applications, templates
- 🐛 **Bug Fixes**: Issues and improvements
- 🔍 **Performance**: Optimizations and profiling

## 📝 API Reference

### Built-in Endpoints

#### `GET /metrics`
Prometheus metrics endpoint for monitoring.

**Response:**
```text
# HELP http_requests_total Total number of HTTP requests
# TYPE http_requests_total counter
http_requests_total{method="GET",path="/api/hello",status="200"} 24

# HELP http_request_duration_seconds HTTP request duration in seconds
# TYPE http_request_duration_seconds histogram
http_request_duration_seconds_bucket{method="GET",path="/api/hello",le="0.005"} 18
...
```

#### `GET /api/hello`
Simple hello world endpoint for testing.

**Response:**
```json
{
  "message": "Hello from FeatherJet!",
  "timestamp": "2025-09-02T10:30:00Z",
  "method": "GET",
  "path": "/api/hello"
}
```

#### `GET /api/status`
Server health and status information.

**Response:**
```json
{
  "status": "healthy",
  "server": "FeatherJet", 
  "version": "1.0.0",
  "timestamp": "2025-09-02T10:30:00Z",
  "uptime": "2h15m30s"
}
```

#### `GET /api/info`
Detailed server configuration and runtime information.

**Response:**
```json
{
  "server": {
    "name": "FeatherJet",
    "version": "1.0.0", 
    "host": "localhost",
    "port": 8081
  },
  "static": {
    "directory": "./public",
    "exists": true,
    "is_directory": true
  },
  "middleware": {
    "cors_enabled": true,
    "compression_enabled": false,
    "request_logging": true
  },
  "timestamp": "2025-09-02T10:30:00Z"
}
```

## 🔒 Security

### Security Features

- **Security Headers**: X-Content-Type-Options, X-Frame-Options, X-XSS-Protection
- **CORS Support**: Configurable Cross-Origin Resource Sharing
- **Rate Limiting**: IP-based rate limiting to prevent abuse
- **Input Validation**: Request validation and sanitization
- **Timeouts**: Configurable request/response timeouts
- **Static File Security**: Directory traversal protection
- **Metrics**: Prometheus metrics for monitoring and alerting

### Security Best Practices

1. **Configure Timeouts**: Set appropriate timeout values
2. **Restrict CORS**: Configure CORS for production use
3. **Use HTTPS**: Deploy behind reverse proxy with TLS
4. **Validate Input**: Sanitize all user input
5. **Monitor Logs**: Enable request logging for security monitoring
6. **Update Dependencies**: Keep Go and dependencies updated

## 📈 Performance

### Benchmarks

Approximate performance on a modern server:

- **Static Files**: ~50,000 requests/second
- **API Endpoints**: ~30,000 requests/second  
- **Memory Usage**: ~10-20MB base
- **Startup Time**: <1 second
- **Response Time**: <1ms for static files, <5ms for APIs

### Performance Tuning

1. **Increase Limits**:
   ```yaml
   server:
     read_timeout: "60s"
     write_timeout: "60s" 
     idle_timeout: "300s"
   ```

2. **Enable Caching**:
   ```yaml
   static:
     cache_max_age: "86400"  # 24 hours
   ```

3. **OS Tuning** (Linux):
   ```bash
   # Increase file descriptor limits
   ulimit -n 65536
   
   # Tune TCP settings
   echo 'net.core.somaxconn = 1024' >> /etc/sysctl.conf
   ```

##  Acknowledgments
- **[Sakshi Pachlaniya](https://github.com/SakshiP3103)**: Core contributor

---

**Start serving with FeatherJet! 🚀**
