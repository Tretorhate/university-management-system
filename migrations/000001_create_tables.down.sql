-- Drop tables in reverse order to respect foreign key dependencies
DROP TABLE IF EXISTS public.enrollments;
DROP TABLE IF EXISTS public.courses;
DROP TABLE IF EXISTS public.teachers;
DROP TABLE IF EXISTS public.students;
DROP TABLE IF EXISTS public.users;

-- Optionally drop the schema if no other objects remain
DROP SCHEMA IF EXISTS public CASCADE;