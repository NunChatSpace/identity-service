CREATE TABLE "action_allows" (
    "id"                        text,
    "created_at"                timestamptz NOT NULL,
    "updated_at"                timestamptz NOT NULL,
    "deleted_at"                timestamptz,
    "name"                      text,
    
    PRIMARY KEY("id")
);