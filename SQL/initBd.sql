CREATE TABLE Character (
                        id VARCHAR(255) PRIMARY KEY,
                        name VARCHAR(255) NOT NULL,
                        class VARCHAR(255) NOT NULL,
                        level NUMERIC(10, 2) NOT NULL,
                        mony NUMERIC(10, 2) NOT NULL
);

CREATE TABLE deadCharacter (
                            id VARCHAR(255) PRIMARY KEY,
                            name VARCHAR(255) NOT NULL,
                            cause_of_death VARCHAR(255) NOT NULL,
                            level NUMERIC(10, 2) NOT NULL
);

INSERT INTO deadCharacter (id, name, cause_of_death,level) VALUES
            ('1', 'Mardyk', 'died of autism', 20),
            ('2', 'Mirvan', 'got married', 6);

INSERT INTO Character (id, name, class, level, mony) VALUES
              ('1', 'Bertrud', 'wizard', 20, 20000),
              ('2', 'Hohus', 'wicher', 6, 5000),
              ('3', 'Feralas', 'figter', 25, 0);


DROP TABLE Character;
DROP TABLE deadCharacter;