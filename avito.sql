
DROP TABLE IF EXISTS hotel; 
CREATE TABLE hotel (
    room_id         SERIAL         NOT NULL,
    price           int         NOT NULL,
    description     CHAR(250)   NOT NULL,
    update          date        NOT NULL,
    PRIMARY KEY (room_id)
);

DROP TABLE IF EXISTS booking; 
CREATE TABLE booking (
    room_id          int        REFERENCES hotel NOT NULL,
    booking_id       SERIAL         NOT NULL,
    date_start       date        NOT NULL,
    date_end         date        NOT NULL,
    PRIMARY KEY (booking_id)
);