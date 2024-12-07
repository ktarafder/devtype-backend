CREATE TABLE leaderboard (
    `rank` INT PRIMARY KEY,                          -- Rank (1, 2, 3)
    user_id INT UNSIGNED NOT NULL,                   -- Foreign key to User table
    total_score DECIMAL(10, 2) NOT NULL,             -- User's total score
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, -- Last update time
    CONSTRAINT fk_user_leaderboard FOREIGN KEY (user_id) REFERENCES users(id)
);
