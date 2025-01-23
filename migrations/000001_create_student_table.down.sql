-- Drop the table and related objects
DROP TRIGGER IF EXISTS update_student_updated_at ON student_information;
DROP FUNCTION IF EXISTS update_updated_at_column();
DROP TABLE IF EXISTS student_information;
DROP TYPE IF EXISTS gender_type;
DROP TYPE IF EXISTS course_type;
DROP TYPE IF EXISTS class_type;

-- Drop the citext extension if not needed by other tables
-- Note: Only drop if you're sure no other tables are using it
DROP EXTENSION IF EXISTS citext;
