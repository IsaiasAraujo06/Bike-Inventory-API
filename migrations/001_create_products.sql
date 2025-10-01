-- Criação da tabela de produtos --
CREATE TABLE products (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    brand TEXT,
    price REAL NOT NULL,
    quantity INTEGER NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- Inserção de produtos iniciais --
INSERT INTO products (name, brand, price, quantity) VALUES
('Esteira elétrica Movement LX 160 G4', 'Movement', 17000.00, 4),
('Esteira elétrica Movement Aria iLed', 'Movement', 23000.00, 1),
('Ergométrica vertical Movement RTU', 'Movement', 18000.00, 5);