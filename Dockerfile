# Используем официальный образ PostgreSQL с Docker Hub
FROM postgres:latest

# Копирование SQL-скрипта инициализации базы данных
COPY migration/001_init_up.sql /docker-entrypoint-initdb.d/

# Открытие порта по умолчанию
EXPOSE 5432
