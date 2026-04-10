# Chirpy API 🐦

Backend-сервис на Go для социальной сети (аналог Twitter), созданный в рамках обучения на Boot.dev. 

## 🚀 Основные возможности

- **REST API**: Эндпоинты для регистрации, логина и управления сообщениями (чирпами).
- **База данных**: PostgreSQL + `sqlc` для типобезопасной работы с SQL.
- **Аутентификация**: 
  - JWT Access/Refresh токены.
  - Хеширование паролей (bcrypt).
- **Безопасность**: Защита вебхуков Polka через API Key.
- **Функционал**:
  - Фильтрация "плохих" слов.
  - Сортировка и фильтрация чирпов по автору (Query params).

## 🛠 Стек технологий

- **Язык**: Go
- **БД**: PostgreSQL
- **Инструменты**: `sqlc`, `go-chi`, `golang-jwt`, `godotenv`.

## 📦 Как запустить

1. **Клонируйте репозиторий:**
   ```bash
   git clone [https://github.com/YOUR_USERNAME/chirpy.git](https://github.com/YOUR_USERNAME/chirpy.git)
