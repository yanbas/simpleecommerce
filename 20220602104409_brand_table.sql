-- +goose Up
-- +goose StatementBegin
CREATE TABLE public.brand (
	id varchar NULL,
	"name" varchar NULL,
	images varchar NULL,
	CONSTRAINT brand_pk PRIMARY KEY (id)
);

INSERT INTO public.brand
(id, "name", images)
VALUES('cb8c42e6-a1f1-411b-a701-d5abbec1c819', 'Brand A', 'https://s3.amazonaws.com/www-inside-design/uploads/2019/05/woolmarkimagelogo-1024x576.png');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
