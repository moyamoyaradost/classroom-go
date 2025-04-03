CREATE TABLE tasks (
    id UUID PRIMARY KEY,
    course_id UUID NOT NULL REFERENCES courses(id),
    description TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
