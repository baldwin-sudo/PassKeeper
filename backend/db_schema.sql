-- Create the users table to store user information
CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT UNIQUE NOT NULL,
 
    hashed_master_password TEXT NOT NULL, -- This will store the hashed master password
    salt TEXT NOT NULL, -- Salt used for hashing the password
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- Create the passwords table, referencing the user who owns the password
CREATE TABLE passwords (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    website TEXT NOT NULL,
    description TEXT,
    email TEXT,
    username TEXT,
    password TEXT NOT NULL, -- This will store the encrypted password
    access_count INTEGER DEFAULT 0, -- Tracks how often each password is accessed
    
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    user_id INTEGER, -- Reference to the user who owns the password
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
