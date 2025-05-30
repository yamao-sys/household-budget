
-- +migrate Up
ALTER TABLE expenses ADD INDEX index_expenses_on_user_id_paid_at_amount(user_id, paid_at, amount);
ALTER TABLE expenses ADD INDEX index_expenses_on_user_id_paid_at_category_amount(user_id, paid_at, category, amount);

-- +migrate Down
ALTER TABLE expenses DROP INDEX index_expenses_on_user_id_paid_at_amount;
ALTER TABLE expenses DROP INDEX index_expenses_on_user_id_paid_at_category_amount;
