-- +goose Up
ALTER TABLE usuario
ADD CONSTRAINT FK_vendedor_id
FOREIGN KEY (vendedor_id) REFERENCES vendedor(id)
ON UPDATE CASCADE,
ADD CONSTRAINT FK_cliente_id
FOREIGN KEY (cliente_id) REFERENCES cliente(id)
ON UPDATE CASCADE;

ALTER TABLE vendedor
ADD CONSTRAINT UC_vendedor_email UNIQUE (email(255));

ALTER TABLE cliente
ADD CONSTRAINT UC_cliente_email UNIQUE (email(255));

ALTER TABLE ke_nivgcia
ADD CONSTRAINT UC_gerencia_coordinacion UNIQUE (kng_codgcia, kng_codcoord);

ALTER TABLE ke_opmv
ADD CONSTRAINT FK_pedido_id 
FOREIGN KEY (pedido_id) REFERENCES ke_opti(id)
ON UPDATE CASCADE,
ADD CONSTRAINT FKOpmv_articulo_id 
FOREIGN KEY (articulo_id) REFERENCES articulo(id)
ON UPDATE CASCADE,
ADD CONSTRAINT UC_ndoc_codart UNIQUE (kti_ndoc, kmv_codart);

ALTER TABLE ke_doclmv
ADD CONSTRAINT FK_doc_id 
FOREIGN KEY (doc_id) REFERENCES ke_doccti(id)
ON UPDATE CASCADE,
ADD CONSTRAINT FKDocl_articulo_id 
FOREIGN KEY (articulo_id) REFERENCES articulo(id)
ON UPDATE CASCADE,
ADD CONSTRAINT UC_doc_codart UNIQUE (documento, codigo);

-- +goose Down
ALTER TABLE usuario
DROP FOREIGN KEY FK_vendedor_id,
DROP FOREIGN KEY FK_cliente_id;

ALTER TABLE ke_opmv
DROP FOREIGN KEY FK_pedido_id,
DROP FOREIGN KEY FKOpmv_articulo_id;

ALTER TABLE ke_doclmv
DROP FOREIGN KEY FK_doc_id,
DROP FOREIGN KEY FKDocl_articulo_id;

