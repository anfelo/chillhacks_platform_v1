CREATE TABLE subjects (
    id UUID PRIMARY KEY,
    title TEXT NOT NULL,
    img_url TEXT NOT NULL
);

CREATE TABLE courses (
    id UUID PRIMARY KEY,
    subject_id UUID NOT NULL REFERENCES subjects (id) ON DELETE CASCADE,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    img_url TEXT NOT NULL,
    slug TEXT NOT NULL
);

CREATE TABLE lessons (
    id UUID PRIMARY KEY,
    course_id UUID NOT NULL REFERENCES courses (id) ON DELETE CASCADE,
    title TEXT NOT NULL,
    slug TEXT NOT NULL,
    category TEXT NOT NULL,
    order INT NOT NULL
);
