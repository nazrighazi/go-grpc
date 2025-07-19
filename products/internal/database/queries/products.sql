-- name: AllProducts :many
SELECT * FROM products WHERE is_active = true AND deleted_at IS NULL;