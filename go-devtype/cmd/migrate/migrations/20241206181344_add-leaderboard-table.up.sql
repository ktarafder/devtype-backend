CREATE TABLE leaderboard (
    leaderboard_id INT AUTO_INCREMENT PRIMARY KEY, -- Unique identifier for leaderboard entries
    rank INT NOT NULL,                             -- User's rank
    ranking_date DATE NOT NULL,                    -- Date of the ranking
    total_score DECIMAL(10, 2) NOT NULL,           -- User's total score
    user_id INT NOT NULL,                          -- Foreign key to User table
    CONSTRAINT fk_user_leaderboard FOREIGN KEY (user_id) REFERENCES User(user_id)
);
