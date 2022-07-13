create table if not exists users (
    user_id uuid NOT NULL DEFAULT gen_random_uuid() PRIMARY KEY,
    user_name varchar(100) NULL,
    CONSTRAINT user_name_unique UNIQUE (user_name)
);