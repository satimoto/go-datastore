-- Nodes
ALTER TABLE nodes 
    RENAME COLUMN rpc_addr TO lsp_addr;  

ALTER TABLE nodes
    DROP IF EXISTS is_lsp;
