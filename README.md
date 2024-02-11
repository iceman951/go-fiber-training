# calcal

## 📘 About

calcal คือ API Service ที่ให้บริการ...

## 🗂️ Table of Contents

- [example\_name](#calcal)
  - [📘 About](#-about)
  - [🗂️ Table of Contents](#️-table-of-contents)
  - [📦 Built With](#-built-with)
  - [⚒ Structure](#-structure)
  - [🏷 Versions](#-versions)
    - [v0.1.0 - YYYY-MM-DD](#v010---yyyy-mm-dd)
  - [📋 Features](#-features)
  - [⚙️ Get Started](#️-get-started)

## 📦 Built With

- [x] Go 1.21 (with GoFiber)
- [x] Viper - Configuration Management
- [x] Zerolog - Log Management (wrapped with baac-tech/zlogwrap)

## ⚒ Structure

```mermaid
  graph LR;
    Requester --> |Called|calcal;
```

## 🏷 Versions

### v0.1.0 - YYYY-MM-DD

- Initialized
- Health API (`/api/health`)

... [more](./CHANGELOG.md)

## 📋 Features

- `/api/health` | for get health status

## ⚙️ Get Started

1. Clone project

   ```bash
   git clone https://ipanda.baac.tech/calcal/calcal.git
   ```

2. Go to project folder

   ```bash
   cd calcal
   ```

3. Set up environment

   ```bash
   export ENV=dev
   ```

4. Run project by command

   ```shell
   # Normal Mode
   go run main.go

   # Test Mode
   go test ./... -v

   # Air (Hot Reloader)
   air

   # Also work with Make command
   make run
   # or #
   make air
   ```

Made with ❤️
Powered by Watsize boilerplate (newGoAPI_v4.0)
