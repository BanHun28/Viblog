SELECT 'CREATE DATABASE viblog'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'viblog')\gexec
