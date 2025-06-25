# Mapnests Go SDK Sample Application

This is a **sample Go application** that demonstrates how to use the [Mapnests Go SDK](git@github.com:mapnests/sample-app-go.git) to access secure Mapnests APIs such as:

- ✅ Search (Geocode)
- ✅ Reverse (Reverse Geocode)
- ✅ Distance Matrix
- ✅ Distance Matrix Details

---

## 🚀 How It Works

This app uses the SDK like a regular Go package:

```go
import mapnests "github.com/mapnests/sdk-go"
```

Then it initializes the SDK like this:

```go
sdk := mapnests.NewClient("YOUR_API_KEY", "PACKAGE_NAME")
```

---

## 📦 Requirements

- Go ≥ 1.18
- Your API key and package identifier

---

## 🛠 Setup

### 1. Clone the SDK and this example

```bash
git clone git@github.com:mapnests/sample-app-go.git
```

### 2. Update package:

```bash
go get -u ./...

```

### 2. Download Go dependencies

```bash
go mod tidy
```

---

## ▶️ Run the Sample

```bash
go run main.go
```
