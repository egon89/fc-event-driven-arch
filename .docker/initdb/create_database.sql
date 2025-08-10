-- postgresql statement to create a database event_driven_arch if it does not exist
SELECT 'CREATE DATABASE event_driven_arch'
WHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = 'event_driven_arch')\gexec
