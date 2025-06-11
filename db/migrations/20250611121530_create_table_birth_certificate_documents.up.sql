CREATE TABLE birth_certificate_documents (
  id INT NOT NULL AUTO_INCREMENT,
  statement_type_item_id INT NOT NULL,
  user_id INT NOT NULL,
  rt_rw_file_path VARCHAR(100) NOT NULL UNIQUE,
  formulir_file_path VARCHAR(100) NOT NULL UNIQUE,
  surat_kelahiran_file_path VARCHAR(100) NOT NULL UNIQUE,
  buku_nikah_file_path VARCHAR(100) NOT NULL UNIQUE,
  kk_file_path VARCHAR(100) NOT NULL UNIQUE,
  ktp_pelapor_file_path VARCHAR(100) NOT NULL UNIQUE,
  ktp_saksi_1_file_path VARCHAR(100) NOT NULL UNIQUE,
  ktp_saksi_2_file_path VARCHAR(100) NOT NULL UNIQUE,
  ktp_orang_tua_file_path VARCHAR(100) NOT NULL UNIQUE,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (id),
  FOREIGN KEY (statement_type_item_id) REFERENCES statement_type_items(id) ON DELETE CASCADE,
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
) ENGINE = InnoDB;