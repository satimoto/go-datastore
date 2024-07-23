-- Nodes
ALTER TABLE nodes 
    ADD COLUMN is_lsp BOOLEAN NOT NULL DEFAULT true;

ALTER TABLE nodes 
    RENAME COLUMN lsp_addr TO rpc_addr;  
