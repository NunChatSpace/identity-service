CREATE TABLE "verify_codes" (
    "id"                        text,
    "created_at"                timestamptz NOT NULL,
    "updated_at"                timestamptz NOT NULL,
    "deleted_at"                timestamptz,
    "verify_code"               text,
    "user_id"                   text,

    PRIMARY KEY("id")
);