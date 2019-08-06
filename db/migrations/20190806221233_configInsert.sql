
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

INSERT INTO miu.configs
values (1, 'sendgridApikey', 'string', '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');

INSERT INTO miu.configs
values (2, 'storageAccountName', 'string', 'miustore');

INSERT INTO miu.configs
values (3, 'storageAccountKey', 'string', 'R6roxVYFKZlmd3UNaZXFKntLEuU3Q8MdiW4y9vY1lrh96ZO9CWUDLidhg/0iaWGCcBUmcqW8xKol5WPUKXrVqw==');

INSERT INTO miu.configs
values (4, 'storageContainerName', 'string', 'miu-clothing');

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DELETE FROM miu.configs WHERE ID IN (1, 2, 3, 4);