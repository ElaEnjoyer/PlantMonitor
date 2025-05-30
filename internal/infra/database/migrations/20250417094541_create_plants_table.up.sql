CREATE TABLE IF NOT EXISTS public.plants
(
    id              serial PRIMARY KEY,
    user_id         integer NOT NULL REFERENCES public.users(id),
    name            text NOT NULL,
    city            text NOT NULL,
    address         text NOT NULL,
    type            text NOT NULL,
    created_at      timestamptz NOT NULL,
    updated_at      timestamptz NOT NULL,
    deleted_at      timestamptz
)
