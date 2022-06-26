ALTER TABLE book DROP CONSTRAINT fk_book_category;

ALTER TABLE book DROP CONSTRAINT fk_book_author;

ALTER TABLE author DROP CONSTRAINT unique_fullname;