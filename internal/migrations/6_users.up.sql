CREATE TABLE "users" (
    "id"                        text,
    "created_at"                timestamptz NOT NULL,
    "updated_at"                timestamptz NOT NULL,
    "deleted_at"                timestamptz,
    "first_name"                text,
    "middle_name"               text,
    "last_name"                 text NOT NULL,
    "contact_id"                text NOT NULL,
    "password"                  text,
    "role_id"                   text,

    PRIMARY KEY("id"),
    CONSTRAINT "fk_users_roles" FOREIGN KEY("role_id") REFERENCES "roles" ("id"),
    CONSTRAINT "fk_users_contacts" FOREIGN KEY("contact_id")  REFERENCES "contacts" ("id")
);
CREATE INDEX "idx_users_deleted_at" ON "users" ("deleted_at");
