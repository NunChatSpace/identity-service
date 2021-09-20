CREATE TABLE "roles"(
    "id"                        text,
    "created_at"                timestamptz NOT NULL,
    "updated_at"                timestamptz NOT NULL,
    "deleted_at"                timestamptz,
    "role_name_id"              text,
    "permission_id"             text,

    PRIMARY KEY("id"),
    CONSTRAINT "fk_roles_role_names" FOREIGN KEY("role_name_id") REFERENCES "role_names" ("id"),
    CONSTRAINT "fk_roles_permissions" FOREIGN KEY("permission_id") REFERENCES "permissions" ("id")
);

CREATE INDEX "idx_roles_deleted_at" ON "roles" ("deleted_at");