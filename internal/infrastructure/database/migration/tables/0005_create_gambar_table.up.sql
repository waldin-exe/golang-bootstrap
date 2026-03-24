CREATE TABLE gambars (
  id SERIAL PRIMARY KEY,
  reference VARCHAR(50) NOT NULL,      
  foreign_id INT NOT NULL,             
  path TEXT NOT NULL,                  
  original_name TEXT,                  
  mime_type VARCHAR(100),              
  uploaded_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
