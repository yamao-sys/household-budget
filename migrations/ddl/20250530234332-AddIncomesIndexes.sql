
-- +migrate Up
ALTER TABLE incomes ADD INDEX index_incomes_on_user_id_received_at_amount(user_id, received_at, amount);
ALTER TABLE incomes ADD INDEX index_incomes_on_user_id_received_at_client_name_amount(user_id, received_at, client_name, amount);

-- +migrate Down
ALTER TABLE incomes DROP INDEX index_incomes_on_user_id_received_at_amount;
ALTER TABLE incomes DROP INDEX index_incomes_on_user_id_received_at_client_name_amount;
