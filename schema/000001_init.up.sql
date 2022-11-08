CREATE TABLE license
(
    id SERIAL NOT NULL UNIQUE,
    uuid VARCHAR(255) NOT NULL,
    hardware_parameters VARCHAR(255) NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT FALSE,
    activated_on DATE,
    expirated_on DATE
);