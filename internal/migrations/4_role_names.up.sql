CREATE TABLE "role_names" (
    "id"                        text,
    "created_at"                timestamptz NOT NULL,
    "updated_at"                timestamptz NOT NULL,
    "deleted_at"                timestamptz,
    "name"                      text,

    PRIMARY KEY("id")
);

INSERT INTO "role_names" ("id", "created_at", "updated_at", "deleted_at", "name")
VALUES 
('D63B59F6-34EB-4FC7-9F9F-75E683CF0D4E', '2021-09-21 11:25:50.794', '2021-09-21 11:25:50.794', NULL, 'admin'),
('A3FFF83C-3C32-4022-BEB7-4D064BB14A69', '2021-09-21 11:25:50.794', '2021-09-21 11:25:50.794', NULL, 'user')