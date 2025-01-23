-- Create enum types
CREATE TYPE gender_type AS ENUM ('male', 'female', 'other');
CREATE TYPE course_type AS ENUM ('science', 'commerce', 'arts');
CREATE TYPE class_type AS ENUM ('six', 'seven', 'eight', 'nine', 'ten');

-- Create citext extension for case-insensitive text fields
CREATE EXTENSION IF NOT EXISTS citext;

-- Create student_information table
CREATE TABLE IF NOT EXISTS student_information (
    id SERIAL PRIMARY KEY,
    student_id BIGINT UNIQUE NOT NULL,
    first_name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    father_name VARCHAR(100) NOT NULL,
    phone VARCHAR(20) NOT NULL,
    gender gender_type NOT NULL,
    course course_type NOT NULL,
    email citext UNIQUE NOT NULL,
    is_active BOOLEAN DEFAULT true,
    class_id INTEGER NOT NULL,
    class class_type NOT NULL,
    password BYTEA NOT NULL,
    father_phone VARCHAR(20) NOT NULL,
    admission_date TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    admission_fee DECIMAL(10,2) NOT NULL,
    total_fee DECIMAL(10,2) NOT NULL,
    remaining_fee DECIMAL(10,2) NOT NULL,
    monthly_fee DECIMAL(10,2) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT positive_admission_fee CHECK (admission_fee >= 0),
    CONSTRAINT positive_total_fee CHECK (total_fee >= 0),
    CONSTRAINT positive_remaining_fee CHECK (remaining_fee >= 0),
    CONSTRAINT positive_monthly_fee CHECK (monthly_fee >= 0)
);

-- Create index on commonly queried fields
CREATE INDEX idx_student_email ON student_information(email);
CREATE INDEX idx_student_student_id ON student_information(student_id);
CREATE INDEX idx_student_class_id ON student_information(class_id);

-- Create a trigger to automatically update updated_at timestamp
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TRIGGER update_student_updated_at
    BEFORE UPDATE ON student_information
    FOR EACH ROW
    EXECUTE FUNCTION update_updated_at_column();
