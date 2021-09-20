CREATE TABLE "contacts" (
    "id"                        text,
    "created_at"                timestamptz NOT NULL,
    "updated_at"                timestamptz NOT NULL,
    "deleted_at"                timestamptz,
    "email"                     text,
    "phone_number"              text,
    "address"                   text NOT NULL,
    PRIMARY KEY("id")
);
CREATE INDEX "idx_contacts_deleted_at" ON "contacts" ("deleted_at");
