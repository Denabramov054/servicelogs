CREATE TABLE comments (
    id serial primary key,
    body text not nul,
    user_id integer REFERENCES users (id),
    video_id int not null,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);