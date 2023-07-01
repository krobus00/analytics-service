-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS analytics (
    id varchar(36) UNIQUE,
    source text NOT NULL,
    medium text NOT NULL,
	campaign text NOT NULL,
    ip CIDR NOT NULL,
    country varchar(2) NOT NULL,
    city text NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS analytics;
-- +goose StatementEnd
