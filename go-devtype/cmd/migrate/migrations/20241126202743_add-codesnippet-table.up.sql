CREATE TABLE IF NOT EXISTS code_snippets (
    snippet_id INT AUTO_INCREMENT PRIMARY KEY,  
    code_language VARCHAR(50) NOT NULL, 
    difficulty_level VARCHAR(20) NOT NULL, 
    snippet_text TEXT NOT NULL  
);
