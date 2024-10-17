-- +goose Up
ALTER TABLE subgrupos
ADD CONSTRAINT FK_grupo_id
FOREIGN KEY (grupo_id) REFERENCES grupos(id)
ON UPDATE CASCADE;

ALTER TABLE subsectores
ADD CONSTRAINT FK_sector_id
FOREIGN KEY (sector_id) REFERENCES sectores(id)
ON UPDATE CASCADE;

ALTER TABLE usuario
ADD CONSTRAINT UC_email UNIQUE (email(255));

ALTER TABLE ke_opti
ADD CONSTRAINT FKOpti_cliente_id FOREIGN KEY (cliente_id) REFERENCES cliempre(id)
ON UPDATE CASCADE;

ALTER TABLE ke_opmv
ADD CONSTRAINT FK_pedido_id FOREIGN KEY (pedido_id) REFERENCES ke_opti(id)
ON UPDATE CASCADE,
ADD CONSTRAINT FKOpmv_articulo_id FOREIGN KEY (articulo_id) REFERENCES articulo(id)
ON UPDATE CASCADE;

ALTER TABLE ke_doccti
ADD CONSTRAINT FKDocc_cliente_id FOREIGN KEY (cliente_id) REFERENCES cliempre(id)
ON UPDATE CASCADE;

ALTER TABLE ke_doclmv
ADD CONSTRAINT FK_doc_id FOREIGN KEY (doc_id) REFERENCES ke_doccti(id)
ON UPDATE CASCADE,
ADD CONSTRAINT FKDocl_articulo_id FOREIGN KEY (articulo_id) REFERENCES articulo(id)
ON UPDATE CASCADE;

-- +goose Down
ALTER TABLE subgrupos
DROP FOREIGN KEY FK_grupo_id;

ALTER TABLE subsectores
DROP FOREIGN KEY FK_sector_id;

ALTER TABLE datos_usuario
DROP FOREIGN KEY FK_user_id;

ALTER TABLE ke_opmv
DROP FOREIGN KEY FK_pedido_id,
DROP FOREIGN KEY FKOpmv_articulo_id;

ALTER TABLE ke_opti
DROP FOREIGN KEY FKOpti_cliente_id;

ALTER TABLE ke_doclmv
DROP FOREIGN KEY FK_doc_id,
DROP FOREIGN KEY FKDocl_articulo_id;

ALTER TABLE ke_doccti
DROP FOREIGN KEY FKDocc_cliente_id;


