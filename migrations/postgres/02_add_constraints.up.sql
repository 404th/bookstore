ALTER TABLE author ADD CONSTRAINT unique_fullname UNIQUE (firstname, secondname);

ALTER TABLE book
    ADD CONSTRAINT fk_book_author
        FOREIGN KEY (author_id)
            REFERENCES author (id);

ALTER TABLE book
    ADD CONSTRAINT fk_book_category
        FOREIGN KEY (category_id)
            REFERENCES book_category(id);