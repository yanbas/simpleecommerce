-- +goose Up
-- +goose StatementBegin
CREATE TABLE public.products (
	id varchar NULL,
	brand_id varchar NULL,
	"name" varchar NULL,
	uom varchar NULL,
	price numeric NULL,
	stock int NULL,
	images varchar NULL,
	sku varchar NULL,
	CONSTRAINT products_pk PRIMARY KEY (id)
);

INSERT INTO public.products
(id, brand_id, "name", uom, price, stock, images, sku)
VALUES('46fdce60-291b-4873-9bb4-cde998ff272a', 'cb8c42e6-a1f1-411b-a701-d5abbec1c819', 'Peptisol', 'pcs', 9000, 100, '', 'P123');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
