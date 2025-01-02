create table space_crafts (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    class ENUM(
        'Star Destroyer',
        'Super Star Destroyer',
        'Cruiser',
        'Frigate',
        'Battleship'
    ) NOT NULL,
    crew INT,
    image TEXT,
    value INT NOT NULL,
    status ENUM(
        'operational',
        'damaged',
        'destroyed',
        'decommissioned'
    ) NOT NULL DEFAULT 'operational',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
