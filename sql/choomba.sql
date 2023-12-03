CREATE TABLE IF NOT EXISTS choomba (
    url TEXT PRIMARY KEY,
    mediatype TEXT,
    tags TEXT,
    auxlinks TEXT,
    summary TEXT,
    savedAt INTEGER,
);