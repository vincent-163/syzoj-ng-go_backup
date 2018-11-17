CREATE TABLE users (
    id bytea PRIMARY KEY NOT NULL,
    user_name varchar NOT NULL,
    auth_info jsonb NOT NULL,
    can_login bool NOT NULL,
    git_password varchar,
    user_profile_info jsonb NOT NULL DEFAULT '{}'
);
CREATE UNIQUE INDEX users_user_name ON users (user_name);
