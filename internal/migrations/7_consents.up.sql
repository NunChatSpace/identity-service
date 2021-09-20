CREATE TABLE "consents" (
    "id"                        text,
    "created_at"                timestamptz NOT NULL,
    "updated_at"                timestamptz NOT NULL,
    "deleted_at"                timestamptz,
    "action_allow_id"           text,
    "user_id"                   text,

    PRIMARY KEY("id"),
    CONSTRAINT "fk_consents_action_allows" FOREIGN KEY("action_allow_id") REFERENCES "action_allows" ("id"),
    CONSTRAINT "fk_consents_users" FOREIGN KEY("user_id") REFERENCES "users" ("id")
);
CREATE INDEX "idx_consents_deleted_at" ON "consents" ("deleted_at");