CREATE TABLE albums (
                        id VARCHAR(255) PRIMARY KEY,
                        title VARCHAR(255) NOT NULL,
                        artist VARCHAR(255) NOT NULL,
                        price NUMERIC(10, 2) NOT NULL
);


INSERT INTO albums (id, title, artist, price) VALUES
              ('1', 'Blue Train', 'John Coltrane', 56.99),
              ('2', 'Jeru', 'Gerry Mulligan', 17.99),
              ('3', 'Sarah Vaughan and Clifford Brown', 'Sarah Vaughan', 39.99);


DROP TABLE albums;
