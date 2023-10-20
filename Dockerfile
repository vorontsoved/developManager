FROM postgres:latest

# Задаем переменные окружения для настройки PostgreSQL
ENV POSTGRES_DB=developManager
ENV POSTGRES_USER=postgres
ENV POSTGRES_PASSWORD=zxc
ENV PS=dcsac

EXPOSE 5432
