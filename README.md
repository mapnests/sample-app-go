# Mapnests Go SDK Sample Application

This is a **sample Go application** that demonstrates how to use the [Mapnests Go SDK](git@github.com:mapnests/sample-app-go.git) to access secure Mapnests APIs such as:

- ✅ Search (Geocode)
- ✅ Reverse (Reverse Geocode)
- ✅ Distance Matrix
- ✅ Distance Matrix Details
- ✅ Autocomplete
- ✅ Autocomplete Without Zone
- ✅ Search By Radius
- ✅ Detailed Search By PlaceId
- ✅ Pairwise Route Summary
- ✅ Multi Source Route Summary

---

## API Registration & Authentication

### Step 1: Register for API Access

**Endpoint:** `https://mapnests.com/sign-up`

- Create an account and register for API access.

### Step 2: Login to Dashboard

**Endpoint:** `https://mapnests.com/sign-in`

- Use your credentials to log in and manage your projects.

## Project Setup

### Step 1: Create a Project

1. Log in to the TNMaps Dashboard.
2. Navigate to the dashboard.
3. Go to the **Projects** section.
4. Click on **"Create Project"**.
5. Provide a **Project Name** and **Project Description**.
6. After project creation, go to the **details section**.
7. In the **Web Section**, add your **Package Name**.
8. Once successfully created, you will receive an **API Key**.

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

- Go ≥ 1.24.2
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
