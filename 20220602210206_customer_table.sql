-- +goose Up
-- +goose StatementBegin
CREATE TABLE public.customer (
	id varchar NULL,
	fullname varchar NULL,
	dob varchar NULL,
	email varchar NULL,
	CONSTRAINT customer_pk PRIMARY KEY (id)
);

INSERT INTO public.customer
(id, fullname, dob, email)
VALUES('9ab116d6-5139-4e64-9bec-cbcac4688eef', 'Aditya', NULL, 'aditya@gmail.com');


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
