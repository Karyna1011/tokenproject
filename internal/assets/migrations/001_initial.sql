-- +migrate Up

CREATE TABLE token
(
    id        bigserial        not null,
    asset        text    not null,
    addresslp     text         not null,
    vault  text         not null
);

-- +migrate Down

DROP TABLE token;
