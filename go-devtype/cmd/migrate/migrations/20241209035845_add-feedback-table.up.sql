CREATE TABLE feedback (
    feedback_id INT AUTO_INCREMENT PRIMARY KEY,       -- Primary Key: Unique identifier for feedback
    improvement_area VARCHAR(255) NOT NULL,          -- Area of improvement
    feedback_text TEXT NOT NULL,                     -- Detailed feedback text
    user_id INT UNSIGNED NOT NULL,                            -- Foreign Key referencing a Users table
    CONSTRAINT fk_user_ FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
