CREATE TABLE users (
    id bytea PRIMARY KEY NOT NULL CHECK(LENGTH(id) = 16),
    name VARCHAR NOT NULL,
    auth_info jsonb NOT NULL,
    profile_info jsonb NOT NULL DEFAULT '{}'::jsonb,
    can_login bool NOT NULL,
    git_password VARCHAR,
	CONSTRAINT users_name_unique UNIQUE(name),
	CONSTRAINT users_name_check CHECK(LENGTH(name) BETWEEN 3 AND 64)
);
CREATE TABLE groups (
    id bytea PRIMARY KEY NOT NULL CHECK(LENGTH(id) = 16),
    name VARCHAR NOT NULL,
    policy_info jsonb NOT NULL DEFAULT '{}'::jsonb,
	CONSTRAINT groups_name_unique UNIQUE(name),
	CONSTRAINT groups_name_check CHECK(LENGTH(name) BETWEEN 3 AND 64)
);
CREATE TABLE group_users (
    group_id bytea NOT NULL REFERENCES groups(id),
    user_id bytea NOT NULL REFERENCES users(id),
    role_info jsonb NOT NULL DEFAULT '{}'::jsonb,
    PRIMARY KEY (group_id, user_id)
);
CREATE TABLE problemsets (
    id bytea PRIMARY KEY NOT NULL CHECK(LENGTH(id) = 16),
    name VARCHAR NOT NULL,
    group_id bytea REFERENCES groups(id),
    type VARCHAR NOT NULL,
    info jsonb NOT NULL DEFAULT '{}'::jsonb,
	CONSTRAINT problemsets_name_unique UNIQUE(group_id, name),
	CONSTRAINT problemsets_name_check CHECK(LENGTH(name) BETWEEN 1 AND 64)
);
CREATE TABLE problemset_users (
    problemset_id bytea NOT NULL REFERENCES problemsets(id),
    user_id bytea NOT NULL REFERENCES users(id),
    info jsonb NOT NULL DEFAULT '{}'::jsonb,
    PRIMARY KEY (problemset_id, user_id)
);
CREATE TABLE problems (
    id bytea PRIMARY KEY NOT NULL CHECK(LENGTH(id) = 16),
    problemset_id bytea NOT NULL REFERENCES problemsets(id),
	name VARCHAR NOT NULL,
    type VARCHAR NOT NULL,
    info jsonb NOT NULL,
    push_token VARCHAR,
	CONSTRAINT problems_name_unique UNIQUE(problemset_id, name),
	CONSTRAINT problems_name_check CHECK(LENGTH(name) BETWEEN 1 AND 64)
);
CREATE TABLE submissions (
    id bytea PRIMARY KEY NOT NULL CHECK(LENGTH(id) = 16),
    problem_id bytea NOT NULL REFERENCES problems(id),
    summary jsonb NOT NULL DEFAULT '{}'::jsonb,
    content jsonb NOT NULL DEFAULT '{}'::jsonb,
    result jsonb NOT NULL DEFAULT '{}'::jsonb
);
CREATE TABLE git_repos (
    id bytea PRIMARY KEY NOT NULL CHECK(LENGTH(id) = 16)
);
