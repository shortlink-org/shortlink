-- Removing the primary key constraint
ALTER TABLE link.link_view
    DROP CONSTRAINT link_view_pk;

-- Resetting the replica identity (assuming it was DEFAULT before)
ALTER TABLE link.link_view REPLICA IDENTITY DEFAULT;
