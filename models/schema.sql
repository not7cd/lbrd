CREATE TABLE IF NOT EXISTS leaderboard (
    id integer primary key autoincrement,
    name text
);

CREATE TABLE IF NOT EXISTS score (
    id integer primary key,
    value int,
    player text,
    last_edit int,
    leaderboard_id int,
    FOREIGN KEY(leaderboard_id) REFERENCES leaderboard(id)
);

INSERT INTO
    leaderboard (name)
VALUES
    ("udane memy");

INSERT INTO
    score (player, value, last_edit, leaderboard_id)
VALUES
    ("memiarz", 100, CURRENT_TIMESTAMP, 1);

INSERT INTO
    score (player, value, last_edit, leaderboard_id)
VALUES
    ("cb1t", 99, CURRENT_TIMESTAMP, 1);