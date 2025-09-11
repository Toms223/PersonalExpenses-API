
CREATE TABLE clients (
    id UUID PRIMARY KEY,
    secret UUID NOT NULL
);


CREATE TABLE client_redirect_uris (
    client_id UUID NOT NULL,
    uri TEXT NOT NULL,
    PRIMARY KEY (client_id, uri),
    FOREIGN KEY (client_id) REFERENCES clients(id) ON DELETE CASCADE
);

CREATE TABLE refresh_tokens (
    value UUID PRIMARY KEY,
    expiry BIGINT NOT NULL,
    user_id INT NOT NULL,
    client_id UUID NOT NULL,
    CONSTRAINT fk_client
    FOREIGN KEY (client_id)
    REFERENCES clients(id)
    ON DELETE CASCADE
);

CREATE TABLE access_tokens (
    value TEXT PRIMARY KEY,
    client_id UUID NOT NULL,
    CONSTRAINT fk_client
    FOREIGN KEY (client_id)
    REFERENCES clients(id)
    ON DELETE CASCADE
);
