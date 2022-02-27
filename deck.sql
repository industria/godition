

CREATE TABLE current (
    product_id      TEXT NOT NULL,
    product_name    TEXT NOT NUll,
    product_type    TEXT NOT NULL,
    edition_id      TEXT NOT NULL
    edition_name    TEXT NOT NULL,
    burn_id         TEXT NOT NULL,
    PRIMARY KEY (product_name, edition_name)
)


CREATE TABLE burn (
    burn_id     TEXT NOT NULL,
    edition_id  TEXT NOT NULL,
    html_hash   TEXT NOT NULL,
    css_hash    TEXT NOT NULL,
    PRIMARY KEY (burn_id)
)

CREATE TABLE deck (
    html_hash       TEXT NOT NULL,
    deck_number     INT NOT NULL,
    html            TEXT NOT NULL,
    PRIMARY KEY (html_hash, deck_number)
)

CREATE TABLE css (
    css_hash    text not null,
    css         text not null,
    PRIMARY KEY (css_hash)
)
