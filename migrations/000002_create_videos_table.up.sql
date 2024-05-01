CREATE TABLE videos (
                                    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                                    title VARCHAR(100),
                                    url VARCHAR(200),
                                    path VARCHAR(200),
                                    is_processed BOOLEAN DEFAULT false,
                                    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                                    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);