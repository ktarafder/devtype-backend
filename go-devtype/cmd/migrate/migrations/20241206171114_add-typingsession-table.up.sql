CREATE TABLE IF NOT EXISTS typing_session (
    session_id INT AUTO_INCREMENT PRIMARY KEY, -- Unique session identifier
    overall_accuracy DECIMAL(5, 2) NOT NULL,  -- Percentage accuracy (e.g., 95.50%)
    overall_speed DECIMAL(5, 2) NOT NULL,     -- Speed in WPM or similar metric
    user_id INT UNSIGNED NOT NULL,                     -- Foreign key to User table
    snippet_id INT NOT NULL,                  -- Foreign key to Code_snippet table
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT fk_snippet FOREIGN KEY (snippet_id) REFERENCES code_snippets(snippet_id)
);