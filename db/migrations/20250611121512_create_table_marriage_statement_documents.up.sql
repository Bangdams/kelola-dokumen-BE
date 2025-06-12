CREATE TABLE marriage_statement_documents (
  id INT NOT NULL AUTO_INCREMENT,
  statement_type_item_id INT NOT NULL,
  user_id INT NOT NULL,
  kk_file_path VARCHAR(100) NOT NULL UNIQUE,
  ktp_file_path VARCHAR(100) NOT NULL UNIQUE,
  ijazah_file_path VARCHAR(100) NOT NULL UNIQUE,
  akta_file_path VARCHAR(100) NOT NULL UNIQUE,
  ktp_saksi_file_path VARCHAR(100) NOT NULL UNIQUE,
  surat_balasan_file_path VARCHAR(255) UNIQUE,
  status ENUM("belum", "selesai") NOT NULL DEFAULT "belum",
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  FOREIGN KEY (statement_type_item_id) REFERENCES statement_type_items(id) ON DELETE CASCADE,
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
) ENGINE = InnoDB;