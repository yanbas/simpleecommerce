-- +goose Up
-- +goose StatementBegin
CREATE TABLE public."order" (
	id varchar NULL,
	order_date varchar NULL,
	total numeric NULL,
	promo numeric NULL,
	subtotal numeric NULL,
	CONSTRAINT order_pk PRIMARY KEY (id)
);

CREATE TABLE public.order_item (
	id varchar NULL,
	order_id varchar NULL,
	product_id varchar NULL,
	price numeric NULL,
	qty int NULL,
	promo numeric NULL,
	total numeric NULL,
	CONSTRAINT order_item_pk PRIMARY KEY (id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
